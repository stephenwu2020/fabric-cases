package genid

import (
	"encoding/binary"
	"strconv"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/pkg/errors"
)

type PersonCounter struct {
	count       uint64
	key         string
	start       uint64
	head        string
	historyHead string
}

func NewPersonCounter() *PersonCounter {
	p := &PersonCounter{
		key:         "person_counter",
		start:       uint64(1000),
		head:        "person_",
		historyHead: "history_",
	}

	return p
}

func (p *PersonCounter) Init(ctx contractapi.TransactionContextInterface) error {
	couterRes, err := ctx.GetStub().GetState(p.key)
	if err != nil {
		return errors.WithMessage(err, "Get couter fail")
	}

	if couterRes != nil {
		p.count = binary.BigEndian.Uint64(couterRes)
	} else {
		p.count = p.start
	}
	return nil
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
