#!/bin/bash

. ./scripts/env.sh

function package(){
  if [ -f "pkg/${CHAINCODE_NAME}.tar.gz" ]; then
    echo "pkg already exist"
  else
    echo "fetch go dependency"
    pushd ${CHAINCODE_PATH}
    #GO111MODULE=on go mod vendor
    GO111MODULE=on go mod tidy -v
    if [ ! $? -eq 0 ]; then
      echo "Fetch go dependency fail! Check your network please!"
      exit 1
    fi
    popd

    echo "Package chaincode..."
    peer lifecycle chaincode package ${CHAINCODE_NAME}.tar.gz \
      --path ${CHAINCODE_PATH} \
      --lang golang \
      --label ${CHAINCODE_LABEL}_1

    cp ${CHAINCODE_NAME}.tar.gz /opt/gopath/src/github.com/hyperledger/fabric/peer/pkg
  fi
}

function install(){
  echo "Install chaincode..."
  GO111MODULE=on

  # install on peer of r1
  CORE_PEER_MSPCONFIGPATH=${PEER_MSP}
  CORE_PEER_ADDRESS=${PEER_ADDR}
  CORE_PEER_LOCALMSPID=${PEER_MSP_ID}
  CORE_PEER_TLS_ROOTCERT_FILE=${PEER_TLS_CA}
  peer lifecycle chaincode install /opt/gopath/src/github.com/hyperledger/fabric/peer/pkg/${CHAINCODE_NAME}.tar.gz

}

function approve(){
  echo "Approve chaincode..."
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
  echo "Commit chaincode..."
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
  echo "Invoke chaincode..."
  if [ -z ${CHAINCODE_INVOKE_OPTIONS} ]; then
    echo "skip..."
    return
  fi
  peer chaincode invoke \
    -o orderer.develop.com:7050 \
    --isInit \
    --tls \
    --cafile $ORDERER_TLS_CA \
    -C $CHANNEL_NAME \
    -n ${CHAINCODE_NAME} \
    --peerAddresses peer0.org1.develop.com:7051 \
    --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.develop.com/peers/peer0.org1.develop.com/tls/ca.crt \
    -c ${CHAINCODE_INVOKE_OPTIONS} \
    --waitForEvent
}

function query(){
  echo "Query chaincode..."
  if [ -z ${CHAINCODE_QUERY_OPTIONS} ]; then
    echo "skip..."
    return
  fi
  peer chaincode query -C $CHANNEL_NAME -n ${CHAINCODE_NAME} -c ${CHAINCODE_QUERY_OPTIONS}
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
  echo "  · package"
  echo "  · install"
  echo "  · approve"
  echo "  · beforeCommit"
  echo "  · commit"
  echo "  · queryCommit"
  echo "  · invoke"
  echo "  · query"
  echo "  · upgrade"
  echo "  · debug"
  echo "  · default"
  echo "flag: "
  echo "  --ccn: chaincode name"
  echo "  --ccp: chaincode path"
  echo "  --ccl: chaincode label"
  echo "  --cci: invoke options"
  echo "  --ccq: query options"
}

# parse mode
MODE=$1
shift

# parse flags
while [[ $# -ge 1 ]]; do
  opt="$1"
  case $opt in
    --ccn)
      CHAINCODE_NAME="$2"
      shift
      ;;
    --ccp)
      CHAINCODE_PATH="$2"
      shift
      ;;
    --ccl)
      CHAINCODE_LABEL="$2"
      shift
      ;;
    --cci)
      CHAINCODE_INVOKE_OPTIONS="$2"
      shift
      ;;
    --ccq)
      CHAINCODE_QUERY_OPTIONS="$2"
      shift
      ;;
    *)
      echo "unkonwn flag: $opt"
      help
      exit 1
  esac
  shift
done

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
