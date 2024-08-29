VALIDATOR_NAME=validator1
CHAIN_ID=rollup-chain
KEY_NAME=rollup-key
CHAINFLAG="--chain-id ${CHAIN_ID}"
TOKEN_AMOUNT="10000000000000000000000000stake"
STAKING_AMOUNT="1000000000stake"


trash ~/.simapp/

rollkit rebuild
rollkit tendermint unsafe-reset-all
rollkit init $VALIDATOR_NAME --chain-id $CHAIN_ID
rollkit keys add $KEY_NAME --keyring-backend test
rollkit genesis add-genesis-account $KEY_NAME $TOKEN_AMOUNT --keyring-backend test
rollkit genesis gentx $KEY_NAME $STAKING_AMOUNT --chain-id $CHAIN_ID --keyring-backend test
rollkit genesis collect-gentxs

ADDRESS=$(jq -r '.address' ~/.simapp/config/priv_validator_key.json)
PUB_KEY=$(jq -r '.pub_key' ~/.simapp/config/priv_validator_key.json)
jq --argjson pubKey "$PUB_KEY" '.consensus["validators"]=[{"address": "'$ADDRESS'", "pub_key": $pubKey, "power": "1000", "name": "Rollkit Sequencer"}]' ~/.simapp/config/genesis.json > temp.json && mv temp.json ~/.simapp/config/genesis.json

rollkit start --rollkit.aggregator --rpc.laddr tcp://127.0.0.1:36657 --grpc.address 127.0.0.1:9290 --p2p.laddr "0.0.0.0:36656" --minimum-gas-prices="0.025stake" --rollkit.da_address "http://localhost:7980"
