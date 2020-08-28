package main

import (
	"log"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	health := new(Health)
	chaincode, err := contractapi.NewChaincode(health)
	if err != nil {
		log.Fatal("Create chaincode failed", err)
	}
	if err := chaincode.Start(); err != nil {
		log.Fatal("Start chaincode failed", err)
	}
}
