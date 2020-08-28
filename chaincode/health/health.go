package main

import (
	"encoding/json"
	"time"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/pkg/errors"
	"github.com/stephenwu2020/fabric-cases/chaincode/health/datatype"
)

type Health struct {
	contractapi.Contract
}

var recordkey = "health_record"

func (h Health) Intro(ctx contractapi.TransactionContextInterface) (*datatype.HealthIntro, error) {
	intro := &datatype.HealthIntro{
		Name:     "Health",
		Function: "Record health data, analyse health situation.",
		Version:  "0.0.1",
		Author:   "Ming",
	}
	return intro, nil
}

func (h Health) AddSleepRecord(ctx contractapi.TransactionContextInterface, sleepT datatype.RecordType, start, end time.Time) error {
	record := datatype.HealthRecord{
		Type:  sleepT,
		Start: start,
		End:   end,
	}
	records, err := h.GetRecords(ctx)
	if err != nil {
		return errors.WithMessage(err, "Get records failed")
	}
	records = append(records, record)
	bytes, err := json.Marshal(&records)
	if err != nil {
		return errors.WithMessage(err, "Marshal recrds failed")
	}
	return ctx.GetStub().PutState(recordkey, bytes)
}

func (h Health) GetRecords(ctx contractapi.TransactionContextInterface) ([]datatype.HealthRecord, error) {
	bytes, err := ctx.GetStub().GetState(recordkey)
	if err != nil {
		return nil, errors.WithMessage(err, "Get records from state failed")
	}
	if bytes == nil {
		return []datatype.HealthRecord{}, nil
	}
	var records []datatype.HealthRecord
	if err := json.Unmarshal(bytes, &records); err != nil {
		return nil, errors.WithMessage(err, "Marshal records failed")
	}
	return records, nil
}
