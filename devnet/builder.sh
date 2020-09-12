#!/bin/bash

MODE=$1

# help 
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
  echo "  - modify"
}

# check binary

if [ ! -e "../bin/configtxgen" ] || [ ! -e "../bin/cryptogen" ] || [ ! -e "../config/core.yaml" ]; then
  echo "Please check README.md, run make, make sure fabric tools and configs downloaded first!"
  exit 1
fi

case "$MODE" in
  "network")
    shift
    scripts/network.sh "$@"
    ;;
  "channel")
    shift 
    docker exec cli scripts/channel.sh "$@"
    ;;
  "chaincode")
    shift
    docker exec cli scripts/chaincode.sh "$@"
    ;;
  "modify")
    shift
    docker exec cli scripts/modify.sh "$@"
    ;;
  "new")
    ./builder.sh network default
    ./builder.sh channel default
    ./builder.sh chaincode default
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
