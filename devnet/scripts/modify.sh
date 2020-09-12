#!/bin/bash

MODE="$1"
. ./scripts/env.sh

function beforeEdit(){
  # fetch config file
  peer channel fetch config config_block.pb \
    -o orderer.develop.com:7050 \
    -c $CHANNEL_NAME \
    --tls \
    --cafile $ORDERER_TLS_CA

  # transform pb file to json format
  configtxlator proto_decode \
    --input config_block.pb \
    --type common.Block | jq .data.data[0].payload.data.config > config.json

  # copy and for edit
  cp config.json modified_config.json

  # check files
  ls -al *.json *.pb
}

function edit(){
  echo "Edit..."
}

function afterEdit(){
  # transform json to pb format
  configtxlator proto_encode --input config.json --type common.Config --output config.pb
  configtxlator proto_encode --input modified_config.json --type common.Config --output modified_config.pb

  # cal difference
  configtxlator compute_update --channel_id $CHANNEL_NAME --original config.pb --updated modified_config.pb --output diff_config.pb

  if [ ! $? -eq 0 ]; then
    echo "Execute fail and exist"
    exit 1
  fi

  # diff pb to json
  configtxlator proto_decode --input diff_config.pb --type common.ConfigUpdate | jq . > diff_config.json

  # add envelope
  echo '{"payload":{"header":{"channel_header":{"channel_id":"'$CHANNEL_NAME'", "type":2}},"data":{"config_update":'$(cat diff_config.json)'}}}' | jq . > diff_config_envelope.json
  configtxlator proto_encode --input diff_config_envelope.json --type common.Envelope --output diff_config_envelope.pb

  # check files
  ls -al *.json *.pb
}

function submit(){
  # sign tx with org1 msp
  peer channel signconfigtx -f diff_config_envelope.pb

  # sign tx again with orderer msp
  export CORE_PEER_ADDRESS=orderer.develop.com:7050
  export CORE_PEER_LOCALMSPID="OrdererMSP"
  export CORE_PEER_TLS_ROOTCERT_FILE=$ORDERER_TLS_CA
  export CORE_PEER_MSPCONFIGPATH=$ORDERER_MSP
  peer channel signconfigtx -f diff_config_envelope.pb

  # submit tx
  peer channel update \
    -f diff_config_envelope.pb \
    -c $CHANNEL_NAME \
    -o orderer.develop.com:7050 \
    --tls --cafile $ORDERER_TLS_CA
}

function clean(){
  rm -rf *.json 
  rm -f *.pb
}

function check(){
  beforeEdit
  cat config.json
}

function help(){
  echo "Usage: modify.sh <cmd>"
  echo "cmd:"
  echo "  before: fetch pb, transform to json"
  echo "  edit:   do some edit"
  echo "  after:  cal diff pb, sign pb"
  echo "  submit: submit tx"
  echo "  check:  print channel config as json format"
  echo "  clean:  clean json and pb files"
}

case $MODE in
  before )
    beforeEdit
    ;;
  edit)
    edit
    ;;
  after )
    afterEdit
    ;;
  submit )
    submit
    ;;
  check )
    check
    ;;
  clean )
    clean
    ;;
  * )
    help
    exit 1
esac