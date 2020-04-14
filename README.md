## Instruction

1. `Make install` to get `bcd` and `bccli`
2. Initialized chain example can find at `start.sh`
3. Run single validator by `bcd start --rpc.laddr=tcp://0.0.0.0:26657 --pruning=nothing`

## Setup relayer

Look how to setup at relayer.sh

## How to get gold

0. Set up channel in gold chain by bccli

```
bccli tx consuming set-channel bandchain consuming <channel_id_of_consuming_goldchin> --from validator --keyring-backend test
bccli tx consuming set-channel band-cosmoshub transfer <channel_id_of_transfer_goldchin> --from validator --keyring-backend test
```

1. Transfer coin from gaia to bandchain

```
(bccli|gaiacli) tx transfer transfer transfer <channel_id_of_gaia> 10000000 <account_in_gold_chain> 800000000transfer/<channel_id_of_gold_chain>/uatom --from <account_in_gaia> --node http://gaia-ibc-hackathon.node.bandchain.org:26657 --keyring-backend test --chain-id band-cosmoshub
```

2. Send buy transaction

```
bccli tx consuming buy <amount_same_unit_as_transfer> --from <account_in_gold_chain> --keyring-backend test
```
