package main

import (
	"encoding/json"
	"fmt"
	"github.com/stephenwu2020/fabric-cases/chaincode/roster/datatype"
	"github.com/stephenwu2020/fabric-cases/chaincode/roster/genid"
	"strconv"
	"time"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/pkg/errors"
)

var timeLayoutStr = "2006-01-02 15:04:05"

type RosterContract struct {
	contractapi.Contract
	personCounter *genid.PersonCounter
}

func (r *RosterContract) Instantiate(ctx contractapi.TransactionContextInterface) error {
	r.personCounter = genid.NewPersonCounter(ctx)
	return nil
}

func (r *RosterContract) Version(ctx contractapi.TransactionContextInterface) (string, error) {
	return "1.0.0", nil
}

/*
**	Person Operations
 */

func (r *RosterContract) AddPerson(ctx contractapi.TransactionContextInterface, name string) error {
	person := &datatype.Person{
		Id:        r.personCounter.GetPersonId(),
		Name:      name,
		GroupTags: []string{},
		HistroyId: r.personCounter.GetHistoryId(),
	}
	bytes, _ := json.Marshal(person)
	if err := ctx.GetStub().PutState(person.Id, bytes); err != nil {
		return errors.WithMessage(err, "AddPerson PutState failed")
	}
	if err := r.personCounter.IncreasePersonId(ctx); err != nil {
		return errors.WithMessage(err, "AddPerson Increate Person Id failed")
	}
	return nil
}

func (r *RosterContract) ModifyPerson(ctx contractapi.TransactionContextInterface, id, name string, age, gender uint8, birth int64, birthPlace string) error {
	bytes, _ := ctx.GetStub().GetState(id)
	if bytes == nil {
		return errors.New("Modify Person failed, id not exist.")
	}
	var person datatype.Person
	if err := json.Unmarshal(bytes, &person); err != nil {
		return errors.WithMessage(err, "Modify Person unmarshal failed")
	}
	person.Name = name
	person.Age = age
	person.Gender = gender
	person.Birth = time.Unix(birth, 0)
	person.BirthPlace = birthPlace
	newBytes, _ := json.Marshal(&person)
	if err := ctx.GetStub().PutState(person.Id, newBytes); err != nil {
		return errors.WithMessage(err, "Modify person put state failed")
	}
	return nil
}

func (r *RosterContract) DelPerson(ctx contractapi.TransactionContextInterface, id string) error {
	return ctx.GetStub().DelState(id)
}

func (r *RosterContract) SearchPerson(ctx contractapi.TransactionContextInterface, name string) ([]datatype.Person, error) {
	queryString := fmt.Sprintf("{\"selector\":{\"name\":{\"$regex\": \"%s\"}}}", name)
	queryResults, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer queryResults.Close()
	var persons []datatype.Person
	for queryResults.HasNext() {
		var person datatype.Person
		queryRsp, err := queryResults.Next()
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(queryRsp.Value, &person)
		if err != nil {
			return nil, err
		}
		persons = append(persons, person)
	}
	return persons, nil
}

/*
** History operations
 */
func (r *RosterContract) AddHistory(ctx contractapi.TransactionContextInterface, historyId, content, comment string) error {
	var history datatype.History
	if bytes, _ := ctx.GetStub().GetState(historyId); bytes != nil {
		json.Unmarshal(bytes, &history)
	} else {
		history = datatype.History{
			Id:      historyId,
			Records: []datatype.Record{},
		}
	}
	record := datatype.Record{
		Id:      strconv.Itoa(len(history.Records)),
		Content: content,
		Comment: comment,
		Time:    time.Now(),
	}
	history.Records = append(history.Records, record)
	newBytes, err := json.Marshal(&history)
	if err != nil {
		return errors.WithMessage(err, "Marshal new history failed")
	}
	return ctx.GetStub().PutState(historyId, newBytes)
}

func (r *RosterContract) GetHistory(ctx contractapi.TransactionContextInterface, historyId string) ([]datatype.Record, error) {
	bytes, _ := ctx.GetStub().GetState(historyId)
	if bytes == nil {
		return nil, errors.New("History not exist")
	}
	var history datatype.History
	if err := json.Unmarshal(bytes, &history); err != nil {
		return nil, errors.WithMessage(err, "Unmarshal history failed")
	}
	return history.Records, nil
}