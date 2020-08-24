package genid

import (
	"encoding/binary"
	"strconv"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type PersonCounter struct {
	count       uint64
	key         string
	start       uint64
	head        string
	historyHead string
}

func NewPersonCounter(ctx contractapi.TransactionContextInterface) *PersonCounter {
	p := &PersonCounter{
		key:         "person_counter",
		start:       uint64(1000),
		head:        "person_",
		historyHead: "history_",
	}

	var id uint64
	couterRes, _ := ctx.GetStub().GetState(p.key)
	if couterRes != nil {
		id = binary.BigEndian.Uint64(couterRes)
	} else {
		id = p.start
	}
	p.count = id
	return p
}

func (p *PersonCounter) GetPersonId() string {
	return p.head + strconv.FormatUint(p.count, 10)
}

func (p *PersonCounter) GetHistoryId() string {
	return p.historyHead + strconv.FormatUint(p.count, 10)
}

func (p *PersonCounter) IncreasePersonId(ctx contractapi.TransactionContextInterface) error {
	p.count++
	bytes := make([]byte, 8)
	binary.BigEndian.PutUint64(bytes, p.count)
	return ctx.GetStub().PutState(p.key, bytes)
}
