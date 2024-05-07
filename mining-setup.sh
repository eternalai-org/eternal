#!/bin/bash

echo "Downloading miner-node"
curl https://github.com/eternalai-org/miner-node/releases/download/v0.5.0/miner-node-linux-amd64 -o miner-node


echo "Enter your account private key"
read ACCOUNT

echo "Enter your lighthouse API"
read LIGHTHOUSE_API

script="
#!/bin/bash
./miner-node -account $ACCOUNT -lighthouse-api $LIGHTHOUSE_API
"

echo "Creating script"
echo "$script" > run-miner-node.sh

echo "Setting up permissions"

chmod +x run-miner-node.sh
chmod +x miner-node
