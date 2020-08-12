# global variable and functions

# tools variables
CRYPTOGEN=../bin/cryptogen
CONFIGTXGEN=../bin/configtxgen
FABRIC_CA_CLIENT=../bin/fabric-ca-client 
COMPOSE_FILE_CA=./docker-compose-ca.yaml
COMPOSE_FILE=./docker-compose.yaml

# network variables, used in cli container
CHANNEL_NAME="c1"
CHAINCODE_NAME="mycc"
CHAINCODE_VERSION="1.0.0"
CHAINCODE_LABEL="abstore"
SEQUENCE=1
FABRIC_CFG_PATH="/etc/hyperledger/fabric"

# orderer variables, used in cli container
ORDERER_TLS_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/develop.com/orderers/orderer.develop.com/msp/tlscacerts/tlsca.develop.com-cert.pem

# peer variables, used in cli container
PEER_MSP=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.develop.com/users/Admin@org1.develop.com/msp
PEER_ADDR=peer0.org1.develop.com:7051
PEER_MSP_ID="Org1MSP"
PEER_TLS_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.develop.com/peers/peer0.org1.develop.com/tls/ca.crt 