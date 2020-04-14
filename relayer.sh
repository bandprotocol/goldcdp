# Remove old config
rm -rf ~/.relayer

rly config init

# Add config after these commands can check your config file at `~/.relayer/config/config.yaml`
rly chains add -f gaia.json
rly chains add -f bandchain.json
rly chains add -f goldchain.json

# Add relayer account (Recover by mnemonic help for developing)
rly keys restore band-consumer relayer "clutch amazing good produce frequent release super evidence jungle voyage design clip title involve offer brain tobacco brown glide wire soft depend stand practice"
rly keys restore bandchain relayer "mix swift essence lawsuit plastic major social copper chicken aisle caution unfold leaf turtle prize remove gravity tourist gym parade number street twelve long"
rly keys restore band-cosmoshub relayer "clutch amazing good produce frequent release super evidence jungle voyage design clip title involve offer brain tobacco brown glide wire soft depend stand practice"

# Update default relayer for each chain
rly ch edit band-consumer key relayer
rly ch edit bandchain key relayer
rly ch edit band-cosmoshub key relayer

# And make sure every relayer have default coin in each chain by
rly q bal band-consumer
rly q bal bandchain
rly q bal band-cosmoshub

# Init lite client and save state for each chain
rly lite init band-consumer -f
rly lite init bandchain -f
rly lite init band-cosmoshub -f

# Create path(specific connection between chain)
rly pth gen  band-consumer transfer band-cosmoshub transfer transfer
rly pth gen  band-consumer consuming bandchain oracle oracle

# Create connection and channel from path
rly tx link transfer
rly tx link oracle

# Seperate run these command in different windows
rly st oracle
rly st transfer
