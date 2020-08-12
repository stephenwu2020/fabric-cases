package main

import (
	"encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

const key = "counter_id"

type Counter struct{}
type CounterID struct {
	ID uint `json:"id"`
}

func (c *Counter) Init(APIstub shim.ChaincodeStubInterface, num uint) {
	counterID := CounterID{ID: num}
	bytes, _ := json.Marshal(&counterID)
	APIstub.PutState(key, bytes)
}

func (c *Counter) GetID(APIstub shim.ChaincodeStubInterface) uint {
	couterIDByte, _ := APIstub.GetState(key)
	counterID := CounterID{}
	json.Unmarshal(couterIDByte, &counterID)
	return counterID.ID
}

func (c *Counter) Increment(APIstub shim.ChaincodeStubInterface) {
	id := c.GetID(APIstub)
	id = id + 1
	newCounterId := CounterID{ID: id}
	bytes, _ := json.Marshal(&newCounterId)
	APIstub.PutState(key, bytes)
}
