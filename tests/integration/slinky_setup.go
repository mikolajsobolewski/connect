package integration

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"sort"
	"strconv"
	"sync"
	"testing"
	"time"

	"cosmossdk.io/math"
	cmtabci "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/libs/rand"
	sdk "github.com/cosmos/cosmos-sdk/types"
	grpctypes "github.com/cosmos/cosmos-sdk/types/grpc"
	govtypesv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	"github.com/pelletier/go-toml"
	interchaintest "github.com/strangelove-ventures/interchaintest/v8"
	"github.com/strangelove-ventures/interchaintest/v8/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v8/ibc"
	"github.com/strangelove-ventures/interchaintest/v8/testutil"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

	compression "github.com/skip-mev/slinky/abci/strategies/codec"
	slinkyabci "github.com/skip-mev/slinky/abci/ve/types"
	oracleconfig "github.com/skip-mev/slinky/oracle/config"
	oracletypes "github.com/skip-mev/slinky/x/oracle/types"
)

const (
	oracleConfigPath = "oracle.toml"
	appConfigPath    = "config/app.toml"
)

var (
	extCommitCodec = compression.NewCompressionExtendedCommitCodec(
		compression.NewDefaultExtendedCommitCodec(),
		compression.NewZStdCompressor(),
	)

	veCodec = compression.NewCompressionVoteExtensionCodec(
		compression.NewDefaultVoteExtensionCodec(),
		compression.NewZLibCompressor(),
	)
)

// construct the network from a spec
// ChainBuilderFromChainSpec creates an interchaintest chain builder factory given a ChainSpec
// and returns the associated chain
func ChainBuilderFromChainSpec(t *testing.T, spec *interchaintest.ChainSpec) *cosmos.CosmosChain {
	// require that NumFullNodes == NumValidators == 4
	require.Equal(t, *spec.NumValidators, 4)

	cf := interchaintest.NewBuiltinChainFactory(zaptest.NewLogger(t), []*interchaintest.ChainSpec{spec})

	chains, err := cf.Chains(t.Name())
	require.NoError(t, err)

	require.Len(t, chains, 1)
	chain := chains[0]

	cosmosChain, ok := chain.(*cosmos.CosmosChain)
	require.True(t, ok)

	return cosmosChain
}

// SetOracleConfigOnApp writes the oracle configuration to the given node's application config.
func SetOracleConfigsOnApp(node *cosmos.ChainNode, oracleConfig oracleconfig.OracleConfig) {
	// read the app config from the node
	bz, err := node.ReadFile(context.Background(), appConfigPath)
	if err != nil {
		panic(err)
	}

	// Unmarshall the app config to update the oracle and metrics file paths.
	var appConfig map[string]interface{}
	err = toml.Unmarshal(bz, &appConfig)
	if err != nil {
		panic(err)
	}

	oracleAppConfig, ok := appConfig["oracle"].(map[string]interface{})
	if !ok {
		panic("oracle config not found")
	}

	// Update the file paths to the oracle and metrics configs.
	oracleAppConfig["enabled"] = true
	oracleAppConfig["oracle_address"] = fmt.Sprintf("%s:%s", node.HostName(), "8080")
	oracleAppConfig["client_timeout"] = "1s"
	oracleAppConfig["metrics_enabled"] = false

	appConfig["oracle"] = oracleAppConfig
	bz, err = toml.Marshal(appConfig)
	if err != nil {
		panic(err)
	}

	// Write back the app config.
	err = node.WriteFile(context.Background(), bz, appConfigPath)
	if err != nil {
		panic(err)
	}

	// marshal the oracle config
	bz, err = toml.Marshal(oracleConfig)
	if err != nil {
		panic(err)
	}

	// write the oracle config to the node
	err = node.WriteFile(context.Background(), bz, oracleConfigPath)
	if err != nil {
		panic(err)
	}
}

// AddSidecarToNode adds the sidecar configured by the given config to the given node. These are configured
// so that the sidecar is started before the node is started.
func AddSidecarToNode(node *cosmos.ChainNode, conf ibc.SidecarConfig) {
	// create the sidecar process
	node.NewSidecarProcess(
		context.Background(),
		true,
		conf.ProcessName,
		node.DockerClient,
		node.NetworkID,
		conf.Image,
		conf.HomeDir,
		conf.Ports,
		conf.StartCmd,
	)
}

