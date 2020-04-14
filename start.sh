rm -rf ~/.bc*

bcd init validator --chain-id band-consumer
bccli keys add validator --keyring-backend test
echo "smile stem oven genius cave resource better lunar nasty moon company ridge brass rather supply used horn three panic put venue analyst leader comic" | bccli keys add requester --recover --keyring-backend test
echo "clutch amazing good produce frequent release super evidence jungle voyage design clip title involve offer brain tobacco brown glide wire soft depend stand practice" | bccli keys add relayer --recover --keyring-backend test


bcd add-genesis-account validator 10000000000000stake --keyring-backend test
bcd add-genesis-account requester 10000000000000stake --keyring-backend test
bcd add-genesis-account relayer 10000000000000stake --keyring-backend test

bccli config chain-id band-consumer
bccli config output json
bccli config indent true
bccli config trust-node true

bcd gentx --name validator --keyring-backend test
bcd collect-gentxs
