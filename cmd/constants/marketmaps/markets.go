package marketmaps

import (
	"encoding/json"
	"errors"
	"fmt"

	mmtypes "github.com/skip-mev/connect/v2/x/marketmap/types"
)

var (
	// CoinMarketCapMarketMap is used to initialize the CoinMarketCap market map. This only includes
	// the markets that are supported by CoinMarketCap.
	CoinMarketCapMarketMap mmtypes.MarketMap
	// CoinMarketCapMarketMapJSON is the JSON representation of the CoinMarketCap MarketMap that can be used
	// to initialize for a genesis state or used by the sidecar as as static market map.
	CoinMarketCapMarketMapJSON = `
{
    "markets": {
	  "W/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "W",
            "Quote": "USD"
          },
          "decimals": 8,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "29587"
          }
        ]
      },
	  "TON/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "TON",
            "Quote": "USD"
          },
          "decimals": 8,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "11419"
          }
        ]
      },
	  "ZRO/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "ZRO",
            "Quote": "USD"
          },
          "decimals": 8,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "26997"
          }
        ]
      },
	  "CHZ/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "CHZ",
            "Quote": "USD"
          },
          "decimals": 8,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "4066"
          }
        ]
      },
	  "ZK/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "ZK",
            "Quote": "USD"
          },
          "decimals": 8,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "24091"
          }
        ]
      },
	  "BODEN/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "BODEN",
            "Quote": "USD"
          },
          "decimals": 8,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "29687"
          }
        ]
      },
	  "ETHFI/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "ETHFI",
            "Quote": "USD"
          },
          "decimals": 8,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "29814"
          }
        ]
      },
      "KHAI/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "KHAI",
            "Quote": "USD"
          },
          "decimals": 8,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "30948"
          }
        ]
      },
      "WAFFLES/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "WAFFLES",
            "Quote": "USD"
          },
          "decimals": 8,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "31442"
          }
        ]
      },
      "HEGE/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "HEGE",
            "Quote": "USD"
          },
          "decimals": 8,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "31044"
          }
        ]
      },
      "WUF/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "WUF",
            "Quote": "USD"
          },
          "decimals": 8,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "30683"
          }
        ]
      },
      "CHAT/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "CHAT",
            "Quote": "USD"
          },
          "decimals": 8,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "29478"
          }
        ]
      },
      "BEER/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "BEER",
            "Quote": "USD"
          },
          "decimals": 8,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "31337"
          }
        ]
      },
      "MANEKI/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "MANEKI",
            "Quote": "USD"
          },
          "decimals": 11,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "30912"
          }
        ]
      },
      "SLERF/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "SLERF",
            "Quote": "USD"
          },
          "decimals": 8,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "29920"
          }
        ]
      },
      "MYRO/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "MYRO",
            "Quote": "USD"
          },
          "decimals": 8,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "28382"
          }
        ]
      },
      "RAY/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "RAY",
            "Quote": "USD"
          },
          "decimals": 8,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "8526"
          }
        ]
      },
      "WIF/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "WIF",
            "Quote": "USD"
          },
          "decimals": 8,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "28752"
          }
        ]
      },
      "MICHI/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "MICHI",
            "Quote": "USD"
          },
          "decimals": 8,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "30943"
          }
        ]
      },
      "MEW/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "MEW",
            "Quote": "USD"
          },
          "decimals": 8,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "30126"
          }
        ]
      },
      "PONKE/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "PONKE",
            "Quote": "USD"
          },
          "decimals": 8,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "29150"
          }
        ]
      },
      "BOME/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "BOME",
            "Quote": "USD"
          },
          "decimals": 8,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "29870"
          }
        ]
      },
      "DJT/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "DJT",
            "Quote": "USD"
          },
          "decimals": 8,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "31891"
          }
        ]
      },
      "POPCAT/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "POPCAT",
            "Quote": "USD"
          },
          "decimals": 8,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "28782"
          }
        ]
      },
      "AAVE/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "AAVE",
            "Quote": "USD"
          },
          "decimals": 8,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "7278"
          }
        ]
      },
      "ADA/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "ADA",
            "Quote": "USD"
          },
          "decimals": 10,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "2010"
          }
        ]
      },
      "AEVO/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "AEVO",
            "Quote": "USD"
          },
          "decimals": 9,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "29676"
          }
        ]
      },
      "ALGO/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "ALGO",
            "Quote": "USD"
          },
          "decimals": 10,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "4030"
          }
        ]
      },
      "APE/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "APE",
            "Quote": "USD"
          },
          "decimals": 9,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "18876"
          }
        ]
      },
      "APT/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "APT",
            "Quote": "USD"
          },
          "decimals": 9,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "21794"
          }
        ]
      },
      "ARB/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "ARB",
            "Quote": "USD"
          },
          "decimals": 9,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "11841"
          }
        ]
      },
      "ARKM/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "ARKM",
            "Quote": "USD"
          },
          "decimals": 9,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "27565"
          }
        ]
      },
      "ASTR/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "ASTR",
            "Quote": "USD"
          },
          "decimals": 10,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "12885"
          }
        ]
      },
      "ATOM/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "ATOM",
            "Quote": "USD"
          },
          "decimals": 9,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "3794"
          }
        ]
      },
      "AVAX/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "AVAX",
            "Quote": "USD"
          },
          "decimals": 8,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "5805"
          }
        ]
      },
      "AXL/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "AXL",
            "Quote": "USD"
          },
          "decimals": 9,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "17799"
          }
        ]
      },
      "BCH/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "BCH",
            "Quote": "USD"
          },
          "decimals": 7,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "1831"
          }
        ]
      },
      "BLUR/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "BLUR",
            "Quote": "USD"
          },
          "decimals": 10,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "23121"
          }
        ]
      },
      "BNB/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "BNB",
            "Quote": "USD"
          },
          "decimals": 7,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "1839"
          }
        ]
      },
      "BONK/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "BONK",
            "Quote": "USD"
          },
          "decimals": 14,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "23095"
          }
        ]
      },
      "BTC/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "BTC",
            "Quote": "USD"
          },
          "decimals": 5,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "1"
          }
        ]
      },
      "BUBBA/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "BUBBA",
            "Quote": "USD"
          },
          "decimals": 12,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "31411"
          }
        ]
      },
      "COMP/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "COMP",
            "Quote": "USD"
          },
          "decimals": 8,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "5692"
          }
        ]
      },
      "CRV/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "CRV",
            "Quote": "USD"
          },
          "decimals": 10,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "6538"
          }
        ]
      },
      "DOGE/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "DOGE",
            "Quote": "USD"
          },
          "decimals": 11,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "74"
          }
        ]
      },
      "DOT/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "DOT",
            "Quote": "USD"
          },
          "decimals": 9,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "6636"
          }
        ]
      },
      "DYDX/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "DYDX",
            "Quote": "USD"
          },
          "decimals": 9,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "28324"
          }
        ]
      },
      "DYM/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "DYM",
            "Quote": "USD"
          },
          "decimals": 9,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "28932"
          }
        ]
      },
      "TREMP/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "TREMP",
            "Quote": "USD"
          },
          "decimals": 11,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "29717"
          }
        ]
      },
      "MOG/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "MOG",
            "Quote": "USD"
          },
          "decimals": 11,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "27659"
          }
        ]
      },
      "MOTHER/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "MOTHER",
            "Quote": "USD"
          },
          "decimals": 11,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "31510"
          }
        ]
      },
      "EOS/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "EOS",
            "Quote": "USD"
          },
          "decimals": 10,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "1765"
          }
        ]
      },
      "ETC/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "ETC",
            "Quote": "USD"
          },
          "decimals": 8,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "1321"
          }
        ]
      },
      "ETH/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "ETH",
            "Quote": "USD"
          },
          "decimals": 6,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "1027"
          }
        ]
      },
      "FET/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "FET",
            "Quote": "USD"
          },
          "decimals": 10,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "3773"
          }
        ]
      },
      "FIL/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "FIL",
            "Quote": "USD"
          },
          "decimals": 9,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "2280"
          }
        ]
      },
      "GRT/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "GRT",
            "Quote": "USD"
          },
          "decimals": 10,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "6719"
          }
        ]
      },
      "HBAR/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "HBAR",
            "Quote": "USD"
          },
          "decimals": 10,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "4642"
          }
        ]
      },
      "ICP/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "ICP",
            "Quote": "USD"
          },
          "decimals": 8,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "8916"
          }
        ]
      },
      "IMX/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "IMX",
            "Quote": "USD"
          },
          "decimals": 9,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "10603"
          }
        ]
      },
      "INJ/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "INJ",
            "Quote": "USD"
          },
          "decimals": 8,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "7226"
          }
        ]
      },
      "JTO/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "JTO",
            "Quote": "USD"
          },
          "decimals": 9,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "28541"
          }
        ]
      },
      "JUP/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "JUP",
            "Quote": "USD"
          },
          "decimals": 10,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "29210"
          }
        ]
      },
      "LDO/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "LDO",
            "Quote": "USD"
          },
          "decimals": 9,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "8000"
          }
        ]
      },
      "LINK/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "LINK",
            "Quote": "USD"
          },
          "decimals": 9,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "1975"
          }
        ]
      },
      "LTC/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "LTC",
            "Quote": "USD"
          },
          "decimals": 8,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "2"
          }
        ]
      },
      "MANA/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "MANA",
            "Quote": "USD"
          },
          "decimals": 10,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "1966"
          }
        ]
      },
      "MATIC/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "MATIC",
            "Quote": "USD"
          },
          "decimals": 10,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "3890"
          }
        ]
      },
      "MKR/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "MKR",
            "Quote": "USD"
          },
          "decimals": 6,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "1518"
          }
        ]
      },
      "NEAR/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "NEAR",
            "Quote": "USD"
          },
          "decimals": 9,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "6535"
          }
        ]
      },
      "NTRN/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "NTRN",
            "Quote": "USD"
          },
          "decimals": 8,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "26680"
          }
        ]
      },
      "OP/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "OP",
            "Quote": "USD"
          },
          "decimals": 9,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "11840"
          }
        ]
      },
      "ORDI/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "ORDI",
            "Quote": "USD"
          },
          "decimals": 8,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "25028"
          }
        ]
      },
      "PEPE/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "PEPE",
            "Quote": "USD"
          },
          "decimals": 16,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "24478"
          }
        ]
      },
      "PYTH/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "PYTH",
            "Quote": "USD"
          },
          "decimals": 10,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "28177"
          }
        ]
      },
      "RUNE/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "RUNE",
            "Quote": "USD"
          },
          "decimals": 9,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "4157"
          }
        ]
      },
      "SEI/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "SEI",
            "Quote": "USD"
          },
          "decimals": 10,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "23149"
          }
        ]
      },
      "SHIB/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "SHIB",
            "Quote": "USD"
          },
          "decimals": 15,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "5994"
          }
        ]
      },
      "SNX/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "SNX",
            "Quote": "USD"
          },
          "decimals": 9,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "2586"
          }
        ]
      },
      "SOL/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "decimals": 8,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "5426"
          }
        ]
      },
      "STRK/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "STRK",
            "Quote": "USD"
          },
          "decimals": 9,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "22691"
          }
        ]
      },
      "STX/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "STX",
            "Quote": "USD"
          },
          "decimals": 9,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "4847"
          }
        ]
      },
      "SUI/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "SUI",
            "Quote": "USD"
          },
          "decimals": 10,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "20947"
          }
        ]
      },
      "TIA/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "TIA",
            "Quote": "USD"
          },
          "decimals": 8,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "22861"
          }
        ]
      },
      "TRX/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "TRX",
            "Quote": "USD"
          },
          "decimals": 11,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "1958"
          }
        ]
      },
      "UNI/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "UNI",
            "Quote": "USD"
          },
          "decimals": 9,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "7083"
          }
        ]
      },
      "USDT/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "USDT",
            "Quote": "USD"
          },
          "decimals": 9,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "825"
          }
        ]
      },
      "WLD/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "WLD",
            "Quote": "USD"
          },
          "decimals": 9,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "13502"
          }
        ]
      },
      "WOO/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "WOO",
            "Quote": "USD"
          },
          "decimals": 10,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "7501"
          }
        ]
      },
      "XLM/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "XLM",
            "Quote": "USD"
          },
          "decimals": 10,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "512"
          }
        ]
      },
      "XRP/USD": {
        "ticker": {
          "currency_pair": {
            "Base": "XRP",
            "Quote": "USD"
          },
          "decimals": 10,
          "min_provider_count": 1,
          "enabled": true
        },
        "provider_configs": [
          {
            "name": "coinmarketcap_api",
            "off_chain_ticker": "52"
          }
        ]
      }
    }
}
	`

	// RaydiumMarketMap is used to initialize the Raydium market map. This only includes
	// the markets that are supported by Raydium.
	RaydiumMarketMap mmtypes.MarketMap
	// RaydiumMarketMapJSON is the JSON representation of the Raydium MarketMap that can be used
	// to initialize for a genesis state or used by the sidecar as as static market map.
	RaydiumMarketMapJSON = `
{
  "markets": {
    "AART,RAYDIUM,F3NEFJBCEJYBTDREJUI1T9DPH5DBGPKKQ7U2GAAMXS5B/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "AART,RAYDIUM,F3NEFJBCEJYBTDREJUI1T9DPH5DBGPKKQ7U2GAAMXS5B",
          "Quote": "USD"
        },
        "decimals": 13,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"13571\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "AART,RAYDIUM,F3NEFJBCEJYBTDREJUI1T9DPH5DBGPKKQ7U2GAAMXS5B/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"HzKUtBVQMJ6SmmLpmbETLTtdLC4V7TimZ3jdELCw8r18\",\"token_decimals\":6},\"quote_token_vault\":{\"token_vault_address\":\"8aLSafgiPVNhqtryJdJwc9BRQgFbFwMYLdwoe2nKthE6\",\"token_decimals\":9},\"amm_info_address\":\"B4qjzQ5gt8YL8EyoDoM1z2xnMAQCDQ2udYxUCTzxKVBC\",\"open_orders_address\":\"4Hv6di87NN1pByfDx8c7QwvnkdZrJ1drazHdsWudzJHw\"}"
        }
      ]
    },
    "ANALOS,RAYDIUM,7IT1GRYYHEOP2NV1DYCWK2MGYLMPHQ47WHPGSWIQCUG5/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "ANALOS,RAYDIUM,7IT1GRYYHEOP2NV1DYCWK2MGYLMPHQ47WHPGSWIQCUG5",
          "Quote": "USD"
        },
        "decimals": 14,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"28805\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "ANALOS,RAYDIUM,7IT1GRYYHEOP2NV1DYCWK2MGYLMPHQ47WHPGSWIQCUG5/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"9ibeYfpgDxyoSYNvsc37EwGua6Z1NQpp7LH6e7CvBabM\",\"token_decimals\":8},\"quote_token_vault\":{\"token_vault_address\":\"5JFikPKzw3JeXKaJZaKTEAHfF3pJAoFJbHXUB5p2Ns5S\",\"token_decimals\":9},\"amm_info_address\":\"69grLw4PcSypZnn3xpsozCJFT8vs8WA5817VUVnzNGTh\",\"open_orders_address\":\"2QoiVyXa8Bfgx35yTcJhQJVvZYouKZutj5r5CEDPgQUm\"}"
        }
      ]
    },
    "APES,RAYDIUM,984GBL7PHCECHTN64NWLDBB49RSQXX7OZPDKEBR1PUMP/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "APES,RAYDIUM,984GBL7PHCECHTN64NWLDBB49RSQXX7OZPDKEBR1PUMP",
          "Quote": "USD"
        },
        "decimals": 11,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"32743\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SOL,RAYDIUM,SO11111111111111111111111111111111111111112/APES,RAYDIUM,984GBL7PHCECHTN64NWLDBB49RSQXX7OZPDKEBR1PUMP",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "invert": true,
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"9zoD3Xe8rfijq9KtgEM2YXsLGbq7pbrRXjkrfD8YvfCB\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"2w43DwFsEoxGQgcNHeC1Z5m88HxwZdaGWeB2G1dVbwDe\",\"token_decimals\":6},\"amm_info_address\":\"Di2FYotKzTV7Kwyyj476KDyk95ispN6jZYRQafwGGFet\",\"open_orders_address\":\"23D2NxnUbSgxmGQcQ15wXUVSbskMc4Pr5bMThxAYhdF1\"}"
        }
      ]
    },
    "ATLAS,RAYDIUM,ATLASXMBPQXBUYBXPSV97USA3FPQYEQZQBUHGIFCUSXX/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "ATLAS,RAYDIUM,ATLASXMBPQXBUYBXPSV97USA3FPQYEQZQBUHGIFCUSXX",
          "Quote": "USD"
        },
        "decimals": 12,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"11212\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "ATLAS,RAYDIUM,ATLASXMBPQXBUYBXPSV97USA3FPQYEQZQBUHGIFCUSXX/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"CBz2GGWNissxivNRXk5UFJHiZUmm1x5pWWuhSaoM8ULK\",\"token_decimals\":8},\"quote_token_vault\":{\"token_vault_address\":\"5VhBQ77aqytGZ8QWtV6eq4p4yNHsWrH3bALKFisTp18n\",\"token_decimals\":9},\"amm_info_address\":\"7Mu9zK6qV3wGp5deSkhCeWqaDnL3kdD4gKL87ui6GtmX\",\"open_orders_address\":\"AAL5LQpnHbSpp7R84kqP7RC8qLEBmK6tdVPwoenhE2sQ\"}"
        }
      ]
    },
    "AURA,RAYDIUM,DTR4D9FTVOTX2569GAL837ZGRB6WNJJ6TKMNX9RDK9B2/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "AURA,RAYDIUM,DTR4D9FTVOTX2569GAL837ZGRB6WNJJ6TKMNX9RDK9B2",
          "Quote": "USD"
        },
        "decimals": 12,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"31843\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SOL,RAYDIUM,SO11111111111111111111111111111111111111112/AURA,RAYDIUM,DTR4D9FTVOTX2569GAL837ZGRB6WNJJ6TKMNX9RDK9B2",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "invert": true,
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"9jbyBXHinaAah2SthksJTYGzTQNRLA7HdT2A7VMF91Wu\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"9v9FpQYd46LS9zHJitTtnPDDQrHfkSdW2PRbbEbKd2gw\",\"token_decimals\":6},\"amm_info_address\":\"9ViX1VductEoC2wERTSp2TuDxXPwAf69aeET8ENPJpsN\",\"open_orders_address\":\"2yWYj3BftHsfy5jCNXXcs3H1gcjAk7KHLCrBcZ3izv7n\"}"
        }
      ]
    },
    "BABY,RAYDIUM,5HMF8JT9PUWOQIFQTB3VR22732ZTKYRLRW9VO7TN3RCZ/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "BABY,RAYDIUM,5HMF8JT9PUWOQIFQTB3VR22732ZTKYRLRW9VO7TN3RCZ",
          "Quote": "USD"
        },
        "decimals": 16,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"29668\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SOL,RAYDIUM,SO11111111111111111111111111111111111111112/BABY,RAYDIUM,5HMF8JT9PUWOQIFQTB3VR22732ZTKYRLRW9VO7TN3RCZ",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "invert": true,
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"7ii3G8iRAVpfFt8NbmisDCVNa8KJxaikL9PQ49r8e6gm\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"6ogxazd3uvk3VSF8n9DRcEp2HNqQDxiqRNYXVKNWPNoU\",\"token_decimals\":6},\"amm_info_address\":\"2prhzdRwWzas2f4g5AAjyRUBcQcdajxd8NAzKcqhv76P\",\"open_orders_address\":\"7usCDD5FXposJD8xAxQDyQ3ByprweUY9AT8scjKswK1W\"}"
        }
      ]
    },
    "BABYTRUMP,RAYDIUM,6NBNHQKD2DH4JSWTLMMCP7LNSH4NH6Y2CNGDQG2NY9ZW/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "BABYTRUMP,RAYDIUM,6NBNHQKD2DH4JSWTLMMCP7LNSH4NH6Y2CNGDQG2NY9ZW",
          "Quote": "USD"
        },
        "decimals": 11,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"27955\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "BABYTRUMP,RAYDIUM,6NBNHQKD2DH4JSWTLMMCP7LNSH4NH6Y2CNGDQG2NY9ZW/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"GfLzieTjxyukRovbE8PFsgnfx4tJiSNy3JXAkSfEa1nu\",\"token_decimals\":8},\"quote_token_vault\":{\"token_vault_address\":\"59DTfm81hYHdrFr4Bu9Z5NaiM8YmPZz34snVTWXEaEng\",\"token_decimals\":9},\"amm_info_address\":\"Azzrh6PHeEqsDtfZujyZbXnMnRx8WnUpcUStvF87J93R\",\"open_orders_address\":\"GN2Xyr6z7oWYMyTD8sTBC4zhAKNkpo5eqfcCgSChQ2Y5\"}"
        }
      ]
    },
    "BAG,RAYDIUM,D8R8XTUCRUHLHEWEGXSWC3G92RHASFICV3YA7B2XWCLV/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "BAG,RAYDIUM,D8R8XTUCRUHLHEWEGXSWC3G92RHASFICV3YA7B2XWCLV",
          "Quote": "USD"
        },
        "decimals": 12,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"30088\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "BAG,RAYDIUM,D8R8XTUCRUHLHEWEGXSWC3G92RHASFICV3YA7B2XWCLV/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"7eLwyCqfhxKLsKeFwcN4JdfspKK22rSC4uQHNy3zWNPB\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"Cr7Yo8Uf5f8pzMsY3ZwgDFNx85nb3UDvPfQxuWG4acxc\",\"token_decimals\":9},\"amm_info_address\":\"Bv7mM5TwLxsukrRrwzEc6TFAj22GAdVCcH5ViAZFNZC\",\"open_orders_address\":\"Du6ZaABu8cxmCAvwoGMixZgZuw57cCQc8xE8yRenaxL4\"}"
        }
      ]
    },
    "BEER,RAYDIUM,AUJTJJ7AMS8LDO3BFZOYXDWT3JBALUBU4VZHZZDTZLMG/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "BEER,RAYDIUM,AUJTJJ7AMS8LDO3BFZOYXDWT3JBALUBU4VZHZZDTZLMG",
          "Quote": "USD"
        },
        "decimals": 14,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"31337\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "BEER,RAYDIUM,AUJTJJ7AMS8LDO3BFZOYXDWT3JBALUBU4VZHZZDTZLMG/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"HK4WKQfUKU2VuYfhjVzUR8Sx2Tkpqjg7VmrPjuNeNM6Q\",\"token_decimals\":6},\"quote_token_vault\":{\"token_vault_address\":\"52MdSmjzjnmRqsimXoPUTQX3zTr11abKqgkGBKbNv7Mg\",\"token_decimals\":9},\"amm_info_address\":\"Cne2WysEXzSLWbdABTG3vYkRNyJyMJ1zLhn26QPrBRZg\",\"open_orders_address\":\"BJDwwanWSMmq4GJEq21LvT7HUQN5kuYDmby8BkhqT687\"}"
        }
      ]
    },
    "BENDOG,RAYDIUM,AHW5N8IQZOBTCBEPKSJZZ61XTAUSZBDCPXTRLG6KUKPK/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "BENDOG,RAYDIUM,AHW5N8IQZOBTCBEPKSJZZ61XTAUSZBDCPXTRLG6KUKPK",
          "Quote": "USD"
        },
        "decimals": 12,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"29574\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "BENDOG,RAYDIUM,AHW5N8IQZOBTCBEPKSJZZ61XTAUSZBDCPXTRLG6KUKPK/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"2Pza1YUczgc4RWLhAgdXSJh4oYUspvhhAiSecFDd7ZJ3\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"2BFpMzi33JtpY4CGUjY7x5JPApy6f2AdkuLZsd1QGqRv\",\"token_decimals\":9},\"amm_info_address\":\"47857wX96Tb4Ud3M3ka949iVRFmUqS33KLBxoVsqgfLK\",\"open_orders_address\":\"H1FPc9WQpA3GPnXMmzSjtt6gMuYuyDqYndBscaHNyCbv\"}"
        }
      ]
    },
    "BILL,RAYDIUM,5O817YNOR97F1H5SUZBGHWC8DP53OURERUFBGJIACFUZ/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "BILL,RAYDIUM,5O817YNOR97F1H5SUZBGHWC8DP53OURERUFBGJIACFUZ",
          "Quote": "USD"
        },
        "decimals": 15,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"32297\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "BILL,RAYDIUM,5O817YNOR97F1H5SUZBGHWC8DP53OURERUFBGJIACFUZ/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"CKFsrVfEcYgUBrJxnchNVdtvDrsXuYkrueshfxxbAoLG\",\"token_decimals\":6},\"quote_token_vault\":{\"token_vault_address\":\"CkrTmNKe5eoZzUVr8eNT763kT7rJi8C4KFkVkedJ2PE2\",\"token_decimals\":9},\"amm_info_address\":\"7AAdTWLTgDCpAE1wXtPpGBstMVQtQhvftKcxkeC8H4FR\",\"open_orders_address\":\"4zpZVHQRgdGWH8qz7Au5CwZwqPiA729PPUkFo1rdp5KF\"}"
        }
      ]
    },
    "BILLY,RAYDIUM,3B5WUURMEI5YATD7ON46HKFEJ3PFMD7T1RKGRSN3PUMP/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "BILLY,RAYDIUM,3B5WUURMEI5YATD7ON46HKFEJ3PFMD7T1RKGRSN3PUMP",
          "Quote": "USD"
        },
        "decimals": 11,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"31914\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SOL,RAYDIUM,SO11111111111111111111111111111111111111112/BILLY,RAYDIUM,3B5WUURMEI5YATD7ON46HKFEJ3PFMD7T1RKGRSN3PUMP",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "invert": true,
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"B79b1mVm7w33Jp4WKi8s4noHw98VvAwPZtD9WUwRzoqa\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"4718at6MKguFJPaL1J8hkxJ23tmW4vn8oRbrKngKZE4m\",\"token_decimals\":6},\"amm_info_address\":\"9uWW4C36HiCTGr6pZW9VFhr9vdXktZ8NA8jVnzQU35pJ\",\"open_orders_address\":\"9jLvPDTie8cbv71XU89busAdzwR3cJmM4TvmkLgvRNTw\"}"
        }
      ]
    },
    "BLUE,RAYDIUM,CWQVQTKUH1IU8ZSFFFVAUXAVZLZQU1E8GYU5D6ECGBNE/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "BLUE,RAYDIUM,CWQVQTKUH1IU8ZSFFFVAUXAVZLZQU1E8GYU5D6ECGBNE",
          "Quote": "USD"
        },
        "decimals": 12,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"32349\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "BLUE,RAYDIUM,CWQVQTKUH1IU8ZSFFFVAUXAVZLZQU1E8GYU5D6ECGBNE/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"Fyp9Rhkxm3V5o7KGN6WzgCwqRZydBczNE8JniQ57YbjH\",\"token_decimals\":6},\"quote_token_vault\":{\"token_vault_address\":\"BU7otgLUAAyKsgzAf5TnBYTNFz6eRBWhBYv1h9MrZSzq\",\"token_decimals\":9},\"amm_info_address\":\"5AKR4Yt7jXpYZ9KLfH5f7GCpgRHFcnGcPuZKrsZSEops\",\"open_orders_address\":\"2wXY5PtMdvhNrQT6JfCTxBY7EeQRskBhifSoCCXCv2H4\"}"
        }
      ]
    },
    "BLUR/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "BLUR",
          "Quote": "USD"
        },
        "decimals": 10,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"23121\"}]}"
      },
      "provider_configs": [
        {
          "name": "coinbase_ws",
          "off_chain_ticker": "BLUR-USD"
        }
      ]
    },
    "BOBAOPPA,RAYDIUM,BOBAM3U8QMQZHY1HWATNVZE9DLXVKGKYK3TD3T8MLVA/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "BOBAOPPA,RAYDIUM,BOBAM3U8QMQZHY1HWATNVZE9DLXVKGKYK3TD3T8MLVA",
          "Quote": "USD"
        },
        "decimals": 13,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"30157\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "BOBAOPPA,RAYDIUM,BOBAM3U8QMQZHY1HWATNVZE9DLXVKGKYK3TD3T8MLVA/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"HnkHkhxgQ3XPjJzNu5BVsypgKFhqDtfeCc4xZAMKs7fz\",\"token_decimals\":6},\"quote_token_vault\":{\"token_vault_address\":\"8ZX85AkabTUMBWWgTRbwVQZgkgkun19MAN4GgFBrpxee\",\"token_decimals\":9},\"amm_info_address\":\"8MHLZAHWtU5VVN1aWN2tz7xhUihkoZc2B9u1akcqgs97\",\"open_orders_address\":\"3TwLrMkLSDxJ3pTg3duNovFG3D7RmyH4URGUTFezhKux\"}"
        }
      ]
    },
    "BODEN,RAYDIUM,3PSH1MJ1F7YUFAD5GH6ZJ7EPE8HHRMKMETGV5TSHQA4O/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "BODEN,RAYDIUM,3PSH1MJ1F7YUFAD5GH6ZJ7EPE8HHRMKMETGV5TSHQA4O",
          "Quote": "USD"
        },
        "decimals": 12,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"29687\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "BODEN,RAYDIUM,3PSH1MJ1F7YUFAD5GH6ZJ7EPE8HHRMKMETGV5TSHQA4O/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"54zedUwxuSnmHHYg9oY1AfykeBDaCF6ZFZDW3ym2Nea4\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"DzpiXKsTUCacKyahLBUC5sfjj2fiWbwCpiCPEgyS3zDC\",\"token_decimals\":9},\"amm_info_address\":\"6UYbX1x8YUcFj8YstPYiZByG7uQzAq2s46ZWphUMkjg5\",\"open_orders_address\":\"9ndGwmmTcFLut1TNjWFA8pDvcrxgmqPEJTZ2Y3jTipva\"}"
        }
      ]
    },
    "BOME,RAYDIUM,UKHH6C7MMYIWCF1B9PNWE25TSPKDDT3H5PQZGZ74J82/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "BOME,RAYDIUM,UKHH6C7MMYIWCF1B9PNWE25TSPKDDT3H5PQZGZ74J82",
          "Quote": "USD"
        },
        "decimals": 12,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"29870\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "BOME,RAYDIUM,UKHH6C7MMYIWCF1B9PNWE25TSPKDDT3H5PQZGZ74J82/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"FBba2XsQVhkoQDMfbNLVmo7dsvssdT39BMzVc2eFfE21\",\"token_decimals\":6},\"quote_token_vault\":{\"token_vault_address\":\"GuXKCb9ibwSeRSdSYqaCL3dcxBZ7jJcj6Y7rDwzmUBu9\",\"token_decimals\":9},\"amm_info_address\":\"DSUvc5qf5LJHHV5e2tD184ixotSnCnwj7i4jJa4Xsrmt\",\"open_orders_address\":\"38p42yoKFWgxw2LCbB96wAKa2LwAxiBArY3fc3eA9yWv\"}"
        }
      ]
    },
    "BONK/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "BONK",
          "Quote": "USD"
        },
        "decimals": 14,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"23095\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SOL,RAYDIUM,SO11111111111111111111111111111111111111112/BONK,RAYDIUM,DEZXAZ8Z7PNRNRJJZ3WXBORGIXCA6XJNB7YAB1PPB263",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "invert": true,
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"B1mmGm5bveLSwYHkQXPJ7mFb5KBFNa7U9Hma3Qdw1qbd\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"DbPFYPjgCFK6KLJFh55mtn5hEroYcm7Dzs2fYBW3GGy2\",\"token_decimals\":5},\"amm_info_address\":\"GGj7YKTJdavHv2F7WcCic2SqEdPcZK1EWFfGDZMbDLo4\",\"open_orders_address\":\"4mhdTamk3wCiZTmoANGvw3iZGW2Lp6jwCjK14qmDf5KC\"}"
        },
        {
          "name": "coinbase_ws",
          "off_chain_ticker": "BONK-USD"
        }
      ]
    },
    "BORK,RAYDIUM,4JZXKSNGTQKCDB36ECZ6A2ANZCUNIGCDEXGTDTM2HXAX/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "BORK,RAYDIUM,4JZXKSNGTQKCDB36ECZ6A2ANZCUNIGCDEXGTDTM2HXAX",
          "Quote": "USD"
        },
        "decimals": 13,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"28883\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "BORK,RAYDIUM,4JZXKSNGTQKCDB36ECZ6A2ANZCUNIGCDEXGTDTM2HXAX/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"7cpepsDvzsfnEkhNoJzCJ8zQXHVv5QCwah2EhsasUDRA\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"CNBAWpbS3NGAD9JrEQUxtFqXYubbhNq4ARum2mdK9i12\",\"token_decimals\":9},\"amm_info_address\":\"9Rc5LrMNdjxePyd7xjZiSTAJURpzoi6GjiCPqnxQopdD\",\"open_orders_address\":\"3ZaGGXJ7ZDGJuNcJDt1mjUPdQ3DfwjQELyoZ4na372UZ\"}"
        }
      ]
    },
    "BRAINLET,RAYDIUM,8NNXWRWVCTNW1UFEABYPFFIMTDCLCCD8XJZHVYSMGWPF/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "BRAINLET,RAYDIUM,8NNXWRWVCTNW1UFEABYPFFIMTDCLCCD8XJZHVYSMGWPF",
          "Quote": "USD"
        },
        "decimals": 11,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"32448\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SOL,RAYDIUM,SO11111111111111111111111111111111111111112/BRAINLET,RAYDIUM,8NNXWRWVCTNW1UFEABYPFFIMTDCLCCD8XJZHVYSMGWPF",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "invert": true,
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"2z19yGJ6dcsBDXPEhfoBX4c3GXhXwyRSoxxXfqFM2YMW\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"CTYrZGYRJAH56N6nbhb7XHcHTdUrf3xULJcSirhChbQi\",\"token_decimals\":6},\"amm_info_address\":\"CW9DFoTWEUiwxyxVGnQFYhbrYEfGkvaqXEgxKZG7d7X1\",\"open_orders_address\":\"FxQebs1z7Xy1xAyQsucpnYAw61co33h7u177vPrUGnFJ\"}"
        }
      ]
    },
    "BTW,RAYDIUM,4YTPZGVONB66BFS6NRCUAAVSLDTYK2FHQ4U92JNJPUMP/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "BTW,RAYDIUM,4YTPZGVONB66BFS6NRCUAAVSLDTYK2FHQ4U92JNJPUMP",
          "Quote": "USD"
        },
        "decimals": 12,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"32664\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SOL,RAYDIUM,SO11111111111111111111111111111111111111112/BTW,RAYDIUM,4YTPZGVONB66BFS6NRCUAAVSLDTYK2FHQ4U92JNJPUMP",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "invert": true,
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"7MoFk4XsL1uyi7WMvN7Fcxmzofuvu83ZKvcgBSeRs4yU\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"D43XeJ1Duo1Vxx3cGg1Zxybbbtme2fieKMpx8EKSeyJh\",\"token_decimals\":6},\"amm_info_address\":\"HhGgHkKgRfrCpub5oxUdituvnre9xWjWy3A5DCCwA5NK\",\"open_orders_address\":\"F5rHouGkVeGbrDdBxSMAx5T7bkiQV9rrhqkg5R4TwgY\"}"
        }
      ]
    },
    "BWB,RAYDIUM,6FVYLVHQSSHWVUSCQ2FJRR1MRECGSHC3QXBWWTGIVFWK/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "BWB,RAYDIUM,6FVYLVHQSSHWVUSCQ2FJRR1MRECGSHC3QXBWWTGIVFWK",
          "Quote": "USD"
        },
        "decimals": 10,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"31503\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SOL,RAYDIUM,SO11111111111111111111111111111111111111112/BWB,RAYDIUM,6FVYLVHQSSHWVUSCQ2FJRR1MRECGSHC3QXBWWTGIVFWK",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "invert": true,
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"4ePLhRVWCAkYwZrxWP8XfDL4qZ8GkLXWkj1mZ4kKvkrP\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"F4PaCs96hXkftVonsBfhkfaok7yKM3vdTr9EMYM2mpFk\",\"token_decimals\":8},\"amm_info_address\":\"AoLJD4ZMjRdHrcv3RxwyTXYHMWhDpoNhzvkD6NjinBRF\",\"open_orders_address\":\"3cYNCgGmC7Xa2kouifMmw85Y9D7zKyLKu6SAN8e1eNKd\"}"
        }
      ]
    },
    "BYTE,RAYDIUM,ARGFK9JJ72QETCMCKXVJCZ9APUATGP1MF9YSVHNPU4UT/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "BYTE,RAYDIUM,ARGFK9JJ72QETCMCKXVJCZ9APUATGP1MF9YSVHNPU4UT",
          "Quote": "USD"
        },
        "decimals": 15,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"28664\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "BYTE,RAYDIUM,ARGFK9JJ72QETCMCKXVJCZ9APUATGP1MF9YSVHNPU4UT/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"CE11TUoJJobrhdQx5nP7wHizm9X8UKZ5ijC7Jnh1z3M5\",\"token_decimals\":8},\"quote_token_vault\":{\"token_vault_address\":\"DbdVr62eSAq56GxRMPsa4WYkfAhqYDeZQibCXUwLqkJG\",\"token_decimals\":9},\"amm_info_address\":\"2dew1kHVZVzirmuMeFW3ZN2f9Az1BUqHuNs1miRvsBk8\",\"open_orders_address\":\"EeKZscVR4Vh8JAtzp1Ah21yfGrqu1yGWigiCGZ3C5Kt9\"}"
        }
      ]
    },
    "CAIR,RAYDIUM,D7Z8T6FADMQDYGHY3LSMN3BZFRABMVVUNAFURKFUWZ8F/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "CAIR,RAYDIUM,D7Z8T6FADMQDYGHY3LSMN3BZFRABMVVUNAFURKFUWZ8F",
          "Quote": "USD"
        },
        "decimals": 11,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"28670\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SOL,RAYDIUM,SO11111111111111111111111111111111111111112/CAIR,RAYDIUM,D7Z8T6FADMQDYGHY3LSMN3BZFRABMVVUNAFURKFUWZ8F",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "invert": true,
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"EvY2KcLhcyeAzGKwKfBgmAV6Bpd5H1wcDeYNjLmXEQXZ\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"AqvBFvFCEMysM85E87wrm4qNLedgdAC4nv1VpF1kPRga\",\"token_decimals\":6},\"amm_info_address\":\"DFLK6TLKMFoQmTCLoN2qMk2jfh1nD22dZbgyBL23ruPh\",\"open_orders_address\":\"Ae2syXzkGDcXCXmzXCBYMa3SrjR8u6oJGzhYDahkrZdv\"}"
        }
      ]
    },
    "CAT,RAYDIUM,3WKZQDH3ZW3TP2PHATAUDU4E1XSEZFHK7QNN8MAPM3S2/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "CAT,RAYDIUM,3WKZQDH3ZW3TP2PHATAUDU4E1XSEZFHK7QNN8MAPM3S2",
          "Quote": "USD"
        },
        "decimals": 13,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"31438\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "CAT,RAYDIUM,3WKZQDH3ZW3TP2PHATAUDU4E1XSEZFHK7QNN8MAPM3S2/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"5ezZ425K5cRiZLYuYtPacFEpnZfx68PS7pFNpxnJWLKg\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"5qRnDE3vcBQrAY3YM2K6gu3UksahWfPtjAGvLqy2cs6r\",\"token_decimals\":9},\"amm_info_address\":\"A8WteoYJdxxir2zN7oGKe27e8xfZyq2C9xUW2vNJXWJt\",\"open_orders_address\":\"HMj9zn5zobGvVkPwYt4boP2F3o9gE5MWcR3n54311sHC\"}"
        }
      ]
    },
    "CATDOG,RAYDIUM,CATTZAWLYADD2EKZVJTJX8TVUBYFROZDKJBKUTJGGDB7/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "CATDOG,RAYDIUM,CATTZAWLYADD2EKZVJTJX8TVUBYFROZDKJBKUTJGGDB7",
          "Quote": "USD"
        },
        "decimals": 14,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"32638\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "CATDOG,RAYDIUM,CATTZAWLYADD2EKZVJTJX8TVUBYFROZDKJBKUTJGGDB7/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"87CSWDJRS9c3CVCnzp6WrohXJLtPjnhPGp9fjhCdUDfE\",\"token_decimals\":8},\"quote_token_vault\":{\"token_vault_address\":\"ALTTB76kXsiwERgJZz6KXTauRyj92n9FhQ4Uoc82pVXp\",\"token_decimals\":9},\"amm_info_address\":\"mh2TEd7H29EZ4YTXPvgzysBWHFqtPQbT2MVy6kDUwSh\",\"open_orders_address\":\"9RzG9FtdSihXanxBDJYmxu4miKnj6wZe4c4oDbEes3RJ\"}"
        }
      ]
    },
    "CATGPT,RAYDIUM,FGF1US3KQU9AXU2X1YWKFIKE8USX42ACVRIURBUAODZQ/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "CATGPT,RAYDIUM,FGF1US3KQU9AXU2X1YWKFIKE8USX42ACVRIURBUAODZQ",
          "Quote": "USD"
        },
        "decimals": 13,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"30990\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "CATGPT,RAYDIUM,FGF1US3KQU9AXU2X1YWKFIKE8USX42ACVRIURBUAODZQ/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"AYAmCRPotwZprbNpPQ1hVGSEpbgWUgWHUbjnjt4bfLo1\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"ok2NPhxx2q3tn8XL289m5irCyGntURLQNhtisLowQ7y\",\"token_decimals\":9},\"amm_info_address\":\"92NvJRnTxkaiHcfRd72B8h1SHyj5ZGtMoeFAQvCdB3vB\",\"open_orders_address\":\"DXA5jH1r5c9QeAZxAYQb6emFGY2eBb3ZgMjSDuTNFZ6n\"}"
        }
      ]
    },
    "CATMAN,RAYDIUM,EAVJDLH8CYTANT3QDITPKGMSPL2HQ1MY5G9R2P6AT6LC/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "CATMAN,RAYDIUM,EAVJDLH8CYTANT3QDITPKGMSPL2HQ1MY5G9R2P6AT6LC",
          "Quote": "USD"
        },
        "decimals": 14,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"29086\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "CATMAN,RAYDIUM,EAVJDLH8CYTANT3QDITPKGMSPL2HQ1MY5G9R2P6AT6LC/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"9fJawgPH4GK2KNpaqd9safRi9mTQj8GrcJEiTcUr998m\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"AnP97EVFrbDyHz4trdY2BZAvA7KeQBYhnUZzWnXf1k5A\",\"token_decimals\":9},\"amm_info_address\":\"utpdtL9XZJdxyeah5sxRcAamNfqmJuTcbu2TZoHkz4c\",\"open_orders_address\":\"3dmTcjBVUfh1kBxLE6Nr7iHgZzpNZ5s2rDThAijziDuU\"}"
        }
      ]
    },
    "CAW,RAYDIUM,9NIFQK8MSPARJSXHYQ3YS2A6ZHMSEUKSB1M7WWDVZ7BJ/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "CAW,RAYDIUM,9NIFQK8MSPARJSXHYQ3YS2A6ZHMSEUKSB1M7WWDVZ7BJ",
          "Quote": "USD"
        },
        "decimals": 17,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"30402\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "CAW,RAYDIUM,9NIFQK8MSPARJSXHYQ3YS2A6ZHMSEUKSB1M7WWDVZ7BJ/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"95eYAV2RnhR3dabJ4KMKVSvHGrJ6pBuCM1PHeo5568N3\",\"token_decimals\":0},\"quote_token_vault\":{\"token_vault_address\":\"3hbccNdoopmvtGrT5shtGFckvNgD9q1vX1gMCnZ6Rtpn\",\"token_decimals\":9},\"amm_info_address\":\"r5ZNvv1VUBhPdJjyF6JftLkTAYiJKvuLRwWR8pTi19A\",\"open_orders_address\":\"BncD2zAqbiLDJxVyRvNqcS4KAemkpkCZYauMyzZMtJc7\"}"
        }
      ]
    },
    "CDT,RAYDIUM,AK3OVNWQNAXPSFOSNCONYJLNJTQDCKRBH4HWHWKB6HFM/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "CDT,RAYDIUM,AK3OVNWQNAXPSFOSNCONYJLNJTQDCKRBH4HWHWKB6HFM",
          "Quote": "USD"
        },
        "decimals": 10,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"14489\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "CDT,RAYDIUM,AK3OVNWQNAXPSFOSNCONYJLNJTQDCKRBH4HWHWKB6HFM/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"HfhTdqpEgm5BvLLq5pxFWBzQfpcfk23mPRGgWArqxYZU\",\"token_decimals\":8},\"quote_token_vault\":{\"token_vault_address\":\"rzqrBKSp49uyCCgYCwuCGHk5SiRgsrJ2iFHDcoCLb2H\",\"token_decimals\":9},\"amm_info_address\":\"L6o9YCETZAWxg4aMq2Nt47KoVipK3YJBqJbJbbiekGd\",\"open_orders_address\":\"EGEqg43424B1qjsRgcknPkB1ouLBf3512bEDeqtYwEVh\"}"
        }
      ]
    },
    "CHAT,RAYDIUM,947TEOG318GUMYJVYHRANRVWPMX7FPBTDQFBOJVSKSG3/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "CHAT,RAYDIUM,947TEOG318GUMYJVYHRANRVWPMX7FPBTDQFBOJVSKSG3",
          "Quote": "USD"
        },
        "decimals": 9,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"29478\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "CHAT,RAYDIUM,947TEOG318GUMYJVYHRANRVWPMX7FPBTDQFBOJVSKSG3/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"FKCCPsYDgEoVpEhyE2XMFAXq5zWFrWHgpQjVEKQk1C54\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"9APPnbdEXbJpktfKSGbbdgYvj6K3ZFRDFwQUabFw6CHP\",\"token_decimals\":9},\"amm_info_address\":\"9kLGUEFwEuFzn9txDfGJ3FimGp9LjMtNPp4GvMLfkZSY\",\"open_orders_address\":\"G9fse9D2feKdSjy4eLDQfuuBfxQDqektwNMG9smVBJr9\"}"
        }
      ]
    },
    "CHEEKS,RAYDIUM,6E6RVIDZAVLRV56NVZYE5UOFRKDG36MF6DTQRMCOPTW9/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "CHEEKS,RAYDIUM,6E6RVIDZAVLRV56NVZYE5UOFRKDG36MF6DTQRMCOPTW9",
          "Quote": "USD"
        },
        "decimals": 12,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"31336\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SOL,RAYDIUM,SO11111111111111111111111111111111111111112/CHEEKS,RAYDIUM,6E6RVIDZAVLRV56NVZYE5UOFRKDG36MF6DTQRMCOPTW9",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "invert": true,
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"5gXgjMgYCMJ6JNb6jAE1i23ypy3hzmGK4hDiCGbV6Tc8\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"7CDUKKCrck8KtDNkTUR9ynjbUo99Zg8Kergj6eukV6tX\",\"token_decimals\":6},\"amm_info_address\":\"FE7BtzcVMFmdQZVR18ywexqkuE4oHqPeZL8dTSon2965\",\"open_orders_address\":\"AcXy7Q2oc44sjjhP8qEr44tSjrtxtPrAamL6Yubi4DE3\"}"
        }
      ]
    },
    "CHEEMS,RAYDIUM,3FOUASGDBVTD6YZ4WVKJGTB76ONJUKZ7GPEBNIR5B8WC/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "CHEEMS,RAYDIUM,3FOUASGDBVTD6YZ4WVKJGTB76ONJUKZ7GPEBNIR5B8WC",
          "Quote": "USD"
        },
        "decimals": 12,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"10269\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "CHEEMS,RAYDIUM,3FOUASGDBVTD6YZ4WVKJGTB76ONJUKZ7GPEBNIR5B8WC/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"G2Quo8iai75CEUDvT2g9PwWUGHnMD8cQZ3fHKDpHgYK5\",\"token_decimals\":4},\"quote_token_vault\":{\"token_vault_address\":\"4m8vE64U9MbnXojPVgRZx8a1mH6P2Y9PAJFY6Z4Y2Eki\",\"token_decimals\":9},\"amm_info_address\":\"CVGxQpCFJYKky38bUEVghucWCAf3THN6hLFtMQtfeEn\",\"open_orders_address\":\"72VkdcbwUTD8FXhGFXG6qbZdBkqbQUoacaCyyXLYhdGD\"}"
        }
      ]
    },
    "CHEX,RAYDIUM,6DKCOWJPJ5MFU5GWDEFDPUUEBASBLK3WLEWHUZQPAA1E/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "CHEX,RAYDIUM,6DKCOWJPJ5MFU5GWDEFDPUUEBASBLK3WLEWHUZQPAA1E",
          "Quote": "USD"
        },
        "decimals": 11,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"8534\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "CHEX,RAYDIUM,6DKCOWJPJ5MFU5GWDEFDPUUEBASBLK3WLEWHUZQPAA1E/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"2qmJC3SVFSkK2fYz7fswkPn7ySqtf9xH2AkLBZ999sbr\",\"token_decimals\":8},\"quote_token_vault\":{\"token_vault_address\":\"AD7qRvY14byTLY19tWZTvXLaMF7AZj4dw4Gb1nHbgaW6\",\"token_decimals\":9},\"amm_info_address\":\"D8JjVpFdXjFvHmsX7LyFy8iHXEqzhbQo576Rt8rZkyiq\",\"open_orders_address\":\"5ePrmurwh3HGL67udZ3jkTJBPeMwtQK9Ey7uUqt12jRw\"}"
        }
      ]
    },
    "CHILL,RAYDIUM,BZPQOPC44OAHU9SB5HEK1GRNJZC4UWFCL4OWUSWOZM3N/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "CHILL,RAYDIUM,BZPQOPC44OAHU9SB5HEK1GRNJZC4UWFCL4OWUSWOZM3N",
          "Quote": "USD"
        },
        "decimals": 17,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"32111\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "CHILL,RAYDIUM,BZPQOPC44OAHU9SB5HEK1GRNJZC4UWFCL4OWUSWOZM3N/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"AuQ5EHn4c5qpdKKMNQkYYM32UJd8G6SgoBw8HTknTugx\",\"token_decimals\":3},\"quote_token_vault\":{\"token_vault_address\":\"G8o8sMwgfoGLtvAacHWQqQhmKWiLzgKgvUHQdDrsLs2D\",\"token_decimals\":9},\"amm_info_address\":\"D75GyXxJRPgSuyiBwKGvb7BG3tSD7FrPK1PVdowJDngp\",\"open_orders_address\":\"EhMqSPd9L8PKWKJQ8xhZDmfnSvDzAfjW4oRNK8s1cMrJ\"}"
        }
      ]
    },
    "CHINU,RAYDIUM,FLRGWXXAX8Q8ECF18WEDF3PLAYORXST5ORPY34D8JFBM/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "CHINU,RAYDIUM,FLRGWXXAX8Q8ECF18WEDF3PLAYORXST5ORPY34D8JFBM",
          "Quote": "USD"
        },
        "decimals": 13,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"29047\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "CHINU,RAYDIUM,FLRGWXXAX8Q8ECF18WEDF3PLAYORXST5ORPY34D8JFBM/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"43tka7kFH5oCXXWDZKYQZUrvz57ebTWN9mA7wqNvSXz6\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"AAsLmo7UqNBxXhSYfyfvYEYqwSEcWsKF9cGmn47NJhN4\",\"token_decimals\":9},\"amm_info_address\":\"2N6SHfcg2U8KPPYujRGMzBjAmW2NZUuWnRWRZVCihBxw\",\"open_orders_address\":\"FQzqRt48mNSmDHD8H5nzGMCVRv4v75FwB6DUyucHkpAp\"}"
        }
      ]
    },
    "CHITAN,RAYDIUM,J95PXHUEYTZTGBF9DPLPYPHXNUTWN4FBRK2JNC5NYKA3/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "CHITAN,RAYDIUM,J95PXHUEYTZTGBF9DPLPYPHXNUTWN4FBRK2JNC5NYKA3",
          "Quote": "USD"
        },
        "decimals": 13,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"32551\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "CHITAN,RAYDIUM,J95PXHUEYTZTGBF9DPLPYPHXNUTWN4FBRK2JNC5NYKA3/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"CASQbPygHHUzZqUaAtUZVBuCuBNJrZusqr7qNTuZGRtp\",\"token_decimals\":6},\"quote_token_vault\":{\"token_vault_address\":\"4DoPRWcDUE3S4KPYL5xQYZNGXP8aUNRh2zs1bVveeitw\",\"token_decimals\":9},\"amm_info_address\":\"Gm8k45peewP18XLX2ftGw14DZxmmWfrpJPxnB2Gd8PkK\",\"open_orders_address\":\"6Zj2c5yDX4CABm89gKTuh6wwEGteB2bVKCAuw95LGqQZ\"}"
        }
      ]
    },
    "CLOUD,RAYDIUM,CLOUDKC4ANE7HEQCPPE3YHNZNRXHMIMJ4MYAUQYHFZAU/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "CLOUD,RAYDIUM,CLOUDKC4ANE7HEQCPPE3YHNZNRXHMIMJ4MYAUQYHFZAU",
          "Quote": "USD"
        },
        "decimals": 10,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"32299\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "CLOUD,RAYDIUM,CLOUDKC4ANE7HEQCPPE3YHNZNRXHMIMJ4MYAUQYHFZAU/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"AwLEF1xnDm9bDmEbT5XMsi785zUfngC2CRDLP7iYdbwV\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"4TkeSJTzGcFqf8QJMERxwbJVZ3qT1bHs82afrPG3v3f7\",\"token_decimals\":9},\"amm_info_address\":\"4AG8E6GgtLG95juo9saVwTZxsaew8an17vLzLjg8z8LJ\",\"open_orders_address\":\"BQWxK5GXxQuuSHMgsE4FVVDK5vPEvuXd16XmfdyPToNX\"}"
        }
      ]
    },
    "COK,RAYDIUM,DNB9DLSXXAARXVEXEHZEH8W8NFMLMNJSUGOADDZSWTOG/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "COK,RAYDIUM,DNB9DLSXXAARXVEXEHZEH8W8NFMLMNJSUGOADDZSWTOG",
          "Quote": "USD"
        },
        "decimals": 15,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"31779\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "COK,RAYDIUM,DNB9DLSXXAARXVEXEHZEH8W8NFMLMNJSUGOADDZSWTOG/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"C4hk6k49gotrWP1b9j2ejPcPo4Lq59jVmfGwB2YYYGds\",\"token_decimals\":6},\"quote_token_vault\":{\"token_vault_address\":\"6mk1jhhWr6yeYxQkcrkia2wLHFyuy1LW6Xmj2MmwJ2x5\",\"token_decimals\":9},\"amm_info_address\":\"1D5GHSzrcaSXLtUYxSCg4vWHdKGd7hFnasYPiPFYFGX\",\"open_orders_address\":\"F8F7FGDKfqVEC4qpnVjigZHB8kijTx8qqpmc1fX8s1dY\"}"
        }
      ]
    },
    "COST,RAYDIUM,AV6QVIGKB7USQYPXJKUVAEM4F599WTRVD75PUWBA9ENM/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "COST,RAYDIUM,AV6QVIGKB7USQYPXJKUVAEM4F599WTRVD75PUWBA9ENM",
          "Quote": "USD"
        },
        "decimals": 12,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"30513\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "COST,RAYDIUM,AV6QVIGKB7USQYPXJKUVAEM4F599WTRVD75PUWBA9ENM/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"FgQifwxmQfjhHvh2ggVxQwb9qwRwHrxwwxxQXASLAnVH\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"9dDGmEfXJXgjcMAp516c5eUd1eBRW3ZgKg6diyBmd1xh\",\"token_decimals\":9},\"amm_info_address\":\"GQdUPA8cUV8WsqEdCfDQtphvztocNCoSBGo1wARtaAXK\",\"open_orders_address\":\"4MQHW9GXiypDCGgjgEGKYB6pLiPPF7v38ki9VpaiUvni\"}"
        }
      ]
    },
    "DADDY,RAYDIUM,4CNK9EPNW5IXFLZATCPJJDB1PUTCRPVVGTQUKM9EPUMP/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "DADDY,RAYDIUM,4CNK9EPNW5IXFLZATCPJJDB1PUTCRPVVGTQUKM9EPUMP",
          "Quote": "USD"
        },
        "decimals": 11,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"31830\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SOL,RAYDIUM,SO11111111111111111111111111111111111111112/DADDY,RAYDIUM,4CNK9EPNW5IXFLZATCPJJDB1PUTCRPVVGTQUKM9EPUMP",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "invert": true,
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"8LqocGsMwPJ7h2s1r8k4Vmc9c222Z4fMae25uz58qb3n\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"BWssmZs8cnEKTfHaXaqHebSUyfmWHuQmY6Ew6kJw59N\",\"token_decimals\":6},\"amm_info_address\":\"zcdAw3jpcqEY8JYVxNVMqs2cU35cyDdy4ot7V8edNhz\",\"open_orders_address\":\"DzkLxhpjAsNX9Wv1FJsnanzcKcAPJGPFy8GxoEJw53qM\"}"
        }
      ]
    },
    "DIMO/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "DIMO",
          "Quote": "USD"
        },
        "decimals": 10,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"22837\"}]}"
      },
      "provider_configs": [
        {
          "name": "coinbase_ws",
          "off_chain_ticker": "DIMO-USD"
        }
      ]
    },
    "DINGO,RAYDIUM,6VYF5JXQ6RFQ4QRGGMG6CO7B1EV1LJ7KSBHBXFQ9E1L3/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "DINGO,RAYDIUM,6VYF5JXQ6RFQ4QRGGMG6CO7B1EV1LJ7KSBHBXFQ9E1L3",
          "Quote": "USD"
        },
        "decimals": 14,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"16185\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "DINGO,RAYDIUM,6VYF5JXQ6RFQ4QRGGMG6CO7B1EV1LJ7KSBHBXFQ9E1L3/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"4V4s3PxgLYSq5nSGSYo1Edsu54EVTfcsmnLPhvDsY2su\",\"token_decimals\":8},\"quote_token_vault\":{\"token_vault_address\":\"2czYCE8NXvWRBnFnvUSkZfXumu5fpRcXFaAUygDRRmXT\",\"token_decimals\":9},\"amm_info_address\":\"7n45btQhNu5expVQrevYxzNX5V7pikmBJvtJkKCfQxAb\",\"open_orders_address\":\"3qgKaQvQTQnQdHWxr5BiXLZsWz5SqLZ3akpGgzUPEKzh\"}"
        }
      ]
    },
    "DLORD,RAYDIUM,3KRWSXRWEUBPSDJ9NKIWZNJSXLQKDPJNGZEEU5MZKKRB/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "DLORD,RAYDIUM,3KRWSXRWEUBPSDJ9NKIWZNJSXLQKDPJNGZEEU5MZKKRB",
          "Quote": "USD"
        },
        "decimals": 14,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"30817\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "DLORD,RAYDIUM,3KRWSXRWEUBPSDJ9NKIWZNJSXLQKDPJNGZEEU5MZKKRB/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"6zEqwR82Nbc2wckFWGpsZL58VRa9HCcbC7RT5QkaY5Ff\",\"token_decimals\":6},\"quote_token_vault\":{\"token_vault_address\":\"DafZwvyWmovRUSDV2Dnje9GW2BpUbDxUt5BJM7VwAhnM\",\"token_decimals\":9},\"amm_info_address\":\"BeuMFQpR3j1oZCNLs3nVfDRwQqdADojfEmn2dGKGUEc3\",\"open_orders_address\":\"DApLjBRd47UtQcHBq8LLY1TQ1LZDSBfHik7BugANNmx9\"}"
        }
      ]
    },
    "DMAGA,RAYDIUM,7D7BRCBYEPFI77VXYSAPMEQRNN1WSBBXNFPJGBH5PUMP/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "DMAGA,RAYDIUM,7D7BRCBYEPFI77VXYSAPMEQRNN1WSBBXNFPJGBH5PUMP",
          "Quote": "USD"
        },
        "decimals": 12,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"32514\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SOL,RAYDIUM,SO11111111111111111111111111111111111111112/DMAGA,RAYDIUM,7D7BRCBYEPFI77VXYSAPMEQRNN1WSBBXNFPJGBH5PUMP",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "invert": true,
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"3p7DbnWekWjtMQkey6fp1emdwXbAAxecnYuD5eU9ZZPH\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"5HuEsLbebeuottMsFW3hHcXTrKVsu3WUNSQnSr6aX4eL\",\"token_decimals\":6},\"amm_info_address\":\"CRSwXnfF21yHDXdYBVWwCwQ5Ni2A4QuK3qGYis5k7x7k\",\"open_orders_address\":\"DSdk5cF6XbZ4DfDWgn7xHDQM2ZTjYYyD2aPjddJx6Lap\"}"
        }
      ]
    },
    "DOAI,RAYDIUM,3VMFEATR9M2PP5JCFNC8C8U6U4EFUBDQ6FQJGPPCNFKS/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "DOAI,RAYDIUM,3VMFEATR9M2PP5JCFNC8C8U6U4EFUBDQ6FQJGPPCNFKS",
          "Quote": "USD"
        },
        "decimals": 11,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"32425\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "DOAI,RAYDIUM,3VMFEATR9M2PP5JCFNC8C8U6U4EFUBDQ6FQJGPPCNFKS/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"7scBfDZQqUFNpKDR9h9QS1zPKXJD3KSBQvjTQBiUpopZ\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"DaBRgufq7KznTnoCe1Hj4A8mTpSi971N8YHxuaVjkDc1\",\"token_decimals\":9},\"amm_info_address\":\"5ATbUQwZYiPTxBWQ5MueYBHsFoamBXFhEQ56tqTF7PjN\",\"open_orders_address\":\"4TKdnREBajgT1pmkBMrn4AmEiz1ZzWunukGqwPiHXLNr\"}"
        }
      ]
    },
    "DODO,RAYDIUM,DODOYOGRFCC7XRWDJFMQKLVCCYJKAVQPNQVPFZQ6KJYS/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "DODO,RAYDIUM,DODOYOGRFCC7XRWDJFMQKLVCCYJKAVQPNQVPFZQ6KJYS",
          "Quote": "USD"
        },
        "decimals": 16,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"32530\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "DODO,RAYDIUM,DODOYOGRFCC7XRWDJFMQKLVCCYJKAVQPNQVPFZQ6KJYS/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"9V5sSuZn43UCibxNK7XFCg5PJwiAytDvf6QfG2FrRUiV\",\"token_decimals\":6},\"quote_token_vault\":{\"token_vault_address\":\"DtEYdRWUHap5dsnKkpnN21FSfamoanQP8hGggnf4w9h6\",\"token_decimals\":9},\"amm_info_address\":\"GV7xTA28Gpy4u7vqcuYBpmXV5G4VeN2nRqYxZVwheru4\",\"open_orders_address\":\"ENNpkSv4UYFwmHLLBS79u8447uCfNxA9QTD5SivEWbbk\"}"
        }
      ]
    },
    "DOKY,RAYDIUM,5RS53FY3Q4T4MLK9ZBQ45CVNGF1RH7NQRNXIP6PA5RYH/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "DOKY,RAYDIUM,5RS53FY3Q4T4MLK9ZBQ45CVNGF1RH7NQRNXIP6PA5RYH",
          "Quote": "USD"
        },
        "decimals": 14,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"31106\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "DOKY,RAYDIUM,5RS53FY3Q4T4MLK9ZBQ45CVNGF1RH7NQRNXIP6PA5RYH/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"GRknDb4EPX4jdCPzrdoF4KQMmr424GD8Ri93Xbg56CMu\",\"token_decimals\":6},\"quote_token_vault\":{\"token_vault_address\":\"E2DZDmGSE3mqQX36QMhdFh64nCgh4qbTjJQUXh8FuU2C\",\"token_decimals\":9},\"amm_info_address\":\"AGmUouqWyRdq7Pb3Y3HCmaWpTEBLHwszwdnabbGiTvpD\",\"open_orders_address\":\"J7iEcnxTbg3WJjzAAmk3W2QqVhmQgiFeTtMxodtNgUxc\"}"
        }
      ]
    },
    "DOLLAR,RAYDIUM,5ANPDX9GPOSBI9JSW2DFFE5QQD3FMXBUDOQUMNDXPUMP/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "DOLLAR,RAYDIUM,5ANPDX9GPOSBI9JSW2DFFE5QQD3FMXBUDOQUMNDXPUMP",
          "Quote": "USD"
        },
        "decimals": 13,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"32474\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SOL,RAYDIUM,SO11111111111111111111111111111111111111112/DOLLAR,RAYDIUM,5ANPDX9GPOSBI9JSW2DFFE5QQD3FMXBUDOQUMNDXPUMP",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "invert": true,
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"9g1ZpGbxokC5Gc4X2KBZ3KHuuNzNhCthQG5oZVhPtCHS\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"6eV9cx3NdszGrRXYsnfVGPhSh84e5epXpougdPGHgnre\",\"token_decimals\":6},\"amm_info_address\":\"J9n4vSqRFnWiERTW2NzWs4TimjxWXRdMgu528Xmy7om7\",\"open_orders_address\":\"9VxXSSgSPNTS53Unjf1j33GJi3rgu1imTYZDha8ZckPg\"}"
        }
      ]
    },
    "DOOGLE,RAYDIUM,F6TSRCJTLBZKDTZYQJTPVQ9WTNWHMMC1WCQGUEGCPUMP/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "DOOGLE,RAYDIUM,F6TSRCJTLBZKDTZYQJTPVQ9WTNWHMMC1WCQGUEGCPUMP",
          "Quote": "USD"
        },
        "decimals": 13,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"32067\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SOL,RAYDIUM,SO11111111111111111111111111111111111111112/DOOGLE,RAYDIUM,F6TSRCJTLBZKDTZYQJTPVQ9WTNWHMMC1WCQGUEGCPUMP",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "invert": true,
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"9rRYtSuSEiRDjK1e7KTYvU4rp6JXrnh5eb4sSmmZbd2e\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"GDpCuz7SxU3Ks5SEnPaMXNMjBhzTz5fzbvdnUo2EVmN2\",\"token_decimals\":6},\"amm_info_address\":\"Hv4wMatEoiREbyZjEW5V3AYHQ7wwdtrMWcp26GLM6hBr\",\"open_orders_address\":\"FdsXY5GFCutsyLqWdzU8bng9HAsiJJ1VCeaXm1wnxkSj\"}"
        }
      ]
    },
    "DOUG,RAYDIUM,BAVUJ8BNTC79A8AHTXQI1EUHCCNQVEU8KSBE4SVCAAHC/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "DOUG,RAYDIUM,BAVUJ8BNTC79A8AHTXQI1EUHCCNQVEU8KSBE4SVCAAHC",
          "Quote": "USD"
        },
        "decimals": 12,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"31670\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "DOUG,RAYDIUM,BAVUJ8BNTC79A8AHTXQI1EUHCCNQVEU8KSBE4SVCAAHC/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"41NGcpmJGjRfJ2MpimqmkQC5EM7MHygYdVeeQjUhK3md\",\"token_decimals\":6},\"quote_token_vault\":{\"token_vault_address\":\"Ew5HrbdQ4HT1nzg1LG1aj3Ke9AeVztxBFFEeXbGqrnhb\",\"token_decimals\":9},\"amm_info_address\":\"EcgRHhYtUL6f3U92f7xkaF2arSQXPjKGTAvwW7qHof2w\",\"open_orders_address\":\"DtfPax8fMWsXWqarfopYbbAo9cu7rCQmfAnGbJZETAa2\"}"
        }
      ]
    },
    "DUKO,RAYDIUM,HLPTM5E6RTGH4EKGDPYFRNRHBJPKMYVDEEREEA2G7RF9/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "DUKO,RAYDIUM,HLPTM5E6RTGH4EKGDPYFRNRHBJPKMYVDEEREEA2G7RF9",
          "Quote": "USD"
        },
        "decimals": 12,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"29494\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "DUKO,RAYDIUM,HLPTM5E6RTGH4EKGDPYFRNRHBJPKMYVDEEREEA2G7RF9/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"HD7aZ6YrqAxVbGNAMEKxozcW1ZDU7pbKfd7XMmZtxyzk\",\"token_decimals\":6},\"quote_token_vault\":{\"token_vault_address\":\"A9J2mXPXfRZ7Sh2ymUgCJM4p9iUjZBcyAfrz49PoBBN4\",\"token_decimals\":9},\"amm_info_address\":\"BGS69Ju7DRRVxw9b2B5TnrMLzVdJcscV8UtKywqNsgwx\",\"open_orders_address\":\"FoBQDGey332Ppv1KiTow8z9oZP8n6mEPLyhedPdG1nUG\"}"
        }
      ]
    },
    "DUST,RAYDIUM,DUSTAWUCRTSGU8HCQRDHDCBUYHCPADMLM2VCCB8VNFNQ/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "DUST,RAYDIUM,DUSTAWUCRTSGU8HCQRDHDCBUYHCPADMLM2VCCB8VNFNQ",
          "Quote": "USD"
        },
        "decimals": 10,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"18802\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "DUST,RAYDIUM,DUSTAWUCRTSGU8HCQRDHDCBUYHCPADMLM2VCCB8VNFNQ/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"823yDgBfDqVtMQn33xqsDqADtqsNZibThd7QFE96gbbi\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"45Z629voh31VufjhAMDrhRtLhptkF3eFTxp2tkJQVfJr\",\"token_decimals\":9},\"amm_info_address\":\"syNSyUTeJf8rohN5LRZkcka4Jh78eQHwoDDrZxaYdzd\",\"open_orders_address\":\"FRPR2nXwifKZoBpDupxnD8m29G72qi32Adyytq9gGNXT\"}"
        }
      ]
    },
    "EGL,RAYDIUM,DCCA1IVQRYFQWSQTZ3ARKFKPD6UXV6AN8JDW5S3EVKQZ/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "EGL,RAYDIUM,DCCA1IVQRYFQWSQTZ3ARKFKPD6UXV6AN8JDW5S3EVKQZ",
          "Quote": "USD"
        },
        "decimals": 12,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"32718\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "EGL,RAYDIUM,DCCA1IVQRYFQWSQTZ3ARKFKPD6UXV6AN8JDW5S3EVKQZ/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"GF5r9VTYwsBMBkekxfwsuoPoXz11tKZiSFBwgbHntkxU\",\"token_decimals\":6},\"quote_token_vault\":{\"token_vault_address\":\"J16yJxfr2SuFioEhx1Fr4mPKH796BC9GjQT4o72SRRbm\",\"token_decimals\":9},\"amm_info_address\":\"2sSHtoatQorkimsc4Pkq9EHrN1Swd9zbMak7ozucMSK1\",\"open_orders_address\":\"52KKqBKn4T8cp9spg2ayUy5mG4tUSSazUojTimTFU37i\"}"
        }
      ]
    },
    "ELON,RAYDIUM,7ZCM8WBN9ALA3O47SOYCTU6ILDJ7WKGG5SV2HE5CGTD5/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "ELON,RAYDIUM,7ZCM8WBN9ALA3O47SOYCTU6ILDJ7WKGG5SV2HE5CGTD5",
          "Quote": "USD"
        },
        "decimals": 16,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"9436\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SOL,RAYDIUM,SO11111111111111111111111111111111111111112/ELON,RAYDIUM,7ZCM8WBN9ALA3O47SOYCTU6ILDJ7WKGG5SV2HE5CGTD5",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "invert": true,
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"GYFqFXJfRQqsT8EJaJsnCkYHBt64hQKvzzQsknre41nP\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"9s6vZyVdAxsa3ZgVHGqfEn7cbkiaCpr1LWUmdgFrjTgf\",\"token_decimals\":4},\"amm_info_address\":\"EF16cKCMkrtwwdkRKDy7oumVV1n5A6sohGszVmipXcdj\",\"open_orders_address\":\"FPV7Aag3VPQf6kU6RauHpQVjnoCtaZh121PUJFSZqihE\"}"
        }
      ]
    },
    "EPIK,RAYDIUM,3BGWJ8B7B9HHX4SGFZ2KJHV9496COVFSMK2YEPEVSBRW/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "EPIK,RAYDIUM,3BGWJ8B7B9HHX4SGFZ2KJHV9496COVFSMK2YEPEVSBRW",
          "Quote": "USD"
        },
        "decimals": 12,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"30571\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SOL,RAYDIUM,SO11111111111111111111111111111111111111112/EPIK,RAYDIUM,3BGWJ8B7B9HHX4SGFZ2KJHV9496COVFSMK2YEPEVSBRW",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "invert": true,
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"9aGBQqKRyC5bbrZsnZJJtp59EqJj7vBkgV3HehgKEu5y\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"ANpMJb9ToMGNivLEdmBNBC2Qcf5ASaZkEdmUddV1FUZB\",\"token_decimals\":6},\"amm_info_address\":\"AZaaQaRhp1ys9VaJBRZYbmPz3JSBSp7m8cSSrLBn4BP9\",\"open_orders_address\":\"FjCKdnpN1t262QGGn6chWYRtoSaY6fuYxyKoqhinyGEK\"}"
        }
      ]
    },
    "EXGO,RAYDIUM,D5YMUBHSVOVYKAGUCGKNK2CM8UYGKTNTXJ62T3C46NXS/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "EXGO,RAYDIUM,D5YMUBHSVOVYKAGUCGKNK2CM8UYGKTNTXJ62T3C46NXS",
          "Quote": "USD"
        },
        "decimals": 11,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"23459\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "EXGO,RAYDIUM,D5YMUBHSVOVYKAGUCGKNK2CM8UYGKTNTXJ62T3C46NXS/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"AFBvfgzLao8bAzZoaDHVxxhnhM86tTES2pQswBj7x6pv\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"4e3DngwhpogjomrQY1s3WJc5Kg8D1bdD1ZVKSeDauDhZ\",\"token_decimals\":9},\"amm_info_address\":\"MFS3mcPqADm6r7vDhwQCL8yZ3qD1P4mgpWSnhpXoh5W\",\"open_orders_address\":\"ERnv1SL3ykHAnZ6Z9EPZuzqLtgdS3R6Mb8kMz4Vg2MpJ\"}"
        }
      ]
    },
    "FALX,RAYDIUM,AFO4NUMBNHDXC7M7P6QJZ1PF3LBQYFG5K1CNRGVE8RVU/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "FALX,RAYDIUM,AFO4NUMBNHDXC7M7P6QJZ1PF3LBQYFG5K1CNRGVE8RVU",
          "Quote": "USD"
        },
        "decimals": 13,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"31384\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "FALX,RAYDIUM,AFO4NUMBNHDXC7M7P6QJZ1PF3LBQYFG5K1CNRGVE8RVU/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"6TtniYJPdHJ764d3rUvk8SokmzyZMYCgqUjSgUAnjma2\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"2eH2DVXQvs5qWwDQjgiSPsdZ19KQsj84RKRfCeorsGem\",\"token_decimals\":9},\"amm_info_address\":\"2hPp2aKd6T6HZmMQW2LkqH7R1wLZDjzZ1bZjhj5nrhrV\",\"open_orders_address\":\"3jf9f9VJdUXQQha6nHJkZVxqBTW5oJUNHYuKLyfDDuMM\"}"
        }
      ]
    },
    "FLUFFI,RAYDIUM,6CEJCG7JO5RV9KFSGKX66RPW19NRSCMCCD2BXFWPUMP/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "FLUFFI,RAYDIUM,6CEJCG7JO5RV9KFSGKX66RPW19NRSCMCCD2BXFWPUMP",
          "Quote": "USD"
        },
        "decimals": 12,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"32825\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SOL,RAYDIUM,SO11111111111111111111111111111111111111112/FLUFFI,RAYDIUM,6CEJCG7JO5RV9KFSGKX66RPW19NRSCMCCD2BXFWPUMP",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "invert": true,
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"EBbWMLL14vttchJpmnE6ChPDyAc4RPuni52BhDvZPmbo\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"EarYm3s8Gc2gU91ZVXSogpn15MEGb3pX6ZNju48WKip3\",\"token_decimals\":6},\"amm_info_address\":\"ALUefQoiZnMzeJqanpYwjGwG44drhA1tYFuXCEpgcEpr\",\"open_orders_address\":\"9Vf2CzzKb6Eys4m3Dkt8KcZiyVXTD6zd4iotBxgSrUg3\"}"
        }
      ]
    },
    "FTT,RAYDIUM,AGFEAD2ET2ZJIF9JAGPDMIXQQVW5I81ABDVKE7PHNFZ3/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "FTT,RAYDIUM,AGFEAD2ET2ZJIF9JAGPDMIXQQVW5I81ABDVKE7PHNFZ3",
          "Quote": "USD"
        },
        "decimals": 9,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"4195\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "FTT,RAYDIUM,AGFEAD2ET2ZJIF9JAGPDMIXQQVW5I81ABDVKE7PHNFZ3/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"CFqS3UG7TwboneHdsr46Jb7qj5b9YzeZn59X22F8HZQx\",\"token_decimals\":6},\"quote_token_vault\":{\"token_vault_address\":\"ArsBckC8nngtBjAfvcGrXSB82M91ouukET8cNeXVMda7\",\"token_decimals\":9},\"amm_info_address\":\"63J7uzJKogFK3d3T2Y2c5NSVjSQrCbpfuBhdqH6H5pSJ\",\"open_orders_address\":\"3aW9jx5QZ2nnucv5LGPDLcXYGnxRytgfZqrBXyewzeew\"}"
        }
      ]
    },
    "GIGA,RAYDIUM,63LFDMNB3MQ8MW9MTZ2TO9BEA2M71KZUUGQ5TIJXCQJ9/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "GIGA,RAYDIUM,63LFDMNB3MQ8MW9MTZ2TO9BEA2M71KZUUGQ5TIJXCQJ9",
          "Quote": "USD"
        },
        "decimals": 11,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"30063\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "GIGA,RAYDIUM,63LFDMNB3MQ8MW9MTZ2TO9BEA2M71KZUUGQ5TIJXCQJ9/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"BpHq5J4A78xLo3Z8WNEEN98F1xMdr5Wq5dyzzsC9Mpdp\",\"token_decimals\":5},\"quote_token_vault\":{\"token_vault_address\":\"Bj5iLRaZpfSYYc1u5jEevQBxW2Wb1wg4RKUpvRfoqSwA\",\"token_decimals\":9},\"amm_info_address\":\"4xxM4cdb6MEsCxM52xvYqkNbzvdeWWsPDZrBcTqVGUar\",\"open_orders_address\":\"9sGKv3mz4EUzNZwdgeRNwJ8cbujErKHQ31cvbuNTSAgF\"}"
        }
      ]
    },
    "GINNAN,RAYDIUM,GINNABFFZL4FUJ9VACTXHA74GDAW8KDPGAHQMTMZPS2F/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "GINNAN,RAYDIUM,GINNABFFZL4FUJ9VACTXHA74GDAW8KDPGAHQMTMZPS2F",
          "Quote": "USD"
        },
        "decimals": 15,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"32524\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "GINNAN,RAYDIUM,GINNABFFZL4FUJ9VACTXHA74GDAW8KDPGAHQMTMZPS2F/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"FkSzskEZFC25GhWX2V7wonvfHgB3HhesiyNKaTjmoUwf\",\"token_decimals\":6},\"quote_token_vault\":{\"token_vault_address\":\"EwcLkzY7hnfkBpSWRPAv4xLQ5W85ErXRGNBFqMmuJC5C\",\"token_decimals\":9},\"amm_info_address\":\"AUPVtmxNccSq5LYSrLRXEjCfAtpfGvDyJfdZCwKyia3G\",\"open_orders_address\":\"79dXhsa2dHiNTnv7cepwmcDqcSHvEgadhWdE7vZZSSvz\"}"
        }
      ]
    },
    "GME,RAYDIUM,8WXTPEU6557ETKP9WHFY1N1ECU6NXDVBAGGHGSMYIHSB/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "GME,RAYDIUM,8WXTPEU6557ETKP9WHFY1N1ECU6NXDVBAGGHGSMYIHSB",
          "Quote": "USD"
        },
        "decimals": 12,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"29241\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "GME,RAYDIUM,8WXTPEU6557ETKP9WHFY1N1ECU6NXDVBAGGHGSMYIHSB/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"2kdQwdXdwxSJdzFEFtuo9tpmA88FVjMH7F5kgBZNHPR5\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"2VzMCxFW3nj7pNDbF6WTJgxoJ8ekJS9hfquxPq8edLTY\",\"token_decimals\":9},\"amm_info_address\":\"9tz6vYKiBDLYx2RnGWC5tESu4pyVE4jD6Tm56352UGte\",\"open_orders_address\":\"Gg5UhU7AyhM142a7pdk6WBtdf9rv1VcmJQXgESLjxmsk\"}"
        }
      ]
    },
    "GOAT,RAYDIUM,59U8QAD2S2GETSY5VS7DJ95YSPNHAVYHETRFZCEB9F7G/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "GOAT,RAYDIUM,59U8QAD2S2GETSY5VS7DJ95YSPNHAVYHETRFZCEB9F7G",
          "Quote": "USD"
        },
        "decimals": 16,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"30647\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "GOAT,RAYDIUM,59U8QAD2S2GETSY5VS7DJ95YSPNHAVYHETRFZCEB9F7G/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"2x8sbxHVkNBC1tjE6PerbyfwnxWLvZ6zVbdN5gTpqbDk\",\"token_decimals\":6},\"quote_token_vault\":{\"token_vault_address\":\"67qaEoGD69BBBkCYakGvvZjRQU1DS4WE5RkuuEDjkTsK\",\"token_decimals\":9},\"amm_info_address\":\"H6WLzofWdE5jLoU7E4oL6EecWWDpndp5iMh9nqiGE3b5\",\"open_orders_address\":\"EezqtLFXL3PjKgn3tgeACkbUH45SfvB6fdYokKACyTxf\"}"
        }
      ]
    },
    "GUAC,RAYDIUM,AZSHEMXD36BJ1EMNXHOWJAJPUXZRKCK57WW4ZGXVA7YR/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "GUAC,RAYDIUM,AZSHEMXD36BJ1EMNXHOWJAJPUXZRKCK57WW4ZGXVA7YR",
          "Quote": "USD"
        },
        "decimals": 17,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"24935\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "GUAC,RAYDIUM,AZSHEMXD36BJ1EMNXHOWJAJPUXZRKCK57WW4ZGXVA7YR/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"92fCMBLKZqkAdUWthyvErXPiZDoDQh3RCv6JJ9g4qNAx\",\"token_decimals\":5},\"quote_token_vault\":{\"token_vault_address\":\"D54WC6GWeX8sYAaXhDFzvZNAtLHjoSUmzeFJ2tSn1RXn\",\"token_decimals\":9},\"amm_info_address\":\"9TBMGVkzW9RdRMR9XggsXYcLUqPEcB26rRGaiyfTXh2X\",\"open_orders_address\":\"28yPPeAummTgs76af1E5aUmYp6VyDuuiBgucvBzwPwEB\"}"
        }
      ]
    },
    "GUMMY,RAYDIUM,H7BTHGB5CVO5FGE5JBDNDPUV8KYKQNZYZA3QZ8SH7YXW/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "GUMMY,RAYDIUM,H7BTHGB5CVO5FGE5JBDNDPUV8KYKQNZYZA3QZ8SH7YXW",
          "Quote": "USD"
        },
        "decimals": 12,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"30803\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "GUMMY,RAYDIUM,H7BTHGB5CVO5FGE5JBDNDPUV8KYKQNZYZA3QZ8SH7YXW/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"33AQRrPaZTckDJQd5DZstiwi11tcMVryu63V8rAHFF7N\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"4DHoHzkMHYqJeNDpkdeL6AGDymLFjJnS4SRsJHoT52Bm\",\"token_decimals\":9},\"amm_info_address\":\"FMiecMsYhPdBf94zZKa7i6inK1GX7aypLf7QewNz1i6w\",\"open_orders_address\":\"FSv96pMp3x5XwFdYgqXUY47o7nSKhA6tvCHX1UZZPWnv\"}"
        }
      ]
    },
    "HABIBI,RAYDIUM,864YJRB3JAVARC4FNUDTPKFXDESYRBB39NWXKZUDXY46/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "HABIBI,RAYDIUM,864YJRB3JAVARC4FNUDTPKFXDESYRBB39NWXKZUDXY46",
          "Quote": "USD"
        },
        "decimals": 11,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"31189\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "HABIBI,RAYDIUM,864YJRB3JAVARC4FNUDTPKFXDESYRBB39NWXKZUDXY46/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"AseQmyRFtmH2KGcBtsnDmVGiH68WP32KEak7VshLddr5\",\"token_decimals\":6},\"quote_token_vault\":{\"token_vault_address\":\"6Vj4gTxdkkhS2DFgyLuoAv1c1iTWZuYnZpw7bhK8oeQj\",\"token_decimals\":9},\"amm_info_address\":\"2ukgjDC99Nk34RfRjWjCoHAuQLtLnz8TLcBrDQk3f2ay\",\"open_orders_address\":\"Ap3oiBWsLbDFwcigjktNvt2WjPQnLReRd28wtqJE1yDF\"}"
        }
      ]
    },
    "HAMI,RAYDIUM,4SP2EUDRQF46RZUN6SYAWZJRXWUPX2T3NJUOKMV766RJ/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "HAMI,RAYDIUM,4SP2EUDRQF46RZUN6SYAWZJRXWUPX2T3NJUOKMV766RJ",
          "Quote": "USD"
        },
        "decimals": 12,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"30550\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "HAMI,RAYDIUM,4SP2EUDRQF46RZUN6SYAWZJRXWUPX2T3NJUOKMV766RJ/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"DRWQsj4hANod7KpxyXDSRUswgDNFuTawnjp5fysx8Ad\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"AAT9KHuEPSJ3wGzZN76V9yq28v1kR6HnfswdRUMSoTK\",\"token_decimals\":9},\"amm_info_address\":\"84Nu43GsHLaCDuVSvgiTwpQDTNhYejKFEkJmeXBFJmqQ\",\"open_orders_address\":\"CGxztvFguH3YpkdFd7XZAdSKqNrF4qg954pN6cFRS47p\"}"
        }
      ]
    },
    "HAMMY,RAYDIUM,26KMQVGDUOB6REFNJ51YAABWWJND8UMTPNQGSHQ64UDR/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "HAMMY,RAYDIUM,26KMQVGDUOB6REFNJ51YAABWWJND8UMTPNQGSHQ64UDR",
          "Quote": "USD"
        },
        "decimals": 12,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"31284\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "HAMMY,RAYDIUM,26KMQVGDUOB6REFNJ51YAABWWJND8UMTPNQGSHQ64UDR/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"324NgHgEDyU9d7TE9dkAkB2GNtqxdEU4PsYRTDL68qoR\",\"token_decimals\":6},\"quote_token_vault\":{\"token_vault_address\":\"4kPJL1LmempALPjjwMWSo6JRBjmKQY7HX3edozqmJBPe\",\"token_decimals\":9},\"amm_info_address\":\"X131b3frGn4b8ue51EyvrnzWuTuBGoM93uRYrNteEFy\",\"open_orders_address\":\"9WZDqKjvpyoAShnp3Dg1725uyo2aQtgp8z7GG9XdB5NM\"}"
        }
      ]
    },
    "HARAMBE,RAYDIUM,FCH1OIXTPRI8ZXBNMDCEADOJW2TOYFHXQDZACQKWDVSP/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "HARAMBE,RAYDIUM,FCH1OIXTPRI8ZXBNMDCEADOJW2TOYFHXQDZACQKWDVSP",
          "Quote": "USD"
        },
        "decimals": 11,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"29088\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "HARAMBE,RAYDIUM,FCH1OIXTPRI8ZXBNMDCEADOJW2TOYFHXQDZACQKWDVSP/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"5f9Fgcp2C9vdrp75GspNKBjzdaxq5uiqpLVkgtWKpDZZ\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"Apgp3SzNB5VpVWbK5q2ucBvCJEsf1gqXL4iUAqvD9pgB\",\"token_decimals\":9},\"amm_info_address\":\"2BJKy9pnzTDvMPdHJhv8qbWejKiLzebD7i2taTyJxAze\",\"open_orders_address\":\"BPv68DZUMxpqvfRye2JoeK1GRkkGs5PEUycmx5b448x2\"}"
        }
      ]
    },
    "HAWKTUAH,RAYDIUM,4GFE6MBDORSY5BLBIUMRGETR6PZCJYFXMDM5EHSGPUMP/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "HAWKTUAH,RAYDIUM,4GFE6MBDORSY5BLBIUMRGETR6PZCJYFXMDM5EHSGPUMP",
          "Quote": "USD"
        },
        "decimals": 13,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"31923\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SOL,RAYDIUM,SO11111111111111111111111111111111111111112/HAWKTUAH,RAYDIUM,4GFE6MBDORSY5BLBIUMRGETR6PZCJYFXMDM5EHSGPUMP",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "invert": true,
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"2xnf676U7A3RkEjwZi7KYBoMfF4DAZ1z8mgHvDCvmsGa\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"J4eexABsJRqvXuQrnBE3tP2HQGhiVzubRuP86x4HaV5a\",\"token_decimals\":6},\"amm_info_address\":\"ErRdtWwYKdz37oGVjdruq2tXNC9NTX78x32pnXWPeQ7L\",\"open_orders_address\":\"35XjpT5BM5AJ3A3uLxoTyePqvQTijBhrnhA8uSatS4rN\"}"
        }
      ]
    },
    "HEGE,RAYDIUM,ULWSJMMPXMNRFPU6BJNK6RPRKXQD5JXUMPPS1FXHXFY/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "HEGE,RAYDIUM,ULWSJMMPXMNRFPU6BJNK6RPRKXQD5JXUMPPS1FXHXFY",
          "Quote": "USD"
        },
        "decimals": 11,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"31044\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "HEGE,RAYDIUM,ULWSJMMPXMNRFPU6BJNK6RPRKXQD5JXUMPPS1FXHXFY/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"5gXSNFNkVo9v1c1EyjqiMQ4d8C9L7RDVNfTSbw4D5tpT\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"EWSXjkGiv8Bg4tF5rxkCoj1dHpBUeHAad1sunpmyqXYU\",\"token_decimals\":9},\"amm_info_address\":\"CJcu7ciRHBHu4BDnpLgAUm1A6iSp9RuhJMG36rjjrxnd\",\"open_orders_address\":\"CsmGHvqHwLscG5hg4pV6UmEQDdiYZf8e1nXGpzCHea9L\"}"
        }
      ]
    },
    "HEHE,RAYDIUM,BREUHVOHXX5FV6Q41UYB3SOJTAUGOGAIAHKBMTCRPUMP/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "HEHE,RAYDIUM,BREUHVOHXX5FV6Q41UYB3SOJTAUGOGAIAHKBMTCRPUMP",
          "Quote": "USD"
        },
        "decimals": 12,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"32386\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SOL,RAYDIUM,SO11111111111111111111111111111111111111112/HEHE,RAYDIUM,BREUHVOHXX5FV6Q41UYB3SOJTAUGOGAIAHKBMTCRPUMP",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "invert": true,
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"EkyNAwXyA65rDDc185gnShwJqsRfPE65R1sXQb9UQLMc\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"FHstwarKRKJpt3uFBgHEcrafZQcyWx4NbBLJeGmMLWcz\",\"token_decimals\":6},\"amm_info_address\":\"23KJaRate7XthAQ7C5XbJJYK5cyG1sNA2ikCPsiAcbVP\",\"open_orders_address\":\"APaKnoWdk59VWiMYK8ZwJubt5e4B3WArj1kodv4A6XnL\"}"
        }
      ]
    },
    "HNT/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "HNT",
          "Quote": "USD"
        },
        "decimals": 9,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"5665\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "HNT,RAYDIUM,HNTYVP6YFM1HG25TN9WGLQM12B8TQMCKNKRDU1OXWUX/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"FGvuJZP3r6YJdMP6Tnas7FKC4PtVdjFFoSSkr2TY1KRz\",\"token_decimals\":8},\"quote_token_vault\":{\"token_vault_address\":\"HkwY9tMctkiiFYh18RJJaFWx8XD5w51q9dzbhFifokE9\",\"token_decimals\":9},\"amm_info_address\":\"91axdGmMoqQQg2MPUXKxT3bdqX4p2RMLXqZSQgXn2JpM\",\"open_orders_address\":\"8ee9PgNUtocGHFXyN6C9BiD1MUetUz3JCcF2CDVgeUMU\"}"
        },
        {
          "name": "coinbase_ws",
          "off_chain_ticker": "HNT-USD"
        },
        {
          "name": "crypto_dot_com_ws",
          "off_chain_ticker": "HNT_USD"
        }
      ]
    },
    "HOGE,RAYDIUM,74GUPK636NRIEDT9KLSGF1PEWG5BB7ITNKJF2BKMDWZJ/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "HOGE,RAYDIUM,74GUPK636NRIEDT9KLSGF1PEWG5BB7ITNKJF2BKMDWZJ",
          "Quote": "USD"
        },
        "decimals": 14,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"8438\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "HOGE,RAYDIUM,74GUPK636NRIEDT9KLSGF1PEWG5BB7ITNKJF2BKMDWZJ/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"BxDVACAtYU22phTmeoD1Nh8vHjYfqYDfQpSgv5FvGDbb\",\"token_decimals\":8},\"quote_token_vault\":{\"token_vault_address\":\"AgBNzPggicSFDafTdfJZvmLkxFeHPCHRAAQgPDs33iUB\",\"token_decimals\":9},\"amm_info_address\":\"L1NbEqMHUsLEQWfYGSkPXrVaA2bqx4FKPVStxM4UHWk\",\"open_orders_address\":\"HCCqChrPkod7TA7H9jz1k34bsvTpTgQTq897wXqyEGuE\"}"
        }
      ]
    },
    "HONEY/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "HONEY",
          "Quote": "USD"
        },
        "decimals": 11,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"22850\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "HONEY,RAYDIUM,4VMSOUT2BWATFWEUDNQM1XEDRLFJGJ7HSWHCPZ4XGBTY/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"CERR6pnjJikRYGgL6mdd1ZgA4e2mTACyrPBxSYvbXbEY\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"7s6PsNgBYkvVBZiAdzA5EotcM9ZEmzZELg7DAbsehd22\",\"token_decimals\":9},\"amm_info_address\":\"ADNTXTwCXwxo9HvHPp3ynPxwPApsPtuvhBMtg1D9CvHN\",\"open_orders_address\":\"4scqEAd65BFEgxZUabQewVQHWztPV36XT8Vx6szmTbz9\"}"
        },
        {
          "name": "coinbase_ws",
          "off_chain_ticker": "HONEY-USD"
        }
      ]
    },
    "HONK,RAYDIUM,3AG1MJ9AKZ9FAKCQ6GAEHPLSX8B2PUBPDKB9IBSDLZNB/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "HONK,RAYDIUM,3AG1MJ9AKZ9FAKCQ6GAEHPLSX8B2PUBPDKB9IBSDLZNB",
          "Quote": "USD"
        },
        "decimals": 12,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"29153\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "HONK,RAYDIUM,3AG1MJ9AKZ9FAKCQ6GAEHPLSX8B2PUBPDKB9IBSDLZNB/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"3NnEMNfVLvuKBGEGx5EioGi2w6ctYoyXBRkVmByTqDYo\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"AapAeyv3qpPjFoibm6V9aGZ1QkaVWXqhvJQYJkb3aQKt\",\"token_decimals\":9},\"amm_info_address\":\"BZivKpJWgQvrA3yYe3ubomufeGVouoYoUhosmBEdqF9y\",\"open_orders_address\":\"7LiFnmWYtKw9siGLmGXpdQ6VJzADJaAXVJ1yoZzsRYWq\"}"
        }
      ]
    },
    "HUND,RAYDIUM,2XPQOKFJITK8YCMDGBKY7CMZRRYF2X9PNIZECYKDUZEV/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "HUND,RAYDIUM,2XPQOKFJITK8YCMDGBKY7CMZRRYF2X9PNIZECYKDUZEV",
          "Quote": "USD"
        },
        "decimals": 12,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"30007\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "HUND,RAYDIUM,2XPQOKFJITK8YCMDGBKY7CMZRRYF2X9PNIZECYKDUZEV/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"BseLLmtAwhdn5VkoUQP2AMpF56wVDk7QaPJMsoFhBCDB\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"F8fHaGYintJd6WdyeRHQ6DMjaLiqh9PGFjM4y89Jotb4\",\"token_decimals\":9},\"amm_info_address\":\"CtXYCgkfMdTnvtiU8Ju6BjntBb3GRhmwYA8FaV1KCLMu\",\"open_orders_address\":\"eaxbqPg9AjJtdEaE71SaSv95jBS4Bh1kb22r5V1JnMH\"}"
        }
      ]
    },
    "IOT,RAYDIUM,IOTEVVZLEYWOTN1QDWNPDDXPWSZN3ZFHEOT3MFL9FNS/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "IOT,RAYDIUM,IOTEVVZLEYWOTN1QDWNPDDXPWSZN3ZFHEOT3MFL9FNS",
          "Quote": "USD"
        },
        "decimals": 12,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"24686\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "IOT,RAYDIUM,IOTEVVZLEYWOTN1QDWNPDDXPWSZN3ZFHEOT3MFL9FNS/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"B1W8zKqMNKFupe2Gtqyk4Qkm3pfCw6N1aYvTPFvnGf3D\",\"token_decimals\":6},\"quote_token_vault\":{\"token_vault_address\":\"HEmXJFLML3UaFBZA1EB5dJ8tHmQLoLTZWreWviUv1BLF\",\"token_decimals\":9},\"amm_info_address\":\"BvvefSzw1xqW7NSYbfSfspRyW9P79f3QgmD1PfqLEMaL\",\"open_orders_address\":\"BZtZMdKrpmzWFsnnMPkwCnXZ1idXUSb9zxYfBJRLD5iq\"}"
        }
      ]
    },
    "IQ50,RAYDIUM,21RWEMLGYEMNONHW7H3XA5PY17X6ZFRCHIRCP9INRBQA/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "IQ50,RAYDIUM,21RWEMLGYEMNONHW7H3XA5PY17X6ZFRCHIRCP9INRBQA",
          "Quote": "USD"
        },
        "decimals": 15,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"30079\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "IQ50,RAYDIUM,21RWEMLGYEMNONHW7H3XA5PY17X6ZFRCHIRCP9INRBQA/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"7oqq2Cywq27NUsTKvnDU9kQFH7Xm7v2FDFX1igGb7HFz\",\"token_decimals\":6},\"quote_token_vault\":{\"token_vault_address\":\"FZPQ7Usv3bKC68UPT1HyyQktdam4gUJDCq54qJmtqd7v\",\"token_decimals\":9},\"amm_info_address\":\"CJVLgaSSuGarPWLx57f79T1EEMKg26fM1o3MMm1afD6J\",\"open_orders_address\":\"3vfQ2WEYByvpFEWHys1vpEBLJZdYKdn1HCC26GvyjE2j\"}"
        }
      ]
    },
    "JASON,RAYDIUM,6SURYVEUDZ5HQAXAB6QRGFBZWVJN8DC7M29EZSVDPUMP/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "JASON,RAYDIUM,6SURYVEUDZ5HQAXAB6QRGFBZWVJN8DC7M29EZSVDPUMP",
          "Quote": "USD"
        },
        "decimals": 12,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"32097\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SOL,RAYDIUM,SO11111111111111111111111111111111111111112/JASON,RAYDIUM,6SURYVEUDZ5HQAXAB6QRGFBZWVJN8DC7M29EZSVDPUMP",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "invert": true,
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"CXQ2LTd4UvvnznPqsrdo2sLwUkfmLXCSagxMPaEbsuhh\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"BPrCqJaU3biQPPvN8aJHCVutci7fGREzgho7G9eariR7\",\"token_decimals\":6},\"amm_info_address\":\"4XmdX5Umx7qt3uvMapdXzWPQefx2WZf4tZ69XAmNyF3X\",\"open_orders_address\":\"DmbCUVv2N9cxsQtqnGcubkeXKyTv7oUFUSGgwpRWoiuE\"}"
        }
      ]
    },
    "JUNGLE,RAYDIUM,9P32YQUCXFZNDOXMMJNZBPQXQGFOU4TA4SB1RAQH9CYW/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "JUNGLE,RAYDIUM,9P32YQUCXFZNDOXMMJNZBPQXQGFOU4TA4SB1RAQH9CYW",
          "Quote": "USD"
        },
        "decimals": 14,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"31298\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "JUNGLE,RAYDIUM,9P32YQUCXFZNDOXMMJNZBPQXQGFOU4TA4SB1RAQH9CYW/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"BbSGnCJaTvkuTY2cifiK7o54yF9U37WjhLkQeFQ9DThM\",\"token_decimals\":8},\"quote_token_vault\":{\"token_vault_address\":\"AhWQ8RR6fFCHW1RVojXotRWtHfJSFEsj4LNNvbqAp8Xc\",\"token_decimals\":9},\"amm_info_address\":\"FTLA9G7cj1MGxa715JZo3dL9SiNrw2suLCd2fRL5P9g7\",\"open_orders_address\":\"8G74o1egQKHP5fbPVQ3NGrB3n95WBqMWuxHW3395BMiz\"}"
        }
      ]
    },
    "JUP/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "JUP",
          "Quote": "USD"
        },
        "decimals": 10,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"29210\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SOL,RAYDIUM,SO11111111111111111111111111111111111111112/JUP,RAYDIUM,JUPYIWRYJFSKUPIHA7HKER8VUTAEFOSYBKEDZNSDVCN",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "invert": true,
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"Acwf28DzuSVX6eYVWMNuPaD5378twu17NafazqF5PH89\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"BegAFgvdVAkdvZv6G8tTjpuGxd1m6TagyBR26APSt5su\",\"token_decimals\":6},\"amm_info_address\":\"5Qtn7FFKmVH4Brx3V1cR8E4mot2V3eSmvFRBpmKjNy35\",\"open_orders_address\":\"6YZhHwBtcZee8gUWP9Q5UAo45ZTWq4fbmcuzZavDCTqr\"}"
        },
        {
          "name": "crypto_dot_com_ws",
          "off_chain_ticker": "JUP_USDT",
          "normalize_by_pair": {
            "Base": "USDT",
            "Quote": "USD"
          }
        }
      ]
    },
    "KAMA,RAYDIUM,HNKKZR1YTFBUUXM6G3IVRS2RY68KHHGV7BNDFF1GCSJB/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "KAMA,RAYDIUM,HNKKZR1YTFBUUXM6G3IVRS2RY68KHHGV7BNDFF1GCSJB",
          "Quote": "USD"
        },
        "decimals": 12,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"32419\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "KAMA,RAYDIUM,HNKKZR1YTFBUUXM6G3IVRS2RY68KHHGV7BNDFF1GCSJB/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"HQijRirHzRKmWNYq4RF6sGVzgWAoXXDZBaVX4ZS12sFh\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"3c8tfAxmjL6VznS9YNAdjfLQ1XJ8ABNDtLsuUQZa61Yg\",\"token_decimals\":9},\"amm_info_address\":\"JC9HSjtyJLTxqQRYHZt642LBnCfPBofZBaBRHgwV4nR\",\"open_orders_address\":\"F4mdtjqumz8dR7tYWzCRgoaJ3iXSUox5whvUMV71aCWE\"}"
        }
      ]
    },
    "KHAI,RAYDIUM,3TWGDVYBL2YPET2LXNWAWSMEOA8AL4DUTNUWAT2PKCJC/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "KHAI,RAYDIUM,3TWGDVYBL2YPET2LXNWAWSMEOA8AL4DUTNUWAT2PKCJC",
          "Quote": "USD"
        },
        "decimals": 10,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"30948\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "KHAI,RAYDIUM,3TWGDVYBL2YPET2LXNWAWSMEOA8AL4DUTNUWAT2PKCJC/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"6g4rhxVTrN6SrtNvimq4QiU8yA5XScvwL6wxaMkegrtJ\",\"token_decimals\":8},\"quote_token_vault\":{\"token_vault_address\":\"7p2PrGMyeetNRqTKFraL7eYo2TbU3apWz6vfqrZFiPcG\",\"token_decimals\":9},\"amm_info_address\":\"ECbK6PSMZ5yQaUYBocsXaVrax2fWADw2ijTqLGPtt9sC\",\"open_orders_address\":\"2DaRg4UycKL9GSVfARBDrcensb89WD5PyyFX9NrMunLc\"}"
        }
      ]
    },
    "KITTENWIF,RAYDIUM,9A8AMDFQXFJ44RRVHKDIXNKWCWT99IGWP5QRRUC73PIN/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "KITTENWIF,RAYDIUM,9A8AMDFQXFJ44RRVHKDIXNKWCWT99IGWP5QRRUC73PIN",
          "Quote": "USD"
        },
        "decimals": 13,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"30629\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "KITTENWIF,RAYDIUM,9A8AMDFQXFJ44RRVHKDIXNKWCWT99IGWP5QRRUC73PIN/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"3vWMav4PFGC6sGqAd3rfMnfdqZ3XxWLtvQNytLYEDXqe\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"G6X88zc1sTjyaJSBWTxPCuNp9XT8mVi4AEKfeHWZbwD3\",\"token_decimals\":9},\"amm_info_address\":\"FHjxJM4nU7YHCbwqsmRs5M89m6BG1FygopXsj4vFSRVy\",\"open_orders_address\":\"DhekiVGFYaNF7dznD2LQYbUuGrPxVXpziAm1SfnaMwd9\"}"
        }
      ]
    },
    "KITTY,RAYDIUM,8YJ15EE2AUQMWBWPXXLTTEBTZYMGN4MTSRKMQVHW1J1G/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "KITTY,RAYDIUM,8YJ15EE2AUQMWBWPXXLTTEBTZYMGN4MTSRKMQVHW1J1G",
          "Quote": "USD"
        },
        "decimals": 13,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"28894\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "KITTY,RAYDIUM,8YJ15EE2AUQMWBWPXXLTTEBTZYMGN4MTSRKMQVHW1J1G/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"GCA6YEzYiFdTCc2JThfydyxEiH7198B3jR9vCEbRBApP\",\"token_decimals\":6},\"quote_token_vault\":{\"token_vault_address\":\"3LhbwmsDbK1bU2ND8nV9N1K35W59shpW1UUdnG8MRCPd\",\"token_decimals\":9},\"amm_info_address\":\"B7DVswJuQk7UQmPPnjpL7WS2Vym9srAHh1WGu3kPu5QS\",\"open_orders_address\":\"HsN1Y5hFGyHQH77AFhcijQgm79rVfTkEeWxi23vwi4zh\"}"
        }
      ]
    },
    "KOKO,RAYDIUM,FSA54YL49WKS7RWOGV9SUCBSGWCWV756JTD349E6H2YW/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "KOKO,RAYDIUM,FSA54YL49WKS7RWOGV9SUCBSGWCWV756JTD349E6H2YW",
          "Quote": "USD"
        },
        "decimals": 15,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"30089\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "KOKO,RAYDIUM,FSA54YL49WKS7RWOGV9SUCBSGWCWV756JTD349E6H2YW/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"3E2XDCuabamvcaA5QdWp2A4ke8rzsMuqYdTHexuszRtL\",\"token_decimals\":6},\"quote_token_vault\":{\"token_vault_address\":\"594urFoEaGhCAxWRbrmYpKXn1QxHXtwR6rYWQJQCYr4t\",\"token_decimals\":9},\"amm_info_address\":\"6d3YhKJSf1kxFiv5rNW8EZkL6vg2E8XgnMwNz3LQos8x\",\"open_orders_address\":\"9jbUvkNr2WcfvcT4ET4qcm7smfaieC3orEkpcnKMp5my\"}"
        }
      ]
    },
    "LAB,RAYDIUM,2PP6EBUVEL9YRTAUUTMGTWYZKRFYQXGM9JE4S8WPDTEY/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "LAB,RAYDIUM,2PP6EBUVEL9YRTAUUTMGTWYZKRFYQXGM9JE4S8WPDTEY",
          "Quote": "USD"
        },
        "decimals": 12,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"31426\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SOL,RAYDIUM,SO11111111111111111111111111111111111111112/LAB,RAYDIUM,2PP6EBUVEL9YRTAUUTMGTWYZKRFYQXGM9JE4S8WPDTEY",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "invert": true,
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"DbawKC9Tc53TzE2hFdanwPY26yDDS8WhuKyWWg1H6dzy\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"5oY1VkiDTuzMZRUeVSVJJdZgiAppHpgQQhRaFQvpPwQK\",\"token_decimals\":6},\"amm_info_address\":\"Bqi49HnqzXmdFe6ARh9EruhkpUWBRoXdWxGjttiwhRZK\",\"open_orders_address\":\"GUXZa7uxoXmCBPiC16DqPJsErzzkkLFMGjPn28aGwxeK\"}"
        }
      ]
    },
    "LABZ,RAYDIUM,4VC7UYQBO9SIW8ZNKPXFW9D3DZYCIVRPDZS9XRTYRJMH/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "LABZ,RAYDIUM,4VC7UYQBO9SIW8ZNKPXFW9D3DZYCIVRPDZS9XRTYRJMH",
          "Quote": "USD"
        },
        "decimals": 11,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"31121\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "LABZ,RAYDIUM,4VC7UYQBO9SIW8ZNKPXFW9D3DZYCIVRPDZS9XRTYRJMH/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"98Gczsrb27gbYpCUbKNV617bxGV6QYvZ8Kez1h63vQ5Z\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"BEFDPRJRmG4qJmPs3GD57S9LBrAnck9Sf7cUQpQy7s14\",\"token_decimals\":9},\"amm_info_address\":\"7iAYqozANKpRjhvCPGGJdJ6jAc3MWdK9hqBV3DX3R1Lx\",\"open_orders_address\":\"4cMPUZNVtCP88fDn6KnAkYoZzmEV4t7xsod8KsxBSmaW\"}"
        }
      ]
    },
    "LADYF,RAYDIUM,3X8GCLIH2HTTJYQEPG7MAZPMBWBGQ5URUMTYDZ5TKMDE/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "LADYF,RAYDIUM,3X8GCLIH2HTTJYQEPG7MAZPMBWBGQ5URUMTYDZ5TKMDE",
          "Quote": "USD"
        },
        "decimals": 14,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"30071\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "LADYF,RAYDIUM,3X8GCLIH2HTTJYQEPG7MAZPMBWBGQ5URUMTYDZ5TKMDE/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"DS8fj1JcQgkf5Qt9T57UwHqJZWcbSwfqpun3kAiemyo8\",\"token_decimals\":6},\"quote_token_vault\":{\"token_vault_address\":\"GRCoAQ1igerxsSKXGFDXxTFG7sHmBRqtXWivEFixV3qW\",\"token_decimals\":9},\"amm_info_address\":\"7m7kxcMVhMpnaXRAhwmVHKLymZ8yC2XCZBXXNA6g4Ni2\",\"open_orders_address\":\"HGi67UEj9TDTV4boo9Nc5dNsJzZEYxvzCxUWdEqpLh6y\"}"
        }
      ]
    },
    "LAIKA,RAYDIUM,EUOQ6CYQFCJCVSLR9WFAUPDW19Y6ZHWECJOZSEI643I1/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "LAIKA,RAYDIUM,EUOQ6CYQFCJCVSLR9WFAUPDW19Y6ZHWECJOZSEI643I1",
          "Quote": "USD"
        },
        "decimals": 14,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"31845\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "LAIKA,RAYDIUM,EUOQ6CYQFCJCVSLR9WFAUPDW19Y6ZHWECJOZSEI643I1/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"3eoSzSxo5XJiFbhDk2AMbRzqXMjntbp22fsjdoswmNLg\",\"token_decimals\":6},\"quote_token_vault\":{\"token_vault_address\":\"CwRJpNPtJBUTnYNVf6AsVDSWViffpuV7eFuoGKDKSaDo\",\"token_decimals\":9},\"amm_info_address\":\"AJT2iW7P7G3m96vrL3ZRUU4kPnJCciSEjkfik1c9ZKYm\",\"open_orders_address\":\"DJyENBeqKAR8xhm2azRJYNQaGXPaWb6dgP5uUTJWvpJ\"}"
        }
      ]
    },
    "LIBERTA,RAYDIUM,EGQCDSZVIK6T1DHCAGJBVFTC9CD4AQIRIHZWEHYQ84LG/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "LIBERTA,RAYDIUM,EGQCDSZVIK6T1DHCAGJBVFTC9CD4AQIRIHZWEHYQ84LG",
          "Quote": "USD"
        },
        "decimals": 12,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"31821\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "LIBERTA,RAYDIUM,EGQCDSZVIK6T1DHCAGJBVFTC9CD4AQIRIHZWEHYQ84LG/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"9Qdgokf9dG7rrXJDDQaNSNYomH2FXaC9L8hKGU25ddzZ\",\"token_decimals\":6},\"quote_token_vault\":{\"token_vault_address\":\"ErPUFurtutrfD1569tAm4ev4X6a3AXHCqcSi7fkV233n\",\"token_decimals\":9},\"amm_info_address\":\"BBWQtTvUaauEzyhPi72DThuiEoN2XkjC9iN2jeVdxRKE\",\"open_orders_address\":\"5TWeMPo56XHf91CckQ8rJPZ5BJdrj7QN4mjPK2GANnud\"}"
        }
      ]
    },
    "LIKE,RAYDIUM,3BRTIVRVSITBMCTGTQWP7HXXPSYBKJN4XLNTPSHQA3ZR/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "LIKE,RAYDIUM,3BRTIVRVSITBMCTGTQWP7HXXPSYBKJN4XLNTPSHQA3ZR",
          "Quote": "USD"
        },
        "decimals": 11,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"10891\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "LIKE,RAYDIUM,3BRTIVRVSITBMCTGTQWP7HXXPSYBKJN4XLNTPSHQA3ZR/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"EUSm9PEyFhw5dcSVC5bqf54Y5q4CgAdzsWFvqFNtqnLD\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"6womVyd2LZrtHKHk9aCaYJAK5RCjxRPieefDejyRH3hz\",\"token_decimals\":9},\"amm_info_address\":\"CS5KfNp4PJaekA7tE44hAkvTLHLxD6kSE7GGgzVRVvhm\",\"open_orders_address\":\"5YxKdBrq6d23W2tgC6wZpmT2vDscoJpZm1EUQZiTmfZW\"}"
        }
      ]
    },
    "LOS,RAYDIUM,44BZGE9EZJGPJRYNMSA64MDKZ1EELDDCDICZRMOYATEZ/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "LOS,RAYDIUM,44BZGE9EZJGPJRYNMSA64MDKZ1EELDDCDICZRMOYATEZ",
          "Quote": "USD"
        },
        "decimals": 14,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"31217\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "LOS,RAYDIUM,44BZGE9EZJGPJRYNMSA64MDKZ1EELDDCDICZRMOYATEZ/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"Gb3jw1d3YYt6oo6NZdfBgHQy29XhgEKF9QYqYKAdYv2a\",\"token_decimals\":6},\"quote_token_vault\":{\"token_vault_address\":\"DNtrRjVZiTd1iuGeCXKta3QZhcy7fF5tzUUScjCGfL5R\",\"token_decimals\":9},\"amm_info_address\":\"EEanQSZv6LyempQFRf6Xs3j2U4GxRuv2uF7WuDuwGKEU\",\"open_orders_address\":\"9ynRmLYtXALFy65akgE5x3ZzaEFmH1iUNL5iRb2sJ3su\"}"
        }
      ]
    },
    "LOVE,RAYDIUM,4QQV4LQUUXAN1EN1XQGRFY65TFLE5STJCFSCQOZQYB8T/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "LOVE,RAYDIUM,4QQV4LQUUXAN1EN1XQGRFY65TFLE5STJCFSCQOZQYB8T",
          "Quote": "USD"
        },
        "decimals": 13,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"31405\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "LOVE,RAYDIUM,4QQV4LQUUXAN1EN1XQGRFY65TFLE5STJCFSCQOZQYB8T/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"FiZgmoBhn5pEhRhQLtNMRu4hUrwhyvhRE95SUn7iATtf\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"EzyyA1yXVBQPzWKXENorixPZdFF3roBNsfKTHa8VBpMV\",\"token_decimals\":9},\"amm_info_address\":\"3P9GpSSugj7ZYYQNhWJkS6HEmeMHHCg3ZwAgq9r32uhL\",\"open_orders_address\":\"yvaWawMC6pPTtP9uSQGauvbPB9kVyJjCdX9gRM4biF4\"}"
        }
      ]
    },
    "MAD,RAYDIUM,MADHPJRN6BD8T78RSY7NUSUNWWA2HU8BYPOBZPRHBHV/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "MAD,RAYDIUM,MADHPJRN6BD8T78RSY7NUSUNWWA2HU8BYPOBZPRHBHV",
          "Quote": "USD"
        },
        "decimals": 14,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"32103\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "MAD,RAYDIUM,MADHPJRN6BD8T78RSY7NUSUNWWA2HU8BYPOBZPRHBHV/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"FLsJ7FP9oPCV99NYwF1HzFbufGGngAEtj9H6952LhRnj\",\"token_decimals\":6},\"quote_token_vault\":{\"token_vault_address\":\"BixUe826rnUARvgGPMMvL2sFWh8Fdcu2VAFf9vcT2a3W\",\"token_decimals\":9},\"amm_info_address\":\"DUScoEC6LFHv3F9L6at2tG7BV8LnG3epoNvWpyBDzcKp\",\"open_orders_address\":\"3ws8ECSNDmjHMnKJEzTAVML8gE9eHiDAQJ2DD2jv8yWY\"}"
        }
      ]
    },
    "MANEKI/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "MANEKI",
          "Quote": "USD"
        },
        "decimals": 12,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"30912\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "MANEKI,RAYDIUM,25HAYBQFODHFWX9AY6RARBGVWGWDDNQCHSXS3JQ3MTDJ/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"5FeTzLNqwrvSzexFujeV62a2v4kmQUrBnCQjJANStMXj\",\"token_decimals\":5},\"quote_token_vault\":{\"token_vault_address\":\"2kjCeDKKK9pCiDqfsbS72q81RZiUnSwoaruuwz1avUWn\",\"token_decimals\":9},\"amm_info_address\":\"2aPsSVxFw6dGRqWWUKfwujN6WVoyxuhjJaPzYaJvGDDR\",\"open_orders_address\":\"9pd9FFJfVjY1aG9dh96ArJB5F2HAyfj2XryjVTHbJhc9\"}"
        },
        {
          "name": "crypto_dot_com_ws",
          "off_chain_ticker": "MANEKI_USD"
        }
      ]
    },
    "MDT,RAYDIUM,8WQBST4QAN2FQBCCH5GDXQ2WJ7VTNWEY4ONLRPUG7TYA/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "MDT,RAYDIUM,8WQBST4QAN2FQBCCH5GDXQ2WJ7VTNWEY4ONLRPUG7TYA",
          "Quote": "USD"
        },
        "decimals": 11,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"2348\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "MDT,RAYDIUM,8WQBST4QAN2FQBCCH5GDXQ2WJ7VTNWEY4ONLRPUG7TYA/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"49fyQqrqGfz1Fc8mmLMafj8KSMbMpyCR5HyibmnW5U53\",\"token_decimals\":8},\"quote_token_vault\":{\"token_vault_address\":\"8FhYugdrsjnFE9wrfSyZyQuUMca17yKN2LmASWrCjSFe\",\"token_decimals\":9},\"amm_info_address\":\"2DCrFk3Y9dkgcB4FLCL6QyMTuzHWwssXC4YYtb2VeDaZ\",\"open_orders_address\":\"73xNSut57AyNsxWvPrX74e6ugALhGVRkNk6ioLG7E3qV\"}"
        }
      ]
    },
    "MEW,RAYDIUM,MEW1GQWJ3NEXG2QGERIKU7FAFJ79PHVQVREQUZSCPP5/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "MEW,RAYDIUM,MEW1GQWJ3NEXG2QGERIKU7FAFJ79PHVQVREQUZSCPP5",
          "Quote": "USD"
        },
        "decimals": 12,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"30126\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "MEW,RAYDIUM,MEW1GQWJ3NEXG2QGERIKU7FAFJ79PHVQVREQUZSCPP5/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"4HqAjFKuQX7tnXEkxsjGQha3G4bcgER8qPXRahn9gj8F\",\"token_decimals\":5},\"quote_token_vault\":{\"token_vault_address\":\"BhNdEGJef9jSqT1iCEkFZ2bYZCdpC1vuiWtqDt87vBVp\",\"token_decimals\":9},\"amm_info_address\":\"879F697iuDJGMevRkRcnW21fcXiAeLJK1ffsw2ATebce\",\"open_orders_address\":\"CV3Gq5M2R7KRU5ey4LpnZYRR7r7vzKoV9Bt4mZ8P6bSB\"}"
        }
      ]
    },
    "MICHI,RAYDIUM,5MBK36SZ7J19AN8JFOCHHQS4OF8G6BWUJBECSXBSOWDP/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "MICHI,RAYDIUM,5MBK36SZ7J19AN8JFOCHHQS4OF8G6BWUJBECSXBSOWDP",
          "Quote": "USD"
        },
        "decimals": 10,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"30943\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SOL,RAYDIUM,SO11111111111111111111111111111111111111112/MICHI,RAYDIUM,5MBK36SZ7J19AN8JFOCHHQS4OF8G6BWUJBECSXBSOWDP",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "invert": true,
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"EzSLA8B6N2oetEnsPcdiYohjqSKP4MUzwwpnzz75DzwH\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"5tg9qKYRAXgqxtrvzM77XTKBLQBJppixjqKCXcfhNAE1\",\"token_decimals\":6},\"amm_info_address\":\"GH8Ers4yzKR3UKDvgVu8cqJfGzU4cU62mTeg9bcJ7ug6\",\"open_orders_address\":\"8pkc323WErsJdV9jQTxT8yqPCgcpNKFfkViaGhDYm1eF\"}"
        }
      ]
    },
    "MINI,RAYDIUM,2JCXACFWT9MVAWBQ5NZKYWCYXQKRCDSYRDXN6HJ22SBP/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "MINI,RAYDIUM,2JCXACFWT9MVAWBQ5NZKYWCYXQKRCDSYRDXN6HJ22SBP",
          "Quote": "USD"
        },
        "decimals": 12,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"31227\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SOL,RAYDIUM,SO11111111111111111111111111111111111111112/MINI,RAYDIUM,2JCXACFWT9MVAWBQ5NZKYWCYXQKRCDSYRDXN6HJ22SBP",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "invert": true,
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"38cprYrVy2Pewni1wxxjrfmBjFVTVfLMgS1uRdU3dJFc\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"J5uHNPFcPNwJ82MoDLfee9YJm33NQyJdqpobkiUhVvjW\",\"token_decimals\":6},\"amm_info_address\":\"HYpXCaAT9YBu7vYa5BURGprsa23hmvdkqXtSUD5gQWdc\",\"open_orders_address\":\"EtQYH5Rm72PtM9hnbg3tk8igeghHsxiuQnbsENn5GYbX\"}"
        }
      ]
    },
    "MOBILE/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "MOBILE",
          "Quote": "USD"
        },
        "decimals": 12,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"24600\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "MOBILE,RAYDIUM,MB1EU7TZEC71KXDPSMSKOUCSSUUOGLV1DRYS1OP2JH6/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"DHiRnbcrJzfXB1hzjSLVJfWoL6YYkcFaQDmskbCShFmC\",\"token_decimals\":6},\"quote_token_vault\":{\"token_vault_address\":\"95TScor3WpYhBDhWqKsXyF81Me5o4Uh1scHH6cvFw2pJ\",\"token_decimals\":9},\"amm_info_address\":\"B5uP8Zincgjc6psTzy3poAXTWEX6ZbHz6nJYgMVzVrxt\",\"open_orders_address\":\"1ChkMpfPqk59pihBAuDUPqZRffDTEF5psfiWAn1uAJ2\"}"
        },
        {
          "name": "coinbase_ws",
          "off_chain_ticker": "MOBILE-USD"
        }
      ]
    },
    "MONGY,RAYDIUM,FSBPYIGZ4BHUXVSPP7XPJYFTPM5PSLJC2WGZAFADPUMP/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "MONGY,RAYDIUM,FSBPYIGZ4BHUXVSPP7XPJYFTPM5PSLJC2WGZAFADPUMP",
          "Quote": "USD"
        },
        "decimals": 12,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"32360\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SOL,RAYDIUM,SO11111111111111111111111111111111111111112/MONGY,RAYDIUM,FSBPYIGZ4BHUXVSPP7XPJYFTPM5PSLJC2WGZAFADPUMP",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "invert": true,
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"HTadu4YsyBvZQyhTj9FDaCRksFsG5rVNDLvpZ4SZvShA\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"AUj6jFQVpkEeFfsUE32sQA4XMEzXvLVCbbLdk6RdnVTy\",\"token_decimals\":6},\"amm_info_address\":\"B2VgYAnXWbmkm2pqJFdioCwnuuNkf9q4H6tHM7hdjEZo\",\"open_orders_address\":\"BvvpxiK5KSJt8BdeoHbtNxuiZ8iVnScLHHYFyun31Yq4\"}"
        }
      ]
    },
    "MOTHER,RAYDIUM,3S8QX1MSMQRBIWKG2CQYX7NIS1OHMGACUC9C4VFVVDPN/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "MOTHER,RAYDIUM,3S8QX1MSMQRBIWKG2CQYX7NIS1OHMGACUC9C4VFVVDPN",
          "Quote": "USD"
        },
        "decimals": 11,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"31510\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SOL,RAYDIUM,SO11111111111111111111111111111111111111112/MOTHER,RAYDIUM,3S8QX1MSMQRBIWKG2CQYX7NIS1OHMGACUC9C4VFVVDPN",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "invert": true,
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"8uQwXPi1sWwUTVbDBnjznmf4mV44CETiNAh3UENvHejV\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"8ZcrNqaDbqy1H4R2DtmGnuZnJ6TKGSsaGmyRGQQeELLv\",\"token_decimals\":6},\"amm_info_address\":\"HcPgh6B2yHNvT6JsEmkrHYT8pVHu9Xiaoxm4Mmn2ibWw\",\"open_orders_address\":\"1z3rLR6AwR8gjVZ8ArqHp9kdaPrNwPvCdrZ6jGy6wwF\"}"
        }
      ]
    },
    "MOUTAI,RAYDIUM,45EGCWCPXYAGBC7KQBIN4NCFGEZWN7F3Y6NACWXQMCWX/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "MOUTAI,RAYDIUM,45EGCWCPXYAGBC7KQBIN4NCFGEZWN7F3Y6NACWXQMCWX",
          "Quote": "USD"
        },
        "decimals": 12,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"30601\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "MOUTAI,RAYDIUM,45EGCWCPXYAGBC7KQBIN4NCFGEZWN7F3Y6NACWXQMCWX/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"4VkmzH14ETcNhSQLTK6AtL1ZP8UmvWpbNCgokDVfiCcD\",\"token_decimals\":6},\"quote_token_vault\":{\"token_vault_address\":\"HSQPdDCxtGo4fTHeZuBGWtQUqHgRsgdz4BVhTCCAtsTv\",\"token_decimals\":9},\"amm_info_address\":\"578CbhKnpAW5NjbmYku6qSaesZZLy3xwFQ8UkDANzd91\",\"open_orders_address\":\"FCQvrj9mrWN5XsPHDSfKf17i8xbzLxW3Esor7nw42nsp\"}"
        }
      ]
    },
    "MUMU,RAYDIUM,5LAFQURVCO6O7KMZ42EQVEJ9LW31STPYGJEEU5SKOMTA/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "MUMU,RAYDIUM,5LAFQURVCO6O7KMZ42EQVEJ9LW31STPYGJEEU5SKOMTA",
          "Quote": "USD"
        },
        "decimals": 14,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"30285\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "MUMU,RAYDIUM,5LAFQURVCO6O7KMZ42EQVEJ9LW31STPYGJEEU5SKOMTA/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"2Re1H89emr8hNacyDTrm1NU8VEhuwaJX7JwcdDqy5Q6g\",\"token_decimals\":6},\"quote_token_vault\":{\"token_vault_address\":\"4VPXFMpndqZhME27vMqtkuGtBo7hVTA9kEvo87zbjXsA\",\"token_decimals\":9},\"amm_info_address\":\"FvMZrD1qC66Zw8VPrW15xN1N5owUPqpQgNQ5oH18mR4E\",\"open_orders_address\":\"BjWyTUxXSNXN1GNzwR7iRhqmdc3XukYpWFfqy1o94DF2\"}"
        }
      ]
    },
    "MYRO,RAYDIUM,HHJPBHRRN4G56VSYLUT8DL5BV31HKXQSRAHTTUCZEZG4/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "MYRO,RAYDIUM,HHJPBHRRN4G56VSYLUT8DL5BV31HKXQSRAHTTUCZEZG4",
          "Quote": "USD"
        },
        "decimals": 11,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"28382\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SOL,RAYDIUM,SO11111111111111111111111111111111111111112/MYRO,RAYDIUM,HHJPBHRRN4G56VSYLUT8DL5BV31HKXQSRAHTTUCZEZG4",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "invert": true,
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"9Edj5621G5PYhno8WddpALmisTk4LwLtpr7KHoMxq9eP\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"AurG7xviXyZtjK5FUfxe2ajdDQhrVTpZKkWgbfuViER9\",\"token_decimals\":9},\"amm_info_address\":\"HCk6LA93xPVsF8g4v6gjkiCd88tLXwZq4eJwiYNHR8da\",\"open_orders_address\":\"ECSFgxVbXhexJKV26CBoBr4iHKQTZ3CiSWrFgbSD2CD8\"}"
        }
      ]
    },
    "NAP,RAYDIUM,4G86CMXGSMDLETRYNAVMFKPHQZKTVDBYGMRADVTR72NU/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "NAP,RAYDIUM,4G86CMXGSMDLETRYNAVMFKPHQZKTVDBYGMRADVTR72NU",
          "Quote": "USD"
        },
        "decimals": 15,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"29892\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "NAP,RAYDIUM,4G86CMXGSMDLETRYNAVMFKPHQZKTVDBYGMRADVTR72NU/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"ABLksYkz92eK1AbZvxwgfma6Zoz1fKnzhgVGpwBWNQyk\",\"token_decimals\":6},\"quote_token_vault\":{\"token_vault_address\":\"6Q1hGQVEzL8dZCjn6Vb5jvJj61vozD8BRxDWQh6ZAAgY\",\"token_decimals\":9},\"amm_info_address\":\"FQed3Ay883zUcGcLaubkV56JJbweiYjxPSTC84yUxqNd\",\"open_orders_address\":\"HSkmWySG2v5vK7LNvmyVnsjdsoPTCv8AP1ZcvUvh1hJV\"}"
        }
      ]
    },
    "NEIRO,RAYDIUM,CTG3ZGYX79ZRE1MTEDVKMKCGNIIFRK1HJ6YIABROPUMP/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "NEIRO,RAYDIUM,CTG3ZGYX79ZRE1MTEDVKMKCGNIIFRK1HJ6YIABROPUMP",
          "Quote": "USD"
        },
        "decimals": 11,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"32463\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SOL,RAYDIUM,SO11111111111111111111111111111111111111112/NEIRO,RAYDIUM,CTG3ZGYX79ZRE1MTEDVKMKCGNIIFRK1HJ6YIABROPUMP",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "invert": true,
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"Aooy1NjoezvhAxrsJGoV35KCfDTxgmQ6TsSb4tUAmjys\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"Axv6REXfurAwU1uSN2To5FPW3ZMbFo1wxhebU7aDnkAm\",\"token_decimals\":6},\"amm_info_address\":\"HvAqakZgurMR2br1eGWPU6EeFcxzmeW8n6Mn7ejEf3DV\",\"open_orders_address\":\"CqsDYUzSk8in7GpkSsfWGQ3BrsTpd8HvYP1UmKY11hi9\"}"
        }
      ]
    },
    "NEON/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "NEON",
          "Quote": "USD"
        },
        "decimals": 10,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"23015\"}]}"
      },
      "provider_configs": [
        {
          "name": "coinbase_ws",
          "off_chain_ticker": "NEON-USD"
        },
        {
          "name": "crypto_dot_com_ws",
          "off_chain_ticker": "NEON_USD"
        },
        {
          "name": "raydium_api",
          "off_chain_ticker": "SOL,RAYDIUM,SO11111111111111111111111111111111111111112/NEON,RAYDIUM,NEONTJSJSUO3REXG9O6VHUMXW62F9V7ZVMU8M8ZUT44",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "invert": true,
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"GabXD8FFL8DWWnLZ9xjRnCi5cZshPf9Upw2qS6qoeEEe\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"BKo7AUm3eP5kayvZ1nfGcnCCD7Aww2Xcjb5WegtGaSSY\",\"token_decimals\":9},\"amm_info_address\":\"FBs4sXbZKG7YBxdZVXtiygKXp4GkQwa9UoQKZev14d7W\",\"open_orders_address\":\"71s5WWu8FC9X2whyg23i1YnWCMYMCxhTCdnHsNt4KQDn\"}"
        }
      ]
    },
    "NEVER,RAYDIUM,CXRHHSQYW8YTDWC4CSJMMGO7UBUJSXZNZRWHTW9ULDRU/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "NEVER,RAYDIUM,CXRHHSQYW8YTDWC4CSJMMGO7UBUJSXZNZRWHTW9ULDRU",
          "Quote": "USD"
        },
        "decimals": 14,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"29302\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "NEVER,RAYDIUM,CXRHHSQYW8YTDWC4CSJMMGO7UBUJSXZNZRWHTW9ULDRU/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"DuU9X27MsBFmpreXhEAWNMD9ZykpGfpDq7qoUB7MAfxx\",\"token_decimals\":6},\"quote_token_vault\":{\"token_vault_address\":\"6DDCndw2rAhuufGvAS9WPt4bFRXCWcb5YYtEUPzh14H6\",\"token_decimals\":9},\"amm_info_address\":\"8LinWMnf5LVEVARqko7eUyydwzkrLdBUZgeNtu4fAFwA\",\"open_orders_address\":\"BqxAQ8ztSEL9Ukggj4TFEL3wmQsPBKkGu9YxqRaMt9f4\"}"
        }
      ]
    },
    "NOS,RAYDIUM,NOSXBVOACTTYDLVKY6CSB4AC8JCDQKKAAWYTX2ZMOO7/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "NOS,RAYDIUM,NOSXBVOACTTYDLVKY6CSB4AC8JCDQKKAAWYTX2ZMOO7",
          "Quote": "USD"
        },
        "decimals": 9,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"16612\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "NOS,RAYDIUM,NOSXBVOACTTYDLVKY6CSB4AC8JCDQKKAAWYTX2ZMOO7/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"EDG1tk5Ld5rzny9cjxJ4eWNsLbW4ZtFSZD2jgG9M5w3X\",\"token_decimals\":6},\"quote_token_vault\":{\"token_vault_address\":\"aQmagQqNSFSbSLw5GmX4d1e69eARcUqvK14PkxW1HU5\",\"token_decimals\":9},\"amm_info_address\":\"3BDWiAL6e9AkCCDymq8ULQL7U7p5nzbNj1ZMxKULvPtK\",\"open_orders_address\":\"C5Vyj1hvx2NNABHvzUkJRE2jYzcGXBLnvAFoPpi1J2si\"}"
        }
      ]
    },
    "NPCS,RAYDIUM,5TODNKIBAK6K697RRYNGTBURU7YZNFZFX7JZSD1UC7PK/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "NPCS,RAYDIUM,5TODNKIBAK6K697RRYNGTBURU7YZNFZFX7JZSD1UC7PK",
          "Quote": "USD"
        },
        "decimals": 12,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"32412\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "NPCS,RAYDIUM,5TODNKIBAK6K697RRYNGTBURU7YZNFZFX7JZSD1UC7PK/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"57na116WY1UAi1P9f9yWoXtu6TrHMpM3dSV2NSv3RGQq\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"ectv52aafMsmu2N6XSwbLgwrcSe8eGMYQ8jLFQ2XGbb\",\"token_decimals\":9},\"amm_info_address\":\"4BKRQ2iL3Rv8mSpDsFM5FNkZ9SGq4iaqrYtgNWjGE3s4\",\"open_orders_address\":\"8ECbqNa1fmvAhpeq7kvzuLTquPhTd2fiYNksAUvqR2md\"}"
        }
      ]
    },
    "NSO,RAYDIUM,HGMFSGNDLQ6VGLXCW4J33NJRWV2ZTH81IEJNVWK9KCHD/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "NSO,RAYDIUM,HGMFSGNDLQ6VGLXCW4J33NJRWV2ZTH81IEJNVWK9KCHD",
          "Quote": "USD"
        },
        "decimals": 13,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"30409\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "NSO,RAYDIUM,HGMFSGNDLQ6VGLXCW4J33NJRWV2ZTH81IEJNVWK9KCHD/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"Hdv5ccEXKdtsLAubpoZaitL54V2FdSPBMLfDofRQ8L8x\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"2HX489QtXG9vV3MeCVmG8NcQTk1mCTikRRpkoKTDvdcG\",\"token_decimals\":9},\"amm_info_address\":\"FYKqgJdZ3Hauux3KE2SZ4B1AeNJtMHSrWVyyLPYqjxMC\",\"open_orders_address\":\"3GgfucjVa2iFBpmWavXPSxCg4g4vcNsZ1zdHnfog9rXW\"}"
        }
      ]
    },
    "NUB,RAYDIUM,GTDZKAQVMZMNTI46ZEWMIXCA4OXF4BZXWQPOKZXPFXZN/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "NUB,RAYDIUM,GTDZKAQVMZMNTI46ZEWMIXCA4OXF4BZXWQPOKZXPFXZN",
          "Quote": "USD"
        },
        "decimals": 12,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"30493\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "NUB,RAYDIUM,GTDZKAQVMZMNTI46ZEWMIXCA4OXF4BZXWQPOKZXPFXZN/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"9uNqUwneLXbQ6YKndciL5aBXTLJhwpyDXkZmaBbWfwWz\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"75DrZt3zmGSFfKaYDm7yHLKMrr35Wy8ffBNN1143PWbj\",\"token_decimals\":9},\"amm_info_address\":\"83G6VzJzLRCnHBsLATj94VCpRimyyqwuN6ZfL11McADL\",\"open_orders_address\":\"CLXBUkh3hMKNDRUZFFKS721Q1NJb11oHrYvV66QMBcVv\"}"
        }
      ]
    },
    "ORACLE,RAYDIUM,3TU6NPNDFVQNOBQKMGJDWZE1LBVE9SRUDQHSSQXWK1HL/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "ORACLE,RAYDIUM,3TU6NPNDFVQNOBQKMGJDWZE1LBVE9SRUDQHSSQXWK1HL",
          "Quote": "USD"
        },
        "decimals": 12,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"32647\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "ORACLE,RAYDIUM,3TU6NPNDFVQNOBQKMGJDWZE1LBVE9SRUDQHSSQXWK1HL/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"Hqf7kRKBReY2yjCxtPv6NEW1c4Ef6rTUfXwNdgoqJFUP\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"CLCpnUyVspF1KE7irXChECFwgGmD48FGaL27JFgq7gYx\",\"token_decimals\":9},\"amm_info_address\":\"EPT9yuKoZJd25WnPH2HoDAth8CCrAKW18L6yoRXtuTRT\",\"open_orders_address\":\"Aa3SDQT3SksgupqFrcFfLCkJMnF2zH7EYhE8mNNddsdv\"}"
        }
      ]
    },
    "ORBS,RAYDIUM,7JNHPPJBBKSTJ7IEMSIGSBCPJGBCKW28UCRXTQGIMNCP/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "ORBS,RAYDIUM,7JNHPPJBBKSTJ7IEMSIGSBCPJGBCKW28UCRXTQGIMNCP",
          "Quote": "USD"
        },
        "decimals": 14,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"3835\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "ORBS,RAYDIUM,7JNHPPJBBKSTJ7IEMSIGSBCPJGBCKW28UCRXTQGIMNCP/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"63WrZMAU7e6Uf2njzMiMbk88fifa99TAp3b3bApCqb9w\",\"token_decimals\":8},\"quote_token_vault\":{\"token_vault_address\":\"57NCuk4cDVwuLXGmW7DRyafJM3Hxktn17LR9tR3qwRXG\",\"token_decimals\":9},\"amm_info_address\":\"6RPQM9fcQ35EupsrJbJU89hwQcd2NCyoGUrFNj8v55Zb\",\"open_orders_address\":\"F91MXzkyfcXN99ENdsTRDjqSAGtxVVHnMaGnSSeCyvmc\"}"
        }
      ]
    },
    "ORBT,RAYDIUM,BGYJASMSZYM9HHIZ1LBU4EJ7KCTRJMSPBN4ZTRU3W5VF/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "ORBT,RAYDIUM,BGYJASMSZYM9HHIZ1LBU4EJ7KCTRJMSPBN4ZTRU3W5VF",
          "Quote": "USD"
        },
        "decimals": 10,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"30448\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "ORBT,RAYDIUM,BGYJASMSZYM9HHIZ1LBU4EJ7KCTRJMSPBN4ZTRU3W5VF/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"HFewPr7uuyVb2HqK1t6SCCjNrfs6Az1oRBf5D8hUYJAa\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"CNUhZpiNNyH7fRRTPpW8kfUhZwCHbuueaEzFViBbEm4A\",\"token_decimals\":9},\"amm_info_address\":\"2UfT57k2oE13nbRxfQgAPrpygu287wvC2a8YDPs3JJU5\",\"open_orders_address\":\"3dpjuBMjLpUaYnmKExgNrj2A9M5us5xr1g1g62kAi9zL\"}"
        }
      ]
    },
    "ORC,RAYDIUM,CZYWQYWNZACQC7NPHTAPHC8CB2H7OU6U9TEHBRVNZJKX/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "ORC,RAYDIUM,CZYWQYWNZACQC7NPHTAPHC8CB2H7OU6U9TEHBRVNZJKX",
          "Quote": "USD"
        },
        "decimals": 11,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"32285\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "ORC,RAYDIUM,CZYWQYWNZACQC7NPHTAPHC8CB2H7OU6U9TEHBRVNZJKX/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"Di9PMgBq2ccR7tfptYcjrszwwtGwERjJitCDTEk6KYSS\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"BjKdgjvErTjxHv6XrnrGmfvpDWXNJEkGFq89Q9d14zk4\",\"token_decimals\":9},\"amm_info_address\":\"DMi4ftMhn3PtY2gqs1C1WfWTgBtuSz9NkQoESQwKWq3u\",\"open_orders_address\":\"8q53zxa9gLwvQxs2nikUTHs13Aki63DMeDUAZAf1YAq3\"}"
        }
      ]
    },
    "OXY,RAYDIUM,Z3DN17YLAGMKFFVOGEFHQ9ZWVCXGQGF3PQNDSNS2G6M/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "OXY,RAYDIUM,Z3DN17YLAGMKFFVOGEFHQ9ZWVCXGQGF3PQNDSNS2G6M",
          "Quote": "USD"
        },
        "decimals": 11,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"8029\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "OXY,RAYDIUM,Z3DN17YLAGMKFFVOGEFHQ9ZWVCXGQGF3PQNDSNS2G6M/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"4Pg7BQRkptCyJUkAnyvb2LjmxU4qhZd5o1QnHctA19E4\",\"token_decimals\":6},\"quote_token_vault\":{\"token_vault_address\":\"GErwmWBkma9wogivGJos7hEckAjDgXDpaF8ZxR2n8UQo\",\"token_decimals\":9},\"amm_info_address\":\"8w1zbJwrtar9baSaxd1aE3HaahRNviJ27J8Eteb4psaq\",\"open_orders_address\":\"BFpfwLRs6kmjwtsYEnG4FKx4HUfedVe7q3RVYH333ckA\"}"
        }
      ]
    },
    "PAJAMAS,RAYDIUM,FVER7SSVY5GQAMAWF7QFB5MNUUMDDBPNPG4NCA4ZHOLW/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "PAJAMAS,RAYDIUM,FVER7SSVY5GQAMAWF7QFB5MNUUMDDBPNPG4NCA4ZHOLW",
          "Quote": "USD"
        },
        "decimals": 12,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"29855\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "PAJAMAS,RAYDIUM,FVER7SSVY5GQAMAWF7QFB5MNUUMDDBPNPG4NCA4ZHOLW/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"2amsF7CaXcxBDU39e8H8Cm4EFJWJqhWhJ4TBgFFvkbMQ\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"F7vbaUrc9z1CMWDqVtpQCpSQ8m5k5s3WkAf7NVAHdemD\",\"token_decimals\":9},\"amm_info_address\":\"BqricZnjjtFg8wuTbckV6NZcTstuR7BZtKJtzH8oV3eK\",\"open_orders_address\":\"8eSiN9JD5WYJVznfu4EWwPnEDMSvwfSx12NyXdhhkUJ9\"}"
        }
      ]
    },
    "PENG,RAYDIUM,A3EME5CETYZPBOWBRUWY3TSE25S6TB18BA9ZPBWK9EFJ/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "PENG,RAYDIUM,A3EME5CETYZPBOWBRUWY3TSE25S6TB18BA9ZPBWK9EFJ",
          "Quote": "USD"
        },
        "decimals": 10,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"29787\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "PENG,RAYDIUM,A3EME5CETYZPBOWBRUWY3TSE25S6TB18BA9ZPBWK9EFJ/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"2g5q7fBGKZm2CXix8JjK4ZFdBTHQ1LerxkseBTqWuDdD\",\"token_decimals\":6},\"quote_token_vault\":{\"token_vault_address\":\"GmLJXUzjQAAU86a91hKesg5P9pKb6p9AZaGBEZLaDySD\",\"token_decimals\":9},\"amm_info_address\":\"AxBDdiMK9hRPLMPM7k6nCPC1gRARgXQHNejfP2LvNGr6\",\"open_orders_address\":\"9E5VWkY1UsbhkXW4Lk1YovkVouWMG57CuCNXUmecrGpC\"}"
        }
      ]
    },
    "PIP,RAYDIUM,HHJOYWUP5AU6PNRVN4S2PWEERWXNZKHXKGYJRJMOBJLW/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "PIP,RAYDIUM,HHJOYWUP5AU6PNRVN4S2PWEERWXNZKHXKGYJRJMOBJLW",
          "Quote": "USD"
        },
        "decimals": 11,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"16996\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "PIP,RAYDIUM,HHJOYWUP5AU6PNRVN4S2PWEERWXNZKHXKGYJRJMOBJLW/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"6vkUw4ZpdU7iKJTu7VbbKokkg1UZntQwX4nfKeyHtrEA\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"Cy4LjRFbnSf26SSufMuDcERixYyHvD55k3QojUcAcyfb\",\"token_decimals\":9},\"amm_info_address\":\"5Fk3tGEUJAnoyo8n48guBs5S5BX1Fso15gKcJ615V9Rs\",\"open_orders_address\":\"7i4VNjrZgBoyyJ3RA8Q6WCGAWbKpe3VhiARSkMcuZhv2\"}"
        }
      ]
    },
    "PONKE,RAYDIUM,5Z3EQYQO9HICES3R84RCDMU2N7ANPDMXRHDK8PSWMRRC/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "PONKE,RAYDIUM,5Z3EQYQO9HICES3R84RCDMU2N7ANPDMXRHDK8PSWMRRC",
          "Quote": "USD"
        },
        "decimals": 10,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"29150\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "PONKE,RAYDIUM,5Z3EQYQO9HICES3R84RCDMU2N7ANPDMXRHDK8PSWMRRC/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"D7rw7fyEzo9EQcozjqAHJwbdbywGcSLw1at5MioZtMZ4\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"8DcvfWidQ53a3SCBrWxBWL2UU5zEBAKEypApiiCCFu2Y\",\"token_decimals\":9},\"amm_info_address\":\"5uTwG3y3F5cx4YkodgTjWEHDrX5HDKZ5bZZ72x8eQ6zE\",\"open_orders_address\":\"ECoptgCPMxXXWtxv3Zg2V3E7SpWp6SGqKqj32AcdWRQK\"}"
        }
      ]
    },
    "POODL,RAYDIUM,7F7DGNNEL1RWBED6EAO5BE8KNURTNLZNRZEVJKCGJGQD/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "POODL,RAYDIUM,7F7DGNNEL1RWBED6EAO5BE8KNURTNLZNRZEVJKCGJGQD",
          "Quote": "USD"
        },
        "decimals": 11,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"32760\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "POODL,RAYDIUM,7F7DGNNEL1RWBED6EAO5BE8KNURTNLZNRZEVJKCGJGQD/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"HKC2LZ6gFpU2VKUGHQUKBU2LcejyryN2DR58Kg2uuyVi\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"8N4NKJbbi7HXVV6u64QWM7Hn9XRF6cV4eL3FygGmBDCk\",\"token_decimals\":9},\"amm_info_address\":\"CbnU6a4gPqjrdQ5aNj6kufheDDRmrZ7apW1osaDPHQbY\",\"open_orders_address\":\"95pS9TF8VJdb8Be6JJu5rss3majbjfZJLofSv2rdKoPN\"}"
        }
      ]
    },
    "POPCAT,RAYDIUM,7GCIHGDB8FE6KNJN2MYTKZZCRJQY3T9GHDC8UHYMW2HR/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "POPCAT,RAYDIUM,7GCIHGDB8FE6KNJN2MYTKZZCRJQY3T9GHDC8UHYMW2HR",
          "Quote": "USD"
        },
        "decimals": 10,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"28782\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "POPCAT,RAYDIUM,7GCIHGDB8FE6KNJN2MYTKZZCRJQY3T9GHDC8UHYMW2HR/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"4Vc6N76UBu26c3jJDKBAbvSD7zPLuQWStBk7QgVEoeoS\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"n6CwMY77wdEftf2VF6uPvbusYoraYUci3nYBPqH1DJ5\",\"token_decimals\":9},\"amm_info_address\":\"FRhB8L7Y9Qq41qZXYLtC2nw8An1RJfLLxRF2x9RwLLMo\",\"open_orders_address\":\"4ShRqC2n3PURN7EiqmB8X4WLR81pQPvGLTPjL9X8SNQp\"}"
        }
      ]
    },
    "PORT3,RAYDIUM,BO5T8ZKE6XEN9ZIREQONQCZUGCSOPWNNXBBCW9HQCLGJ/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "PORT3,RAYDIUM,BO5T8ZKE6XEN9ZIREQONQCZUGCSOPWNNXBBCW9HQCLGJ",
          "Quote": "USD"
        },
        "decimals": 11,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"29030\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "PORT3,RAYDIUM,BO5T8ZKE6XEN9ZIREQONQCZUGCSOPWNNXBBCW9HQCLGJ/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"7sAGX9xpqKgQUFUUi9vsCUUuU9ERYdt1F3oudiQJqAws\",\"token_decimals\":6},\"quote_token_vault\":{\"token_vault_address\":\"J9LnfRDt9WKBM91PkE3aqP75X1WrkCYXETd27WnR9JAA\",\"token_decimals\":9},\"amm_info_address\":\"8B5ckFm32BcjX63ExchSNnNUrp9u2KGYqro5HBsK6pc4\",\"open_orders_address\":\"6VNL21y6xYLw3WaVnX7EyuXiSre8A9AR547bwoiBq56c\"}"
        }
      ]
    },
    "POS,RAYDIUM,B8VV6AN7XFF3BARB1CMU7TMFKNJJES2WVY7JWQIRC6K6/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "POS,RAYDIUM,B8VV6AN7XFF3BARB1CMU7TMFKNJJES2WVY7JWQIRC6K6",
          "Quote": "USD"
        },
        "decimals": 13,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"27856\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SOL,RAYDIUM,SO11111111111111111111111111111111111111112/POS,RAYDIUM,B8VV6AN7XFF3BARB1CMU7TMFKNJJES2WVY7JWQIRC6K6",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "invert": true,
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"GCembWqoRcadDXHHycMGgL9f2nihBcyMCUab7d6bPKqu\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"Ge9xfjs5UqDQa3Lf6CtFCQfXqrQ75r75daNokduPdGid\",\"token_decimals\":6},\"amm_info_address\":\"4JK54eNmo1R5HrtRaCiwwwY94h1dNugteEnvKnWqMveL\",\"open_orders_address\":\"HUgDbbvgmxMcDG7g1HqNiSf2Ce987wrDMQngfywufaXF\"}"
        }
      ]
    },
    "POWSCHE,RAYDIUM,8CKISHHJDHJV4LUOIRMLUHQG58CUKBYJRTCP4Z3MCXNF/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "POWSCHE,RAYDIUM,8CKISHHJDHJV4LUOIRMLUHQG58CUKBYJRTCP4Z3MCXNF",
          "Quote": "USD"
        },
        "decimals": 11,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"30645\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "POWSCHE,RAYDIUM,8CKISHHJDHJV4LUOIRMLUHQG58CUKBYJRTCP4Z3MCXNF/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"8eYoGvRJqx85yaE7zPEhvystbej183W1g15w3QmaWbCv\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"FdkGtq5KF2LyHw8BprgWTxRKYXU2rcENU8h6cYcn5rTw\",\"token_decimals\":9},\"amm_info_address\":\"4rhMnZ7hfhs4mPjtHN9FJEtcmm1tF8VDYCPh5rEfgTqe\",\"open_orders_address\":\"7bWSAxPu6bVKq43gXGkB33f9cYWGbzC38zmMN1q8SGB8\"}"
        }
      ]
    },
    "PUNDU,RAYDIUM,WSKZSKQEW3ZSMRHPAEVFVZB6PUULZWOV9MJWZSFDEPC/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "PUNDU,RAYDIUM,WSKZSKQEW3ZSMRHPAEVFVZB6PUULZWOV9MJWZSFDEPC",
          "Quote": "USD"
        },
        "decimals": 12,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"30109\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "PUNDU,RAYDIUM,WSKZSKQEW3ZSMRHPAEVFVZB6PUULZWOV9MJWZSFDEPC/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"FQoagUqLxpNq69dpqFLrKm1gySC92NLKMkVgtdHWMKtt\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"GpnU5CAFUyyodrtXTBfvWK7ewsmcJGNksr4Fe49AvpM5\",\"token_decimals\":9},\"amm_info_address\":\"7yEXWTjLyXwBEjMhNwP9dWVJp8G68JvY9KXGT83sDCaM\",\"open_orders_address\":\"CmQ1XbjSeg5opqnJ1nnVf3TBzXUA2KbeWgysCVwKaCSN\"}"
        }
      ]
    },
    "PYTH,RAYDIUM,HZ1JOVNIVVGRGNIIYVEOZEVGZ58XAU3RKWX8EACQBCT3/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "PYTH,RAYDIUM,HZ1JOVNIVVGRGNIIYVEOZEVGZ58XAU3RKWX8EACQBCT3",
          "Quote": "USD"
        },
        "decimals": 10,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"28177\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SOL,RAYDIUM,SO11111111111111111111111111111111111111112/PYTH,RAYDIUM,HZ1JOVNIVVGRGNIIYVEOZEVGZ58XAU3RKWX8EACQBCT3",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "invert": true,
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"HLX4E1BbDfoX4e5rzhw7Wwd1pArPzXd8XdZ3DaBXY9BR\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"5vqxzbPytvVm58vPkA2WYWvoAYfa8AoxD9i1UXCdqzCN\",\"token_decimals\":6},\"amm_info_address\":\"59VtDHQrcDKswbaStjTsjNiqLngjX72UHFqChnpAb93p\",\"open_orders_address\":\"8Cv8eC4PscdjZmiN1v9oVU7KtDXLd6zbGJkiofjN2Jcf\"}"
        }
      ]
    },
    "PYUSD/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "PYUSD",
          "Quote": "USD"
        },
        "decimals": 10,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"27772\"}]}"
      },
      "provider_configs": [
        {
          "name": "coinbase_ws",
          "off_chain_ticker": "PYUSD-USD"
        },
        {
          "name": "crypto_dot_com_ws",
          "off_chain_ticker": "SOL_PYUSD",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "invert": true
        }
      ]
    },
    "PZP,RAYDIUM,2CNFY3WZU715MDBJ1CEWFMNCSKNQKREPJTMW4O1PSJ9J/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "PZP,RAYDIUM,2CNFY3WZU715MDBJ1CEWFMNCSKNQKREPJTMW4O1PSJ9J",
          "Quote": "USD"
        },
        "decimals": 10,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"22556\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "PZP,RAYDIUM,2CNFY3WZU715MDBJ1CEWFMNCSKNQKREPJTMW4O1PSJ9J/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"3yfMeS89fmnDypBBr3xDQ5ERHkFCJ5T5VotDuDftRKGd\",\"token_decimals\":8},\"quote_token_vault\":{\"token_vault_address\":\"B3Ra1VmMYDuMtPwJryqpdGLcVeLFNNCgSNfAWGo6V65w\",\"token_decimals\":9},\"amm_info_address\":\"DCYmxmtkmkWnLj8fP1bCEgZMJJz5zm6ebZ59aECPS62D\",\"open_orders_address\":\"865zWLMxpXfx1oGCFvENNDUVQFhsXWyEKejkPDieETi9\"}"
        }
      ]
    },
    "RADX,RAYDIUM,FVFZXEEWZRXOOU8N8PJGPWKHHZOUDJPFW6RZVAEXQGVY/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "RADX,RAYDIUM,FVFZXEEWZRXOOU8N8PJGPWKHHZOUDJPFW6RZVAEXQGVY",
          "Quote": "USD"
        },
        "decimals": 14,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"31378\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "RADX,RAYDIUM,FVFZXEEWZRXOOU8N8PJGPWKHHZOUDJPFW6RZVAEXQGVY/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"6tPu3hgGmEyBohHSuxeCrNss15oRA3CSeEGiGQcrDsP6\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"8ykP6FdSUius14zGR9GVnzB98vnkzTfYDgyK6ngvsegp\",\"token_decimals\":9},\"amm_info_address\":\"BPGrrEspGcLUeFp5MyK9TbBNxgXGr23FdR71mM89KGx1\",\"open_orders_address\":\"BZ8D2yk4VWXaz91m6awi4ErUW3KKPdLVcjHQrPFD5rqo\"}"
        }
      ]
    },
    "RARE/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "RARE",
          "Quote": "USD"
        },
        "decimals": 10,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"11294\"}]}"
      },
      "provider_configs": [
        {
          "name": "coinbase_ws",
          "off_chain_ticker": "RARE-USD"
        }
      ]
    },
    "RAY,RAYDIUM,4K3DYJZVZP8EMZWUXBBCJEVWSKKK59S5ICNLY3QRKX6R/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "RAY,RAYDIUM,4K3DYJZVZP8EMZWUXBBCJEVWSKKK59S5ICNLY3QRKX6R",
          "Quote": "USD"
        },
        "decimals": 9,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"8526\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SOL,RAYDIUM,SO11111111111111111111111111111111111111112/RAY,RAYDIUM,4K3DYJZVZP8EMZWUXBBCJEVWSKKK59S5ICNLY3QRKX6R",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "invert": true,
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"2jVkTjKnb8i2q5wQ2EeaAjHqYNcTbdM117C88rz2MRBL\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"7ENCdzn4Jnn3prbZLReLbw8a3bLp4oUMf69C7XA7WZcy\",\"token_decimals\":6},\"amm_info_address\":\"AQptcJhCg5k1BQpTtFDVvuZAekhm5eS49oneMfwZW9V5\",\"open_orders_address\":\"FBGa4iqmKDZTenSCEEYQDw3KN8CLgTowAjMd24bftLfd\"}"
        }
      ]
    },
    "RBT,RAYDIUM,65NTNUJGHME4PQVKQYJYKKP1BJAKK4A8Q66SD2YBWUGF/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "RBT,RAYDIUM,65NTNUJGHME4PQVKQYJYKKP1BJAKK4A8Q66SD2YBWUGF",
          "Quote": "USD"
        },
        "decimals": 15,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"9808\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "RBT,RAYDIUM,65NTNUJGHME4PQVKQYJYKKP1BJAKK4A8Q66SD2YBWUGF/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"8ri5wkpUKPRwJXtboPBxZEXZGJTNEUSYzrqzotDzL3CF\",\"token_decimals\":6},\"quote_token_vault\":{\"token_vault_address\":\"PnC8Na2qubk5QhFhjase7Ue1PgK8b3Je8RTVGzkSASu\",\"token_decimals\":9},\"amm_info_address\":\"DpKZKyq6Hd1twPZhdhyED19hbdrBPhMto7rv89QoZFU3\",\"open_orders_address\":\"EKYkmdTF5kBwYs3mr9oz5B9nPhjsbSqmQqFgMufGB5Jc\"}"
        }
      ]
    },
    "RENDER/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "RENDER",
          "Quote": "USD"
        },
        "decimals": 9,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"5690\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SOL,RAYDIUM,SO11111111111111111111111111111111111111112/RENDER,RAYDIUM,RNDRIZKT3MK1IIMDXRDWABCF7ZG7AR5T4NUD4EKHBOF",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "invert": true,
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"HTzYooXyc7QXfn7S2QRMQ1JtgLHDwYdB85t8Fh1RxdCn\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"G9zpMudRgBPxff8MvmEHFf8qnCpfEYFqwwydFEdg5L5G\",\"token_decimals\":8},\"amm_info_address\":\"AmJjgt9ccLooiNpyVtB5NRNHX32vwa7rxvV7Jti57354\",\"open_orders_address\":\"9vMAddmhZW5PbGtGyW4LoobVu9UA4cGVHARmLq82ArUt\"}"
        },
        {
          "name": "coinbase_ws",
          "off_chain_ticker": "RENDER-USD"
        }
      ]
    },
    "RETARDIO,RAYDIUM,6OGZHHZDRQR9PGV6HZ2MNZE7URZBMAFYBBWUYP1FHITX/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "RETARDIO,RAYDIUM,6OGZHHZDRQR9PGV6HZ2MNZE7URZBMAFYBBWUYP1FHITX",
          "Quote": "USD"
        },
        "decimals": 11,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"31921\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "RETARDIO,RAYDIUM,6OGZHHZDRQR9PGV6HZ2MNZE7URZBMAFYBBWUYP1FHITX/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"HXzTvbuKKPyNMmLKJb8vaSUaRZsVS2J2AAsDuDm36rNC\",\"token_decimals\":6},\"quote_token_vault\":{\"token_vault_address\":\"HNcAAdLKHSRnwdmmWCYnP5Zcd11sfGpAoCuWFtugt2ma\",\"token_decimals\":9},\"amm_info_address\":\"5eLRsN6qDQTQSBF8KdW4B8mVpeeAzHCCwaDptzMyszxH\",\"open_orders_address\":\"5TcDuBbtU8Q6LagcM8wfw1Ux2MWgCC6Q1FY22FVDZnXX\"}"
        }
      ]
    },
    "REXHAT,RAYDIUM,ACV2T3MWLUQMIIQCSAFVM35ZWPQKMLRFRTAW3716FZVI/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "REXHAT,RAYDIUM,ACV2T3MWLUQMIIQCSAFVM35ZWPQKMLRFRTAW3716FZVI",
          "Quote": "USD"
        },
        "decimals": 12,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"32101\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "REXHAT,RAYDIUM,ACV2T3MWLUQMIIQCSAFVM35ZWPQKMLRFRTAW3716FZVI/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"8FbGyxRFAvKz4vj2TrZQu6EisH3cH9aqWRMirfFgL8xU\",\"token_decimals\":6},\"quote_token_vault\":{\"token_vault_address\":\"EPsxbG9zceguTMNUB5hrscsBDKyCS9gfTiexB24wCYkY\",\"token_decimals\":9},\"amm_info_address\":\"CyDX4CtUgz4MSdVhnfZuKFzbZRZfRnBoSEw24qS1qWVz\",\"open_orders_address\":\"6Qmq2oUVxcRGSx95ni6DWLC79is8MJXQByxGSCFnqsYC\"}"
        }
      ]
    },
    "RNT,RAYDIUM,2FUFHZYD47MAPV9WCFXH5GNQWFXTQCYU9XAN4THBPUMP/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "RNT,RAYDIUM,2FUFHZYD47MAPV9WCFXH5GNQWFXTQCYU9XAN4THBPUMP",
          "Quote": "USD"
        },
        "decimals": 12,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"31705\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SOL,RAYDIUM,SO11111111111111111111111111111111111111112/RNT,RAYDIUM,2FUFHZYD47MAPV9WCFXH5GNQWFXTQCYU9XAN4THBPUMP",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "invert": true,
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"CLuBFFfERr2NqZL46T3Ng6TzDmv1edWdU5HhG8XHm3BE\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"FwjAy3zL3ErTx37JAHkbQSoFPL6wLRmAW8qZ2ZSVP9kD\",\"token_decimals\":6},\"amm_info_address\":\"9LfXeYQgTXJWhyTQhykCSnfUDd1ffCYA1LcSdcwaRLBk\",\"open_orders_address\":\"GGpNUZJLNmM2oBhAAF5FbxooEW7dER4UxcAR9vchEd8f\"}"
        }
      ]
    },
    "ROA,RAYDIUM,5TB5D6DGJMXXHYMNKFJNG237X6PZGEWTZGPUUH62YQJ7/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "ROA,RAYDIUM,5TB5D6DGJMXXHYMNKFJNG237X6PZGEWTZGPUUH62YQJ7",
          "Quote": "USD"
        },
        "decimals": 11,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"23799\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SOL,RAYDIUM,SO11111111111111111111111111111111111111112/ROA,RAYDIUM,5TB5D6DGJMXXHYMNKFJNG237X6PZGEWTZGPUUH62YQJ7",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "invert": true,
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"ApH7fvKMcvSds1qRvAADnfoBo3XTFtJ6Tj8ALTQSGsV8\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"EQPfCAi1MgsBgXVeY9a7r7igEP4T1sYn1MFVwBbWdJyE\",\"token_decimals\":9},\"amm_info_address\":\"4KWuoRvN2JnWsL9Qe89amJ24eRmSRanLAwfepuqXSBsQ\",\"open_orders_address\":\"6CvY9jPeMNHgSR3WsdanubokU9qgsPnDWBv3wFhT7284\"}"
        }
      ]
    },
    "RODAI,RAYDIUM,GDBYLSNKHKLXTZVEO8QRGKVMREXEEZUYVHPSFUZ9TDKC/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "RODAI,RAYDIUM,GDBYLSNKHKLXTZVEO8QRGKVMREXEEZUYVHPSFUZ9TDKC",
          "Quote": "USD"
        },
        "decimals": 18,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"29305\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "RODAI,RAYDIUM,GDBYLSNKHKLXTZVEO8QRGKVMREXEEZUYVHPSFUZ9TDKC/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"9hHfCgD53kK6cXtDFQEeP2eAcfKc8KmPb6G4TCwzXcT4\",\"token_decimals\":5},\"quote_token_vault\":{\"token_vault_address\":\"GkUnWKeBLD2TDkvqr72fhXjnzMbUkJ2dNCrhAU48224C\",\"token_decimals\":9},\"amm_info_address\":\"CzWqL4M1CzQiw45djF1xdrWHQmuJbJoZPY38ezKGb6q8\",\"open_orders_address\":\"2S2z5e9zSDdw7ZKWbs7oqMirhJQHEv9gMMPxzSuBiYkb\"}"
        }
      ]
    },
    "SBABE,RAYDIUM,D9MFKGNZHNQGRTZKVNJ44YVOLTJMFBZRAHXIUKCAZRE4/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "SBABE,RAYDIUM,D9MFKGNZHNQGRTZKVNJ44YVOLTJMFBZRAHXIUKCAZRE4",
          "Quote": "USD"
        },
        "decimals": 14,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"29355\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SBABE,RAYDIUM,D9MFKGNZHNQGRTZKVNJ44YVOLTJMFBZRAHXIUKCAZRE4/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"7VCVNUnjfzytDk2ExvKJbDdN3vXZBm1KbXe4snuNrzEL\",\"token_decimals\":6},\"quote_token_vault\":{\"token_vault_address\":\"7PGY6iCzYsvP56HfZgJWWWCoSEFB2PAifwkZnzLxZuf6\",\"token_decimals\":9},\"amm_info_address\":\"HqUJPMaitfPAR2oQyzsVezirpaiuToD5kL3hxto9N7cP\",\"open_orders_address\":\"HwyfNuNEEz9bHY2NRj9bZ7cPpvD4JQjeh7xVGTVdjED5\"}"
        }
      ]
    },
    "SBAE,RAYDIUM,BWWWURBODJGBOVFETC3FC6OSBQKDOE62E1XQZ7X4PUMP/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "SBAE,RAYDIUM,BWWWURBODJGBOVFETC3FC6OSBQKDOE62E1XQZ7X4PUMP",
          "Quote": "USD"
        },
        "decimals": 13,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"31905\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SOL,RAYDIUM,SO11111111111111111111111111111111111111112/SBAE,RAYDIUM,BWWWURBODJGBOVFETC3FC6OSBQKDOE62E1XQZ7X4PUMP",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "invert": true,
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"APcG13ZnDBe8uRc4hwD2sFv9ud1WquhgQkvVjucgRLXc\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"26cmLKdGB3AieeywCtsnxzVJSigTksJTPxgjK8cMx6kS\",\"token_decimals\":6},\"amm_info_address\":\"H1FtEraKbT6CSJZnVyb5drumCSACXzdDrZXGD6kzUXY9\",\"open_orders_address\":\"C7zjJdKgscrLXDzZMKURRyP5K1hMbsCoXgbyAzNh7Dmv\"}"
        }
      ]
    },
    "SC,RAYDIUM,6D7NAB2XSLD7CAUWU1WKK6KBSJOHJMP2QZH9GEFVI5UI/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "SC,RAYDIUM,6D7NAB2XSLD7CAUWU1WKK6KBSJOHJMP2QZH9GEFVI5UI",
          "Quote": "USD"
        },
        "decimals": 11,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"30309\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SOL,RAYDIUM,SO11111111111111111111111111111111111111112/SC,RAYDIUM,6D7NAB2XSLD7CAUWU1WKK6KBSJOHJMP2QZH9GEFVI5UI",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "invert": true,
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"DQPzUPZ5GFVVyNSf7tvMLu5NkN2UxUt1VijXGMxJfU8q\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"5ZgybxRW9AMd2U3P3wpzsgCY8UUDCBR488ievHuNiG8B\",\"token_decimals\":6},\"amm_info_address\":\"BSzedbEvWRqVksaF558epPWCM16avEpyhm2HgSq9WZyy\",\"open_orders_address\":\"H3Pbn9sMF8wA7hUGgV18YTt6LdR6ALpKaNf4y6j8ep9Q\"}"
        }
      ]
    },
    "SCF,RAYDIUM,GIG7HR61RVM4CSUXJMGICOYSFQTDIWXTQF64MSRPPUMP/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "SCF,RAYDIUM,GIG7HR61RVM4CSUXJMGICOYSFQTDIWXTQF64MSRPPUMP",
          "Quote": "USD"
        },
        "decimals": 11,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"32615\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SOL,RAYDIUM,SO11111111111111111111111111111111111111112/SCF,RAYDIUM,GIG7HR61RVM4CSUXJMGICOYSFQTDIWXTQF64MSRPPUMP",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "invert": true,
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"4dSwMRxxaTWeTtPNafKY67XaBCUdkVd5J9PrUmUVVZBH\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"4SRQ3FyFH2eU3F7VHk77aBuHpYW8sPvM7qc7Fx25Z6qv\",\"token_decimals\":6},\"amm_info_address\":\"6USpEBbN94DUYLUi4a2wo3AZDCyozon1PLGYu27jzPkX\",\"open_orders_address\":\"GRv7YS45zPAnYUA1VpHGErmauWELJGqLM9xbKADYfw8k\"}"
        }
      ]
    },
    "SELFIE,RAYDIUM,9WPTUKH8FKUCNEPRWOPYLH3AK9GSJPHFDENBQ2X1CZDP/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "SELFIE,RAYDIUM,9WPTUKH8FKUCNEPRWOPYLH3AK9GSJPHFDENBQ2X1CZDP",
          "Quote": "USD"
        },
        "decimals": 11,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"31601\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SOL,RAYDIUM,SO11111111111111111111111111111111111111112/SELFIE,RAYDIUM,9WPTUKH8FKUCNEPRWOPYLH3AK9GSJPHFDENBQ2X1CZDP",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "invert": true,
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"HGDoCdba9yPpKvyYptWv747mG2ti8oVr8Cz88gV9TMdW\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"HmEomvDg2BjV8bvdb1DWH52WEju6KDTpG5CZBBqW2Zgb\",\"token_decimals\":6},\"amm_info_address\":\"Dfk133hHxjAA1yPryNkoPERGJ5DMpUtm79YeY1p1Wiyh\",\"open_orders_address\":\"7xGGsWHaXoPw4mJaJKoUatrQbUSVyy3TjvniWVxTBfbc\"}"
        }
      ]
    },
    "SHDW,RAYDIUM,SHDWYBXIHQICJ6YEKG2GUR7WQKLELAMK1GHZCK9PL6Y/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "SHDW,RAYDIUM,SHDWYBXIHQICJ6YEKG2GUR7WQKLELAMK1GHZCK9PL6Y",
          "Quote": "USD"
        },
        "decimals": 10,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"16868\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SHDW,RAYDIUM,SHDWYBXIHQICJ6YEKG2GUR7WQKLELAMK1GHZCK9PL6Y/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"GVzektnh6TDocY4FTdLm3F5Aha5XBfXmYKadc82BHJEV\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"7naoUqLwTWYpQSDj259mRb7hAqCAmiKG3og1MWYoMqBu\",\"token_decimals\":9},\"amm_info_address\":\"DY9DeyKj9T6yCXf8UM6FGMGeh7arfmmePf9E9jzdaymg\",\"open_orders_address\":\"5gTJDU2uA81Nsf6dWGrmo7Z6dfZ5pkPbr8kdcFJ4BGVV\"}"
        }
      ]
    },
    "SIGMA,RAYDIUM,5SVG3T9CNQSM2KEWZBRQ6HASQH1OGFJQTTLXYUIBPUMP/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "SIGMA,RAYDIUM,5SVG3T9CNQSM2KEWZBRQ6HASQH1OGFJQTTLXYUIBPUMP",
          "Quote": "USD"
        },
        "decimals": 12,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"32498\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SOL,RAYDIUM,SO11111111111111111111111111111111111111112/SIGMA,RAYDIUM,5SVG3T9CNQSM2KEWZBRQ6HASQH1OGFJQTTLXYUIBPUMP",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "invert": true,
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"AKJKbaJnRAjd11yxhsqdhFPbq2MgGwCBGyRUXVFddujn\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"3dHkGkQneQPTwKkKMoUUkQ5qf5mjBkFLgUAzvwe8o2Mu\",\"token_decimals\":6},\"amm_info_address\":\"424kbbJyt6VkSn7GeKT9Vh5yetuTR1sbeyoya2nmBJpw\",\"open_orders_address\":\"H3fZcGRYozcsZ2oz3UpX1ZCwJYQAWN9ESqtEsgk4vEwS\"}"
        }
      ]
    },
    "SILLY,RAYDIUM,7EYNHQOR9YM3N7UOAKROA44UY8JEAZV3QYOUOV87AWMS/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "SILLY,RAYDIUM,7EYNHQOR9YM3N7UOAKROA44UY8JEAZV3QYOUOV87AWMS",
          "Quote": "USD"
        },
        "decimals": 12,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"28789\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SILLY,RAYDIUM,7EYNHQOR9YM3N7UOAKROA44UY8JEAZV3QYOUOV87AWMS/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"E2uenGi8vD5sJaS8wNqBuMRSd7GiL5re8F3Kw1NPEjcq\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"yEvvCAhk9uLsQgnGRMETJgAnQ9abPQfqGwhMWxdFmRi\",\"token_decimals\":9},\"amm_info_address\":\"DsD69qYsFvMX4cBvHbssGneB2aYwECkL3ehYjQ6NH6aq\",\"open_orders_address\":\"C6j2vFirdbE1nCamBQHiTfMVQQXiKzc1kQTUksTYLuHm\"}"
        }
      ]
    },
    "SKBDI,RAYDIUM,DPAQFQ5SFNOQW2SH9WMMMASFL9LNU6RDTDQWE1TAB2TB/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "SKBDI,RAYDIUM,DPAQFQ5SFNOQW2SH9WMMMASFL9LNU6RDTDQWE1TAB2TB",
          "Quote": "USD"
        },
        "decimals": 11,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"32415\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SOL,RAYDIUM,SO11111111111111111111111111111111111111112/SKBDI,RAYDIUM,DPAQFQ5SFNOQW2SH9WMMMASFL9LNU6RDTDQWE1TAB2TB",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "invert": true,
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"Ee1ZQAGEQPe335w65sCXtQkHPawDZZeALyBEoUo9sCDv\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"5ptCqLRuMHsadSLNiSWTfRG4BLmKUCKt4pKif4aiM7u3\",\"token_decimals\":9},\"amm_info_address\":\"BnkDVQKwr7DWx9zmPcqtgCpTfKm7wzp2ydzrE218nrtV\",\"open_orders_address\":\"7CUEckmcRR9Aeu1XR5m7BgQn9AAE3tH3VpCqt94YfgEr\"}"
        }
      ]
    },
    "SKID,RAYDIUM,9X2RHTKRBZW3SLYE9E88CBD1KZ5RFU1F4JTSN4ARH43D/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "SKID,RAYDIUM,9X2RHTKRBZW3SLYE9E88CBD1KZ5RFU1F4JTSN4ARH43D",
          "Quote": "USD"
        },
        "decimals": 11,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"30011\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SKID,RAYDIUM,9X2RHTKRBZW3SLYE9E88CBD1KZ5RFU1F4JTSN4ARH43D/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"94bG7JWrDtnsXCZyuLFNT4iycqtBBNU5a1BQ676ZTbh4\",\"token_decimals\":6},\"quote_token_vault\":{\"token_vault_address\":\"4DMRWND9m4FwjQHupqiC7et1kXBR84h8G8e62pstCvij\",\"token_decimals\":9},\"amm_info_address\":\"43dfuacTABqQsgMaWDTuFTi8r9TL3xbtr99y6ZFQ641V\",\"open_orders_address\":\"GeY7bc5FxcTxRyJG6cKry8mzk4e2WU5rW62xA4DU43vb\"}"
        }
      ]
    },
    "SKL/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "SKL",
          "Quote": "USD"
        },
        "decimals": 11,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"5691\"}]}"
      },
      "provider_configs": [
        {
          "name": "coinbase_ws",
          "off_chain_ticker": "SKL-USD"
        }
      ]
    },
    "SLERF,RAYDIUM,7BGBVYJRZX1YKZ4OH9MJB8ZSCATKKWB8DZFX7LOIVKM3/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "SLERF,RAYDIUM,7BGBVYJRZX1YKZ4OH9MJB8ZSCATKKWB8DZFX7LOIVKM3",
          "Quote": "USD"
        },
        "decimals": 10,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"29920\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SLERF,RAYDIUM,7BGBVYJRZX1YKZ4OH9MJB8ZSCATKKWB8DZFX7LOIVKM3/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"9A2ZsPz5Zg6jKN4o4KRMjTVPmkH51wYWFLmt4KBRy1Rq\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"5Zumc1SYPmQ89nqwXqzogeuhdJ85iEMpSk35A4P87pmD\",\"token_decimals\":9},\"amm_info_address\":\"AgFnRLUScRD2E4nWQxW73hdbSN7eKEUb2jHX7tx9YTYc\",\"open_orders_address\":\"FT5Ptk37g5r6D9BKt3hne8ovHZ1g56oJBvuZRwn3zS3j\"}"
        }
      ]
    },
    "SLIM,RAYDIUM,XXXXA1SKNGWFTW2KFN8XAUW9XQ8HBZ5KVTCSESTT9FW/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "SLIM,RAYDIUM,XXXXA1SKNGWFTW2KFN8XAUW9XQ8HBZ5KVTCSESTT9FW",
          "Quote": "USD"
        },
        "decimals": 10,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"9741\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SLIM,RAYDIUM,XXXXA1SKNGWFTW2KFN8XAUW9XQ8HBZ5KVTCSESTT9FW/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"6FoSD24CM2MyadTwVUqgZQ17kXozfMa3DfusbnuqYduy\",\"token_decimals\":6},\"quote_token_vault\":{\"token_vault_address\":\"EDL73XTnmr56U4ohW5uXXh6LJwsQQdoRLragMYEWLGPn\",\"token_decimals\":9},\"amm_info_address\":\"8idN93ZBpdtMp4672aS4GGMDy7LdVWCCXH7FKFdMw9P4\",\"open_orders_address\":\"E9t69DajWSrPC2acSjPb2EnLhFjXaDzcWsfZkEu5i26i\"}"
        }
      ]
    },
    "SLOTH,RAYDIUM,HQ7DAOIUXZC2K1DR7KXRHCCNTXVEYGNVOUEXTXE8DMBH/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "SLOTH,RAYDIUM,HQ7DAOIUXZC2K1DR7KXRHCCNTXVEYGNVOUEXTXE8DMBH",
          "Quote": "USD"
        },
        "decimals": 12,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"31163\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SLOTH,RAYDIUM,HQ7DAOIUXZC2K1DR7KXRHCCNTXVEYGNVOUEXTXE8DMBH/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"Bs7VsZxQYHndLFnfDRRmJ4D44gCoTv7vNoDF2s5s11cV\",\"token_decimals\":6},\"quote_token_vault\":{\"token_vault_address\":\"5xQzJAvJ7Ut4qoTwiKECnaMDUhZFivx96EFomcBbUShq\",\"token_decimals\":9},\"amm_info_address\":\"7mtJbVNEtejYmCLRriwQhymZdzn4wGRFTvTZ5721b4BD\",\"open_orders_address\":\"A7k1mZQNNNKCakhHZN9bQqLzowDmApHTb4564uw5tAVU\"}"
        }
      ]
    },
    "SMILEK,RAYDIUM,7X4FGIFFEQZS1TIUNVJZPT47GTLXAJ8JFN8G1HYYU6JH/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "SMILEK,RAYDIUM,7X4FGIFFEQZS1TIUNVJZPT47GTLXAJ8JFN8G1HYYU6JH",
          "Quote": "USD"
        },
        "decimals": 15,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"10186\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SMILEK,RAYDIUM,7X4FGIFFEQZS1TIUNVJZPT47GTLXAJ8JFN8G1HYYU6JH/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"9Lcn4RxwEMNospvyLEBKricos91iosTfd2VsATRDGD8e\",\"token_decimals\":6},\"quote_token_vault\":{\"token_vault_address\":\"3xzYaYaCtqiEB3YkFo6Hz2S7csR9Rc814UELH15zqhew\",\"token_decimals\":9},\"amm_info_address\":\"DwwE6q9rPHKEoooMjv99hVUvA4CPUZJhws4CqrrBD7XT\",\"open_orders_address\":\"ETgKCCQUbVhTL8mgK6zHptHDk5JJ3P8NpTTSK3wUv3VM\"}"
        }
      ]
    },
    "SMOG,RAYDIUM,FS66V5XYTJAFO14LIPZ5HT93EUMAHMYIPCFQHLPU4SS8/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "SMOG,RAYDIUM,FS66V5XYTJAFO14LIPZ5HT93EUMAHMYIPCFQHLPU4SS8",
          "Quote": "USD"
        },
        "decimals": 11,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"29358\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SMOG,RAYDIUM,FS66V5XYTJAFO14LIPZ5HT93EUMAHMYIPCFQHLPU4SS8/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"5QM7Ec3tMo3Rh6xSeyxC46RMewDjVAeuCwweM7XbftVH\",\"token_decimals\":6},\"quote_token_vault\":{\"token_vault_address\":\"2Chre6WpZ2zXVAX9yaxFHmGK6Fbr5yRuRGrMWYgpProu\",\"token_decimals\":9},\"amm_info_address\":\"G6oNRyjP2WAVDUTYFPZsgBBxxMetyQEi2AYbmzRNe7Mc\",\"open_orders_address\":\"3STsVkoPDkGpbHhjRkQQPYBDTiDUQn5Ak6gmYvdcZazo\"}"
        }
      ]
    },
    "SMOLE,RAYDIUM,9TTYEZ3XIRUYJ6CQAR495HBBKJU6SUWDV6AMQ9MVBYYS/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "SMOLE,RAYDIUM,9TTYEZ3XIRUYJ6CQAR495HBBKJU6SUWDV6AMQ9MVBYYS",
          "Quote": "USD"
        },
        "decimals": 14,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"30049\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SMOLE,RAYDIUM,9TTYEZ3XIRUYJ6CQAR495HBBKJU6SUWDV6AMQ9MVBYYS/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"VDZ9kwvKRbqhNdsoRZyLVzAAQMbGY9akHbtM6YugViS\",\"token_decimals\":6},\"quote_token_vault\":{\"token_vault_address\":\"HiLcngHP5y1Jno53tuuNeFHKWhyyZp3XuxtKPszD6rG2\",\"token_decimals\":9},\"amm_info_address\":\"5EgCcjkuE42YyTZY4QG8qTioUwNh6agTvJuNRyEqcqV1\",\"open_orders_address\":\"FeKBjZ5rBvHPyppHf11qjYxwaQuiympppCTQ5pC6om3F\"}"
        }
      ]
    },
    "SNIBBU,RAYDIUM,CRABDNXPBBJTFPKZMCNWT667PUEQCEQSUWYAFNM8JSX8/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "SNIBBU,RAYDIUM,CRABDNXPBBJTFPKZMCNWT667PUEQCEQSUWYAFNM8JSX8",
          "Quote": "USD"
        },
        "decimals": 16,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"32490\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SNIBBU,RAYDIUM,CRABDNXPBBJTFPKZMCNWT667PUEQCEQSUWYAFNM8JSX8/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"5ZLJXkRWVazNr88jiXNys9SNwBmHUdisZrLXioXuJbip\",\"token_decimals\":6},\"quote_token_vault\":{\"token_vault_address\":\"CcWHNztemPCNNS84yk3dXrHcUHmXfabuL3uokdvxL5kk\",\"token_decimals\":9},\"amm_info_address\":\"Cj7Cq1aTWn8UPmA8jo37MkrHD44n4KCRT5CDEENs3eQ5\",\"open_orders_address\":\"8syHS4CPVaNrhGKBHGAzDY3zuezKxA2MSa85wmQVSJ5P\"}"
        }
      ]
    },
    "SOL/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "SOL",
          "Quote": "USD"
        },
        "decimals": 7,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"5426\"}]}"
      },
      "provider_configs": [
        {
          "name": "coinbase_ws",
          "off_chain_ticker": "SOL-USD"
        },
        {
          "name": "bitfinex_ws",
          "off_chain_ticker": "SOLUSD"
        },
        {
          "name": "crypto_dot_com_ws",
          "off_chain_ticker": "SOL_USD"
        }
      ]
    },
    "SOLAMA,RAYDIUM,AVLHAHDCDQ4M4VHM4UG63OH7XC8JTK49DM5HOE9SAZQR/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "SOLAMA,RAYDIUM,AVLHAHDCDQ4M4VHM4UG63OH7XC8JTK49DM5HOE9SAZQR",
          "Quote": "USD"
        },
        "decimals": 11,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"29015\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SOLAMA,RAYDIUM,AVLHAHDCDQ4M4VHM4UG63OH7XC8JTK49DM5HOE9SAZQR/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"B8QnKAywAQNKwJxkvoyzg3W4Z3dSdVkB6AxYUbL9LohY\",\"token_decimals\":6},\"quote_token_vault\":{\"token_vault_address\":\"4ywvnnjj5aN55f8QX8JJCU9SPe29L8xSSv9EWUgnJWW5\",\"token_decimals\":9},\"amm_info_address\":\"CR8FJB9jqGvtNYnYUuxyyS41WXue2HZRfhNavkaV8CS4\",\"open_orders_address\":\"7yC1jm2zvg85MbzWY62jCh52TmWPjJdxKrq8BgXQ7YFS\"}"
        }
      ]
    },
    "SOLCAT,RAYDIUM,E99FN4TCRB1TQPHXK1DU7PRXJI6HMZXETYPNJRO19FWZ/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "SOLCAT,RAYDIUM,E99FN4TCRB1TQPHXK1DU7PRXJI6HMZXETYPNJRO19FWZ",
          "Quote": "USD"
        },
        "decimals": 11,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"31561\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SOLCAT,RAYDIUM,E99FN4TCRB1TQPHXK1DU7PRXJI6HMZXETYPNJRO19FWZ/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"Hme9rfTdYXqPq8MSbw97tjg19DjniAjZGiCUc7h5FAe9\",\"token_decimals\":6},\"quote_token_vault\":{\"token_vault_address\":\"EVBXVefxmGNZ88kQAqriow74DRDXCJ3FK2yikHQHuEF\",\"token_decimals\":9},\"amm_info_address\":\"7DkcZR6RCFFxqasXXwB4P4Z31gGTAedQ7wgQY4dLnHo4\",\"open_orders_address\":\"FMikdD62szyWDsfQjRhZHS6V7x4jqprX1HAePrA4diTA\"}"
        }
      ]
    },
    "SOLCEX,RAYDIUM,AMJZRN1TBQWQFNAJHFEBB7UGBBQBJB7FZXANGGDFPK6K/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "SOLCEX,RAYDIUM,AMJZRN1TBQWQFNAJHFEBB7UGBBQBJB7FZXANGGDFPK6K",
          "Quote": "USD"
        },
        "decimals": 12,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"30750\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SOLCEX,RAYDIUM,AMJZRN1TBQWQFNAJHFEBB7UGBBQBJB7FZXANGGDFPK6K/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"EjRCAHV82cbRXg1vaE1Xz3qW6yU71WxjnQuRbKLyVdDJ\",\"token_decimals\":6},\"quote_token_vault\":{\"token_vault_address\":\"4qKwzJTRvG4EThyC8xKjUPqN5Hkg87SfMBpkv4RpHbeL\",\"token_decimals\":9},\"amm_info_address\":\"4Ro3pG1XZgSJENfgCccNgQqrHYVqhHjwcL27oHmXMMTG\",\"open_orders_address\":\"GspXWvJuNDZ6dZumQVzm1zfiQ3wXbvYfygDcHjMVoK2Y\"}"
        }
      ]
    },
    "SOLS,RAYDIUM,2WME8EVKW8QSFSK2B3QEX4S64AC6WXHPXB3GRDCKEKIO/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "SOLS,RAYDIUM,2WME8EVKW8QSFSK2B3QEX4S64AC6WXHPXB3GRDCKEKIO",
          "Quote": "USD"
        },
        "decimals": 10,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"28719\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SOL,RAYDIUM,SO11111111111111111111111111111111111111112/SOLS,RAYDIUM,2WME8EVKW8QSFSK2B3QEX4S64AC6WXHPXB3GRDCKEKIO",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "invert": true,
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"GVCDnBBt7rkGzxqpF9HHubJvBje4TFoLjR56HyHbATgZ\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"5XzSGEpth3SyqC9h3oBap45gfuocjyceub18qbPLBLZ2\",\"token_decimals\":9},\"amm_info_address\":\"8XBzPr3TxvfcEu9HVa6ix3wXXbY8s44YZScKSanJyjzv\",\"open_orders_address\":\"3pbDHAo4Hhk51vpgx3Wfq1hTqrRC6DCG3GnqP8t8FXjP\"}"
        }
      ]
    },
    "SOLZILLA,RAYDIUM,31IQSAHFA4CMIIRU7REYGBZUAWG4R4AH7Y4ADU9ZFXJP/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "SOLZILLA,RAYDIUM,31IQSAHFA4CMIIRU7REYGBZUAWG4R4AH7Y4ADU9ZFXJP",
          "Quote": "USD"
        },
        "decimals": 19,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"28866\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SOLZILLA,RAYDIUM,31IQSAHFA4CMIIRU7REYGBZUAWG4R4AH7Y4ADU9ZFXJP/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"G5AvMhBK7ChXqXJga2GaLeypPENEhrbJUxAkNeHRGAku\",\"token_decimals\":4},\"quote_token_vault\":{\"token_vault_address\":\"3xFiovB9NSxx2DgDGrwm7yAEHYCVi8WCeuwmz4jnY9z2\",\"token_decimals\":9},\"amm_info_address\":\"CLmZVV7miLweVzric8jkCKNbyiHKqigW88mFZsEZjh1n\",\"open_orders_address\":\"2e5stgoJLN3NFPKB6RC5CZ61rKntepvKZcQi7xQ7TH8a\"}"
        }
      ]
    },
    "SOY,RAYDIUM,4G3KNXWAA2UQHDPAQTJWQM1SREXCUD7LKT14V2OES7RV/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "SOY,RAYDIUM,4G3KNXWAA2UQHDPAQTJWQM1SREXCUD7LKT14V2OES7RV",
          "Quote": "USD"
        },
        "decimals": 12,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"32350\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SOL,RAYDIUM,SO11111111111111111111111111111111111111112/SOY,RAYDIUM,4G3KNXWAA2UQHDPAQTJWQM1SREXCUD7LKT14V2OES7RV",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "invert": true,
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"AK2b7BqQgSzHESy4vTW5qQwNBpGYN83EupSrffLWVamU\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"6QZy9Ur76mchXctFFrNCg5HJZAmXN1Np4Nzm3NpSYZ4o\",\"token_decimals\":6},\"amm_info_address\":\"DtTkLBvYUaYBZ7PC4vCwWfu56Zkgbf7ycEXxLhAP7Xx8\",\"open_orders_address\":\"HfU3BU3KFK3iKf4b5HgBrYfWCe64KeK2K8SRYdyV9qED\"}"
        }
      ]
    },
    "SPIKE,RAYDIUM,SPIKEYAQOAGYYBANPXRO8NLSYLU93SR56N352JJRLN5/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "SPIKE,RAYDIUM,SPIKEYAQOAGYYBANPXRO8NLSYLU93SR56N352JJRLN5",
          "Quote": "USD"
        },
        "decimals": 14,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"32042\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SPIKE,RAYDIUM,SPIKEYAQOAGYYBANPXRO8NLSYLU93SR56N352JJRLN5/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"6eQAYsvmNXUdapqDDNsGusCXqsipFRjVEsyC2MTZpmC6\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"8iiXNrUcDzNPMX8K9Am8aN8xNpuDtyLZGzQj5XVwBJuB\",\"token_decimals\":9},\"amm_info_address\":\"2u7dCQP8J5EBTVGSzfJxLSAmmjGM2rMpWE45g98HgJPH\",\"open_orders_address\":\"CWL88Hpredsoc15vdU3ZrBnPQsq9Peasv1gx4ihRqV4f\"}"
        }
      ]
    },
    "SPWN,RAYDIUM,5U9QQCPHQXAJCEV9UYZFJD5ZHN93VUPK1ANNKXNUFPNT/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "SPWN,RAYDIUM,5U9QQCPHQXAJCEV9UYZFJD5ZHN93VUPK1ANNKXNUFPNT",
          "Quote": "USD"
        },
        "decimals": 15,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"10285\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SPWN,RAYDIUM,5U9QQCPHQXAJCEV9UYZFJD5ZHN93VUPK1ANNKXNUFPNT/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"2teNjejMpdr2nbZ8qyjcmdKDNSwrXEWmgzSineUnTtcZ\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"3RpmEhmXHZ81SBWQEPYZTJxAmxgjkhw6inF7aaG1sfPs\",\"token_decimals\":9},\"amm_info_address\":\"HMhKqQ1fVXEhPdXeDHUWpNa7ipWqYxdRmSbkRBeggFwX\",\"open_orders_address\":\"GEf8z7JMDZse8nDSQchQ55DLe7Q2mnXMYySdWYmcd1wZ\"}"
        }
      ]
    },
    "SPX,RAYDIUM,J3NKXXXZCNNIMJKW9HYB2K4LUXGWB6T1FTPTQVSV3KFR/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "SPX,RAYDIUM,J3NKXXXZCNNIMJKW9HYB2K4LUXGWB6T1FTPTQVSV3KFR",
          "Quote": "USD"
        },
        "decimals": 12,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"28081\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SPX,RAYDIUM,J3NKXXXZCNNIMJKW9HYB2K4LUXGWB6T1FTPTQVSV3KFR/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"rFkn84fuSQ3ovZECNLF7JnX5xy1to6KR7LM6TSQtKgT\",\"token_decimals\":8},\"quote_token_vault\":{\"token_vault_address\":\"DGUVmREQ4hpRmkf9sRP4Fb48az5ACCLoxH4FMrvEPrZN\",\"token_decimals\":9},\"amm_info_address\":\"9t1H1uDJ558iMPNkEPSN1fqkpC4XSPQ6cqSf6uEsTfTR\",\"open_orders_address\":\"AxTb8hkmnVcAtSJFY96dPhD1biYS7Q1vS5AjZGJPNP4P\"}"
        }
      ]
    },
    "SRM,RAYDIUM,SRMUAPVNDXXOKK5GT7XD5CUUGXMBCOAZ2LHEUAOKWRT/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "SRM,RAYDIUM,SRMUAPVNDXXOKK5GT7XD5CUUGXMBCOAZ2LHEUAOKWRT",
          "Quote": "USD"
        },
        "decimals": 11,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"6187\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SRM,RAYDIUM,SRMUAPVNDXXOKK5GT7XD5CUUGXMBCOAZ2LHEUAOKWRT/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"BCNYwsnz3yXvi4mY5e9w2RmZvwUW3pefzYQ4tsoNdDhp\",\"token_decimals\":6},\"quote_token_vault\":{\"token_vault_address\":\"7BXPSUXeBVqJGyxW3yvkNxnJjYHuC8mnhyFCDp2abAs6\",\"token_decimals\":9},\"amm_info_address\":\"EvWJC2mnmu9C9aQrsJLXw8FhUcwBzFEUQsP1E5Y6a5N7\",\"open_orders_address\":\"9ot4bg8aT2FRKfiRrM2fSPHEr7M1ihBqm1iT4771McqR\"}"
        }
      ]
    },
    "STASH,RAYDIUM,EWMFSJGDCE7CXDAYZ3HBCAA7NSFHTNDDYSXX3SHCO2HS/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "STASH,RAYDIUM,EWMFSJGDCE7CXDAYZ3HBCAA7NSFHTNDDYSXX3SHCO2HS",
          "Quote": "USD"
        },
        "decimals": 13,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"31420\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "STASH,RAYDIUM,EWMFSJGDCE7CXDAYZ3HBCAA7NSFHTNDDYSXX3SHCO2HS/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"CLF3xQkVM2JxP9vZQ1kXgykVDV4fw1gvPcn6mz3sxDyN\",\"token_decimals\":6},\"quote_token_vault\":{\"token_vault_address\":\"FqvQz7X2ePNsAfNRxZfuGPZ2kANJJNou7BLwmACL6qBh\",\"token_decimals\":9},\"amm_info_address\":\"huZfRPDD1Tt72rjkVK9NKcPEbGgAfVKVdP2sDB8hVPF\",\"open_orders_address\":\"4ppw9uZ9CwRmGaFtQtvK3c8iAC3QqGoHYYmcZAvVpVu9\"}"
        }
      ]
    },
    "STEP,RAYDIUM,STEPASCQOEIOFXXWGNH2SLBDFP9D8RVKZ2YP39IDPYT/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "STEP,RAYDIUM,STEPASCQOEIOFXXWGNH2SLBDFP9D8RVKZ2YP39IDPYT",
          "Quote": "USD"
        },
        "decimals": 11,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"9443\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "STEP,RAYDIUM,STEPASCQOEIOFXXWGNH2SLBDFP9D8RVKZ2YP39IDPYT/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"74cQrLwSHqWLe9cbtZsq9Lyc1hoeLkXjzXjpGitijqKX\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"Hr7yHnCCWH7Rz5FDc5rKMRga5v4zLvn4r1s9ZppQfM4H\",\"token_decimals\":9},\"amm_info_address\":\"DN1Rqx5AE5jHV8pTPiwUcSYVAK15YrLGkfVdU8GxhWn1\",\"open_orders_address\":\"2bzaNYPMZXAXZRTKMPicB76wwfiuZrvxo9qbXvzgg4Mh\"}"
        }
      ]
    },
    "STOG,RAYDIUM,AHNZ7VYYQ5JHXBITQL8TUN7CIGG66EVCNU7EKOKX99FZ/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "STOG,RAYDIUM,AHNZ7VYYQ5JHXBITQL8TUN7CIGG66EVCNU7EKOKX99FZ",
          "Quote": "USD"
        },
        "decimals": 13,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"30993\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "STOG,RAYDIUM,AHNZ7VYYQ5JHXBITQL8TUN7CIGG66EVCNU7EKOKX99FZ/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"FjMnZcNN1P8M1VaVy9gZSj25h6yLCNcTWWUWACYfPNNm\",\"token_decimals\":6},\"quote_token_vault\":{\"token_vault_address\":\"7mNgEj1gkpAC6ySzPWdS2qK1TeJz1wG4wYkekaHzc9j5\",\"token_decimals\":9},\"amm_info_address\":\"5uzT4p6GRm78HkhXvdsSR712188yN2kFzRQSnDqCE2JJ\",\"open_orders_address\":\"3NfUDFqRQW2gHv3q9suHSDHoApTRmK6stdhqGaS1Ebi2\"}"
        }
      ]
    },
    "SYP,RAYDIUM,FNKE9N6AGJQONWRBZXY4RW6LZVAO7QWBONUBID7EDUMZ/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "SYP,RAYDIUM,FNKE9N6AGJQONWRBZXY4RW6LZVAO7QWBONUBID7EDUMZ",
          "Quote": "USD"
        },
        "decimals": 13,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"12042\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SYP,RAYDIUM,FNKE9N6AGJQONWRBZXY4RW6LZVAO7QWBONUBID7EDUMZ/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"4iuHfu5rPzdsnjBEPAdGvnK3brF3JiqpwtXerko1o6U4\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"5FvQrUmnCN4o1HBsA3XqbCDPypvyroJ9MBSYH5goxFGC\",\"token_decimals\":9},\"amm_info_address\":\"D95EzH4ZsGLikvYzp7kmz1RM1xNMo1MXXiXaedQesA2m\",\"open_orders_address\":\"34Ggyj2dNyQUWDaGUaMKVvyQDoTHEupD4o2m1mPFaPVf\"}"
        }
      ]
    },
    "TAI,RAYDIUM,HAX9LTGSQKZE1YFYCHNBLTFH8GYBQKTKFWKKG2SP6GDD/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "TAI,RAYDIUM,HAX9LTGSQKZE1YFYCHNBLTFH8GYBQKTKFWKKG2SP6GDD",
          "Quote": "USD"
        },
        "decimals": 10,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"20605\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "TAI,RAYDIUM,HAX9LTGSQKZE1YFYCHNBLTFH8GYBQKTKFWKKG2SP6GDD/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"9CUyaxiToHD93rcY5wYT76YchLXEivDxV3mefWU2kAxi\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"6yB7tBvMLAPTXRNUMD3YTZSxTaGDiBv7m9r8JtPHbDe7\",\"token_decimals\":9},\"amm_info_address\":\"99a2qtfLAxv9LVhKesjEawjKy8Pwb4cypC5ccuCB4VUS\",\"open_orders_address\":\"CBiqcrAAzWXFdCkjL4K4GCGJ5rYgkQJLkCv5yCd696F2\"}"
        }
      ]
    },
    "TIME,RAYDIUM,ED5WBEYAYTLM4WRGNOHPXJEWNIAIKEFIOVMJYZH6K31M/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "TIME,RAYDIUM,ED5WBEYAYTLM4WRGNOHPXJEWNIAIKEFIOVMJYZH6K31M",
          "Quote": "USD"
        },
        "decimals": 15,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"32379\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "TIME,RAYDIUM,ED5WBEYAYTLM4WRGNOHPXJEWNIAIKEFIOVMJYZH6K31M/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"5ckdqMkyg1QqE4L1Tm2xcNJYqAmNREZtg2Rk319P2pWk\",\"token_decimals\":6},\"quote_token_vault\":{\"token_vault_address\":\"Hi8KQULtewanECvY991NCHP8kpR75ufwjBqDVTBDyRQv\",\"token_decimals\":9},\"amm_info_address\":\"qqb54Ljd2VkWkRpgu5aZTGfpvKwhnfZ417vucZncTfP\",\"open_orders_address\":\"EcAFQ7AikduNoGG6f5hBKSADZc3av5VTJUndDKxc9jNr\"}"
        }
      ]
    },
    "TMANIA,RAYDIUM,HUPSPKKI5QDNF5WAU7JTETHKEMHNI6XQ23TUNZRKZWUI/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "TMANIA,RAYDIUM,HUPSPKKI5QDNF5WAU7JTETHKEMHNI6XQ23TUNZRKZWUI",
          "Quote": "USD"
        },
        "decimals": 13,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"31627\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "TMANIA,RAYDIUM,HUPSPKKI5QDNF5WAU7JTETHKEMHNI6XQ23TUNZRKZWUI/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"DkxCSgpmvm7c1TjiwGA7bECenVfai4mQQgC67gBUbgdG\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"2jpRdiZRzNfsoEoDe88quqRuSNrXb8AMeby9mnjVYbZT\",\"token_decimals\":9},\"amm_info_address\":\"HhbGicRCbxdANLGTPdtHCXWFVX2wMTZsNEdKuZUstj4\",\"open_orders_address\":\"CgJiJNA4WpBdCs7rfV7RH3x2bVchaSeJXGntWnkYwYCE\"}"
        }
      ]
    },
    "TNSR,RAYDIUM,TNSRXCUXOT9XBG3DE7PIJYTDYU7KSKLQCPDDXNEJAS6/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "TNSR,RAYDIUM,TNSRXCUXOT9XBG3DE7PIJYTDYU7KSKLQCPDDXNEJAS6",
          "Quote": "USD"
        },
        "decimals": 10,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"30449\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "TNSR,RAYDIUM,TNSRXCUXOT9XBG3DE7PIJYTDYU7KSKLQCPDDXNEJAS6/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"6KG3pRkYRAiYhdPThU1qpTBdgMW9gPc7v7k5Pp5V7q7z\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"Czuw75LXk9XEwRWjDQ6AH64fJSkjhta4THtXSWSMGXXr\",\"token_decimals\":9},\"amm_info_address\":\"H5xso6syEjUSAhfZ5WsTyaDmkE3MwUmWcjSLSdMKREd\",\"open_orders_address\":\"E1uACjdhVR4q1BZ8yYiDfB4c5uXjd36soNeRnwu66o8S\"}"
        }
      ]
    },
    "TOOKER,RAYDIUM,9EYSCPIYSGNEIMNQPZAZR7JN9GVFXFYZGTEJ85HV9L6U/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "TOOKER,RAYDIUM,9EYSCPIYSGNEIMNQPZAZR7JN9GVFXFYZGTEJ85HV9L6U",
          "Quote": "USD"
        },
        "decimals": 12,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"30959\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "TOOKER,RAYDIUM,9EYSCPIYSGNEIMNQPZAZR7JN9GVFXFYZGTEJ85HV9L6U/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"Cki9WdL3sCoNY3cLmfG4iqSbvB8g1Fr9tw8qa5tP1m3Y\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"2vTTh5pGbzc6guAJmt78XnTcXVBEZEWmGBkXkSNZwN59\",\"token_decimals\":9},\"amm_info_address\":\"3vGHsKVKNapB4hSapzKNwtiJ6DA8Ytd9SsMFSoAk154B\",\"open_orders_address\":\"5dzcxMHjuNU5LZyEXBhoWWKuxw51Z3626TTf2FTfLJjb\"}"
        }
      ]
    },
    "TORSY,RAYDIUM,5YQCKGEKWHJMP9LW5AUF2UJRDUBMJAHCYNGJA8M7EBW8/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "TORSY,RAYDIUM,5YQCKGEKWHJMP9LW5AUF2UJRDUBMJAHCYNGJA8M7EBW8",
          "Quote": "USD"
        },
        "decimals": 12,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"32288\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "TORSY,RAYDIUM,5YQCKGEKWHJMP9LW5AUF2UJRDUBMJAHCYNGJA8M7EBW8/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"82L6PrJA2MZfam2JeqGfen1uyDAsbrFva7WtTb5bDCUR\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"28L7P5HT1MS8FXrddTscuCWjdDLatdbKujtP18t5cZtK\",\"token_decimals\":9},\"amm_info_address\":\"7isNHVnuAfAjKWkUEPwgcsV1LZqu9iFG48UyL24iHYAa\",\"open_orders_address\":\"GKPaS9ACddzoG16YtQPdAVK48GztTi5uYmaTtWaDPTfm\"}"
        }
      ]
    },
    "TREMP,RAYDIUM,FU1Q8VJPZNURMQSCISJP8BAKKIDGSLMOUB8CBDF8TKQV/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "TREMP,RAYDIUM,FU1Q8VJPZNURMQSCISJP8BAKKIDGSLMOUB8CBDF8TKQV",
          "Quote": "USD"
        },
        "decimals": 10,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"29717\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "TREMP,RAYDIUM,FU1Q8VJPZNURMQSCISJP8BAKKIDGSLMOUB8CBDF8TKQV/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"G2XNC6Rt2G7JZQWhqpJriYwZyxd2L52KSDbDNBCYCpvx\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"9DfnSR9h3hrmgy5pjqBP3SrVQRWPfjSqZZBrNNYGoyaN\",\"token_decimals\":9},\"amm_info_address\":\"5o9kGvozArYNWfbYTZD1WDRkPkkDr6LdpQbUUqM57nFJ\",\"open_orders_address\":\"kTgLvRcrvhxJy9KZFureP8fU5L11BzFrRvUEUa1joai\"}"
        }
      ]
    },
    "TWT,RAYDIUM,HZNPQL7RT9GXF9EWOWSWZC5DFJZQ41XTQGEA7P3VZAAD/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "TWT,RAYDIUM,HZNPQL7RT9GXF9EWOWSWZC5DFJZQ41XTQGEA7P3VZAAD",
          "Quote": "USD"
        },
        "decimals": 10,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"5964\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "TWT,RAYDIUM,HZNPQL7RT9GXF9EWOWSWZC5DFJZQ41XTQGEA7P3VZAAD/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"HEzhdKDVyeSqtWPsscZuVCHefAe7azbjVioxAphbVbrG\",\"token_decimals\":8},\"quote_token_vault\":{\"token_vault_address\":\"EK2Ng32TNoMJ8SDNrK4FHsPVsVxr1RbGetJN1NMdDhbo\",\"token_decimals\":9},\"amm_info_address\":\"GNtoBHpQv5Apyi8TmR9cB2KAmPvMWkTkZgiYyp7tDa8H\",\"open_orders_address\":\"A4nNXSYGiMKA9zU99xwJV8sct5YQWAk8tSFWJRXkVVoy\"}"
        }
      ]
    },
    "UFO,RAYDIUM,99MYNE1MVNSSV2KAKZZDFNNWF5XB4SASJNACRXDBPUMP/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "UFO,RAYDIUM,99MYNE1MVNSSV2KAKZZDFNNWF5XB4SASJNACRXDBPUMP",
          "Quote": "USD"
        },
        "decimals": 13,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"32893\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SOL,RAYDIUM,SO11111111111111111111111111111111111111112/UFO,RAYDIUM,99MYNE1MVNSSV2KAKZZDFNNWF5XB4SASJNACRXDBPUMP",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "invert": true,
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"AoGtrDoEXNPGGXCxgQEkmPeEzyNPPj1gAqsVKaCixhpy\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"AU9tmB914x8cjhYZeaZLPLLZEKjTfQBKY1QvQUvjdDoP\",\"token_decimals\":6},\"amm_info_address\":\"8b73ueUkSBWHMFqUsmPDvDj8UPtBBotU6v52kuh7BgTA\",\"open_orders_address\":\"6JeiDpcxnJedgWVEeqYZqBSZZdgMPovVCFNZ8pV25Dwc\"}"
        }
      ]
    },
    "UPDOG,RAYDIUM,HJ39RRZ6YS22KDB3USXDGNSL7RKIQMSC3YL8AS3SUUKU/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "UPDOG,RAYDIUM,HJ39RRZ6YS22KDB3USXDGNSL7RKIQMSC3YL8AS3SUUKU",
          "Quote": "USD"
        },
        "decimals": 13,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"29672\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "UPDOG,RAYDIUM,HJ39RRZ6YS22KDB3USXDGNSL7RKIQMSC3YL8AS3SUUKU/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"GCHyJ7iVV5rC4yQt89Nq1yfocQiNAri3g8WDvpC85bsi\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"Fs1CAE1pMo6G2235y3dAh167iAGwnAAcHyv6sbCqCJLX\",\"token_decimals\":9},\"amm_info_address\":\"CyscPWWs6G9w9Au1Bg34DxfHE4PsyDQtdMmFeZC5pQYf\",\"open_orders_address\":\"4R2PpxyHoYmC8WSmJjy8VwAPQiPQnL6Rmw4bpiVcSkyG\"}"
        }
      ]
    },
    "USA,RAYDIUM,69KDRLYP5DTRKPHRAASZAQBWMAWZF9GUKJZFZMXZCBAS/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "USA,RAYDIUM,69KDRLYP5DTRKPHRAASZAQBWMAWZF9GUKJZFZMXZCBAS",
          "Quote": "USD"
        },
        "decimals": 16,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"32110\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "USA,RAYDIUM,69KDRLYP5DTRKPHRAASZAQBWMAWZF9GUKJZFZMXZCBAS/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"CfhV9UTxrxxMYSomtQvy3PPYQUssWsvhG9gKdAqg9HoV\",\"token_decimals\":6},\"quote_token_vault\":{\"token_vault_address\":\"GpnqjVZtgt6t6yWmm1YaCc5iWhk56HZ7mpeaq5ZpKZop\",\"token_decimals\":9},\"amm_info_address\":\"HKprCtGbnh1j8xeQggzWhhVd3kwDUdphqPqDP8vMay8b\",\"open_orders_address\":\"F99Y9rg1o2wVanVcTSAqbfH2RykyU25eFqdVHi1PECzi\"}"
        }
      ]
    },
    "USDC/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "USDC",
          "Quote": "USD"
        },
        "decimals": 10,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"3408\"}]}"
      },
      "provider_configs": [
        {
          "name": "bitstamp_api",
          "off_chain_ticker": "usdcusd"
        },
        {
          "name": "coinbase_ws",
          "off_chain_ticker": "USDT-USDC",
          "normalize_by_pair": {
            "Base": "USDT",
            "Quote": "USD"
          },
          "invert": true
        },
        {
          "name": "raydium_api",
          "off_chain_ticker": "USDC,RAYDIUM,EPJFWDD5AUFQSSQEM2QN1XZYBAPC8G4WEGGKZWYTDT1V/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"HjUib8gsdfqbpbrqWiLR1MqQs7PBvcjsSQ68EbdjWa8w\",\"token_decimals\":6},\"quote_token_vault\":{\"token_vault_address\":\"3S4yEXag74xrv7pJBA2yV1G1SNyrCLYuNqWNbtQ4xF3h\",\"token_decimals\":9},\"amm_info_address\":\"AbbG2aR8iNhy2prC32iDRW7pKJjzqhUtri8rV5HboHUY\",\"open_orders_address\":\"6zfCCqa3DhCDb4MJCbGH2J7U2CwXizav8xfywDYxN8XU\"}"
        }
      ]
    },
    "USDT/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "USDT",
          "Quote": "USD"
        },
        "decimals": 9,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"825\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "USDT,RAYDIUM,ES9VMFRZACERMJFRF4H2FYD4KCONKY11MCCE8BENWNYB/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"9vzKGzPHujLUZxA41ChjzPo8DUi76c2g16HxUFPxAKiT\",\"token_decimals\":6},\"quote_token_vault\":{\"token_vault_address\":\"4QfN42eMJgiPXWC9vtmqAz2T83ViqgXMsnuiQCuG1XdM\",\"token_decimals\":9},\"amm_info_address\":\"74bfmNqNe89JSMn7uwnBmuZ6k9J2j6CSuqC6q9F5hU7e\",\"open_orders_address\":\"C6AyY81ZWJFZvUnZgau9NEd4q1EN4qATRqyqhABhG24H\"}"
        },
        {
          "name": "crypto_dot_com_ws",
          "off_chain_ticker": "USDT_USD"
        },
        {
          "name": "coinbase_ws",
          "off_chain_ticker": "USDT-USD"
        },
        {
          "name": "bitfinex_ws",
          "off_chain_ticker": "USTUSD"
        },
        {
          "name": "bitstamp_api",
          "off_chain_ticker": "usdtusd"
        }
      ]
    },
    "W,RAYDIUM,85VBFQZC9TZKFAPTBWJVUW7YBZJY52A6MJTPGJSTQAMQ/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "W,RAYDIUM,85VBFQZC9TZKFAPTBWJVUW7YBZJY52A6MJTPGJSTQAMQ",
          "Quote": "USD"
        },
        "decimals": 10,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"29587\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "W,RAYDIUM,85VBFQZC9TZKFAPTBWJVUW7YBZJY52A6MJTPGJSTQAMQ/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"5wNvrrBgKHf5UjiJ7mqiTAd2Zx4i6b1MhVgMcCLH2xww\",\"token_decimals\":6},\"quote_token_vault\":{\"token_vault_address\":\"J3zbbWpAhfoQLh246MjuoxmNwgV5o9XTtPhte9p6JJEW\",\"token_decimals\":9},\"amm_info_address\":\"BT1Zt5Y8dzjTyTgPo5FTnVdWutTQrTu4qBMXDrxXkBN9\",\"open_orders_address\":\"CKxrATiNT4RPeMLvEQNa43nzCtwEhuvBLMnhCzoiLVwU\"}"
        }
      ]
    },
    "WAFFLES,RAYDIUM,8DOS8NZMGVZEAACXALKBK5FZTW4UUORP4YT8NEAXFDMB/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "WAFFLES,RAYDIUM,8DOS8NZMGVZEAACXALKBK5FZTW4UUORP4YT8NEAXFDMB",
          "Quote": "USD"
        },
        "decimals": 11,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"31442\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SOL,RAYDIUM,SO11111111111111111111111111111111111111112/WAFFLES,RAYDIUM,8DOS8NZMGVZEAACXALKBK5FZTW4UUORP4YT8NEAXFDMB",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "invert": true,
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"8XxDfCFYb1mp5jfESDEbNGvFR8B7Y8Ujoe3E3B6bj282\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"DtbPxH9DbpEdY8GmrdhNzWt4Yu4GrLb8RMkpsJgiSZuB\",\"token_decimals\":6},\"amm_info_address\":\"FJ6MdHqFwmnzx2g19s6X8NDbF7gZCnU2yE1rKd9vbnwf\",\"open_orders_address\":\"HLWUm5kz5fX6o2KWQawiwPDBkj2NzBdJDGp2rhi4yHKb\"}"
        }
      ]
    },
    "WALTER,RAYDIUM,FV56CMR7FHEYPKYMKFMVIKV48UPO51TI9KAXSSQQTDLU/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "WALTER,RAYDIUM,FV56CMR7FHEYPKYMKFMVIKV48UPO51TI9KAXSSQQTDLU",
          "Quote": "USD"
        },
        "decimals": 12,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"31537\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SOL,RAYDIUM,SO11111111111111111111111111111111111111112/WALTER,RAYDIUM,FV56CMR7FHEYPKYMKFMVIKV48UPO51TI9KAXSSQQTDLU",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "invert": true,
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"AiCFCoVvhXJ62bsbjoaV3fntKvN9k3Hz1Wds1kf1Gguc\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"CgD8kLxdmy7tx668GN5o5Vn5Qo79L3M8Dr68FPAeqAnd\",\"token_decimals\":6},\"amm_info_address\":\"412Mr1t8g1xSzW4wBaCV8J8KDFrhff46aNqGMSoK1asL\",\"open_orders_address\":\"8PpUm7L1Mtj65MBMwqMfgWdHWcdJFv7cDu6Fj1QAsNx7\"}"
        }
      ]
    },
    "WAM,RAYDIUM,7HDEO5QCIUF8S2VFSX6URJKDNVADBU3DDCXW4ZJDCMIN/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "WAM,RAYDIUM,7HDEO5QCIUF8S2VFSX6URJKDNVADBU3DDCXW4ZJDCMIN",
          "Quote": "USD"
        },
        "decimals": 12,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"14133\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "WAM,RAYDIUM,7HDEO5QCIUF8S2VFSX6URJKDNVADBU3DDCXW4ZJDCMIN/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"7EcEbCS2GqEkFnQX15d1g48WG1UycPtTrhSFQ3LfQva3\",\"token_decimals\":8},\"quote_token_vault\":{\"token_vault_address\":\"8oFhomqBCWjNiLGr7hWcQAJUivP8Dt1WksfSGjZGzaH6\",\"token_decimals\":9},\"amm_info_address\":\"DybAcYN1oXweV5JYSU9TY8S6KD6mfnoZMphNRUFchmo7\",\"open_orders_address\":\"CtPB6y9h5UfRDLEBSGd8ujHxdNv4EciZPLnnpwGaRx84\"}"
        }
      ]
    },
    "WDOG,RAYDIUM,GYKMDFCUMZVRQFCH1G579BGJUZSRIJJ3LBUWV79RPUMP/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "WDOG,RAYDIUM,GYKMDFCUMZVRQFCH1G579BGJUZSRIJJ3LBUWV79RPUMP",
          "Quote": "USD"
        },
        "decimals": 11,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"32618\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SOL,RAYDIUM,SO11111111111111111111111111111111111111112/WDOG,RAYDIUM,GYKMDFCUMZVRQFCH1G579BGJUZSRIJJ3LBUWV79RPUMP",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "invert": true,
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"8bWmUEf9zwA3RUF33p9XJvVT8ENjawt9SzBPv4vZ8j2J\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"BZKM8EjA33B6icmyprYxzdzaNVE743u2J1oSkLwnmKTR\",\"token_decimals\":6},\"amm_info_address\":\"25tXTutLkjtcUX3kqoeRvc7AuBYM7fckBWoVqnQnyDGQ\",\"open_orders_address\":\"75ffxZBo1kgWLwNkhffx7n4ww6kahXji9EpL546UPYkg\"}"
        }
      ]
    },
    "WEN,RAYDIUM,WENWENVQQNYA429UBCDR81ZMD69BRWQAABYY6P3LCPK/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "WEN,RAYDIUM,WENWENVQQNYA429UBCDR81ZMD69BRWQAABYY6P3LCPK",
          "Quote": "USD"
        },
        "decimals": 14,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"29175\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "WEN,RAYDIUM,WENWENVQQNYA429UBCDR81ZMD69BRWQAABYY6P3LCPK/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"46Bk4BaXzAsLYbrf3UQvGZLPdEwt5nkeb2wHjiqpa5VG\",\"token_decimals\":5},\"quote_token_vault\":{\"token_vault_address\":\"G86oqnqDH4NxsnhJEtMBPURwUfuJ8yd43ecsgg3kgRW4\",\"token_decimals\":9},\"amm_info_address\":\"5WGx6mE9Xww3ocYzSenGVQMJLCVVwK7ePnYV6cXcpJtK\",\"open_orders_address\":\"7bT5Jgzbf4apkm7TrUZtZGhufnjzQR2Js1bBcQsjUtzD\"}"
        }
      ]
    },
    "WHALES,RAYDIUM,GTH3WG3NERJWCF7VGCOXEXKGXSHVYHX5GTATEEM5JAS1/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "WHALES,RAYDIUM,GTH3WG3NERJWCF7VGCOXEXKGXSHVYHX5GTATEEM5JAS1",
          "Quote": "USD"
        },
        "decimals": 10,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"29282\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "WHALES,RAYDIUM,GTH3WG3NERJWCF7VGCOXEXKGXSHVYHX5GTATEEM5JAS1/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"P9uSMnNEGHDP7Dhu7fKWfRViAGGHjEMv6urC8c2qG4k\",\"token_decimals\":6},\"quote_token_vault\":{\"token_vault_address\":\"CktEbT37HFRtwXVjwPEVfXHdcTAnqnmCvkgAw9SEN7zf\",\"token_decimals\":9},\"amm_info_address\":\"DczmyvnV8hR7d8zvy6bAoc2itZbFvLAx9iG2D7gyyt9e\",\"open_orders_address\":\"5JAwqabcp6KnfUe88RiaMgdpE3nw6CQu4NyAfbGmNEz2\"}"
        }
      ]
    },
    "WINR,RAYDIUM,CSXCTA8USVWKDRHE7KHLU5GGWZYALKOHSZ1MKBVZ4W3M/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "WINR,RAYDIUM,CSXCTA8USVWKDRHE7KHLU5GGWZYALKOHSZ1MKBVZ4W3M",
          "Quote": "USD"
        },
        "decimals": 11,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"23681\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SOL,RAYDIUM,SO11111111111111111111111111111111111111112/WINR,RAYDIUM,CSXCTA8USVWKDRHE7KHLU5GGWZYALKOHSZ1MKBVZ4W3M",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "invert": true,
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"BFRANtoGUzsmhD2zFrF6RbMuoHZD719WuxRXDYB42k5Z\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"kuqEGi9ASemeTyD6HNrdeLLqfWwZPjHf6NM5HNvMwCQ\",\"token_decimals\":8},\"amm_info_address\":\"93o4y72oH55JbKvuU3QSibznocGYKdiexEE3zJ1WAhAe\",\"open_orders_address\":\"2FxC7VHDsPftdWguZALb53Wvj6gnpQyb449crgoGHrYG\"}"
        }
      ]
    },
    "WLKN,RAYDIUM,ECQCUYV57C4V6ROPXKVUIDWTX1SP8Y8FP5AETOYL8AZ/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "WLKN,RAYDIUM,ECQCUYV57C4V6ROPXKVUIDWTX1SP8Y8FP5AETOYL8AZ",
          "Quote": "USD"
        },
        "decimals": 13,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"18775\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "WLKN,RAYDIUM,ECQCUYV57C4V6ROPXKVUIDWTX1SP8Y8FP5AETOYL8AZ/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"9eAcfMeokystVVPvrhhXMVVGPgWLoCsvHySAUytEzdDE\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"4VyRMjuympymYqmYbVwXYBWxKrsg2irRXeGmakgc1eVf\",\"token_decimals\":9},\"amm_info_address\":\"5begBbgNM8T7ZPrG2dY3dasdGoGaQzkyPDsforaANGde\",\"open_orders_address\":\"9hiYzzRAywXVHxTqRV7JPu1y2pZiNGCiGpDHamZkW8N\"}"
        }
      ]
    },
    "WOLF,RAYDIUM,FAF89929NI9FBG4GMVZTCA7EW6NFG877JQN6MIZT3GVW/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "WOLF,RAYDIUM,FAF89929NI9FBG4GMVZTCA7EW6NFG877JQN6MIZT3GVW",
          "Quote": "USD"
        },
        "decimals": 12,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"31847\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "WOLF,RAYDIUM,FAF89929NI9FBG4GMVZTCA7EW6NFG877JQN6MIZT3GVW/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"FoMMZq2mt5KTa2yErsLRHGseFCNpvrzijH7zNSmsJHJs\",\"token_decimals\":8},\"quote_token_vault\":{\"token_vault_address\":\"4UynYZ9WhssfHBGTThY9prTVmrDjM8SDHkRsfQCR6wKW\",\"token_decimals\":9},\"amm_info_address\":\"EskpS4o6sWFAmf7w8PDwWJQfx8LRU7hHMjPxCcyX4Eq1\",\"open_orders_address\":\"CD4CVVusXHxt3faGYcUCBEBxviZGL2dgrCEjvBGkgGnL\"}"
        }
      ]
    },
    "YOUNES,RAYDIUM,ANAUIZ2JJRVNTYW8CD7USQ4LRWB1PTGHYRGMWPXTPUMP/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "YOUNES,RAYDIUM,ANAUIZ2JJRVNTYW8CD7USQ4LRWB1PTGHYRGMWPXTPUMP",
          "Quote": "USD"
        },
        "decimals": 12,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"32462\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SOL,RAYDIUM,SO11111111111111111111111111111111111111112/YOUNES,RAYDIUM,ANAUIZ2JJRVNTYW8CD7USQ4LRWB1PTGHYRGMWPXTPUMP",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "invert": true,
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"ARAyM5H4CgDZPWN9SDMUPy4QYbDiKLh4RDGG5nkzHZie\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"48Lj7zh43J3KJcasyig8fdEENvZwt7ctgYiXzxYzVSjR\",\"token_decimals\":6},\"amm_info_address\":\"6hY8WPTi9fHF5wgB2BVaYm1diANnZ628rSJZBZiANSvJ\",\"open_orders_address\":\"4zN6iG5xVDedGm24vdhU1vVKkfzpc5DgW611cdSisNKB\"}"
        }
      ]
    },
    "ZACK,RAYDIUM,8VCAUBXEJDTAXN6JNX5UAQTYTZLMXALG9U1BVFCAJTX7/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "ZACK,RAYDIUM,8VCAUBXEJDTAXN6JNX5UAQTYTZLMXALG9U1BVFCAJTX7",
          "Quote": "USD"
        },
        "decimals": 12,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"31367\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "SOL,RAYDIUM,SO11111111111111111111111111111111111111112/ZACK,RAYDIUM,8VCAUBXEJDTAXN6JNX5UAQTYTZLMXALG9U1BVFCAJTX7",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "invert": true,
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"5D1kvdok7eDhsR63ytmYAQgqquiJ5d38t3uCRSmYAdVF\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"5ynfwNHWJJjvbKY5f8SokBpnc82VSza5FskYwvQupM7V\",\"token_decimals\":6},\"amm_info_address\":\"7896DcX977xMJboS6BJvgkK4sB5p2FhctJx81DntbyCX\",\"open_orders_address\":\"FBn7gcBnXLoxdir2X1xsuy3XV8j1TLrRAxHnenadUGCq\"}"
        }
      ]
    },
    "ZAP,RAYDIUM,HXPOEHMT1VKEQJKCEPCQTJ6YYGN6XQQ1FKTY3PJX4YRX/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "ZAP,RAYDIUM,HXPOEHMT1VKEQJKCEPCQTJ6YYGN6XQQ1FKTY3PJX4YRX",
          "Quote": "USD"
        },
        "decimals": 12,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"2363\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "ZAP,RAYDIUM,HXPOEHMT1VKEQJKCEPCQTJ6YYGN6XQQ1FKTY3PJX4YRX/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"AUndu2tP6obZVFAtMXtR4FLigQ6jVvrCvpkk9pHQtBKA\",\"token_decimals\":8},\"quote_token_vault\":{\"token_vault_address\":\"TBxsQ3xvCQBJcxN3LzAiWD9XTXsc1XxoRtT99DMUPeX\",\"token_decimals\":9},\"amm_info_address\":\"4W9BeAbMBQnnFdBbVxU5fiK2jsauJhog2pabPExBmVQX\",\"open_orders_address\":\"3Xe8C88NA2B75cUAnAe6afsXtsg5DCgnLdeWy7bz5ds7\"}"
        }
      ]
    },
    "ZERO,RAYDIUM,93RC484OMK5T9H89RZT5QIAXKHGP9JSCXFFFRIHNBE57/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "ZERO,RAYDIUM,93RC484OMK5T9H89RZT5QIAXKHGP9JSCXFFFRIHNBE57",
          "Quote": "USD"
        },
        "decimals": 10,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"28803\"}]}"
      },
      "provider_configs": [
        {
          "name": "raydium_api",
          "off_chain_ticker": "ZERO,RAYDIUM,93RC484OMK5T9H89RZT5QIAXKHGP9JSCXFFFRIHNBE57/SOL,RAYDIUM,SO11111111111111111111111111111111111111112",
          "normalize_by_pair": {
            "Base": "SOL",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"base_token_vault\":{\"token_vault_address\":\"H9E57gyNAanNREFMVz3HDbHNTrPvuHdEDEiw2sFcBHAH\",\"token_decimals\":9},\"quote_token_vault\":{\"token_vault_address\":\"FWJavr6oJSJeGSBeSmvrTJkFrvHNeWTT2tfyAGJQP74M\",\"token_decimals\":9},\"amm_info_address\":\"8ekufKZHJa4UaHkTSfVhP9wq8o42XqcczWNmWjhJppfW\",\"open_orders_address\":\"C88mbtCpTmzRbTNeVMZjxzLDroGZaY2Qd4gChZME53ti\"}"
        }
      ]
    }
  }
}`

	// CoreMarketMap is used to initialize the Core market map.
	CoreMarketMap mmtypes.MarketMap
	// CoreMarketMapJSON is the JSON representation of the Core MarketMap that can be used
	// to initialize for a genesis state or used by the sidecar as as static market map.
	CoreMarketMapJSON = `
	{
		"markets": {
		  "AAVE/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "AAVE",
				"Quote": "USD"
			  },
			  "decimals": 8,
			  "min_provider_count": 3,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "mexc_ws",
				"off_chain_ticker": "AAVEUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "huobi_ws",
				"off_chain_ticker": "aaveusdt",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "kraken_api",
				"off_chain_ticker": "AAVEUSD"
			  },
			  {
				"name": "okx_ws",
				"off_chain_ticker": "AAVE-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "kucoin_ws",
				"off_chain_ticker": "AAVE-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "binance_ws",
				"off_chain_ticker": "AAVEUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "coinbase_ws",
				"off_chain_ticker": "AAVE-USD"
			  },
			  {
				"name": "crypto_dot_com_ws",
				"off_chain_ticker": "AAVE_USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  }
			]
		  },
		  "ADA/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "ADA",
				"Quote": "USD"
			  },
			  "decimals": 10,
			  "min_provider_count": 3,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "binance_ws",
				"off_chain_ticker": "ADAUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "bybit_ws",
				"off_chain_ticker": "ADAUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "coinbase_ws",
				"off_chain_ticker": "ADA-USD"
			  },
			  {
				"name": "gate_ws",
				"off_chain_ticker": "ADA_USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "huobi_ws",
				"off_chain_ticker": "adausdt",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "kraken_api",
				"off_chain_ticker": "ADAUSD"
			  },
			  {
				"name": "kucoin_ws",
				"off_chain_ticker": "ADA-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "mexc_ws",
				"off_chain_ticker": "ADAUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "okx_ws",
				"off_chain_ticker": "ADA-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "crypto_dot_com_ws",
				"off_chain_ticker": "ADA_USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  }
			]
		  },
		  "AEVO/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "AEVO",
				"Quote": "USD"
			  },
			  "decimals": 9,
			  "min_provider_count": 3,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "binance_ws",
				"off_chain_ticker": "AEVOUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "bybit_ws",
				"off_chain_ticker": "AEVOUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "gate_ws",
				"off_chain_ticker": "AEVO_USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "mexc_ws",
				"off_chain_ticker": "AEVOUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "okx_ws",
				"off_chain_ticker": "AEVO-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  }
			]
		  },
		  "ALGO/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "ALGO",
				"Quote": "USD"
			  },
			  "decimals": 10,
			  "min_provider_count": 3,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "binance_ws",
				"off_chain_ticker": "ALGOUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "bybit_ws",
				"off_chain_ticker": "ALGOUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "coinbase_ws",
				"off_chain_ticker": "ALGO-USD"
			  },
			  {
				"name": "kraken_api",
				"off_chain_ticker": "ALGOUSD"
			  },
			  {
				"name": "kucoin_ws",
				"off_chain_ticker": "ALGO-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "mexc_ws",
				"off_chain_ticker": "ALGOUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "okx_ws",
				"off_chain_ticker": "ALGO-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  }
			]
		  },
		  "APE/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "APE",
				"Quote": "USD"
			  },
			  "decimals": 9,
			  "min_provider_count": 3,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "binance_ws",
				"off_chain_ticker": "APEUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "coinbase_ws",
				"off_chain_ticker": "APE-USD"
			  },
			  {
				"name": "gate_ws",
				"off_chain_ticker": "APE_USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "kraken_api",
				"off_chain_ticker": "APEUSD"
			  },
			  {
				"name": "kucoin_ws",
				"off_chain_ticker": "APE-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "mexc_ws",
				"off_chain_ticker": "APEUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "okx_ws",
				"off_chain_ticker": "APE-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "crypto_dot_com_ws",
				"off_chain_ticker": "APE_USD"
			  }
			]
		  },
		  "APT/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "APT",
				"Quote": "USD"
			  },
			  "decimals": 9,
			  "min_provider_count": 3,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "binance_ws",
				"off_chain_ticker": "APTUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "bybit_ws",
				"off_chain_ticker": "APTUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "coinbase_ws",
				"off_chain_ticker": "APT-USD"
			  },
			  {
				"name": "gate_ws",
				"off_chain_ticker": "APT_USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "huobi_ws",
				"off_chain_ticker": "aptusdt",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "kucoin_ws",
				"off_chain_ticker": "APT-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "mexc_ws",
				"off_chain_ticker": "APTUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "okx_ws",
				"off_chain_ticker": "APT-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "crypto_dot_com_ws",
				"off_chain_ticker": "APT_USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  }
			]
		  },
		  "ARB/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "ARB",
				"Quote": "USD"
			  },
			  "decimals": 9,
			  "min_provider_count": 3,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "binance_ws",
				"off_chain_ticker": "ARBUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "bybit_ws",
				"off_chain_ticker": "ARBUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "coinbase_ws",
				"off_chain_ticker": "ARB-USD"
			  },
			  {
				"name": "gate_ws",
				"off_chain_ticker": "ARB_USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "huobi_ws",
				"off_chain_ticker": "arbusdt",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "kucoin_ws",
				"off_chain_ticker": "ARB-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "mexc_ws",
				"off_chain_ticker": "ARBUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "okx_ws",
				"off_chain_ticker": "ARB-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "crypto_dot_com_ws",
				"off_chain_ticker": "ARB_USD"
			  }
			]
		  },
		  "ARKM/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "ARKM",
				"Quote": "USD"
			  },
			  "decimals": 9,
			  "min_provider_count": 3,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "binance_ws",
				"off_chain_ticker": "ARKMUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "bybit_ws",
				"off_chain_ticker": "ARKMUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "gate_ws",
				"off_chain_ticker": "ARKM_USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "kucoin_ws",
				"off_chain_ticker": "ARKM-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "mexc_ws",
				"off_chain_ticker": "ARKMUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  }
			]
		  },
		  "ASTR/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "ASTR",
				"Quote": "USD"
			  },
			  "decimals": 10,
			  "min_provider_count": 3,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "binance_ws",
				"off_chain_ticker": "ASTRUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "gate_ws",
				"off_chain_ticker": "ASTR_USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "kraken_api",
				"off_chain_ticker": "ASTRUSD"
			  },
			  {
				"name": "kucoin_ws",
				"off_chain_ticker": "ASTR-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "mexc_ws",
				"off_chain_ticker": "ASTRUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "okx_ws",
				"off_chain_ticker": "ASTR-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  }
			]
		  },
		  "ATOM/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "ATOM",
				"Quote": "USD"
			  },
			  "decimals": 9,
			  "min_provider_count": 3,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "binance_ws",
				"off_chain_ticker": "ATOMUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "bybit_ws",
				"off_chain_ticker": "ATOMUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "coinbase_ws",
				"off_chain_ticker": "ATOM-USD"
			  },
			  {
				"name": "gate_ws",
				"off_chain_ticker": "ATOM_USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "kraken_api",
				"off_chain_ticker": "ATOMUSD"
			  },
			  {
				"name": "kucoin_ws",
				"off_chain_ticker": "ATOM-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "mexc_ws",
				"off_chain_ticker": "ATOMUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "okx_ws",
				"off_chain_ticker": "ATOM-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "crypto_dot_com_ws",
				"off_chain_ticker": "ATOM_USD"
			  }
			]
		  },
		  "AVAX/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "AVAX",
				"Quote": "USD"
			  },
			  "decimals": 8,
			  "min_provider_count": 3,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "binance_ws",
				"off_chain_ticker": "AVAXUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "bybit_ws",
				"off_chain_ticker": "AVAXUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "coinbase_ws",
				"off_chain_ticker": "AVAX-USD"
			  },
			  {
				"name": "gate_ws",
				"off_chain_ticker": "AVAX_USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "huobi_ws",
				"off_chain_ticker": "avaxusdt",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "kraken_api",
				"off_chain_ticker": "AVAXUSD"
			  },
			  {
				"name": "kucoin_ws",
				"off_chain_ticker": "AVAX-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "okx_ws",
				"off_chain_ticker": "AVAX-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "crypto_dot_com_ws",
				"off_chain_ticker": "AVAX_USD"
			  }
			]
		  },
		  "AXL/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "AXL",
				"Quote": "USD"
			  },
			  "decimals": 9,
			  "min_provider_count": 3,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "binance_ws",
				"off_chain_ticker": "AXLUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "bybit_ws",
				"off_chain_ticker": "AXLUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "coinbase_ws",
				"off_chain_ticker": "AXL-USD"
			  },
			  {
				"name": "gate_ws",
				"off_chain_ticker": "WAXL_USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "kraken_api",
				"off_chain_ticker": "WAXLUSD"
			  }
			]
		  },
		  "BCH/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "BCH",
				"Quote": "USD"
			  },
			  "decimals": 7,
			  "min_provider_count": 3,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "binance_ws",
				"off_chain_ticker": "BCHUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "bybit_ws",
				"off_chain_ticker": "BCHUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "coinbase_ws",
				"off_chain_ticker": "BCH-USD"
			  },
			  {
				"name": "gate_ws",
				"off_chain_ticker": "BCH_USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "huobi_ws",
				"off_chain_ticker": "bchusdt",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "kraken_api",
				"off_chain_ticker": "BCHUSD"
			  },
			  {
				"name": "kucoin_ws",
				"off_chain_ticker": "BCH-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "mexc_ws",
				"off_chain_ticker": "BCHUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "okx_ws",
				"off_chain_ticker": "BCH-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "crypto_dot_com_ws",
				"off_chain_ticker": "BCH_USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  }
			]
		  },
		  "BLUR/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "BLUR",
				"Quote": "USD"
			  },
			  "decimals": 10,
			  "min_provider_count": 3,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coinbase_ws",
				"off_chain_ticker": "BLUR-USD"
			  },
			  {
				"name": "gate_ws",
				"off_chain_ticker": "BLUR_USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "kraken_api",
				"off_chain_ticker": "BLURUSD"
			  },
			  {
				"name": "kucoin_ws",
				"off_chain_ticker": "BLUR-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "mexc_ws",
				"off_chain_ticker": "BLURUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "okx_ws",
				"off_chain_ticker": "BLUR-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  }
			]
		  },
		  "BNB/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "BNB",
				"Quote": "USD"
			  },
			  "decimals": 7,
			  "min_provider_count": 3,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "okx_ws",
				"off_chain_ticker": "BNB-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "kucoin_ws",
				"off_chain_ticker": "BNB-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "mexc_ws",
				"off_chain_ticker": "BNBUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "binance_ws",
				"off_chain_ticker": "BNBUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "bybit_ws",
				"off_chain_ticker": "BNBUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "gate_ws",
				"off_chain_ticker": "BNB_USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  }
			]
		  },
		  "BONK/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "BONK",
				"Quote": "USD"
			  },
			  "decimals": 14,
			  "min_provider_count": 3,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "binance_ws",
				"off_chain_ticker": "BONKUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "bybit_ws",
				"off_chain_ticker": "BONKUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "coinbase_ws",
				"off_chain_ticker": "BONK-USD"
			  },
			  {
				"name": "kucoin_ws",
				"off_chain_ticker": "BONK-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "okx_ws",
				"off_chain_ticker": "BONK-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "mexc_ws",
				"off_chain_ticker": "BONKUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "crypto_dot_com_ws",
				"off_chain_ticker": "BONK_USD"
			  }
			]
		  },
		  "BTC/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "BTC",
				"Quote": "USD"
			  },
			  "decimals": 5,
			  "min_provider_count": 3,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "binance_ws",
				"off_chain_ticker": "BTCUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "bybit_ws",
				"off_chain_ticker": "BTCUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "coinbase_ws",
				"off_chain_ticker": "BTC-USD"
			  },
			  {
				"name": "huobi_ws",
				"off_chain_ticker": "btcusdt",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "kraken_api",
				"off_chain_ticker": "XXBTZUSD"
			  },
			  {
				"name": "kucoin_ws",
				"off_chain_ticker": "BTC-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "mexc_ws",
				"off_chain_ticker": "BTCUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "okx_ws",
				"off_chain_ticker": "BTC-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "crypto_dot_com_ws",
				"off_chain_ticker": "BTC_USD"
			  }
			]
		  },
		  "COMP/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "COMP",
				"Quote": "USD"
			  },
			  "decimals": 8,
			  "min_provider_count": 3,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "binance_ws",
				"off_chain_ticker": "COMPUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "coinbase_ws",
				"off_chain_ticker": "COMP-USD"
			  },
			  {
				"name": "gate_ws",
				"off_chain_ticker": "COMP_USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "kraken_api",
				"off_chain_ticker": "COMPUSD"
			  },
			  {
				"name": "mexc_ws",
				"off_chain_ticker": "COMPUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "okx_ws",
				"off_chain_ticker": "COMP-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  }
			]
		  },
		  "CRV/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "CRV",
				"Quote": "USD"
			  },
			  "decimals": 10,
			  "min_provider_count": 3,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "binance_ws",
				"off_chain_ticker": "CRVUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "coinbase_ws",
				"off_chain_ticker": "CRV-USD"
			  },
			  {
				"name": "gate_ws",
				"off_chain_ticker": "CRV_USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "kraken_api",
				"off_chain_ticker": "CRVUSD"
			  },
			  {
				"name": "kucoin_ws",
				"off_chain_ticker": "CRV-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "mexc_ws",
				"off_chain_ticker": "CRVUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "okx_ws",
				"off_chain_ticker": "CRV-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  }
			]
		  },
		  "DOGE/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "DOGE",
				"Quote": "USD"
			  },
			  "decimals": 11,
			  "min_provider_count": 3,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "binance_ws",
				"off_chain_ticker": "DOGEUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "bybit_ws",
				"off_chain_ticker": "DOGEUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "coinbase_ws",
				"off_chain_ticker": "DOGE-USD"
			  },
			  {
				"name": "gate_ws",
				"off_chain_ticker": "DOGE_USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "huobi_ws",
				"off_chain_ticker": "dogeusdt",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "kraken_api",
				"off_chain_ticker": "XDGUSD"
			  },
			  {
				"name": "kucoin_ws",
				"off_chain_ticker": "DOGE-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "mexc_ws",
				"off_chain_ticker": "DOGEUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "okx_ws",
				"off_chain_ticker": "DOGE-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "crypto_dot_com_ws",
				"off_chain_ticker": "DOGE_USD"
			  }
			]
		  },
		  "DOT/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "DOT",
				"Quote": "USD"
			  },
			  "decimals": 9,
			  "min_provider_count": 3,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "binance_ws",
				"off_chain_ticker": "DOTUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "bybit_ws",
				"off_chain_ticker": "DOTUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "coinbase_ws",
				"off_chain_ticker": "DOT-USD"
			  },
			  {
				"name": "gate_ws",
				"off_chain_ticker": "DOT_USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "kraken_api",
				"off_chain_ticker": "DOTUSD"
			  },
			  {
				"name": "kucoin_ws",
				"off_chain_ticker": "DOT-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "mexc_ws",
				"off_chain_ticker": "DOTUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "okx_ws",
				"off_chain_ticker": "DOT-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "crypto_dot_com_ws",
				"off_chain_ticker": "DOT_USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  }
			]
		  },
		  "DYDX/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "DYDX",
				"Quote": "USD"
			  },
			  "decimals": 9,
			  "min_provider_count": 3,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "binance_ws",
				"off_chain_ticker": "DYDXUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "bybit_ws",
				"off_chain_ticker": "DYDXUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "gate_ws",
				"off_chain_ticker": "DYDX_USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "kucoin_ws",
				"off_chain_ticker": "DYDX-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "mexc_ws",
				"off_chain_ticker": "DYDXUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "okx_ws",
				"off_chain_ticker": "DYDX-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  }
			]
		  },
		  "DYM/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "DYM",
				"Quote": "USD"
			  },
			  "decimals": 9,
			  "min_provider_count": 3,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "binance_ws",
				"off_chain_ticker": "DYMUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "bybit_ws",
				"off_chain_ticker": "DYMUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "gate_ws",
				"off_chain_ticker": "DYM_USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "kucoin_ws",
				"off_chain_ticker": "DYM-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "mexc_ws",
				"off_chain_ticker": "DYMUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  }
			]
		  },
		  "EOS/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "EOS",
				"Quote": "USD"
			  },
			  "decimals": 10,
			  "min_provider_count": 3,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "binance_ws",
				"off_chain_ticker": "EOSUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "bybit_ws",
				"off_chain_ticker": "EOSUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "coinbase_ws",
				"off_chain_ticker": "EOS-USD"
			  },
			  {
				"name": "gate_ws",
				"off_chain_ticker": "EOS_USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "kraken_api",
				"off_chain_ticker": "EOSUSD"
			  },
			  {
				"name": "kucoin_ws",
				"off_chain_ticker": "EOS-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "okx_ws",
				"off_chain_ticker": "EOS-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "mexc_ws",
				"off_chain_ticker": "EOSUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  }
			]
		  },
		  "ETC/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "ETC",
				"Quote": "USD"
			  },
			  "decimals": 8,
			  "min_provider_count": 3,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "binance_ws",
				"off_chain_ticker": "ETCUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "coinbase_ws",
				"off_chain_ticker": "ETC-USD"
			  },
			  {
				"name": "gate_ws",
				"off_chain_ticker": "ETC_USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "huobi_ws",
				"off_chain_ticker": "etcusdt",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "kucoin_ws",
				"off_chain_ticker": "ETC-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "mexc_ws",
				"off_chain_ticker": "ETCUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "okx_ws",
				"off_chain_ticker": "ETC-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  }
			]
		  },
		  "ETH/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "ETH",
				"Quote": "USD"
			  },
			  "decimals": 6,
			  "min_provider_count": 3,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "binance_ws",
				"off_chain_ticker": "ETHUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "bybit_ws",
				"off_chain_ticker": "ETHUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "coinbase_ws",
				"off_chain_ticker": "ETH-USD"
			  },
			  {
				"name": "huobi_ws",
				"off_chain_ticker": "ethusdt",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "kraken_api",
				"off_chain_ticker": "XETHZUSD"
			  },
			  {
				"name": "kucoin_ws",
				"off_chain_ticker": "ETH-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "mexc_ws",
				"off_chain_ticker": "ETHUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "okx_ws",
				"off_chain_ticker": "ETH-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "crypto_dot_com_ws",
				"off_chain_ticker": "ETH_USD"
			  }
			]
		  },
		  "FET/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "FET",
				"Quote": "USD"
			  },
			  "decimals": 10,
			  "min_provider_count": 3,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "binance_ws",
				"off_chain_ticker": "FETUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "coinbase_ws",
				"off_chain_ticker": "FET-USD"
			  },
			  {
				"name": "kraken_api",
				"off_chain_ticker": "FETUSD"
			  },
			  {
				"name": "kucoin_ws",
				"off_chain_ticker": "FET-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "okx_ws",
				"off_chain_ticker": "FET-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "mexc_ws",
				"off_chain_ticker": "FETUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "crypto_dot_com_ws",
				"off_chain_ticker": "FET_USD"
			  }
			]
		  },
		  "FIL/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "FIL",
				"Quote": "USD"
			  },
			  "decimals": 9,
			  "min_provider_count": 3,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "binance_ws",
				"off_chain_ticker": "FILUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "coinbase_ws",
				"off_chain_ticker": "FIL-USD"
			  },
			  {
				"name": "gate_ws",
				"off_chain_ticker": "FIL_USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "huobi_ws",
				"off_chain_ticker": "filusdt",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "kraken_api",
				"off_chain_ticker": "FILUSD"
			  },
			  {
				"name": "mexc_ws",
				"off_chain_ticker": "FILUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "okx_ws",
				"off_chain_ticker": "FIL-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  }
			]
		  },
		  "GRT/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "GRT",
				"Quote": "USD"
			  },
			  "decimals": 10,
			  "min_provider_count": 3,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "binance_ws",
				"off_chain_ticker": "GRTUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "bybit_ws",
				"off_chain_ticker": "GRTUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "coinbase_ws",
				"off_chain_ticker": "GRT-USD"
			  },
			  {
				"name": "gate_ws",
				"off_chain_ticker": "GRT_USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "kraken_api",
				"off_chain_ticker": "GRTUSD"
			  },
			  {
				"name": "kucoin_ws",
				"off_chain_ticker": "GRT-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "mexc_ws",
				"off_chain_ticker": "GRTUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "okx_ws",
				"off_chain_ticker": "GRT-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "crypto_dot_com_ws",
				"off_chain_ticker": "GRT_USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  }
			]
		  },
		  "HBAR/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "HBAR",
				"Quote": "USD"
			  },
			  "decimals": 10,
			  "min_provider_count": 3,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "binance_ws",
				"off_chain_ticker": "HBARUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "bitstamp_api",
				"off_chain_ticker": "hbarusd"
			  },
			  {
				"name": "bybit_ws",
				"off_chain_ticker": "HBARUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "coinbase_ws",
				"off_chain_ticker": "HBAR-USD"
			  },
			  {
				"name": "kucoin_ws",
				"off_chain_ticker": "HBAR-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "mexc_ws",
				"off_chain_ticker": "HBARUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "okx_ws",
				"off_chain_ticker": "HBAR-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "crypto_dot_com_ws",
				"off_chain_ticker": "HBAR_USD"
			  }
			]
		  },
		  "ICP/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "ICP",
				"Quote": "USD"
			  },
			  "decimals": 8,
			  "min_provider_count": 3,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "binance_ws",
				"off_chain_ticker": "ICPUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "bybit_ws",
				"off_chain_ticker": "ICPUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "coinbase_ws",
				"off_chain_ticker": "ICP-USD"
			  },
			  {
				"name": "kraken_api",
				"off_chain_ticker": "ICPUSD"
			  },
			  {
				"name": "kucoin_ws",
				"off_chain_ticker": "ICP-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "okx_ws",
				"off_chain_ticker": "ICP-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "mexc_ws",
				"off_chain_ticker": "ICPUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "crypto_dot_com_ws",
				"off_chain_ticker": "ICP_USD"
			  }
			]
		  },
		  "IMX/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "IMX",
				"Quote": "USD"
			  },
			  "decimals": 9,
			  "min_provider_count": 3,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "binance_ws",
				"off_chain_ticker": "IMXUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "coinbase_ws",
				"off_chain_ticker": "IMX-USD"
			  },
			  {
				"name": "kraken_api",
				"off_chain_ticker": "IMXUSD"
			  },
			  {
				"name": "kucoin_ws",
				"off_chain_ticker": "IMX-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "mexc_ws",
				"off_chain_ticker": "IMXUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "okx_ws",
				"off_chain_ticker": "IMX-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  }
			]
		  },
		  "INJ/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "INJ",
				"Quote": "USD"
			  },
			  "decimals": 8,
			  "min_provider_count": 3,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "binance_ws",
				"off_chain_ticker": "INJUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "bybit_ws",
				"off_chain_ticker": "INJUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "coinbase_ws",
				"off_chain_ticker": "INJ-USD"
			  },
			  {
				"name": "kraken_api",
				"off_chain_ticker": "INJUSD"
			  },
			  {
				"name": "kucoin_ws",
				"off_chain_ticker": "INJ-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "okx_ws",
				"off_chain_ticker": "INJ-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "mexc_ws",
				"off_chain_ticker": "INJUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "crypto_dot_com_ws",
				"off_chain_ticker": "INJ_USD"
			  }
			]
		  },
		  "JTO/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "JTO",
				"Quote": "USD"
			  },
			  "decimals": 9,
			  "min_provider_count": 3,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coinbase_ws",
				"off_chain_ticker": "JTO-USD"
			  },
			  {
				"name": "binance_ws",
				"off_chain_ticker": "JTOUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "bybit_ws",
				"off_chain_ticker": "JTOUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "mexc_ws",
				"off_chain_ticker": "JTOUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "okx_ws",
				"off_chain_ticker": "JTO-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "kucoin_ws",
				"off_chain_ticker": "JTO-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  }
			]
		  },
		  "JUP/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "JUP",
				"Quote": "USD"
			  },
			  "decimals": 10,
			  "min_provider_count": 3,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "kucoin_ws",
				"off_chain_ticker": "JUP-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "okx_ws",
				"off_chain_ticker": "JUP-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "binance_ws",
				"off_chain_ticker": "JUPUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "bybit_ws",
				"off_chain_ticker": "JUPUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "gate_ws",
				"off_chain_ticker": "JUP_USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "mexc_ws",
				"off_chain_ticker": "JUPUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "crypto_dot_com_ws",
				"off_chain_ticker": "JUP_USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  }
			]
		  },
		  "LDO/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "LDO",
				"Quote": "USD"
			  },
			  "decimals": 9,
			  "min_provider_count": 3,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "binance_ws",
				"off_chain_ticker": "LDOUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "coinbase_ws",
				"off_chain_ticker": "LDO-USD"
			  },
			  {
				"name": "kraken_api",
				"off_chain_ticker": "LDOUSD"
			  },
			  {
				"name": "kucoin_ws",
				"off_chain_ticker": "LDO-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "mexc_ws",
				"off_chain_ticker": "LDOUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "okx_ws",
				"off_chain_ticker": "LDO-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "crypto_dot_com_ws",
				"off_chain_ticker": "LDO_USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  }
			]
		  },
		  "LINK/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "LINK",
				"Quote": "USD"
			  },
			  "decimals": 9,
			  "min_provider_count": 3,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "binance_ws",
				"off_chain_ticker": "LINKUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "bybit_ws",
				"off_chain_ticker": "LINKUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "coinbase_ws",
				"off_chain_ticker": "LINK-USD"
			  },
			  {
				"name": "kraken_api",
				"off_chain_ticker": "LINKUSD"
			  },
			  {
				"name": "kucoin_ws",
				"off_chain_ticker": "LINK-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "mexc_ws",
				"off_chain_ticker": "LINKUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "okx_ws",
				"off_chain_ticker": "LINK-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "crypto_dot_com_ws",
				"off_chain_ticker": "LINK_USD"
			  }
			]
		  },
		  "LTC/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "LTC",
				"Quote": "USD"
			  },
			  "decimals": 8,
			  "min_provider_count": 3,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "binance_ws",
				"off_chain_ticker": "LTCUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "bybit_ws",
				"off_chain_ticker": "LTCUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "coinbase_ws",
				"off_chain_ticker": "LTC-USD"
			  },
			  {
				"name": "huobi_ws",
				"off_chain_ticker": "ltcusdt",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "kraken_api",
				"off_chain_ticker": "XLTCZUSD"
			  },
			  {
				"name": "kucoin_ws",
				"off_chain_ticker": "LTC-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "mexc_ws",
				"off_chain_ticker": "LTCUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "okx_ws",
				"off_chain_ticker": "LTC-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "crypto_dot_com_ws",
				"off_chain_ticker": "LTC_USD"
			  }
			]
		  },
		  "MANA/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "MANA",
				"Quote": "USD"
			  },
			  "decimals": 10,
			  "min_provider_count": 3,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "binance_ws",
				"off_chain_ticker": "MANAUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "coinbase_ws",
				"off_chain_ticker": "MANA-USD"
			  },
			  {
				"name": "gate_ws",
				"off_chain_ticker": "MANA_USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "kraken_api",
				"off_chain_ticker": "MANAUSD"
			  },
			  {
				"name": "kucoin_ws",
				"off_chain_ticker": "MANA-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "mexc_ws",
				"off_chain_ticker": "MANAUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "okx_ws",
				"off_chain_ticker": "MANA-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  }
			]
		  },
		  "MATIC/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "MATIC",
				"Quote": "USD"
			  },
			  "decimals": 10,
			  "min_provider_count": 3,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "binance_ws",
				"off_chain_ticker": "MATICUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "bybit_ws",
				"off_chain_ticker": "MATICUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "coinbase_ws",
				"off_chain_ticker": "MATIC-USD"
			  },
			  {
				"name": "gate_ws",
				"off_chain_ticker": "MATIC_USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "huobi_ws",
				"off_chain_ticker": "maticusdt",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "kraken_api",
				"off_chain_ticker": "MATICUSD"
			  },
			  {
				"name": "kucoin_ws",
				"off_chain_ticker": "MATIC-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "mexc_ws",
				"off_chain_ticker": "MATICUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "okx_ws",
				"off_chain_ticker": "MATIC-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "crypto_dot_com_ws",
				"off_chain_ticker": "MATIC_USD"
			  }
			]
		  },
		  "MKR/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "MKR",
				"Quote": "USD"
			  },
			  "decimals": 6,
			  "min_provider_count": 3,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "binance_ws",
				"off_chain_ticker": "MKRUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "coinbase_ws",
				"off_chain_ticker": "MKR-USD"
			  },
			  {
				"name": "kraken_api",
				"off_chain_ticker": "MKRUSD"
			  },
			  {
				"name": "kucoin_ws",
				"off_chain_ticker": "MKR-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "mexc_ws",
				"off_chain_ticker": "MKRUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "okx_ws",
				"off_chain_ticker": "MKR-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  }
			]
		  },
		  "NEAR/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "NEAR",
				"Quote": "USD"
			  },
			  "decimals": 9,
			  "min_provider_count": 3,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "binance_ws",
				"off_chain_ticker": "NEARUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "coinbase_ws",
				"off_chain_ticker": "NEAR-USD"
			  },
			  {
				"name": "gate_ws",
				"off_chain_ticker": "NEAR_USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "huobi_ws",
				"off_chain_ticker": "nearusdt",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "kucoin_ws",
				"off_chain_ticker": "NEAR-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "mexc_ws",
				"off_chain_ticker": "NEARUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "okx_ws",
				"off_chain_ticker": "NEAR-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "crypto_dot_com_ws",
				"off_chain_ticker": "NEAR_USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  }
			]
		  },
		  "NTRN/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "NTRN",
				"Quote": "USD"
			  },
			  "decimals": 8,
			  "min_provider_count": 2,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "binance_ws",
				"off_chain_ticker": "NTRNUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "gate_ws",
				"off_chain_ticker": "NTRN_USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "kucoin_ws",
				"off_chain_ticker": "NTRN-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "okx_ws",
				"off_chain_ticker": "NTRN-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  }
			]
		  },
		  "OP/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "OP",
				"Quote": "USD"
			  },
			  "decimals": 9,
			  "min_provider_count": 3,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "binance_ws",
				"off_chain_ticker": "OPUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "coinbase_ws",
				"off_chain_ticker": "OP-USD"
			  },
			  {
				"name": "gate_ws",
				"off_chain_ticker": "OP_USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "kucoin_ws",
				"off_chain_ticker": "OP-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "mexc_ws",
				"off_chain_ticker": "OPUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "okx_ws",
				"off_chain_ticker": "OP-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "crypto_dot_com_ws",
				"off_chain_ticker": "OP_USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  }
			]
		  },
		  "ORDI/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "ORDI",
				"Quote": "USD"
			  },
			  "decimals": 8,
			  "min_provider_count": 3,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "binance_ws",
				"off_chain_ticker": "ORDIUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "bybit_ws",
				"off_chain_ticker": "ORDIUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "gate_ws",
				"off_chain_ticker": "ORDI_USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "huobi_ws",
				"off_chain_ticker": "ordiusdt",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "kucoin_ws",
				"off_chain_ticker": "ORDI-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "okx_ws",
				"off_chain_ticker": "ORDI-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "mexc_ws",
				"off_chain_ticker": "ORDIUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  }
			]
		  },
		  "PEPE/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "PEPE",
				"Quote": "USD"
			  },
			  "decimals": 16,
			  "min_provider_count": 3,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "binance_ws",
				"off_chain_ticker": "PEPEUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "bybit_ws",
				"off_chain_ticker": "PEPEUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "gate_ws",
				"off_chain_ticker": "PEPE_USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "kraken_api",
				"off_chain_ticker": "PEPEUSD"
			  },
			  {
				"name": "kucoin_ws",
				"off_chain_ticker": "PEPE-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "mexc_ws",
				"off_chain_ticker": "PEPEUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "okx_ws",
				"off_chain_ticker": "PEPE-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "crypto_dot_com_ws",
				"off_chain_ticker": "PEPE_USD"
			  }
			]
		  },
		  "PYTH/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "PYTH",
				"Quote": "USD"
			  },
			  "decimals": 10,
			  "min_provider_count": 3,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "binance_ws",
				"off_chain_ticker": "PYTHUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "bybit_ws",
				"off_chain_ticker": "PYTHUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "gate_ws",
				"off_chain_ticker": "PYTH_USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "kucoin_ws",
				"off_chain_ticker": "PYTH-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "okx_ws",
				"off_chain_ticker": "PYTH-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "mexc_ws",
				"off_chain_ticker": "PYTHUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  }
			]
		  },
		  "RUNE/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "RUNE",
				"Quote": "USD"
			  },
			  "decimals": 9,
			  "min_provider_count": 3,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "binance_ws",
				"off_chain_ticker": "RUNEUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "gate_ws",
				"off_chain_ticker": "RUNE_USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "kraken_api",
				"off_chain_ticker": "RUNEUSD"
			  },
			  {
				"name": "kucoin_ws",
				"off_chain_ticker": "RUNE-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "mexc_ws",
				"off_chain_ticker": "RUNEUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  }
			]
		  },
		  "SEI/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "SEI",
				"Quote": "USD"
			  },
			  "decimals": 10,
			  "min_provider_count": 3,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "binance_ws",
				"off_chain_ticker": "SEIUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "bybit_ws",
				"off_chain_ticker": "SEIUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "coinbase_ws",
				"off_chain_ticker": "SEI-USD"
			  },
			  {
				"name": "gate_ws",
				"off_chain_ticker": "SEI_USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "huobi_ws",
				"off_chain_ticker": "seiusdt",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "kucoin_ws",
				"off_chain_ticker": "SEI-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "mexc_ws",
				"off_chain_ticker": "SEIUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  }
			]
		  },
		  "SHIB/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "SHIB",
				"Quote": "USD"
			  },
			  "decimals": 15,
			  "min_provider_count": 3,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "binance_ws",
				"off_chain_ticker": "SHIBUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "bybit_ws",
				"off_chain_ticker": "SHIBUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "coinbase_ws",
				"off_chain_ticker": "SHIB-USD"
			  },
			  {
				"name": "gate_ws",
				"off_chain_ticker": "SHIB_USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "kraken_api",
				"off_chain_ticker": "SHIBUSD"
			  },
			  {
				"name": "kucoin_ws",
				"off_chain_ticker": "SHIB-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "mexc_ws",
				"off_chain_ticker": "SHIBUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "okx_ws",
				"off_chain_ticker": "SHIB-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "crypto_dot_com_ws",
				"off_chain_ticker": "SHIB_USD"
			  }
			]
		  },
		  "SNX/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "SNX",
				"Quote": "USD"
			  },
			  "decimals": 9,
			  "min_provider_count": 3,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "binance_ws",
				"off_chain_ticker": "SNXUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "bybit_ws",
				"off_chain_ticker": "SNXUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "coinbase_ws",
				"off_chain_ticker": "SNX-USD"
			  },
			  {
				"name": "kraken_api",
				"off_chain_ticker": "SNXUSD"
			  },
			  {
				"name": "mexc_ws",
				"off_chain_ticker": "SNXUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "okx_ws",
				"off_chain_ticker": "SNX-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  }
			]
		  },
		  "SOL/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "SOL",
				"Quote": "USD"
			  },
			  "decimals": 8,
			  "min_provider_count": 3,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "binance_ws",
				"off_chain_ticker": "SOLUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "bybit_ws",
				"off_chain_ticker": "SOLUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "coinbase_ws",
				"off_chain_ticker": "SOL-USD"
			  },
			  {
				"name": "huobi_ws",
				"off_chain_ticker": "solusdt",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "kraken_api",
				"off_chain_ticker": "SOLUSD"
			  },
			  {
				"name": "kucoin_ws",
				"off_chain_ticker": "SOL-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "mexc_ws",
				"off_chain_ticker": "SOLUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "okx_ws",
				"off_chain_ticker": "SOL-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "crypto_dot_com_ws",
				"off_chain_ticker": "SOL_USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  }
			]
		  },
		  "STRK/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "STRK",
				"Quote": "USD"
			  },
			  "decimals": 9,
			  "min_provider_count": 3,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "binance_ws",
				"off_chain_ticker": "STRKUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "bybit_ws",
				"off_chain_ticker": "STRKUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "kraken_api",
				"off_chain_ticker": "STRKUSD"
			  },
			  {
				"name": "kucoin_ws",
				"off_chain_ticker": "STRK-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "okx_ws",
				"off_chain_ticker": "STRK-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "gate_ws",
				"off_chain_ticker": "STRK_USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  }
			]
		  },
		  "STX/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "STX",
				"Quote": "USD"
			  },
			  "decimals": 9,
			  "min_provider_count": 3,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "binance_ws",
				"off_chain_ticker": "STXUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "bybit_ws",
				"off_chain_ticker": "STXUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "coinbase_ws",
				"off_chain_ticker": "STX-USD"
			  },
			  {
				"name": "gate_ws",
				"off_chain_ticker": "STX_USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "kraken_api",
				"off_chain_ticker": "STXUSD"
			  },
			  {
				"name": "kucoin_ws",
				"off_chain_ticker": "STX-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "okx_ws",
				"off_chain_ticker": "STX-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "mexc_ws",
				"off_chain_ticker": "STXUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "crypto_dot_com_ws",
				"off_chain_ticker": "STX_USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  }
			]
		  },
		  "SUI/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "SUI",
				"Quote": "USD"
			  },
			  "decimals": 10,
			  "min_provider_count": 3,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "binance_ws",
				"off_chain_ticker": "SUIUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "bybit_ws",
				"off_chain_ticker": "SUIUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "coinbase_ws",
				"off_chain_ticker": "SUI-USD"
			  },
			  {
				"name": "gate_ws",
				"off_chain_ticker": "SUI_USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "huobi_ws",
				"off_chain_ticker": "suiusdt",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "kucoin_ws",
				"off_chain_ticker": "SUI-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "mexc_ws",
				"off_chain_ticker": "SUIUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "okx_ws",
				"off_chain_ticker": "SUI-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "crypto_dot_com_ws",
				"off_chain_ticker": "SUI_USD"
			  }
			]
		  },
		  "TIA/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "TIA",
				"Quote": "USD"
			  },
			  "decimals": 8,
			  "min_provider_count": 3,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "binance_ws",
				"off_chain_ticker": "TIAUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "bybit_ws",
				"off_chain_ticker": "TIAUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "coinbase_ws",
				"off_chain_ticker": "TIA-USD"
			  },
			  {
				"name": "huobi_ws",
				"off_chain_ticker": "tiausdt",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "kraken_api",
				"off_chain_ticker": "TIAUSD"
			  },
			  {
				"name": "kucoin_ws",
				"off_chain_ticker": "TIA-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "mexc_ws",
				"off_chain_ticker": "TIAUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "okx_ws",
				"off_chain_ticker": "TIA-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "crypto_dot_com_ws",
				"off_chain_ticker": "TIA_USD"
			  }
			]
		  },
		  "TRX/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "TRX",
				"Quote": "USD"
			  },
			  "decimals": 11,
			  "min_provider_count": 3,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "binance_ws",
				"off_chain_ticker": "TRXUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "bybit_ws",
				"off_chain_ticker": "TRXUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "gate_ws",
				"off_chain_ticker": "TRX_USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "huobi_ws",
				"off_chain_ticker": "trxusdt",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "kraken_api",
				"off_chain_ticker": "TRXUSD"
			  },
			  {
				"name": "kucoin_ws",
				"off_chain_ticker": "TRX-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "mexc_ws",
				"off_chain_ticker": "TRXUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "okx_ws",
				"off_chain_ticker": "TRX-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  }
			]
		  },
		  "UNI/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "UNI",
				"Quote": "USD"
			  },
			  "decimals": 9,
			  "min_provider_count": 3,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "binance_ws",
				"off_chain_ticker": "UNIUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "bybit_ws",
				"off_chain_ticker": "UNIUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "coinbase_ws",
				"off_chain_ticker": "UNI-USD"
			  },
			  {
				"name": "gate_ws",
				"off_chain_ticker": "UNI_USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "kraken_api",
				"off_chain_ticker": "UNIUSD"
			  },
			  {
				"name": "kucoin_ws",
				"off_chain_ticker": "UNI-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "okx_ws",
				"off_chain_ticker": "UNI-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "crypto_dot_com_ws",
				"off_chain_ticker": "UNI_USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  }
			]
		  },
		  "USDT/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "USDT",
				"Quote": "USD"
			  },
			  "decimals": 9,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "binance_ws",
				"off_chain_ticker": "USDCUSDT",
				"invert": true
			  },
			  {
				"name": "bybit_ws",
				"off_chain_ticker": "USDCUSDT",
				"invert": true
			  },
			  {
				"name": "coinbase_ws",
				"off_chain_ticker": "USDT-USD"
			  },
			  {
				"name": "huobi_ws",
				"off_chain_ticker": "ethusdt",
				"normalize_by_pair": {
				  "Base": "ETH",
				  "Quote": "USD"
				},
				"invert": true
			  },
			  {
				"name": "kraken_api",
				"off_chain_ticker": "USDTZUSD"
			  },
			  {
				"name": "kucoin_ws",
				"off_chain_ticker": "BTC-USDT",
				"normalize_by_pair": {
				  "Base": "BTC",
				  "Quote": "USD"
				},
				"invert": true
			  },
			  {
				"name": "okx_ws",
				"off_chain_ticker": "USDC-USDT",
				"invert": true
			  },
			  {
				"name": "crypto_dot_com_ws",
				"off_chain_ticker": "USDT_USD"
			  }
			]
		  },
		  "WLD/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "WLD",
				"Quote": "USD"
			  },
			  "decimals": 9,
			  "min_provider_count": 3,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "binance_ws",
				"off_chain_ticker": "WLDUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "bybit_ws",
				"off_chain_ticker": "WLDUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "gate_ws",
				"off_chain_ticker": "WLD_USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "huobi_ws",
				"off_chain_ticker": "wldusdt",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "kucoin_ws",
				"off_chain_ticker": "WLD-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "mexc_ws",
				"off_chain_ticker": "WLDUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "okx_ws",
				"off_chain_ticker": "WLD-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "crypto_dot_com_ws",
				"off_chain_ticker": "WLD_USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  }
			]
		  },
		  "WOO/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "WOO",
				"Quote": "USD"
			  },
			  "decimals": 10,
			  "min_provider_count": 3,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "binance_ws",
				"off_chain_ticker": "WOOUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "gate_ws",
				"off_chain_ticker": "WOO_USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "kucoin_ws",
				"off_chain_ticker": "WOO-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "okx_ws",
				"off_chain_ticker": "WOO-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "mexc_ws",
				"off_chain_ticker": "WOOUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  }
			]
		  },
		  "XLM/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "XLM",
				"Quote": "USD"
			  },
			  "decimals": 10,
			  "min_provider_count": 3,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "binance_ws",
				"off_chain_ticker": "XLMUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "bybit_ws",
				"off_chain_ticker": "XLMUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "coinbase_ws",
				"off_chain_ticker": "XLM-USD"
			  },
			  {
				"name": "kraken_api",
				"off_chain_ticker": "XXLMZUSD"
			  },
			  {
				"name": "kucoin_ws",
				"off_chain_ticker": "XLM-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "mexc_ws",
				"off_chain_ticker": "XLMUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "okx_ws",
				"off_chain_ticker": "XLM-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "crypto_dot_com_ws",
				"off_chain_ticker": "XLM_USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  }
			]
		  },
		  "XRP/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "XRP",
				"Quote": "USD"
			  },
			  "decimals": 10,
			  "min_provider_count": 3,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "binance_ws",
				"off_chain_ticker": "XRPUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "bybit_ws",
				"off_chain_ticker": "XRPUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "coinbase_ws",
				"off_chain_ticker": "XRP-USD"
			  },
			  {
				"name": "gate_ws",
				"off_chain_ticker": "XRP_USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "huobi_ws",
				"off_chain_ticker": "xrpusdt",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "kraken_api",
				"off_chain_ticker": "XXRPZUSD"
			  },
			  {
				"name": "kucoin_ws",
				"off_chain_ticker": "XRP-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "mexc_ws",
				"off_chain_ticker": "XRPUSDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "okx_ws",
				"off_chain_ticker": "XRP-USDT",
				"normalize_by_pair": {
				  "Base": "USDT",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "crypto_dot_com_ws",
				"off_chain_ticker": "XRP_USD"
			  }
			]
		  }
		}
	}
	  `

	// UniswapV3BaseMarketMap is used to initialize the Base market map. This only includes
	// the markets that are supported by uniswapv3 on Base.
	UniswapV3BaseMarketMap mmtypes.MarketMap

	// UniswapV3BaseMarketMapJSON is the JSON representation of UniswapV3BaseMarketMap.
	UniswapV3BaseMarketMapJSON = `
	{
		"markets": {
		  "BRETT/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "BRETT",
				"Quote": "USD"
			  },
			  "decimals": 18,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "uniswapv3_api-base",
				"off_chain_ticker": "BRETT/ETH",
				"metadata_JSON": "{\"address\":\"0xBA3F945812a83471d709BCe9C3CA699A19FB46f7\",\"base_decimals\":18,\"quote_decimals\":18,\"invert\":true}",
				"normalize_by_pair": {
					"Base": "ETH",
					"Quote": "USD"
				}
			  }
			]
		  },
		  "ETH/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "ETH",
				"Quote": "USD"
			  },
			  "decimals": 18,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "uniswapv3_api-base",
				"off_chain_ticker": "ETH/USDT",
				"metadata_JSON": "{\"address\":\"0xd92E0767473D1E3FF11Ac036f2b1DB90aD0aE55F\",\"base_decimals\":18,\"quote_decimals\":6,\"invert\":false}",
				"normalize_by_pair": {
					"Base": "USDT",
					"Quote": "USD"
				}
			  }
			]
		  },
		  "USDT/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "USDT",
				"Quote": "USD"
			  },
			  "decimals": 6,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "uniswapv3_api-base",
				"off_chain_ticker": "USDT/ETH",
				"metadata_JSON": "{\"address\":\"0xd92E0767473D1E3FF11Ac036f2b1DB90aD0aE55F\",\"base_decimals\":6,\"quote_decimals\":18,\"invert\":true}",
				"normalize_by_pair": {
				  "Base": "ETH",
				  "Quote": "USD"
				}
			  },
			  {
				"name": "uniswapv3_api-base",
				"off_chain_ticker": "USDT/USDC",
				"metadata_JSON": "{\"address\":\"0xD56da2B74bA826f19015E6B7Dd9Dae1903E85DA1\",\"base_decimals\":6,\"quote_decimals\":6,\"invert\":true}"
			  }
			]
		  }
	    }
	}
	`

	// CoinGeckoMarketMap is used to initialize the CoinGecko market map. This only includes
	// the markets that are supported by CoinGecko & are included in the Core market map.
	CoinGeckoMarketMap mmtypes.MarketMap

	// CoinGeckoMarketMapJSON is the JSON representation of CoinGeckoMarketMap.
	CoinGeckoMarketMapJSON = `
	{
		"markets": {
		  "KHAI/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "KHAI",
				"Quote": "USD"
			  },
			  "decimals": 8,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "kitten-haimer/usd"
			  }
			]
		  },
		  "WAFFLES/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "WAFFLES",
				"Quote": "USD"
			  },
			  "decimals": 8,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "waffles/usd"
			  }
			]
		  },
		  "HEGE/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "HEGE",
				"Quote": "USD"
			  },
			  "decimals": 8,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "hege/usd"
			  }
			]
		  },
		  "WUF/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "WUF",
				"Quote": "USD"
			  },
			  "decimals": 8,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "wuffi/usd"
			  }
			]
		  },
		  "CHAT/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "CHAT",
				"Quote": "USD"
			  },
			  "decimals": 8,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "solchat/usd"
			  }
			]
		  },
		  "BEER/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "BEER",
				"Quote": "USD"
			  },
			  "decimals": 8,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "beercoin-2/usd"
			  }
			]
		  },
		  "MANEKI/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "MANEKI",
				"Quote": "USD"
			  },
			  "decimals": 11,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "maneki/usd"
			  }
			]
		  },
		  "SLERF/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "SLERF",
				"Quote": "USD"
			  },
			  "decimals": 8,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "slerf/usd"
			  }
			]
		  },
		  "MYRO/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "MYRO",
				"Quote": "USD"
			  },
			  "decimals": 8,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "myro/usd"
			  }
			]
		  },
		  "RAY/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "RAY",
				"Quote": "USD"
			  },
			  "decimals": 8,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "raydium/usd"
			  }
			]
		  },
		  "WIF/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "WIF",
				"Quote": "USD"
			  },
			  "decimals": 8,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "dogwifcoin/usd"
			  }
			]
		  },
		  "MICHI/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "MICHI",
				"Quote": "USD"
			  },
			  "decimals": 8,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "michicoin/usd"
			  }
			]
		  },
		  "MEW/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "MEW",
				"Quote": "USD"
			  },
			  "decimals": 8,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "cat-in-a-dogs-world/usd"
			  }
			]
		  },
		  "PONKE/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "PONKE",
				"Quote": "USD"
			  },
			  "decimals": 8,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "ponke/usd"
			  }
			]
		  },
		  "USA/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "USA",
				"Quote": "USD"
			  },
			  "decimals": 8,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "american-coin/usd"
			  }
			]
		  },
		  "BOME/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "BOME",
				"Quote": "USD"
			  },
			  "decimals": 8,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "book-of-meme/usd"
			  }
			]
		  },
		  "GME/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "GME",
				"Quote": "USD"
			  },
			  "decimals": 8,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "gme/usd"
			  }
			]
		  },
		  "DJT/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "DJT",
				"Quote": "USD"
			  },
			  "decimals": 8,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "trumpcoin/usd"
			  }
			]
		  },
		  "POPCAT/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "POPCAT",
				"Quote": "USD"
			  },
			  "decimals": 8,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "popcat/usd"
			  }
			]
		  },
		  "$RETIRE/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "$RETIRE",
				"Quote": "USD"
			  },
			  "decimals": 8,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "retire-on-sol/usd"
			  }
			]
		  },
		  "AAVE/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "AAVE",
				"Quote": "USD"
			  },
			  "decimals": 8,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "aave/usd"
			  }
			]
		  },
		  "ADA/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "ADA",
				"Quote": "USD"
			  },
			  "decimals": 10,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "cardano/usd"
			  }
			]
		  },
		  "AEVO/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "AEVO",
				"Quote": "USD"
			  },
			  "decimals": 9,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "aevo-exchange/usd"
			  }
			]
		  },
		  "ALGO/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "ALGO",
				"Quote": "USD"
			  },
			  "decimals": 10,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "algorand/usd"
			  }
			]
		  },
		  "APE/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "APE",
				"Quote": "USD"
			  },
			  "decimals": 9,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "apecoin/usd"
			  }
			]
		  },
		  "APT/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "APT",
				"Quote": "USD"
			  },
			  "decimals": 9,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "aptos/usd"
			  }
			]
		  },
		  "ARB/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "ARB",
				"Quote": "USD"
			  },
			  "decimals": 9,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "arbitrum/usd"
			  }
			]
		  },
		  "ARKM/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "ARKM",
				"Quote": "USD"
			  },
			  "decimals": 9,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "arkham/usd"
			  }
			]
		  },
		  "ASTR/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "ASTR",
				"Quote": "USD"
			  },
			  "decimals": 10,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "astar/usd"
			  }
			]
		  },
		  "ATOM/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "ATOM",
				"Quote": "USD"
			  },
			  "decimals": 9,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "cosmos/usd"
			  }
			]
		  },
		  "AVAX/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "AVAX",
				"Quote": "USD"
			  },
			  "decimals": 8,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "avalanche-2/usd"
			  }
			]
		  },
		  "AXL/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "AXL",
				"Quote": "USD"
			  },
			  "decimals": 9,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "axelar/usd"
			  }
			]
		  },
		  "BAZINGA/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "BAZINGA",
				"Quote": "USD"
			  },
			  "decimals": 11,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "bazinga-2/usd"
			  }
			]
		  },
		  "BCH/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "BCH",
				"Quote": "USD"
			  },
			  "decimals": 7,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "bitcoin-cash/usd"
			  }
			]
		  },
		  "BENDOG/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "BENDOG",
				"Quote": "USD"
			  },
			  "decimals": 11,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "ben-the-dog/usd"
			  }
			]
		  },
		  "BLUR/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "BLUR",
				"Quote": "USD"
			  },
			  "decimals": 10,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "blur/usd"
			  }
			]
		  },
		  "BNB/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "BNB",
				"Quote": "USD"
			  },
			  "decimals": 7,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "binancecoin/usd"
			  }
			]
		  },
		  "BONK/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "BONK",
				"Quote": "USD"
			  },
			  "decimals": 14,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "bonk/usd"
			  }
			]
		  },
		  "BTC/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "BTC",
				"Quote": "USD"
			  },
			  "decimals": 5,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "bitcoin/usd"
			  }
			]
		  },
		  "BUBBA/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "BUBBA",
				"Quote": "USD"
			  },
			  "decimals": 12,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "bubba/usd"
			  }
			]
		  },
		  "CATGPT/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "CATGPT",
				"Quote": "USD"
			  },
			  "decimals": 12,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "catgpt/usd"
			  }
			]
		  },
		  "CHEESE/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "CHEESE",
				"Quote": "USD"
			  },
			  "decimals": 11,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "cheese-2/usd"
			  }
			]
		  },
		  "CHUD/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "CHUD",
				"Quote": "USD"
			  },
			  "decimals": 12,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "chudjak/usd"
			  }
			]
		  },
		  "COK/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "COK",
				"Quote": "USD"
			  },
			  "decimals": 14,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "catownkimono/usd"
			  }
			]
		  },
		  "COST/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "COST",
				"Quote": "USD"
			  },
			  "decimals": 11,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "costco-hot-dog/usd"
			  }
			]
		  },
		  "COMP/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "COMP",
				"Quote": "USD"
			  },
			  "decimals": 8,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "compound-governance-token/usd"
			  }
			]
		  },
		  "CRV/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "CRV",
				"Quote": "USD"
			  },
			  "decimals": 10,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "curve-dao-token/usd"
			  }
			]
		  },
		  "DOGE/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "DOGE",
				"Quote": "USD"
			  },
			  "decimals": 11,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "dogecoin/usd"
			  }
			]
		  },
		  "DOT/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "DOT",
				"Quote": "USD"
			  },
			  "decimals": 9,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "polkadot/usd"
			  }
			]
		  },
		  "DEVIN/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "DEVIN",
				"Quote": "USD"
			  },
			  "decimals": 11,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "devin-on-solana/usd"
			  }
			]
		  },
		  "DUKO/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "DUKO",
				"Quote": "USD"
			  },
			  "decimals": 12,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "duko/usd"
			  }
			]
		  },
		  "DYDX/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "DYDX",
				"Quote": "USD"
			  },
			  "decimals": 9,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "dydx/usd"
			  }
			]
		  },
		  "DYM/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "DYM",
				"Quote": "USD"
			  },
			  "decimals": 9,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "dymension/usd"
			  }
			]
		  },
		  "EGG/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "EGG",
				"Quote": "USD"
			  },
			  "decimals": 11,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "justanegg-2/usd"
			  }
			]
		  },
		  "FALX/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "FALX",
				"Quote": "USD"
			  },
			  "decimals": 12,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "falx/usd"
			  }
			]
		  },
		  "GOL/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "GOL",
				"Quote": "USD"
			  },
			  "decimals": 11,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "golazo-world/usd"
			  }
			]
		  },
		  "GUMMY/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "GUMMY",
				"Quote": "USD"
			  },
			  "decimals": 11,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "gummy/usd"
			  }
			]
		  },
		  "HABIBI/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "HABIBI",
				"Quote": "USD"
			  },
			  "decimals": 11,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "habibi-sol/usd"
			  }
			]
		  },
		  "HAMMY/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "HAMMY",
				"Quote": "USD"
			  },
			  "decimals": 11,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "sad-hamster/usd"
			  }
			]
		  },
		  "HARAMBE/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "HARAMBE",
				"Quote": "USD"
			  },
			  "decimals": 11,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "harambe-2/usd"
			  }
			]
		  },
		  "KITTY/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "KITTY",
				"Quote": "USD"
			  },
			  "decimals": 11,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "roaring-kitty-solana/usd"
			  }
			]
		  },
		  "MUMU/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "MUMU",
				"Quote": "USD"
			  },
			  "decimals": 14,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "mumu-the-bull-3/usd"
			  }
			]
		  },
		  "NUB/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "NUB",
				"Quote": "USD"
			  },
			  "decimals": 11,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "sillynubcat/usd"
			  }
			]
		  },
		  "PAJAMAS/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "PAJAMAS",
				"Quote": "USD"
			  },
			  "decimals": 11,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "pajamas-cat/usd"
			  }
			]
		  },
		  "PENG/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "PENG",
				"Quote": "USD"
			  },
			  "decimals": 10,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "peng/usd"
			  }
			]
		  },
		  "PUNDU/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "PUNDU",
				"Quote": "USD"
			  },
			  "decimals": 11,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "pundu/usd"
			  }
			]
		  },
		  "RETARDIO/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "RETARDIO",
				"Quote": "USD"
			  },
			  "decimals": 11,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "retardio/usd"
			  }
			]
		  },
		  "SELFIE/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "SELFIE",
				"Quote": "USD"
			  },
			  "decimals": 11,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "selfiedogcoin/usd"
			  }
			]
		  },
		  "SLOTH/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "SLOTH",
				"Quote": "USD"
			  },
			  "decimals": 11,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "slothana/usd"
			  }
			]
		  },
		  "SPEED/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "SPEED",
				"Quote": "USD"
			  },
			  "decimals": 11,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "real-fast/usd"
			  }
			]
		  },
		  "SPIKE/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "SPIKE",
				"Quote": "USD"
			  },
			  "decimals": 11,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "spike/usd"
			  }
			]
		  },
		  "STACKS/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "STACKS",
				"Quote": "USD"
			  },
			  "decimals": 11,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "slap-city/usd"
			  }
			]
		  },
		  "TREMP/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "TREMP",
				"Quote": "USD"
			  },
			  "decimals": 11,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "donald-tremp/usd"
			  }
			]
		  },
		  "MOG/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "MOG",
				"Quote": "USD"
			  },
			  "decimals": 11,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "mog-coin/usd"
			  }
			]
		  },
		  "TRUMP/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "TRUMP",
				"Quote": "USD"
			  },
			  "decimals": 11,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "maga/usd"
			  }
			]
		  },
		  "MOTHER/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "MOTHER",
				"Quote": "USD"
			  },
			  "decimals": 11,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "mother-iggy/usd"
			  }
			]
		  },
		  "EOS/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "EOS",
				"Quote": "USD"
			  },
			  "decimals": 10,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "eos/usd"
			  }
			]
		  },
		  "ETC/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "ETC",
				"Quote": "USD"
			  },
			  "decimals": 8,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "ethereum-classic/usd"
			  }
			]
		  },
		  "ETH/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "ETH",
				"Quote": "USD"
			  },
			  "decimals": 6,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "ethereum/usd"
			  }
			]
		  },
		  "FET/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "FET",
				"Quote": "USD"
			  },
			  "decimals": 10,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "fetch-ai/usd"
			  }
			]
		  },
		  "FIL/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "FIL",
				"Quote": "USD"
			  },
			  "decimals": 9,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "filecoin/usd"
			  }
			]
		  },
		  "GRT/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "GRT",
				"Quote": "USD"
			  },
			  "decimals": 10,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "the-graph/usd"
			  }
			]
		  },
		  "HBAR/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "HBAR",
				"Quote": "USD"
			  },
			  "decimals": 10,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "hedera-hashgraph/usd"
			  }
			]
		  },
		  "ICP/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "ICP",
				"Quote": "USD"
			  },
			  "decimals": 8,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "internet-computer/usd"
			  }
			]
		  },
		  "IMX/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "IMX",
				"Quote": "USD"
			  },
			  "decimals": 9,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "immutable-x/usd"
			  }
			]
		  },
		  "INJ/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "INJ",
				"Quote": "USD"
			  },
			  "decimals": 8,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "injective-protocol/usd"
			  }
			]
		  },
		  "JTO/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "JTO",
				"Quote": "USD"
			  },
			  "decimals": 9,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "jito-governance-token/usd"
			  }
			]
		  },
		  "JUP/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "JUP",
				"Quote": "USD"
			  },
			  "decimals": 10,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "jupiter-exchange-solana/usd"
			  }
			]
		  },
		  "LDO/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "LDO",
				"Quote": "USD"
			  },
			  "decimals": 9,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "lido-dao/usd"
			  }
			]
		  },
		  "LINK/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "LINK",
				"Quote": "USD"
			  },
			  "decimals": 9,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "chainlink/usd"
			  }
			]
		  },
		  "LTC/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "LTC",
				"Quote": "USD"
			  },
			  "decimals": 8,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "litecoin/usd"
			  }
			]
		  },
		  "MANA/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "MANA",
				"Quote": "USD"
			  },
			  "decimals": 10,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "decentraland/usd"
			  }
			]
		  },
		  "MATIC/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "MATIC",
				"Quote": "USD"
			  },
			  "decimals": 10,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "matic-network/usd"
			  }
			]
		  },
		  "MKR/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "MKR",
				"Quote": "USD"
			  },
			  "decimals": 6,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "maker/usd"
			  }
			]
		  },
		  "NEAR/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "NEAR",
				"Quote": "USD"
			  },
			  "decimals": 9,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "near/usd"
			  }
			]
		  },
		  "NTRN/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "NTRN",
				"Quote": "USD"
			  },
			  "decimals": 8,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "neutron-3/usd"
			  }
			]
		  },
		  "OP/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "OP",
				"Quote": "USD"
			  },
			  "decimals": 9,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "optimism/usd"
			  }
			]
		  },
		  "ORDI/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "ORDI",
				"Quote": "USD"
			  },
			  "decimals": 8,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "ordinals/usd"
			  }
			]
		  },
		  "PEPE/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "PEPE",
				"Quote": "USD"
			  },
			  "decimals": 16,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "pepe/usd"
			  }
			]
		  },
		  "PYTH/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "PYTH",
				"Quote": "USD"
			  },
			  "decimals": 10,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "pyth-network/usd"
			  }
			]
		  },
		  "RUNE/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "RUNE",
				"Quote": "USD"
			  },
			  "decimals": 9,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "thorchain/usd"
			  }
			]
		  },
		  "SEI/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "SEI",
				"Quote": "USD"
			  },
			  "decimals": 10,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "sei-network/usd"
			  }
			]
		  },
		  "SHIB/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "SHIB",
				"Quote": "USD"
			  },
			  "decimals": 15,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "shiba-inu/usd"
			  }
			]
		  },
		  "SNX/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "SNX",
				"Quote": "USD"
			  },
			  "decimals": 9,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "havven/usd"
			  }
			]
		  },
		  "SOL/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "SOL",
				"Quote": "USD"
			  },
			  "decimals": 8,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "solana/usd"
			  }
			]
		  },
		  "STRK/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "STRK",
				"Quote": "USD"
			  },
			  "decimals": 9,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "starknet/usd"
			  }
			]
		  },
		  "STX/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "STX",
				"Quote": "USD"
			  },
			  "decimals": 9,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "blockstack/usd"
			  }
			]
		  },
		  "SUI/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "SUI",
				"Quote": "USD"
			  },
			  "decimals": 10,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "sui/usd"
			  }
			]
		  },
		  "TIA/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "TIA",
				"Quote": "USD"
			  },
			  "decimals": 8,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "celestia/usd"
			  }
			]
		  },
		  "TRX/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "TRX",
				"Quote": "USD"
			  },
			  "decimals": 11,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "tron/usd"
			  }
			]
		  },
		  "UNI/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "UNI",
				"Quote": "USD"
			  },
			  "decimals": 9,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "uniswap/usd"
			  }
			]
		  },
		  "USDT/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "USDT",
				"Quote": "USD"
			  },
			  "decimals": 9,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "tether/usd"
			  }
			]
		  },
		  "WLD/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "WLD",
				"Quote": "USD"
			  },
			  "decimals": 9,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "worldcoin-wld/usd"
			  }
			]
		  },
		  "WOO/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "WOO",
				"Quote": "USD"
			  },
			  "decimals": 10,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "woo-network/usd"
			  }
			]
		  },
		  "XLM/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "XLM",
				"Quote": "USD"
			  },
			  "decimals": 10,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "stellar/usd"
			  }
			]
		  },
		  "XRP/USD": {
			"ticker": {
			  "currency_pair": {
				"Base": "XRP",
				"Quote": "USD"
			  },
			  "decimals": 10,
			  "min_provider_count": 1,
			  "enabled": true
			},
			"provider_configs": [
			  {
				"name": "coingecko_api",
				"off_chain_ticker": "ripple/usd"
			  }
			]
		  }
		}
	  } 
	`

	// OsmosisMarketMap is used to initialize the osmosis market map. This only includes
	// the markets that are supported by osmosis.
	OsmosisMarketMap mmtypes.MarketMap

	// OsmosisMarketMapJSON is the JSON representation of OsmosisMarketMap.
	OsmosisMarketMapJSON = `
{
    "markets": {
        "STARS/USD": {
            "ticker": {
                "currency_pair": {
                    "Base": "STARS",
                    "Quote": "USD"
                },
                "decimals": 18,
                "min_provider_count": 1,
                "enabled": true,
                "metadata_JSON": "{\"reference_price\":1,\"liquidity\":0,\"aggregate_ids\":[]}"
            },
            "provider_configs": [
                {
                    "name": "osmosis_api",
                    "off_chain_ticker": "STARS/OSMO",
                    "metadata_JSON": "{\"pool_id\":1096,\"base_token_denom\":\"ibc/987C17B11ABC2B20019178ACE62929FE9840202CE79498E29FE8E5CB02B7C0A4\",\"quote_token_denom\":\"uosmo\"}",
                    "normalize_by_pair": {
                        "Base": "OSMO",
                        "Quote": "USD"
                    }
                }
            ]
        },
        "USDT/USD": {
            "ticker": {
                "currency_pair": {
                    "Base": "USDT",
                    "Quote": "USD"
                },
                "decimals": 9,
                "min_provider_count": 1,
                "enabled": true
            },
            "provider_configs": [
                {
                    "name": "binance_ws",
                    "off_chain_ticker": "USDCUSDT",
                    "invert": true
                },
                {
                    "name": "bybit_ws",
                    "off_chain_ticker": "USDCUSDT",
                    "invert": true
                },
                {
                    "name": "coinbase_ws",
                    "off_chain_ticker": "USDT-USD"
                },
                {
                    "name": "kraken_api",
                    "off_chain_ticker": "USDTZUSD"
                },
                {
                    "name": "okx_ws",
                    "off_chain_ticker": "USDC-USDT",
                    "invert": true
                },
                {
                    "name": "crypto_dot_com_ws",
                    "off_chain_ticker": "USDT_USD"
                }
            ]
        },
        "OSMO/USD": {
            "ticker": {
                "currency_pair": {
                    "Base": "OSMO",
                    "Quote": "USD"
                },
                "decimals": 8,
                "min_provider_count": 1,
                "enabled": true,
                "metadata_JSON": "{\"reference_price\":1,\"liquidity\":0,\"aggregate_ids\":[]}"
            },
            "provider_configs": [
                {
                    "name": "coinbase_ws",
                    "off_chain_ticker": "OSMO-USD"
                },
                {
                    "name": "huobi_ws",
                    "off_chain_ticker": "osmousdt",
                    "normalize_by_pair": {
                        "Base": "USDT",
                        "Quote": "USD"
                    }
                },
                {
                    "name": "binance_api",
                    "off_chain_ticker": "OSMOUSDT",
                    "normalize_by_pair": {
                        "Base": "USDT",
                        "Quote": "USD"
                    }
                }
            ]
        }
    }
}
	`

	// PolymarketMarketMap is used to initialize the Polymarket market map. This only includes one prediction market
	// with one outcome token.
	PolymarketMarketMap mmtypes.MarketMap

	// PolymarketMarketMapJSON is the JSON representation of PolymarketMarketMap.
	PolymarketMarketMapJSON = ` 
{
   "markets":{
      "WILL_BERNIE_SANDERS_WIN_THE_2024_US_PRESIDENTIAL_ELECTION?YES/USD":{
         "ticker":{
            "currency_pair":{
               "Base":"WILL_BERNIE_SANDERS_WIN_THE_2024_US_PRESIDENTIAL_ELECTION?YES",
               "Quote":"USD"
            },
            "decimals":4,
            "min_provider_count":1,
            "enabled":true
         },
         "provider_configs":[
            {
               "name":"polymarket_api",
               "off_chain_ticker":"0x08f5fe8d0d29c08a96f0bc3dfb52f50e0caf470d94d133d95d38fa6c847e0925/95128817762909535143571435260705470642391662537976312011260538371392879420759"
            }
         ]
      },
      "WILL_INSIDE_OUT_2_GROSS_MOST_IN_2024?YES/USD":{
         "ticker":{
            "currency_pair":{
               "Base":"WILL_INSIDE_OUT_2_GROSS_MOST_IN_2024?YES",
               "Quote":"USD"
            },
            "decimals":4,
            "min_provider_count":1,
            "enabled":true
         },
         "provider_configs":[
            {
               "name":"polymarket_api",
               "off_chain_ticker":"0x1ab07117f9f698f28490f57754d6fe5309374230c95867a7eba572892a11d710/50107902083284751016545440401692219408556171231461347396738260657226842527986"
            }
         ]
      }
   }
}`

	UniswapV3EthMarketMap mmtypes.MarketMap

	UniswapV3EthMarketMapJSON = `
{
  "markets": {
    "AERO/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "AERO",
          "Quote": "USD"
        },
        "decimals": 10,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"29270\"}]}"
      },
      "provider_configs": [
        {
          "name": "coinbase_ws",
          "off_chain_ticker": "AERO-USD"
        },
        {
          "name": "uniswapv3_api-base",
          "off_chain_ticker": "AERO,UNISWAP-V3-BASE,0X940181A94A35A4569E4529A3CDFB74E38FD98631/WETH,UNISWAP-V3-BASE,0X4200000000000000000000000000000000000006",
          "normalize_by_pair": {
            "Base": "WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"address\":\"0x3d5d143381916280ff91407febeb52f2b60f33cf\",\"base_decimals\":18,\"quote_decimals\":18,\"invert\":false}"
        }
      ]
    },
    "ANDY,UNISWAP_V3,0X68BBED6A47194EFF1CF514B50EA91895597FC91E/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "ANDY,UNISWAP_V3,0X68BBED6A47194EFF1CF514B50EA91895597FC91E",
          "Quote": "USD"
        },
        "decimals": 14,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"29879\"}]}"
      },
      "provider_configs": [
        {
          "name": "uniswapv3_api-ethereum",
          "off_chain_ticker": "ANDY,UNISWAP_V3,0X68BBED6A47194EFF1CF514B50EA91895597FC91E/WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
          "normalize_by_pair": {
            "Base": "WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"address\":\"0x2b8574482d62d8df670dfd3be15f2f092941284e\",\"base_decimals\":18,\"quote_decimals\":18,\"invert\":false}"
        }
      ]
    },
    "APE/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "APE",
          "Quote": "USD"
        },
        "decimals": 10,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"18876\"}]}"
      },
      "provider_configs": [
        {
          "name": "uniswapv3_api-ethereum",
          "off_chain_ticker": "APE,UNISWAP_V3,0X4D224452801ACED8B2F0AEBE155379BB5D594381/WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
          "normalize_by_pair": {
            "Base": "WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"address\":\"0xac4b3dacb91461209ae9d41ec517c2b9cb1b7daf\",\"base_decimals\":18,\"quote_decimals\":18,\"invert\":false}"
        },
        {
          "name": "crypto_dot_com_ws",
          "off_chain_ticker": "APE_USDT",
          "normalize_by_pair": {
            "Base": "USDT",
            "Quote": "USD"
          }
        },
        {
          "name": "coinbase_ws",
          "off_chain_ticker": "APE-USD"
        }
      ]
    },
    "ATH,UNISWAP_V3,0XBE0ED4138121ECFC5C0E56B40517DA27E6C5226B/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "ATH,UNISWAP_V3,0XBE0ED4138121ECFC5C0E56B40517DA27E6C5226B",
          "Quote": "USD"
        },
        "decimals": 11,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"30083\"}]}"
      },
      "provider_configs": [
        {
          "name": "uniswapv3_api-ethereum",
          "off_chain_ticker": "ATH,UNISWAP_V3,0XBE0ED4138121ECFC5C0E56B40517DA27E6C5226B/WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
          "normalize_by_pair": {
            "Base": "WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"address\":\"0xd31d41dffa3589bb0c0183e46a1eed983a5e5978\",\"base_decimals\":18,\"quote_decimals\":18,\"invert\":false}"
        }
      ]
    },
    "BGUY,UNISWAP-V3-BASE,0X8931EE05EC111325C1700B68E5EF7B887E00661D/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "BGUY,UNISWAP-V3-BASE,0X8931EE05EC111325C1700B68E5EF7B887E00661D",
          "Quote": "USD"
        },
        "decimals": 14,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"31984\"}]}"
      },
      "provider_configs": [
        {
          "name": "uniswapv3_api-base",
          "off_chain_ticker": "BGUY,UNISWAP-V3-BASE,0X8931EE05EC111325C1700B68E5EF7B887E00661D/WETH,UNISWAP-V3-BASE,0X4200000000000000000000000000000000000006",
          "normalize_by_pair": {
            "Base": "WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"address\":\"0xd209a26cb9d242f41839f96952879771a72d284b\",\"base_decimals\":9,\"quote_decimals\":18,\"invert\":false}"
        }
      ]
    },
    "BOOE,UNISWAP_V3,0X289FF00235D2B98B0145FF5D4435D3E92F9540A6/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "BOOE,UNISWAP_V3,0X289FF00235D2B98B0145FF5D4435D3E92F9540A6",
          "Quote": "USD"
        },
        "decimals": 11,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"31171\"}]}"
      },
      "provider_configs": [
        {
          "name": "uniswapv3_api-ethereum",
          "off_chain_ticker": "BOOE,UNISWAP_V3,0X289FF00235D2B98B0145FF5D4435D3E92F9540A6/WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
          "normalize_by_pair": {
            "Base": "WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"address\":\"0xf0230b6f9604e3edc13e125806e4c5446c34dfdf\",\"base_decimals\":18,\"quote_decimals\":18,\"invert\":false}"
        }
      ]
    },
    "BRETT,UNISWAP-V3-BASE,0X532F27101965DD16442E59D40670FAF5EBB142E4/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "BRETT,UNISWAP-V3-BASE,0X532F27101965DD16442E59D40670FAF5EBB142E4",
          "Quote": "USD"
        },
        "decimals": 11,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"29743\"}]}"
      },
      "provider_configs": [
        {
          "name": "uniswapv3_api-base",
          "off_chain_ticker": "BRETT,UNISWAP-V3-BASE,0X532F27101965DD16442E59D40670FAF5EBB142E4/WETH,UNISWAP-V3-BASE,0X4200000000000000000000000000000000000006",
          "normalize_by_pair": {
            "Base": "WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"address\":\"0x36a46dff597c5a444bbc521d26787f57867d2214\",\"base_decimals\":18,\"quote_decimals\":18,\"invert\":false}"
        }
      ]
    },
    "BYTES,UNISWAP_V3,0XA19F5264F7D7BE11C451C093D8F92592820BEA86/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "BYTES,UNISWAP_V3,0XA19F5264F7D7BE11C451C093D8F92592820BEA86",
          "Quote": "USD"
        },
        "decimals": 9,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"28297\"}]}"
      },
      "provider_configs": [
        {
          "name": "uniswapv3_api-ethereum",
          "off_chain_ticker": "BYTES,UNISWAP_V3,0XA19F5264F7D7BE11C451C093D8F92592820BEA86/WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
          "normalize_by_pair": {
            "Base": "WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"address\":\"0xfeb09c7e130a4b87b27ebd648ec485657b688b34\",\"base_decimals\":18,\"quote_decimals\":18,\"invert\":false}"
        }
      ]
    },
    "CBETH/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "CBETH",
          "Quote": "USD"
        },
        "decimals": 6,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"21535\"}]}"
      },
      "provider_configs": [
        {
          "name": "uniswapv3_api-base",
          "off_chain_ticker": "CBETH,UNISWAP-V3-BASE,0X2AE3F1EC7F1F5012CFEAB0185BFC7AA3CF0DEC22/WETH,UNISWAP-V3-BASE,0X4200000000000000000000000000000000000006",
          "normalize_by_pair": {
            "Base": "WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"address\":\"0x10648ba41b8565907cfa1496765fa4d95390aa0d\",\"base_decimals\":18,\"quote_decimals\":18,\"invert\":false}"
        },
        {
          "name": "uniswapv3_api-ethereum",
          "off_chain_ticker": "CBETH,UNISWAP_V3,0XBE9895146F7AF43049CA1C1AE358B0541EA49704/WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
          "normalize_by_pair": {
            "Base": "WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"address\":\"0x177622e79acece98c39f6e12fa78ac7fc8a8bf62\",\"base_decimals\":18,\"quote_decimals\":18,\"invert\":false}"
        }
      ]
    },
    "DAI/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "DAI",
          "Quote": "USD"
        },
        "decimals": 10,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"4943\"}]}"
      },
      "provider_configs": [
        {
          "name": "coinbase_ws",
          "off_chain_ticker": "DAI-USD"
        },
        {
          "name": "uniswapv3_api-ethereum",
          "off_chain_ticker": "DAI,UNISWAP_V3,0X6B175474E89094C44DA98B954EEDEAC495271D0F/USDT,UNISWAP_V3,0XDAC17F958D2EE523A2206206994597C13D831EC7",
          "normalize_by_pair": {
            "Base": "USDT",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"address\":\"0x48da0965ab2d2cbf1c17c09cfb5cbe67ad5b1406\",\"base_decimals\":18,\"quote_decimals\":6,\"invert\":false}"
        }
      ]
    },
    "DEVVE,UNISWAP_V3,0X8248270620AA532E4D64316017BE5E873E37CC09/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "DEVVE,UNISWAP_V3,0X8248270620AA532E4D64316017BE5E873E37CC09",
          "Quote": "USD"
        },
        "decimals": 10,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"29461\"}]}"
      },
      "provider_configs": [
        {
          "name": "uniswapv3_api-ethereum",
          "off_chain_ticker": "DEVVE,UNISWAP_V3,0X8248270620AA532E4D64316017BE5E873E37CC09/WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
          "normalize_by_pair": {
            "Base": "WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"address\":\"0x18bbe20f81bdcb340325e28a6ee6bb426b7ccbc1\",\"base_decimals\":18,\"quote_decimals\":18,\"invert\":false}"
        }
      ]
    },
    "DSYNC,UNISWAP_V3,0XF94E7D0710709388BCE3161C32B4EEA56D3F91CC/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "DSYNC,UNISWAP_V3,0XF94E7D0710709388BCE3161C32B4EEA56D3F91CC",
          "Quote": "USD"
        },
        "decimals": 10,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"29884\"}]}"
      },
      "provider_configs": [
        {
          "name": "uniswapv3_api-ethereum",
          "off_chain_ticker": "DSYNC,UNISWAP_V3,0XF94E7D0710709388BCE3161C32B4EEA56D3F91CC/WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
          "normalize_by_pair": {
            "Base": "WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"address\":\"0xbe24f9952bf2a15bc4d2661151049d9588d6c0cb\",\"base_decimals\":18,\"quote_decimals\":18,\"invert\":false}"
        }
      ]
    },
    "FACTR,UNISWAP_V3,0XE0BCEEF36F3A6EFDD5EEBFACD591423F8549B9D5/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "FACTR,UNISWAP_V3,0XE0BCEEF36F3A6EFDD5EEBFACD591423F8549B9D5",
          "Quote": "USD"
        },
        "decimals": 11,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"13255\"}]}"
      },
      "provider_configs": [
        {
          "name": "uniswapv3_api-ethereum",
          "off_chain_ticker": "FACTR,UNISWAP_V3,0XE0BCEEF36F3A6EFDD5EEBFACD591423F8549B9D5/WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
          "normalize_by_pair": {
            "Base": "WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"address\":\"0xcab82c42b5195624671b92b1b0d91adf171f103c\",\"base_decimals\":18,\"quote_decimals\":18,\"invert\":false}"
        }
      ]
    },
    "FLT,UNISWAP_V3,0X236501327E701692A281934230AF0B6BE8DF3353/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "FLT,UNISWAP_V3,0X236501327E701692A281934230AF0B6BE8DF3353",
          "Quote": "USD"
        },
        "decimals": 10,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"30097\"}]}"
      },
      "provider_configs": [
        {
          "name": "uniswapv3_api-ethereum",
          "off_chain_ticker": "FLT,UNISWAP_V3,0X236501327E701692A281934230AF0B6BE8DF3353/WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
          "normalize_by_pair": {
            "Base": "WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"address\":\"0x983b92ea1f0a20844466f3e4bb988c1de145293e\",\"base_decimals\":18,\"quote_decimals\":18,\"invert\":false}"
        }
      ]
    },
    "FTM/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "FTM",
          "Quote": "USD"
        },
        "decimals": 10,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"3513\"}]}"
      },
      "provider_configs": [
        {
          "name": "bitfinex_ws",
          "off_chain_ticker": "FTMUSD"
        },
        {
          "name": "crypto_dot_com_ws",
          "off_chain_ticker": "FTM_USDT",
          "normalize_by_pair": {
            "Base": "USDT",
            "Quote": "USD"
          }
        },
        {
          "name": "uniswapv3_api-ethereum",
          "off_chain_ticker": "FTM,UNISWAP_V3,0X4E15361FD6B4BB609FA63C81A2BE19D873717870/WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
          "normalize_by_pair": {
            "Base": "WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"address\":\"0x3b685307c8611afb2a9e83ebc8743dc20480716e\",\"base_decimals\":18,\"quote_decimals\":18,\"invert\":false}"
        }
      ]
    },
    "FUN,UNISWAP_V3,0X419D0D8BDD9AF5E606AE2232ED285AFF190E711B/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "FUN,UNISWAP_V3,0X419D0D8BDD9AF5E606AE2232ED285AFF190E711B",
          "Quote": "USD"
        },
        "decimals": 12,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"1757\"}]}"
      },
      "provider_configs": [
        {
          "name": "uniswapv3_api-ethereum",
          "off_chain_ticker": "FUN,UNISWAP_V3,0X419D0D8BDD9AF5E606AE2232ED285AFF190E711B/USDT,UNISWAP_V3,0XDAC17F958D2EE523A2206206994597C13D831EC7",
          "normalize_by_pair": {
            "Base": "USDT",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"address\":\"0x1a349a3397a8431eed8d94a05f88f9001117fcaa\",\"base_decimals\":8,\"quote_decimals\":6,\"invert\":false}"
        }
      ]
    },
    "HEX,UNISWAP_V3,0X2B591E99AFE9F32EAA6214F7B7629768C40EEB39/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "HEX,UNISWAP_V3,0X2B591E99AFE9F32EAA6214F7B7629768C40EEB39",
          "Quote": "USD"
        },
        "decimals": 12,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"28928\"}]}"
      },
      "provider_configs": [
        {
          "name": "uniswapv3_api-ethereum",
          "off_chain_ticker": "HEX,UNISWAP_V3,0X2B591E99AFE9F32EAA6214F7B7629768C40EEB39/WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
          "normalize_by_pair": {
            "Base": "WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"address\":\"0x9e0905249ceefffb9605e034b534544684a58be6\",\"base_decimals\":8,\"quote_decimals\":18,\"invert\":false}"
        }
      ]
    },
    "INJ/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "INJ",
          "Quote": "USD"
        },
        "decimals": 8,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"7226\"}]}"
      },
      "provider_configs": [
        {
          "name": "coinbase_ws",
          "off_chain_ticker": "INJ-USD"
        },
        {
          "name": "uniswapv3_api-ethereum",
          "off_chain_ticker": "INJ,UNISWAP_V3,0XE28B3B32B6C345A34FF64674606124DD5ACECA30/WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
          "normalize_by_pair": {
            "Base": "WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"address\":\"0x6c063a6e8cd45869b5eb75291e65a3de298f3aa8\",\"base_decimals\":18,\"quote_decimals\":18,\"invert\":false}"
        }
      ]
    },
    "KIZUNA,UNISWAP_V3,0X470C8950C0C3AA4B09654BC73B004615119A44B5/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "KIZUNA,UNISWAP_V3,0X470C8950C0C3AA4B09654BC73B004615119A44B5",
          "Quote": "USD"
        },
        "decimals": 18,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"28217\"}]}"
      },
      "provider_configs": [
        {
          "name": "uniswapv3_api-ethereum",
          "off_chain_ticker": "KIZUNA,UNISWAP_V3,0X470C8950C0C3AA4B09654BC73B004615119A44B5/WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
          "normalize_by_pair": {
            "Base": "WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"address\":\"0xb7f27e5ebf97d88f37e16eddecc59523361a60e1\",\"base_decimals\":18,\"quote_decimals\":18,\"invert\":false}"
        }
      ]
    },
    "LDO/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "LDO",
          "Quote": "USD"
        },
        "decimals": 10,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"8000\"}]}"
      },
      "provider_configs": [
        {
          "name": "uniswapv3_api-ethereum",
          "off_chain_ticker": "LDO,UNISWAP_V3,0X5A98FCBEA516CF06857215779FD812CA3BEF1B32/WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
          "normalize_by_pair": {
            "Base": "WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"address\":\"0xa3f558aebaecaf0e11ca4b2199cc5ed341edfd74\",\"base_decimals\":18,\"quote_decimals\":18,\"invert\":false}"
        },
        {
          "name": "coinbase_ws",
          "off_chain_ticker": "LDO-USD"
        }
      ]
    },
    "LINK/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "LINK",
          "Quote": "USD"
        },
        "decimals": 8,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"1975\"}]}"
      },
      "provider_configs": [
        {
          "name": "crypto_dot_com_ws",
          "off_chain_ticker": "LINK_USDT",
          "normalize_by_pair": {
            "Base": "USDT",
            "Quote": "USD"
          }
        },
        {
          "name": "bitstamp_api",
          "off_chain_ticker": "linkusd"
        },
        {
          "name": "coinbase_ws",
          "off_chain_ticker": "LINK-USD"
        },
        {
          "name": "uniswapv3_api-ethereum",
          "off_chain_ticker": "LINK,UNISWAP_V3,0X514910771AF9CA656AF840DFF83E8264ECF986CA/WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
          "normalize_by_pair": {
            "Base": "WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"address\":\"0xa6cc3c2531fdaa6ae1a3ca84c2855806728693e8\",\"base_decimals\":18,\"quote_decimals\":18,\"invert\":false}"
        }
      ]
    },
    "MAGA,UNISWAP_V3,0XD29DA236DD4AAC627346E1BBA06A619E8C22D7C5/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "MAGA,UNISWAP_V3,0XD29DA236DD4AAC627346E1BBA06A619E8C22D7C5",
          "Quote": "USD"
        },
        "decimals": 14,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"31305\"}]}"
      },
      "provider_configs": [
        {
          "name": "uniswapv3_api-ethereum",
          "off_chain_ticker": "MAGA,UNISWAP_V3,0XD29DA236DD4AAC627346E1BBA06A619E8C22D7C5/WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
          "normalize_by_pair": {
            "Base": "WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"address\":\"0x961ec3bb28c9e98a040c4bded38917aa96b791be\",\"base_decimals\":9,\"quote_decimals\":18,\"invert\":false}"
        }
      ]
    },
    "MKR/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "MKR",
          "Quote": "USD"
        },
        "decimals": 6,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"1518\"}]}"
      },
      "provider_configs": [
        {
          "name": "coinbase_ws",
          "off_chain_ticker": "MKR-USD"
        },
        {
          "name": "uniswapv3_api-ethereum",
          "off_chain_ticker": "MKR,UNISWAP_V3,0X9F8F72AA9304C8B593D555F12EF6589CC3A579A2/WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
          "normalize_by_pair": {
            "Base": "WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"address\":\"0xe8c6c9227491c0a8156a0106a0204d881bb7e531\",\"base_decimals\":18,\"quote_decimals\":18,\"invert\":false}"
        }
      ]
    },
    "MYRIA,UNISWAP_V3,0XA0EF786BF476FE0810408CABA05E536AC800FF86/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "MYRIA,UNISWAP_V3,0XA0EF786BF476FE0810408CABA05E536AC800FF86",
          "Quote": "USD"
        },
        "decimals": 12,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"22289\"}]}"
      },
      "provider_configs": [
        {
          "name": "uniswapv3_api-ethereum",
          "off_chain_ticker": "MYRIA,UNISWAP_V3,0XA0EF786BF476FE0810408CABA05E536AC800FF86/WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
          "normalize_by_pair": {
            "Base": "WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"address\":\"0xbf85f94d3233ee588f0907a9147fbb59d7246b54\",\"base_decimals\":18,\"quote_decimals\":18,\"invert\":false}"
        }
      ]
    },
    "NEIRO,UNISWAP_V3,0X812BA41E071C7B7FA4EBCFB62DF5F45F6FA853EE/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "NEIRO,UNISWAP_V3,0X812BA41E071C7B7FA4EBCFB62DF5F45F6FA853EE",
          "Quote": "USD"
        },
        "decimals": 14,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"32521\"}]}"
      },
      "provider_configs": [
        {
          "name": "uniswapv3_api-ethereum",
          "off_chain_ticker": "NEIRO,UNISWAP_V3,0X812BA41E071C7B7FA4EBCFB62DF5F45F6FA853EE/WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
          "normalize_by_pair": {
            "Base": "WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"address\":\"0x79a6683d82f25535ff3fd2753e03e0961060e882\",\"base_decimals\":9,\"quote_decimals\":18,\"invert\":false}"
        }
      ]
    },
    "NEXO,UNISWAP_V3,0XB62132E35A6C13EE1EE0F84DC5D40BAD8D815206/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "NEXO,UNISWAP_V3,0XB62132E35A6C13EE1EE0F84DC5D40BAD8D815206",
          "Quote": "USD"
        },
        "decimals": 10,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"2694\"}]}"
      },
      "provider_configs": [
        {
          "name": "uniswapv3_api-ethereum",
          "off_chain_ticker": "NEXO,UNISWAP_V3,0XB62132E35A6C13EE1EE0F84DC5D40BAD8D815206/WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
          "normalize_by_pair": {
            "Base": "WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"address\":\"0x4c54ff7f1c424ff5487a32aad0b48b19cbaf087f\",\"base_decimals\":18,\"quote_decimals\":18,\"invert\":false}"
        }
      ]
    },
    "NFAI,UNISWAP_V3,0X17C50D62E6E8D20D2DC18E9AD79C43263D0720D9/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "NFAI,UNISWAP_V3,0X17C50D62E6E8D20D2DC18E9AD79C43263D0720D9",
          "Quote": "USD"
        },
        "decimals": 11,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"23267\"}]}"
      },
      "provider_configs": [
        {
          "name": "uniswapv3_api-ethereum",
          "off_chain_ticker": "NFAI,UNISWAP_V3,0X17C50D62E6E8D20D2DC18E9AD79C43263D0720D9/WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
          "normalize_by_pair": {
            "Base": "WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"address\":\"0x856180c014d712099e7d4649752821c81531ce57\",\"base_decimals\":18,\"quote_decimals\":18,\"invert\":false}"
        }
      ]
    },
    "NPC,UNISWAP-V3-BASE,0XB166E8B140D35D9D8226E40C09F757BAC5A4D87D/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "NPC,UNISWAP-V3-BASE,0XB166E8B140D35D9D8226E40C09F757BAC5A4D87D",
          "Quote": "USD"
        },
        "decimals": 11,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"27960\"}]}"
      },
      "provider_configs": [
        {
          "name": "uniswapv3_api-base",
          "off_chain_ticker": "WETH,UNISWAP-V3-BASE,0X4200000000000000000000000000000000000006/NPC,UNISWAP-V3-BASE,0XB166E8B140D35D9D8226E40C09F757BAC5A4D87D",
          "normalize_by_pair": {
            "Base": "WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
            "Quote": "USD"
          },
          "invert": true,
          "metadata_JSON": "{\"address\":\"0x96e727835c5ec965aae8cbbad34901e3be33befb\",\"base_decimals\":18,\"quote_decimals\":18,\"invert\":false}"
        }
      ]
    },
    "NXRA,UNISWAP_V3,0X644192291CC835A93D6330B24EA5F5FEDD0EEF9E/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "NXRA,UNISWAP_V3,0X644192291CC835A93D6330B24EA5F5FEDD0EEF9E",
          "Quote": "USD"
        },
        "decimals": 11,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"23825\"}]}"
      },
      "provider_configs": [
        {
          "name": "uniswapv3_api-ethereum",
          "off_chain_ticker": "NXRA,UNISWAP_V3,0X644192291CC835A93D6330B24EA5F5FEDD0EEF9E/WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
          "normalize_by_pair": {
            "Base": "WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"address\":\"0x755f01736f93c91585b840c2179c560b754d69f3\",\"base_decimals\":18,\"quote_decimals\":18,\"invert\":false}"
        }
      ]
    },
    "OCEAN/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "OCEAN",
          "Quote": "USD"
        },
        "decimals": 10,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"3911\"}]}"
      },
      "provider_configs": [
        {
          "name": "uniswapv3_api-ethereum",
          "off_chain_ticker": "OCEAN,UNISWAP_V3,0X967DA4048CD07AB37855C090AAF366E4CE1B9F48/WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
          "normalize_by_pair": {
            "Base": "WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"address\":\"0x283e2e83b7f3e297c4b7c02114ab0196b001a109\",\"base_decimals\":18,\"quote_decimals\":18,\"invert\":false}"
        },
        {
          "name": "coinbase_ws",
          "off_chain_ticker": "OCEAN-USD"
        }
      ]
    },
    "OWO,UNISWAP-V3-BASE,0X5D559EA7BB2DAE4B694A079CB8328A2145FD32F6/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "OWO,UNISWAP-V3-BASE,0X5D559EA7BB2DAE4B694A079CB8328A2145FD32F6",
          "Quote": "USD"
        },
        "decimals": 11,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"32669\"}]}"
      },
      "provider_configs": [
        {
          "name": "uniswapv3_api-base",
          "off_chain_ticker": "OWO,UNISWAP-V3-BASE,0X5D559EA7BB2DAE4B694A079CB8328A2145FD32F6/WETH,UNISWAP-V3-BASE,0X4200000000000000000000000000000000000006",
          "normalize_by_pair": {
            "Base": "WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"address\":\"0x09c149c856e6fb6e40aa39209142411b554b1a41\",\"base_decimals\":18,\"quote_decimals\":18,\"invert\":false}"
        }
      ]
    },
    "PEIPEI,UNISWAP_V3,0X3FFEEA07A27FAB7AD1DF5297FA75E77A43CB5790/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "PEIPEI,UNISWAP_V3,0X3FFEEA07A27FAB7AD1DF5297FA75E77A43CB5790",
          "Quote": "USD"
        },
        "decimals": 16,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"31632\"}]}"
      },
      "provider_configs": [
        {
          "name": "uniswapv3_api-ethereum",
          "off_chain_ticker": "PEIPEI,UNISWAP_V3,0X3FFEEA07A27FAB7AD1DF5297FA75E77A43CB5790/WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
          "normalize_by_pair": {
            "Base": "WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"address\":\"0x9e0fc414e8d5c45b0890c32ab9329ac90b3ab534\",\"base_decimals\":18,\"quote_decimals\":18,\"invert\":false}"
        }
      ]
    },
    "PENDLE,UNISWAP_V3,0X808507121B80C02388FAD14726482E061B8DA827/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "PENDLE,UNISWAP_V3,0X808507121B80C02388FAD14726482E061B8DA827",
          "Quote": "USD"
        },
        "decimals": 9,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"9481\"}]}"
      },
      "provider_configs": [
        {
          "name": "uniswapv3_api-ethereum",
          "off_chain_ticker": "PENDLE,UNISWAP_V3,0X808507121B80C02388FAD14726482E061B8DA827/WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
          "normalize_by_pair": {
            "Base": "WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"address\":\"0x57af956d3e2cca3b86f3d8c6772c03ddca3eaacb\",\"base_decimals\":18,\"quote_decimals\":18,\"invert\":false}"
        }
      ]
    },
    "PEPE/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "PEPE",
          "Quote": "USD"
        },
        "decimals": 15,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"24478\"}]}"
      },
      "provider_configs": [
        {
          "name": "uniswapv3_api-ethereum",
          "off_chain_ticker": "PEPE,UNISWAP_V3,0X6982508145454CE325DDBE47A25D4EC3D2311933/WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
          "normalize_by_pair": {
            "Base": "WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"address\":\"0x11950d141ecb863f01007add7d1a342041227b58\",\"base_decimals\":18,\"quote_decimals\":18,\"invert\":false}"
        },
        {
          "name": "crypto_dot_com_ws",
          "off_chain_ticker": "PEPE_USD"
        },
        {
          "name": "bitstamp_api",
          "off_chain_ticker": "pepeusd"
        }
      ]
    },
    "POL/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "POL",
          "Quote": "USD"
        },
        "decimals": 10,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"28321\"}]}"
      },
      "provider_configs": [
        {
          "name": "uniswapv3_api-ethereum",
          "off_chain_ticker": "POL,UNISWAP_V3,0X455E53CBB86018AC2B8092FDCD39D8444AFFC3F6/WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
          "normalize_by_pair": {
            "Base": "WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"address\":\"0x88ada7b1dd1c728ea87404cea7a0780139eb35d1\",\"base_decimals\":18,\"quote_decimals\":18,\"invert\":false}"
        },
        {
          "name": "coinbase_ws",
          "off_chain_ticker": "POL-USD"
        },
        {
          "name": "bitfinex_ws",
          "off_chain_ticker": "POLUSD"
        }
      ]
    },
    "PORTAL,UNISWAP_V3,0X1BBE973BEF3A977FC51CBED703E8FFDEFE001FED/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "PORTAL,UNISWAP_V3,0X1BBE973BEF3A977FC51CBED703E8FFDEFE001FED",
          "Quote": "USD"
        },
        "decimals": 10,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"29555\"}]}"
      },
      "provider_configs": [
        {
          "name": "uniswapv3_api-ethereum",
          "off_chain_ticker": "PORTAL,UNISWAP_V3,0X1BBE973BEF3A977FC51CBED703E8FFDEFE001FED/WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
          "normalize_by_pair": {
            "Base": "WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"address\":\"0xf9bc02a0f79ee8b6982a754979c9dbd909ccee10\",\"base_decimals\":18,\"quote_decimals\":18,\"invert\":false}"
        }
      ]
    },
    "POWR,UNISWAP_V3,0X595832F8FC6BF59C85C527FEC3740A1B7A361269/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "POWR,UNISWAP_V3,0X595832F8FC6BF59C85C527FEC3740A1B7A361269",
          "Quote": "USD"
        },
        "decimals": 10,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"2132\"}]}"
      },
      "provider_configs": [
        {
          "name": "uniswapv3_api-ethereum",
          "off_chain_ticker": "POWR,UNISWAP_V3,0X595832F8FC6BF59C85C527FEC3740A1B7A361269/WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
          "normalize_by_pair": {
            "Base": "WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"address\":\"0xe3fe800b0de664bf0bca8ad58ecbc73b259047b0\",\"base_decimals\":6,\"quote_decimals\":18,\"invert\":false}"
        }
      ]
    },
    "QNT/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "QNT",
          "Quote": "USD"
        },
        "decimals": 8,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"3155\"}]}"
      },
      "provider_configs": [
        {
          "name": "uniswapv3_api-ethereum",
          "off_chain_ticker": "QNT,UNISWAP_V3,0X4A220E6096B25EADB88358CB44068A3248254675/WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
          "normalize_by_pair": {
            "Base": "WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"address\":\"0x24ee2c6b9597f035088cda8575e9d5e15a84b9df\",\"base_decimals\":18,\"quote_decimals\":18,\"invert\":false}"
        },
        {
          "name": "coinbase_ws",
          "off_chain_ticker": "QNT-USD"
        }
      ]
    },
    "RETH,UNISWAP_V3,0XAE78736CD615F374D3085123A210448E74FC6393/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "RETH,UNISWAP_V3,0XAE78736CD615F374D3085123A210448E74FC6393",
          "Quote": "USD"
        },
        "decimals": 6,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"15060\"}]}"
      },
      "provider_configs": [
        {
          "name": "uniswapv3_api-ethereum",
          "off_chain_ticker": "RETH,UNISWAP_V3,0XAE78736CD615F374D3085123A210448E74FC6393/WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
          "normalize_by_pair": {
            "Base": "WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"address\":\"0x553e9c493678d8606d6a5ba284643db2110df823\",\"base_decimals\":18,\"quote_decimals\":18,\"invert\":false}"
        }
      ]
    },
    "RIO,UNISWAP_V3,0XF21661D0D1D76D3ECB8E1B9F1C923DBFFFAE4097/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "RIO,UNISWAP_V3,0XF21661D0D1D76D3ECB8E1B9F1C923DBFFFAE4097",
          "Quote": "USD"
        },
        "decimals": 10,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"4166\"}]}"
      },
      "provider_configs": [
        {
          "name": "uniswapv3_api-ethereum",
          "off_chain_ticker": "RIO,UNISWAP_V3,0XF21661D0D1D76D3ECB8E1B9F1C923DBFFFAE4097/WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
          "normalize_by_pair": {
            "Base": "WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"address\":\"0x5b7e3e37a1aa6369386e5939053779abd3597508\",\"base_decimals\":18,\"quote_decimals\":18,\"invert\":false}"
        }
      ]
    },
    "RLB,UNISWAP_V3,0X046EEE2CC3188071C02BFC1745A6B17C656E3F3D/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "RLB,UNISWAP_V3,0X046EEE2CC3188071C02BFC1745A6B17C656E3F3D",
          "Quote": "USD"
        },
        "decimals": 11,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"15271\"}]}"
      },
      "provider_configs": [
        {
          "name": "uniswapv3_api-ethereum",
          "off_chain_ticker": "RLB,UNISWAP_V3,0X046EEE2CC3188071C02BFC1745A6B17C656E3F3D/USDT,UNISWAP_V3,0XDAC17F958D2EE523A2206206994597C13D831EC7",
          "normalize_by_pair": {
            "Base": "USDT",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"address\":\"0x33676385160f9d8f03a0db2821029882f7c79e93\",\"base_decimals\":18,\"quote_decimals\":6,\"invert\":false}"
        }
      ]
    },
    "RPL/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "RPL",
          "Quote": "USD"
        },
        "decimals": 8,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"2943\"}]}"
      },
      "provider_configs": [
        {
          "name": "uniswapv3_api-ethereum",
          "off_chain_ticker": "RPL,UNISWAP_V3,0XD33526068D116CE69F19A9EE46F0BD304F21A51F/WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
          "normalize_by_pair": {
            "Base": "WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"address\":\"0xe42318ea3b998e8355a3da364eb9d48ec725eb45\",\"base_decimals\":18,\"quote_decimals\":18,\"invert\":false}"
        },
        {
          "name": "coinbase_ws",
          "off_chain_ticker": "RPL-USD"
        }
      ]
    },
    "RSR,UNISWAP_V3,0X320623B8E4FF03373931769A31FC52A4E78B5D70/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "RSR,UNISWAP_V3,0X320623B8E4FF03373931769A31FC52A4E78B5D70",
          "Quote": "USD"
        },
        "decimals": 12,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"3964\"}]}"
      },
      "provider_configs": [
        {
          "name": "uniswapv3_api-ethereum",
          "off_chain_ticker": "RSR,UNISWAP_V3,0X320623B8E4FF03373931769A31FC52A4E78B5D70/WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
          "normalize_by_pair": {
            "Base": "WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"address\":\"0x32d9259e6792b2150fd50395d971864647fa27b2\",\"base_decimals\":18,\"quote_decimals\":18,\"invert\":false}"
        }
      ]
    },
    "SHIB/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "SHIB",
          "Quote": "USD"
        },
        "decimals": 14,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"5994\"}]}"
      },
      "provider_configs": [
        {
          "name": "crypto_dot_com_ws",
          "off_chain_ticker": "SHIB_USDT",
          "normalize_by_pair": {
            "Base": "USDT",
            "Quote": "USD"
          }
        },
        {
          "name": "uniswapv3_api-ethereum",
          "off_chain_ticker": "SHIB,UNISWAP_V3,0X95AD61B0A150D79219DCF64E1E6CC01F0B64C4CE/WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
          "normalize_by_pair": {
            "Base": "WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"address\":\"0x2f62f2b4c5fcd7570a709dec05d68ea19c82a9ec\",\"base_decimals\":18,\"quote_decimals\":18,\"invert\":false}"
        },
        {
          "name": "coinbase_ws",
          "off_chain_ticker": "SHIB-USD"
        }
      ]
    },
    "SHOOT,UNISWAP-V3-BASE,0XA9F5031B54C44C3603B4300FDE9B8F5CD18AD06F/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "SHOOT,UNISWAP-V3-BASE,0XA9F5031B54C44C3603B4300FDE9B8F5CD18AD06F",
          "Quote": "USD"
        },
        "decimals": 12,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"30365\"}]}"
      },
      "provider_configs": [
        {
          "name": "uniswapv3_api-base",
          "off_chain_ticker": "SHOOT,UNISWAP-V3-BASE,0XA9F5031B54C44C3603B4300FDE9B8F5CD18AD06F/WETH,UNISWAP-V3-BASE,0X4200000000000000000000000000000000000006",
          "normalize_by_pair": {
            "Base": "WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"address\":\"0xdfc55d03ad1200c1ab9d76772956bb4a532500a0\",\"base_decimals\":18,\"quote_decimals\":18,\"invert\":false}"
        }
      ]
    },
    "SKOP,UNISWAP-V3-BASE,0X6D3B8C76C5396642960243FEBF736C6BE8B60562/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "SKOP,UNISWAP-V3-BASE,0X6D3B8C76C5396642960243FEBF736C6BE8B60562",
          "Quote": "USD"
        },
        "decimals": 11,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"32022\"}]}"
      },
      "provider_configs": [
        {
          "name": "uniswapv3_api-base",
          "off_chain_ticker": "SKOP,UNISWAP-V3-BASE,0X6D3B8C76C5396642960243FEBF736C6BE8B60562/WETH,UNISWAP-V3-BASE,0X4200000000000000000000000000000000000006",
          "normalize_by_pair": {
            "Base": "WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"address\":\"0xe9ed60539a8ea7a4da04ebfa524e631b1fd48525\",\"base_decimals\":18,\"quote_decimals\":18,\"invert\":false}"
        }
      ]
    },
    "SMT,UNISWAP_V3,0XB17548C7B510427BAAC4E267BEA62E800B247173/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "SMT,UNISWAP_V3,0XB17548C7B510427BAAC4E267BEA62E800B247173",
          "Quote": "USD"
        },
        "decimals": 10,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"11821\"}]}"
      },
      "provider_configs": [
        {
          "name": "uniswapv3_api-ethereum",
          "off_chain_ticker": "SMT,UNISWAP_V3,0XB17548C7B510427BAAC4E267BEA62E800B247173/WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
          "normalize_by_pair": {
            "Base": "WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"address\":\"0x1becf1ac50f31c3441181563f9d350ddf72a2bfa\",\"base_decimals\":18,\"quote_decimals\":18,\"invert\":false}"
        }
      ]
    },
    "SPEC,UNISWAP-V3-BASE,0X96419929D7949D6A801A6909C145C8EEF6A40431/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "SPEC,UNISWAP-V3-BASE,0X96419929D7949D6A801A6909C145C8EEF6A40431",
          "Quote": "USD"
        },
        "decimals": 9,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"32925\"}]}"
      },
      "provider_configs": [
        {
          "name": "uniswapv3_api-base",
          "off_chain_ticker": "SPEC,UNISWAP-V3-BASE,0X96419929D7949D6A801A6909C145C8EEF6A40431/WETH,UNISWAP-V3-BASE,0X4200000000000000000000000000000000000006",
          "normalize_by_pair": {
            "Base": "WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"address\":\"0xee6f3c5f418d1097c50c4698d535edb33bd72931\",\"base_decimals\":18,\"quote_decimals\":18,\"invert\":false}"
        }
      ]
    },
    "TEL,UNISWAP_V3,0X467BCCD9D29F223BCE8043B84E8C8B282827790F/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "TEL,UNISWAP_V3,0X467BCCD9D29F223BCE8043B84E8C8B282827790F",
          "Quote": "USD"
        },
        "decimals": 12,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"2394\"}]}"
      },
      "provider_configs": [
        {
          "name": "uniswapv3_api-ethereum",
          "off_chain_ticker": "TEL,UNISWAP_V3,0X467BCCD9D29F223BCE8043B84E8C8B282827790F/WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
          "normalize_by_pair": {
            "Base": "WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"address\":\"0xf359492d26764481002ed88bd2acae83ca50b5c9\",\"base_decimals\":2,\"quote_decimals\":18,\"invert\":false}"
        }
      ]
    },
    "TOSHI,UNISWAP-V3-BASE,0XAC1BD2486AAF3B5C0FC3FD868558B082A531B2B4/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "TOSHI,UNISWAP-V3-BASE,0XAC1BD2486AAF3B5C0FC3FD868558B082A531B2B4",
          "Quote": "USD"
        },
        "decimals": 14,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"27750\"}]}"
      },
      "provider_configs": [
        {
          "name": "uniswapv3_api-base",
          "off_chain_ticker": "TOSHI,UNISWAP-V3-BASE,0XAC1BD2486AAF3B5C0FC3FD868558B082A531B2B4/WETH,UNISWAP-V3-BASE,0X4200000000000000000000000000000000000006",
          "normalize_by_pair": {
            "Base": "WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"address\":\"0x5aa4ad647580bfe86258d300bc9852f4434e2c61\",\"base_decimals\":18,\"quote_decimals\":18,\"invert\":false}"
        }
      ]
    },
    "TURBO,UNISWAP_V3,0XA35923162C49CF95E6BF26623385EB431AD920D3/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "TURBO,UNISWAP_V3,0XA35923162C49CF95E6BF26623385EB431AD920D3",
          "Quote": "USD"
        },
        "decimals": 12,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"24911\"}]}"
      },
      "provider_configs": [
        {
          "name": "uniswapv3_api-ethereum",
          "off_chain_ticker": "TURBO,UNISWAP_V3,0XA35923162C49CF95E6BF26623385EB431AD920D3/WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
          "normalize_by_pair": {
            "Base": "WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"address\":\"0x8107fca5494375fc743a9fc4d4844353a1af3d94\",\"base_decimals\":18,\"quote_decimals\":18,\"invert\":false}"
        }
      ]
    },
    "UNI/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "UNI",
          "Quote": "USD"
        },
        "decimals": 9,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"7083\"}]}"
      },
      "provider_configs": [
        {
          "name": "coinbase_ws",
          "off_chain_ticker": "UNI-USD"
        },
        {
          "name": "bitstamp_api",
          "off_chain_ticker": "uniusd"
        },
        {
          "name": "uniswapv3_api-ethereum",
          "off_chain_ticker": "UNI,UNISWAP_V3,0X1F9840A85D5AF5BF1D1762F925BDADDC4201F984/WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
          "normalize_by_pair": {
            "Base": "WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"address\":\"0x1d42064fc4beb5f8aaf85f4617ae8b3b5b8bd801\",\"base_decimals\":18,\"quote_decimals\":18,\"invert\":false}"
        },
        {
          "name": "crypto_dot_com_ws",
          "off_chain_ticker": "UNI_USDT",
          "normalize_by_pair": {
            "Base": "USDT",
            "Quote": "USD"
          }
        }
      ]
    },
    "USDC/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "USDC",
          "Quote": "USD"
        },
        "decimals": 9,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"3408\"}]}"
      },
      "provider_configs": [
        {
          "name": "coinbase_ws",
          "off_chain_ticker": "USDT-USDC",
          "normalize_by_pair": {
            "Base": "USDT",
            "Quote": "USD"
          },
          "invert": true
        },
        {
          "name": "bitstamp_api",
          "off_chain_ticker": "usdcusd"
        },
        {
          "name": "uniswapv3_api-ethereum",
          "off_chain_ticker": "WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2/USDC,UNISWAP_V3,0XA0B86991C6218B36C1D19D4A2E9EB0CE3606EB48",
          "normalize_by_pair": {
            "Base": "WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
            "Quote": "USD"
          },
          "invert": true,
          "metadata_JSON": "{\"address\":\"0x88e6a0c2ddd26feeb64f039a2c41296fcb3f5640\",\"base_decimals\":18,\"quote_decimals\":6,\"invert\":false}"
        }
      ]
    },
    "USDE,UNISWAP_V3,0X4C9EDD5852CD905F086C759E8383E09BFF1E68B3/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "USDE,UNISWAP_V3,0X4C9EDD5852CD905F086C759E8383E09BFF1E68B3",
          "Quote": "USD"
        },
        "decimals": 10,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"29470\"}]}"
      },
      "provider_configs": [
        {
          "name": "uniswapv3_api-ethereum",
          "off_chain_ticker": "USDE,UNISWAP_V3,0X4C9EDD5852CD905F086C759E8383E09BFF1E68B3/USDT,UNISWAP_V3,0XDAC17F958D2EE523A2206206994597C13D831EC7",
          "normalize_by_pair": {
            "Base": "USDT",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"address\":\"0x435664008f38b0650fbc1c9fc971d0a3bc2f1e47\",\"base_decimals\":18,\"quote_decimals\":6,\"invert\":false}"
        }
      ]
    },
    "USDT/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "USDT",
          "Quote": "USD"
        },
        "decimals": 9,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"825\"}]}"
      },
      "provider_configs": [
        {
          "name": "crypto_dot_com_ws",
          "off_chain_ticker": "USDT_USD"
        },
        {
          "name": "bitfinex_ws",
          "off_chain_ticker": "USTUSD"
        },
        {
          "name": "bitstamp_api",
          "off_chain_ticker": "usdtusd"
        },
        {
          "name": "coinbase_ws",
          "off_chain_ticker": "USDT-USD"
        }
      ]
    },
    "VRA,UNISWAP_V3,0XF411903CBC70A74D22900A5DE66A2DDA66507255/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "VRA,UNISWAP_V3,0XF411903CBC70A74D22900A5DE66A2DDA66507255",
          "Quote": "USD"
        },
        "decimals": 12,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"3816\"}]}"
      },
      "provider_configs": [
        {
          "name": "uniswapv3_api-ethereum",
          "off_chain_ticker": "VRA,UNISWAP_V3,0XF411903CBC70A74D22900A5DE66A2DDA66507255/WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
          "normalize_by_pair": {
            "Base": "WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"address\":\"0x98409d8ca9629fbe01ab1b914ebf304175e384c8\",\"base_decimals\":18,\"quote_decimals\":18,\"invert\":false}"
        }
      ]
    },
    "WBTC/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "WBTC",
          "Quote": "USD"
        },
        "decimals": 5,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"3717\"}]}"
      },
      "provider_configs": [
        {
          "name": "uniswapv3_api-ethereum",
          "off_chain_ticker": "WBTC,UNISWAP_V3,0X2260FAC5E5542A773AA44FBCFEDF7C193BC2C599/USDT,UNISWAP_V3,0XDAC17F958D2EE523A2206206994597C13D831EC7",
          "normalize_by_pair": {
            "Base": "USDT",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"address\":\"0x9db9e0e53058c89e5b94e29621a205198648425b\",\"base_decimals\":8,\"quote_decimals\":6,\"invert\":false}"
        },
        {
          "name": "coinbase_ws",
          "off_chain_ticker": "WBTC-USD"
        }
      ]
    },
    "WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
          "Quote": "USD"
        },
        "decimals": 6,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"2396\"}]}"
      },
      "provider_configs": [
        {
          "name": "uniswapv3_api-ethereum",
          "off_chain_ticker": "WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2/USDT,UNISWAP_V3,0XDAC17F958D2EE523A2206206994597C13D831EC7",
          "normalize_by_pair": {
            "Base": "USDT",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"address\":\"0xc7bbec68d12a0d1830360f8ec58fa599ba1b0e9b\",\"base_decimals\":18,\"quote_decimals\":6,\"invert\":false}"
        }
      ]
    },
    "WLD,UNISWAP_V3,0X163F8C2467924BE0AE7B5347228CABF260318753/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "WLD,UNISWAP_V3,0X163F8C2467924BE0AE7B5347228CABF260318753",
          "Quote": "USD"
        },
        "decimals": 9,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"13502\"}]}"
      },
      "provider_configs": [
        {
          "name": "uniswapv3_api-ethereum",
          "off_chain_ticker": "WLD,UNISWAP_V3,0X163F8C2467924BE0AE7B5347228CABF260318753/WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
          "normalize_by_pair": {
            "Base": "WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"address\":\"0xc4472dcd0e42ffccc1dbb0b9b3855688c22f3a0f\",\"base_decimals\":18,\"quote_decimals\":18,\"invert\":false}"
        }
      ]
    },
    "WOLF,UNISWAP_V3,0X67466BE17DF832165F8C80A5A120CCC652BD7E69/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "WOLF,UNISWAP_V3,0X67466BE17DF832165F8C80A5A120CCC652BD7E69",
          "Quote": "USD"
        },
        "decimals": 14,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"30902\"}]}"
      },
      "provider_configs": [
        {
          "name": "uniswapv3_api-ethereum",
          "off_chain_ticker": "WOLF,UNISWAP_V3,0X67466BE17DF832165F8C80A5A120CCC652BD7E69/WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
          "normalize_by_pair": {
            "Base": "WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"address\":\"0x6981449ddaa030f83bc5ac9fde1c19544521906e\",\"base_decimals\":18,\"quote_decimals\":18,\"invert\":false}"
        }
      ]
    },
    "WSTETH,UNISWAP_V3,0X7F39C581F595B53C5CB19BD0B3F8DA6C935E2CA0/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "WSTETH,UNISWAP_V3,0X7F39C581F595B53C5CB19BD0B3F8DA6C935E2CA0",
          "Quote": "USD"
        },
        "decimals": 6,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"12409\"}]}"
      },
      "provider_configs": [
        {
          "name": "uniswapv3_api-ethereum",
          "off_chain_ticker": "WSTETH,UNISWAP_V3,0X7F39C581F595B53C5CB19BD0B3F8DA6C935E2CA0/WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
          "normalize_by_pair": {
            "Base": "WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"address\":\"0x109830a1aaad605bbf02a9dfa7b0b92ec2fb7daa\",\"base_decimals\":18,\"quote_decimals\":18,\"invert\":false}"
        }
      ]
    },
    "ZRO/USD": {
      "ticker": {
        "currency_pair": {
          "Base": "ZRO",
          "Quote": "USD"
        },
        "decimals": 9,
        "min_provider_count": 1,
        "enabled": true,
        "metadata_JSON": "{\"aggregate_ids\":[{\"venue\":\"coinmarketcap\",\"ID\":\"26997\"}]}"
      },
      "provider_configs": [
        {
          "name": "uniswapv3_api-ethereum",
          "off_chain_ticker": "ZRO,UNISWAP_V3,0X6985884C4392D348587B19CB9EAAF157F13271CD/WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
          "normalize_by_pair": {
            "Base": "WETH,UNISWAP_V3,0XC02AAA39B223FE8D0A0E5C4F27EAD9083C756CC2",
            "Quote": "USD"
          },
          "metadata_JSON": "{\"address\":\"0x360acf12e72044ba3eaaa654e51e4725c699dcb1\",\"base_decimals\":18,\"quote_decimals\":18,\"invert\":false}"
        },
        {
          "name": "coinbase_ws",
          "off_chain_ticker": "ZRO-USD"
        }
      ]
    }
  }
}`
)

func init() {
	err := errors.Join(
		unmarshalValidate("CoinMarketCap", CoinMarketCapMarketMapJSON, &CoinMarketCapMarketMap),
		unmarshalValidate("Raydium", RaydiumMarketMapJSON, &RaydiumMarketMap),
		unmarshalValidate("Core", CoreMarketMapJSON, &CoreMarketMap),
		unmarshalValidate("UniswapV3Base", UniswapV3BaseMarketMapJSON, &UniswapV3BaseMarketMap),
		unmarshalValidate("CoinGecko", CoinGeckoMarketMapJSON, &CoinGeckoMarketMap),
		unmarshalValidate("Osmosis", OsmosisMarketMapJSON, &OsmosisMarketMap),
		unmarshalValidate("Polymarket", PolymarketMarketMapJSON, &PolymarketMarketMap),
		unmarshalValidate("Uniswap", UniswapV3EthMarketMapJSON, &UniswapV3EthMarketMap),
	)
	if err != nil {
		panic(err)
	}
}

// unmarshalValidate unmarshalls data into mm and then calls ValidateBasic.
func unmarshalValidate(name, data string, mm *mmtypes.MarketMap) error {
	if err := json.Unmarshal([]byte(data), mm); err != nil {
		return fmt.Errorf("failed to unmarshal %sMarketMap: %w", name, err)
	}
	if err := mm.ValidateBasic(); err != nil {
		return fmt.Errorf("%sMarketMap failed validation: %w", name, err)
	}
	return nil
}
