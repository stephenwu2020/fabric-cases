#!/bin/bash

MODE=$1
. ./scripts/env.sh

function package(){
  if [ -f "pkg/${CHAINCODE_NAME}.tar.gz" ]; then
    echo "pkg already exist"
  else
    echo "fetch go dependency"
    pushd /opt/gopath/src/github.com/hyperledger/chaincode/abstore/go
    GO111MODULE=on go mod vendor
    popd

    echo "package chaincode"
    peer lifecycle chaincode package ${CHAINCODE_NAME}.tar.gz \
      --path /opt/gopath/src/github.com/hyperledger/chaincode/abstore/go \
      --lang golang \
      --label ${CHAINCODE_LABEL}_1

    cp ${CHAINCODE_NAME}.tar.gz /opt/gopath/src/github.com/hyperledger/fabric/peer/pkg
  fi
}

function install(){
  GO111MODULE=on

  # install on peer of r1
  CORE_PEER_MSPCONFIGPATH=${PEER_MSP}
  CORE_PEER_ADDRESS=${PEER_ADDR}
  CORE_PEER_LOCALMSPID=${PEER_MSP_ID}
  CORE_PEER_TLS_ROOTCERT_FILE=${PEER_TLS_CA}
  peer lifecycle chaincode install /opt/gopath/src/github.com/hyperledger/fabric/peer/pkg/${CHAINCODE_NAME}.tar.gz

}

function approve(){
  # query
  peer lifecycle chaincode queryinstalled >&log.txt
  cat log.txt
  PACKAGE_ID=`sed -n '/Package/{s/^Package ID: //; s/, Label:.*$//; $p;}' log.txt`
  echo PackageID is ${PACKAGE_ID}

  # approve r1
  CORE_PEER_MSPCONFIGPATH=${PEER_MSP}
  CORE_PEER_ADDRESS=${PEER_ADDR}
  CORE_PEER_LOCALMSPID=${PEER_MSP_ID}
  CORE_PEER_TLS_ROOTCERT_FILE=${PEER_TLS_CA}
  peer lifecycle chaincode approveformyorg \
    --channelID $CHANNEL_NAME \
    --name ${CHAINCODE_NAME} \
    --version ${CHAINCODE_VERSION} \
    --init-required \
    --package-id $PACKAGE_ID \
    --sequence ${SEQUENCE} \
    --tls \
    --cafile $ORDERER_TLS_CA
 
}

function beforeCommit(){
  # checkcommitreadiness
  peer lifecycle chaincode checkcommitreadiness \
    --channelID $CHANNEL_NAME \
    --name ${CHAINCODE_NAME} \
    --version ${CHAINCODE_VERSION} \
    --sequence ${SEQUENCE} \
    --init-required \
    --output json
}

function commit(){
  peer lifecycle chaincode commit \
    -o orderer.develop.com:7050 \
    --channelID $CHANNEL_NAME \
    --name ${CHAINCODE_NAME} \
    --version ${CHAINCODE_VERSION} \
    --sequence ${SEQUENCE} \
    --init-required \
    --tls true \
    --cafile $ORDERER_TLS_CA \
    --peerAddresses peer0.org1.develop.com:7051 \
    --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.develop.com/peers/peer0.org1.develop.com/tls/ca.crt 
}

function queryCommit(){
  peer lifecycle chaincode querycommitted -C $CHANNEL_NAME
}

function invoke(){
  peer chaincode invoke \
    -o orderer.develop.com:7050 \
    --isInit \
    --tls \
    --cafile $ORDERER_TLS_CA \
    -C $CHANNEL_NAME \
    -n ${CHAINCODE_NAME} \
    --peerAddresses peer0.org1.develop.com:7051 \
    --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.develop.com/peers/peer0.org1.develop.com/tls/ca.crt \
    -c '{"Args":["Init","a","100","b","100"]}' \
    --waitForEvent
}

function query(){
  peer chaincode query -C $CHANNEL_NAME -n ${CHAINCODE_NAME} -c '{"Args":["query","a"]}'
}

function upgrade(){
  echo "upgrade chaincode..."
  commitInfo=$( queryCommit )
  version=$(echo ${commitInfo} | cut -d ',' -f 2 | awk '{print $2}')
  sequence=$(echo ${commitInfo} | cut -d ',' -f 3 | awk '{print $2}')
  versionNew=$(increVersion ${version})
  sequenceNew=$(increSequence ${sequence})

  echo "options: version:${versionNew}, sequence:${sequenceNew}"

  CHAINCODE_VERSION=$versionNew
  SEQUENCE=$sequenceNew
  rm -rf /opt/gopath/src/github.com/hyperledger/fabric/peer/pkg/${CHAINCODE_NAME}.tar.gz
  package
  install
  approve
  commit
  invoke
  query
  echo "chaincode upgraded."
}

function increVersion() {
  v=$1
  major=$(echo $v | cut -d'.' -f1)
  minor=$(echo $v | cut -d'.' -f2)
  maintenance=$(echo $v | cut -d'.' -f3)
  maintenance=$((maintenance+1))
  new="$major.$minor.$maintenance"
  echo $new
}

function increSequence() {
  s=$1
  echo $((s+1))
}

function debug(){
  increVersion 1.0.0
  increVersion 1.2.3
  increVersion 1.2.10
  increVersion 1.2.100
  increVersion 1.2.999
  increSequence 1
  increSequence 9
  increSequence 100
  increSequence 999 
}

function help(){
  echo "Usage: "
  echo "  chaincode.sh <cmd>"
  echo "cmd: "
  echo "  - package"
  echo "  - install"
  echo "  - approve"
  echo "  - beforeCommit"
  echo "  - commit"
  echo "  - queryCommit"
  echo "  - invoke"
  echo "  - query"
  echo "  - upgrade"
  echo "  - debug"
  echo "  - default"
}

case "$MODE" in
  "package")
    package
    ;;
  "install")
    install
    ;;
  "approve")
    approve
    ;;
  "beforeCommit")
    beforeCommit
    ;;
  "commit")
    commit
    ;;
  "queryCommit")
    queryCommit
    ;;
  "invoke")
    invoke
    ;;
  "query")
    query
    ;;
  "upgrade")
    upgrade
    ;;
  "debug")
    debug
    ;;
  "default")
    package
    install
    approve
    beforeCommit
    commit
    queryCommit
    invoke
    query    
    ;;
  *)
    help
    exit 1
esac