// spin up the network (with side-cars enabled)
// BuildPOBInterchain creates a new Interchain testing env with the configured POB CosmosChain
func BuildPOBInterchain(t *testing.T, ctx context.Context, chain ibc.Chain) *interchaintest.Interchain {
	ic := interchaintest.NewInterchain()
	ic.AddChain(chain)

	// create docker network
	client, networkID := interchaintest.DockerSetup(t)

	ctx, cancel := context.WithTimeout(ctx, 2*time.Minute)
	defer cancel()

	// build the interchain
	err := ic.Build(ctx, nil, interchaintest.InterchainBuildOptions{
		SkipPathCreation: true,
		Client:           client,
		NetworkID:        networkID,
		TestName:         t.Name(),
	})
	require.NoError(t, err)

	return ic
}

// SetOracleConfigsOnOracle writes the oracle and metrics configs to the given node's
// oracle sidecar.
func SetOracleConfigsOnOracle(
	oracle *cosmos.SidecarProcess,
	oracleCfg oracleconfig.OracleConfig,
) {
	// marshal the oracle config
	bz, err := toml.Marshal(oracleCfg)
	if err != nil {
		panic(err)
	}

	// write the oracle config to the node
	err = oracle.WriteFile(context.Background(), bz, oracleConfigPath)
	if err != nil {
		panic(err)
	}
}

// RestartOracle restarts the oracle sidecar for a given node
func RestartOracle(node *cosmos.ChainNode) error {
	if len(node.Sidecars) != 1 {
		panic("expected node to have oracle sidecar")
	}

	oracle := node.Sidecars[0]

	if err := oracle.StopContainer(context.Background()); err != nil {
		return err
	}

	return oracle.StartContainer(context.Background())
}

// StopOracle stops the oracle sidecar for a given node
func StopOracle(node *cosmos.ChainNode) error {
	if len(node.Sidecars) != 1 {
		panic("expected node to have oracle sidecar")
	}

	oracle := node.Sidecars[0]

	return oracle.StopContainer(context.Background())
}

// StartOracle starts the oracle sidecar for a given node
func StartOracle(node *cosmos.ChainNode) error {
	if len(node.Sidecars) != 1 {
		panic("expected node to have oracle sidecar")
	}

	oracle := node.Sidecars[0]

	return oracle.StartContainer(context.Background())
}

// GetChainGRPC gets a GRPC client of the given chain
//
// NOTICE: this client must be closed after use
func GetChainGRPC(chain *cosmos.CosmosChain) (cc *grpc.ClientConn, close func(), err error) {
	// get grpc address
	grpcAddr := chain.GetHostGRPCAddress()

	// create the client
	cc, err = grpc.Dial(grpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, err
	}

	return cc, func() { cc.Close() }, nil
}

