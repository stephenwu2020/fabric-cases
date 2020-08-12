#!/bin/bash

MODE=$1

function execNetwork(){
  scripts/network.sh $1
}

# notice: chaincode.sh running in cli container
function execChaincode(){
  docker exec cli scripts/chaincode.sh $1
}

# notice: chaincode.sh running in cli container
function execChannel(){
  docker exec cli scripts/channel.sh $1
}

function help(){
  echo "Usage: "
  echo "  ./builder.sh <cmd>"
  echo "cmd: "
  echo "  - new"
  echo "  - destroy"
  echo "  - up"
  echo "  - down"
  echo "  - network"
  echo "  - channel"
  echo "  - chaincode"
}

case "$MODE" in
  "network")
    execNetwork $2
    ;;
  "channel")
    execChannel $2
    ;;
  "chaincode")
    execChaincode $2
    ;;
  "new")
    ./builder.sh network custom
    ./builder.sh channel custom
    ./builder.sh chaincode custom
    ;;
  "destroy")
    ./builder.sh network clear
    ;;
  "up")
    ./builder.sh network up
    ;;
  "down")
    ./builder.sh network down 
    ;;
  *)
    help
    exit 1
esac
