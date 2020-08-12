#!/bin/bash

MODE=$1
. ./scripts/env.sh

function help(){
  echo "Usage: "
  echo "  network.sh <cmd>"
  echo "cmd: "
  echo "  - crypto"
  echo "  - ca"
  echo "  - genesis"
  echo "  - ccp"
  echo "  - up"
  echo "  - createChanTx"
  echo "  - down"
  echo "  - clear"
  echo "  - custom"
}

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
   "custom")
    clear
    #genCrypto
    genCryptoCA
    genCCP
    genGenesis
    createChanTx
    up
    ;;
  *)
    help
    exit 1
esac