// QueryCurrencyPair queries the chain for the given CurrencyPair, this method returns the grpc response from the module
func QueryCurrencyPairs(chain *cosmos.CosmosChain) (*oracletypes.GetAllCurrencyPairsResponse, error) {
	// get grpc address
	grpcAddr := chain.GetHostGRPCAddress()

	// create the client
	cc, err := grpc.Dial(grpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer cc.Close()

	// create the oracle client
	client := oracletypes.NewQueryClient(cc)

	// query the currency pairs
	return client.GetAllCurrencyPairs(context.Background(), &oracletypes.GetAllCurrencyPairsRequest{})
}

// QueryCurrencyPair queries the price for the given currency-pair given a desired height to query from
func QueryCurrencyPair(chain *cosmos.CosmosChain, cp oracletypes.CurrencyPair, height uint64) (*oracletypes.QuotePrice, int64, error) {
	grpcAddr := chain.GetHostGRPCAddress()

	// create the client
	cc, err := grpc.Dial(grpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, 0, err
	}
	defer cc.Close()

	// create the oracle client
	client := oracletypes.NewQueryClient(cc)

	ctx := context.Background()

	md := metadata.New(map[string]string{
		grpctypes.GRPCBlockHeightHeader: strconv.FormatInt(int64(height), 10),
	})

	ctx = metadata.NewOutgoingContext(ctx, md)

	// query the currency pairs
	res, err := client.GetPrice(ctx, &oracletypes.GetPriceRequest{
		CurrencyPairSelector: &oracletypes.GetPriceRequest_CurrencyPair{
			CurrencyPair: &cp,
		},
	})
	if err != nil {
		return nil, 0, err
	}

	return res.Price, int64(res.Nonce), nil
}

// Submit proposal creates and submits a proposal to the chain
func SubmitProposal(chain *cosmos.CosmosChain, deposit sdk.Coin, submitter string, msgs ...sdk.Msg) (string, error) {
	// build the proposal
	rand := rand.Str(10)
	prop, err := chain.BuildProposal(msgs, rand, rand, rand, deposit.String(), submitter, false)
	if err != nil {
		return "", err
	}

	// submit the proposal
	tx, err := chain.SubmitProposal(context.Background(), submitter, prop)
	return tx.ProposalID, err
}

// PassProposal given a proposal id, vote for the proposal and wait for it to pass
func PassProposal(chain *cosmos.CosmosChain, propId string, timeout time.Duration) error {
	if err := WaitForProposalStatus(chain, propId, timeout, govtypesv1.StatusVotingPeriod); err != nil {
		return fmt.Errorf("proposal did not enter voting period: %v", err)
	}

	// have all nodes vote on the proposal
	wg := sync.WaitGroup{}
	for _, node := range chain.Nodes() {
		wg.Add(1)
		go func(node *cosmos.ChainNode) {
			defer wg.Done()
			node.VoteOnProposal(context.Background(), validatorKey, propId, yes)
		}(node)
	}
	wg.Wait()

	// wait for the proposal to pass
	if err := WaitForProposalStatus(chain, propId, timeout, govtypesv1.StatusPassed); err != nil {
		return fmt.Errorf("proposal did not pass: %v", err)
	}
	return nil
}

// AddCurrencyPairs creates + submits the proposal to add the given currency-pairs to state, votes for the prop w/ all nodes,
// and waits for the proposal to pass.
func AddCurrencyPairs(chain *cosmos.CosmosChain, authority, denom string, deposit int64, timeout time.Duration, user cosmos.User, cps ...oracletypes.CurrencyPair) error {
	propId, err := SubmitProposal(chain, sdk.NewCoin(denom, math.NewInt(deposit)), user.KeyName(), []sdk.Msg{&oracletypes.MsgAddCurrencyPairs{
		Authority:     authority,
		CurrencyPairs: cps,
	}}...)
	if err != nil {
		return err
	}

	return PassProposal(chain, propId, timeout)
}

// RemoveCurrencyPairs creates + submits the proposal to remove the given currency-pairs from state, votes for the prop w/ all nodes,
// and waits for the proposal to pass.
func RemoveCurrencyPairs(chain *cosmos.CosmosChain, authority, denom string, deposit int64, timeout time.Duration, user cosmos.User, cpIDs ...string) error {
	propId, err := SubmitProposal(chain, sdk.NewCoin(denom, math.NewInt(deposit)), user.KeyName(), []sdk.Msg{&oracletypes.MsgRemoveCurrencyPairs{
		Authority:       authority,
		CurrencyPairIds: cpIDs,
	}}...)
	if err != nil {
		return err
	}

	return PassProposal(chain, propId, timeout)
}

// QueryProposal queries the chain for a given proposal
func QueryProposal(chain *cosmos.CosmosChain, propID string) (*govtypesv1.QueryProposalResponse, error) {
	// get grpc address
	grpcAddr := chain.GetHostGRPCAddress()

	// create the client
	cc, err := grpc.Dial(grpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer cc.Close()

	// create the oracle client
	client := govtypesv1.NewQueryClient(cc)

	propId, err := strconv.ParseUint(propID, 10, 64)
	if err != nil {
		return nil, err
	}
	// query the currency pairs
	return client.Proposal(context.Background(), &govtypesv1.QueryProposalRequest{
		ProposalId: propId,
	})
}

// WaitForVotingPeriod, waits for the deposit period for the proposal to end
func WaitForProposalStatus(chain *cosmos.CosmosChain, propID string, timeout time.Duration, status govtypesv1.ProposalStatus) error {
	return testutil.WaitForCondition(timeout, 1*time.Second, func() (bool, error) {
		prop, err := QueryProposal(chain, propID)
		if err != nil {
			return false, err
		}

		return prop.Proposal.Status == status, nil
	})
}

// WaitForHeight waits for the giuve height to be reached
func WaitForHeight(chain *cosmos.CosmosChain, height uint64, timeout time.Duration) error {
	return testutil.WaitForCondition(timeout, 100*time.Millisecond, func() (bool, error) {
		h, err := chain.Height(context.Background())
		if err != nil {
			return false, err
		}

		return h >= height, nil
	})
}

// ExpectVoteExtensions waits for empty oracle update waits for a pre-determined number of blocks for an extended commit with the given oracle-vote extensions provided
// per validator. This method returns the height at which the condition was satisfied.
//
// Notice: the height returned is safe for querying, i.e the prices will have been written to state if a quorum reported
func ExpectVoteExtensions(chain *cosmos.CosmosChain, timeout time.Duration, ves []slinkyabci.OracleVoteExtension) (uint64, error) {
	client := chain.Nodes()[0].Client

	var blockHeight uint64
	if err := testutil.WaitForCondition(timeout, 100*time.Millisecond, func() (bool, error) {
		var err error

		blockHeight, err = chain.Height(context.Background())
		if err != nil {
			return false, err
		}

		height := int64(blockHeight)
		// get the block
		block, err := client.Block(context.Background(), &height)
		if err != nil {
			return false, err
		}

		// get the oracle update
		if len(block.Block.Txs) == 0 {
			return false, fmt.Errorf("block is invalid: no oracle transaction")
		}

		// attempt to unmarshal extended commit info
		extendedCommitInfo, err := extCommitCodec.Decode(block.Block.Txs[0])
		if err != nil {
			return false, err
		}

		sort.Sort(validatorVotes(extendedCommitInfo.Votes))

		// iterate through all votes (votes in the extended-commit are deterministically ordered by voting power -> address)
		for i, vote := range extendedCommitInfo.Votes {
			// get the oracle vote extension
			gotVe, err := veCodec.Decode(vote.VoteExtension)
			if err != nil {
				return false, err
			}

			if len(ves[i].Prices) != len(gotVe.Prices) {
				return false, nil
			}

			// check that the vote extension is correct
			for ticker, price := range gotVe.Prices {
				if !bytes.Equal(price, ves[i].Prices[ticker]) {
					return false, nil
				}
			}
		}

		return true, nil
	}); err != nil {
		return 0, err
	}

	// we want to wait for the application state to reflect the proposed state from blockHeight
	return blockHeight, WaitForHeight(chain, blockHeight+1, timeout)
}

// wrapper around extendedVoteInfo for use in sorting (to make ordering deterministic in tests)
type validatorVotes []cmtabci.ExtendedVoteInfo

func (vv validatorVotes) Len() int { return len(vv) }

func (vv validatorVotes) Swap(i, j int) { vv[i], vv[j] = vv[j], vv[i] }

// order the votes by the number of reports first, then by the contents of the vote-extensions.
func (vv validatorVotes) Less(i, j int) bool {
	// break ties by vote-extension data
	var (
		iPrice, jPrice, iTotalPrice, jTotalPrice int
	)

	ve, err := veCodec.Decode(vv[i].VoteExtension)
	if err == nil {
		iPrice = len(ve.Prices)

		for _, priceBz := range ve.Prices {
			var price big.Int
			if err := price.GobDecode(priceBz); err != nil {
				panic(err)
			}

			iTotalPrice += int(price.Int64())
		}
	}

	ve, err = veCodec.Decode(vv[j].VoteExtension)
	if err == nil {
		jPrice = len(ve.Prices)

		for _, priceBz := range ve.Prices {
			var price big.Int
			if err := price.GobDecode(priceBz); err != nil {
				panic(err)
			}

			jTotalPrice += int(price.Int64())
		}
	}

	// check if the number of prices is different
	if iPrice != jPrice {
		return iPrice < jPrice
	}

	// break ties by the sum of the prices for each validator
	return iTotalPrice < jTotalPrice
}
