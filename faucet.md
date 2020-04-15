## bandchain

```
curl --location --request POST 'http://54.169.14.201/faucet/request' \
--header 'Content-Type: application/json' \
--data-raw '{
	"address": "band17md6xa3jykwcfj6nfw448k34qwxlxnylf2e54f",
	"amount": 10000000
}'
```

Can change balance via `http://54.169.14.201/rest/bank/balances/band17md6xa3jykwcfj6nfw448k34qwxlxnylf2e54f`

## gaia

```
curl --location --request POST 'http://gaia-ibc-hackathon.node.bandchain.org:8000' \
--header 'Content-Type: application/javascript' \
--data-raw '{
 "address": "cosmos1sjllsnramtg3ewxqwwrwjxfgc4n4ef9u0tvx7u",
 "chain-id": "band-cosmoshub"
}'
```
