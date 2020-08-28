package health

import "github.com/hyperledger/fabric-contract-api-go/contractapi"

type Health struct {
	contractapi.Contract
}

func (h Health) Intro(ctx contractapi.TransactionContextInterface) (*HealthIntro, error) {
	intro := &HealthIntro{
		Name:     "Health",
		Function: "Record health data, analyse health situation.",
		Version:  "0.0.1",
		Author:   "Ming",
	}
	return intro, nil
}
