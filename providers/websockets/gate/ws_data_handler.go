package gate

import (
	"encoding/json"
	"fmt"

	"go.uber.org/zap"

	"github.com/skip-mev/slinky/oracle/config"
	"github.com/skip-mev/slinky/oracle/types"
	"github.com/skip-mev/slinky/providers/base/websocket/handlers"
	mmtypes "github.com/skip-mev/slinky/x/mm2/types"
)

var _ types.PriceWebSocketDataHandler = (*WebSocketHandler)(nil)

// WebSocketHandler implements the WebSocketDataHandler interface. This is used to
// handle messages received from the Gate.io websocket API.
type WebSocketHandler struct {
	logger *zap.Logger

	// market is the config for the Gate.io API.
	market types.ProviderMarketMap
	// ws is the config for the Gate.io websocket.
	ws config.WebSocketConfig
}

// NewWebSocketDataHandler returns a new Gate.io PriceWebSocketDataHandler.
func NewWebSocketDataHandler(
	logger *zap.Logger,
	market types.ProviderMarketMap,
	ws config.WebSocketConfig,
) (types.PriceWebSocketDataHandler, error) {
	if err := market.ValidateBasic(); err != nil {
		return nil, fmt.Errorf("invalid market config for %s: %w", Name, err)
	}

	if market.Name != Name {
		return nil, fmt.Errorf("expected market config name %s, got %s", Name, market.Name)
	}

	if ws.Name != Name {
		return nil, fmt.Errorf("expected websocket config name %s, got %s", Name, ws.Name)
	}

	if !ws.Enabled {
		return nil, fmt.Errorf("websocket config for %s is not enabled", Name)
	}

	if err := ws.ValidateBasic(); err != nil {
		return nil, fmt.Errorf("invalid websocket config for %s: %w", Name, err)
	}

	return &WebSocketHandler{
		logger: logger,
		market: market,
		ws:     ws,
	}, nil
}

// HandleMessage is used to handle a message received from the data provider. The Gate.io
// provider sends two types of messages:
//
//  1. Subscribe response message. The subscribe response message is used to determine if
//     the subscription was successful.
//  2. Ticker stream message. This is sent when a ticker update is received from the
//     Gate.io websocket API.
func (h *WebSocketHandler) HandleMessage(
	message []byte,
) (types.PriceResponse, []handlers.WebsocketEncodedMessage, error) {
	var (
		resp         types.PriceResponse
		baseMessage  BaseMessage
		subResponse  SubscribeResponse
		tickerStream TickerStream
	)

	// Attempt to unmarshal the message into a base message. This is used to determine the type
	// of message that was received.
	if err := json.Unmarshal(message, &baseMessage); err != nil {
		return resp, nil, err
	}

	switch Event(baseMessage.Event) {
	case EventSubscribe:
		if err := json.Unmarshal(message, &subResponse); err != nil {
			return resp, nil, err
		}

		// handle subscription
		updateMsg, err := h.parseSubscribeResponse(subResponse)
		return resp, updateMsg, err

	case EventUpdate:
		if err := json.Unmarshal(message, &tickerStream); err != nil {
			return resp, nil, err
		}

		// update pair info
		resp, err := h.parseTickerStream(tickerStream)
		return resp, nil, err

	default:
		return resp, nil, fmt.Errorf("unknown message type %s", baseMessage.Event)
	}
}

// CreateMessages is used to create an initial subscription message to send to the data provider.
// Only the tickers that are specified in the config are subscribed to. The only channel that is
// subscribed to is the tickers channel - which supports spot markets.
func (h *WebSocketHandler) CreateMessages(
	tickers []mmtypes.Ticker,
) ([]handlers.WebsocketEncodedMessage, error) {
	instruments := make([]string, 0)

	for _, ticker := range tickers {
		market, ok := h.market.TickerConfigs[ticker]
		if !ok {
			h.logger.Debug("market not found for currency pair", zap.String("ticker", ticker.String()))
			continue
		}

		instruments = append(instruments, market.OffChainTicker)
	}

	return NewSubscribeRequest(instruments)
}

// HeartBeatMessages is not used for Gate.io.
func (h *WebSocketHandler) HeartBeatMessages() ([]handlers.WebsocketEncodedMessage, error) {
	return nil, nil
}

// Copy is used to create a copy of the WebSocketHandler.
func (h *WebSocketHandler) Copy() types.PriceWebSocketDataHandler {
	return &WebSocketHandler{
		logger: h.logger,
		market: h.market,
		ws:     h.ws,
	}
}
