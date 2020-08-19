package main

import (
	"encoding/json"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

const key = "counter_id"

type Counter struct{}
type CounterID struct {
	ID uint `json:"id"`
}

func (c *Counter) Init(ctx contractapi.TransactionContextInterface, num uint) {
	counterID := CounterID{ID: num}
	bytes, _ := json.Marshal(&counterID)
	ctx.GetStub().PutState(key, bytes)
}

func (c *Counter) GetID(ctx contractapi.TransactionContextInterface) uint {
	couterIDByte, _ := ctx.GetStub().GetState(key)
	counterID := CounterID{}
	json.Unmarshal(couterIDByte, &counterID)
	return counterID.ID
}

func (c *Counter) Increase(ctx contractapi.TransactionContextInterface) error {
	id := c.GetID(ctx)
	id = id + 1
	newCounterId := CounterID{ID: id}
	bytes, _ := json.Marshal(&newCounterId)
	return ctx.GetStub().PutState(key, bytes)
}
