#!/bin/bash

function up(){
  docker-compose up -d
}

function down(){
  docker-compose down
}

function clean(){
  down
  rm -rf data
}

# create the three system databases manually
function setup(){
  curl localhost:5984
  curl -X PUT http://admin:password@localhost:5984/_users
  curl -X PUT http://admin:password@localhost:5984/_replicator
  curl -X PUT http://admin:password@localhost:5984/_global_changes
}

function logs(){
    docker logs -f couchdb
}


function help() {
  echo "Usage: ./start.sh <cmd>"
  echo "cmd:"
  echo "  up"
  echo "  setup"
  echo "  down"
  echo "  logs"
  echo "  clean"
}

MODE="$1"
case $MODE in 
  up )
    up 
    ;;
  setup )
    setup
    ;;
  down )
    down
    ;;
  logs )
    logs
    ;;
  clean )
    clean
    ;;
  * )
    help
    exit 1
esac