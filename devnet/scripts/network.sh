#!/bin/bash

. ./scripts/env.sh

function genCrypto(){
  ${CRYPTOGEN} generate --config=./crypto-config.yaml --output="organizations"
}

function genCryptoCA(){
  #echo "generate cryto use ca"
  docker-compose -f $COMPOSE_FILE_CA up -d 2>&1
  sleep 10
  . ./scripts/registerEnroll.sh
  createOrg1
  createOrderer
}

function genGenesis(){
  ${CONFIGTXGEN} -profile Genesis -channelID ordererchannel -outputBlock ./system-genesis-block/genesis.block
}

function createChanTx(){
  ${CONFIGTXGEN} -profile CC1 -outputCreateChannelTx ./channel-artifacts/${CHANNEL_NAME}.tx -channelID $CHANNEL_NAME
  ${CONFIGTXGEN} -profile CC1 \
    -outputAnchorPeersUpdate ./channel-artifacts/Org1Anchors.tx \
    -channelID $CHANNEL_NAME \
    -asOrg Org1
}

function genCCP(){
  ./scripts/ccp.sh
}

function up(){
  docker-compose -f $COMPOSE_FILE -f $COMPOSE_FILE_CA up -d
}

function down(){
  docker-compose -f $COMPOSE_FILE -f $COMPOSE_FILE_CA down
}

function clear(){
  down
  rm -rf organizations system-genesis-block channel-artifacts
  rm -rf productions/orderer productions/peer
  rm -rf pkg/mycc.tar.gz 
}

function help(){
  echo "Usage: "
  echo "  network.sh <cmd>"
  echo "cmd: "
  echo "  · crypto"
  echo "  · ca"
  echo "  · genesis"
  echo "  · ccp"
  echo "  · up"
  echo "  · createChanTx"
  echo "  · down"
  echo "  · clear"
  echo "  · default"
  echo "flag:"
  echo "  --ca: use fabric-ca"
}

# parse mode
MODE=$1
shift

# default flags
FALG_CA="false"

# parse flags
while [[ $# -ge 1 ]]; do
  opt="$1"
  case $opt in
    --ca)
      FALG_CA="true"
      echo "enable fabric-ca"
      ;;
    *)
      echo "unkonwn flag: $opt"
      help
      exit 1
  esac
  shift
done

case "$MODE" in
  "crypto")
    genCrypto
    ;;
  "ca")
    genCryptoCA
    ;;
  "genesis")
    genGenesis
    ;;
  "createChanTx")
    createChanTx
    ;;
  "ccp")
    genCCP
    ;;
  "up")
    up
    ;;
  "down")
    down
    ;;
  "clear")
    clear
    ;;
  "default")
    clear
    if [ $FALG_CA = "true" ]; then
      genCryptoCA
      genCCP
    else
      genCrypto
    fi
    genGenesis
    createChanTx
    up
    ;;
  *)
    help
    exit 1
esac
