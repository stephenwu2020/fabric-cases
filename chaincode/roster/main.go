package main

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"

	"fmt"
)

func main() {
	rosterContract := NewRosterContract()
	chaincode, err := contractapi.NewChaincode(rosterContract)

	if err != nil {
		fmt.Printf("Error create roster chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting roster chaincode: %s", err.Error())
	}
}
