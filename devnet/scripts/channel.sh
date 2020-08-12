#!/bin/bash

MODE=$1
. ./scripts/env.sh

function help(){
  echo "Usage: "
  echo "  channel.sh <cmd>"
  echo "cmd: "
  echo "  - create"
  echo "  - join"
  echo "  - anchor"
  echo "  - info"
  echo "  - default"
}

function createChan(){
  peer channel create \
    -o orderer.develop.com:7050 \
    -c $CHANNEL_NAME \
    -f /opt/gopath/src/github.com/hyperledger/fabric/peer/channel-artifacts/${CHANNEL_NAME}.tx \
    --outputBlock /opt/gopath/src/github.com/hyperledger/fabric/peer/channel-artifacts/${CHANNEL_NAME}.block \
    --tls \
    --cafile $ORDERER_TLS_CA
}

function joinChan(){
  # r1 join
  CORE_PEER_MSPCONFIGPATH=${PEER_MSP}
  CORE_PEER_ADDRESS=${PEER_ADDR}
  CORE_PEER_LOCALMSPID=${PEER_MSP_ID}
  CORE_PEER_TLS_ROOTCERT_FILE=${PEER_TLS_CA}
  peer channel join \
    -b /opt/gopath/src/github.com/hyperledger/fabric/peer/channel-artifacts/${CHANNEL_NAME}.block

}

function showChanInfo(){
  peer channel list
  peer channel getinfo -c ${CHANNEL_NAME}
}

function setAnchor(){
 
  # anchor update
  CORE_PEER_MSPCONFIGPATH=${PEER_MSP}
  CORE_PEER_ADDRESS=${PEER_ADDR}
  CORE_PEER_LOCALMSPID=${PEER_MSP_ID}
  CORE_PEER_TLS_ROOTCERT_FILE=${PEER_TLS_CA}
  peer channel update \
    -o orderer.develop.com:7050 \
    -c $CHANNEL_NAME \
    -f /opt/gopath/src/github.com/hyperledger/fabric/peer/channel-artifacts/Org1Anchors.tx \
    --tls \
    --cafile $ORDERER_TLS_CA

  # show info
  showChanInfo
}

function wait(){
  delay=$1
  if [ "$delay" == "" ]; then
    delay=5
  fi

  echo "wait for $delay seconds..."
  sleep $delay
}

if [ "$MODE" == "create" ]; then
  createChan    
elif [ "$MODE" == "join" ]; then
  joinChan
elif [ "$MODE" == "anchor" ]; then
  setAnchor
elif [ "$MODE" == "info" ]; then
  showChanInfo
elif [ "$MODE" == "default" ]; then
  wait 5
  createChan
  wait 5
  joinChan
  wait 5
  setAnchor
  wait 5
  showChanInfo
else        
  help
  exit 1
fi
