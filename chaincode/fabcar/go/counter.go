package main

import (
	"encoding/json"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

const key = "counter_id"

type Counter struct {
	CID CounterID
}

type CounterID struct {
	ID uint `json:"id"`
}

func (c *Counter) Init(ctx contractapi.TransactionContextInterface, num uint) {
	counterID := CounterID{ID: num}
	c.CID = counterID
	bytes, _ := json.Marshal(&counterID)
	ctx.GetStub().PutState(key, bytes)
}

func (c *Counter) GetID(ctx contractapi.TransactionContextInterface) uint {
	// Not initialize, read from stub
	if c.CID.ID == 0 {
		couterIDByte, _ := ctx.GetStub().GetState(key)
		counterID := CounterID{}
		json.Unmarshal(couterIDByte, &counterID)
		return counterID.ID
	} else {
		return c.CID.ID
	}
}

func (c *Counter) Increase(ctx contractapi.TransactionContextInterface) error {
	id := c.GetID(ctx)
	id = id + 1
	newCounterId := CounterID{ID: id}
	c.CID = newCounterId
	bytes, _ := json.Marshal(&newCounterId)
	return ctx.GetStub().PutState(key, bytes)
}
