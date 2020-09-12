#!/bin/bash

# variabls
MODE="$1"
COMMAND=wrk
BASE_URL='http://localhost:8000/api/v3'
GENERAL_OPTIONS='--latency -t 10 -c 10 -d 3s'

# Check if wrk command installed.
if ! command -v $COMMAND &> /dev/null
then
    echo "$COMMAND not found, please install it first!"
    exit
fi

function help(){
  echo "Usage: ./benchmark.sh <cmd>"
  echo "cmd:"
  echo "  hello"
  echo "  queryAllCars"
  echo "  queryCar"
  echo "  createCar"
  echo "  changeCarOwner"
}

case $MODE in
  hello)
    wrk "$BASE_URL/$MODE" $GENERAL_OPTIONS
    ;;
  queryAllCars)
    wrk "$BASE_URL/$MODE" -s $MODE.lua $GENERAL_OPTIONS
    ;;
  queryCar)
    wrk "$BASE_URL/$MODE" -s $MODE.lua $GENERAL_OPTIONS
    ;;
  createCar)
    wrk "$BASE_URL/$MODE" -s $MODE.lua $GENERAL_OPTIONS
    ;;
  changeCarOwner)
    echo "Not implement."
    ;;
  *)
    help
    exit 1
esac