package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/stephenwu2020/fabric-cases/chaincode/roster/datatype"
	"github.com/stephenwu2020/fabric-cases/chaincode/roster/genid"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/pkg/errors"
)

var timeLayoutStr = "2006-01-02 15:04:05"

type RosterContract struct {
	contractapi.Contract
	personCounter *genid.PersonCounter
}

func NewRosterContract() *RosterContract {
	counter := genid.NewPersonCounter()
	return &RosterContract{
		personCounter: counter,
	}
}

func (r *RosterContract) Instantiate(ctx contractapi.TransactionContextInterface) error {
	r.personCounter.Init(ctx)
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
		Birth:     time.Now(),
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

func (r *RosterContract) ModifyPerson(ctx contractapi.TransactionContextInterface, id, name, age, gender, birth, birthPlace string) error {
	bytes, _ := ctx.GetStub().GetState(id)
	if bytes == nil {
		return errors.New("Modify Person failed, id not exist.")
	}
	var person datatype.Person
	if err := json.Unmarshal(bytes, &person); err != nil {
		return errors.WithMessage(err, "Modify Person unmarshal fail")
	}
	uintAge, err := strconv.ParseUint(age, 10, 8)
	if err != nil {
		return errors.WithMessage(err, "Parse age fail")
	}
	uintGender, err := strconv.ParseUint(gender, 10, 8)
	if err != nil {
		return errors.WithMessage(err, "Parse gender fail")
	}
	intBirth, err := strconv.ParseInt(birth, 10, 64)
	if err != nil {
		return errors.WithMessage(err, "Parse brith fail")
	}
	person.Name = name
	person.Age = uint8(uintAge)
	person.Gender = uint8(uintGender)
	person.Birth = time.Unix(intBirth, 0)
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
	persons := []datatype.Person{}
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

func (r *RosterContract) GetPersonById(ctx contractapi.TransactionContextInterface, id string) (*datatype.Person, error) {
	bytes, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, errors.WithMessage(err, "Get person by id fail")
	}
	person := &datatype.Person{}
	if err := json.Unmarshal(bytes, person); err != nil {
		return nil, errors.WithMessage(err, "Marsha person fail")
	}
	return person, nil
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

func (r *RosterContract) GetHistory(ctx contractapi.TransactionContextInterface, historyId string) (*datatype.History, error) {
	bytes, _ := ctx.GetStub().GetState(historyId)
	var history datatype.History
	if bytes == nil {
		history = datatype.History{
			Id:      historyId,
			Records: []datatype.Record{},
		}
		return &history, nil
	}
	if err := json.Unmarshal(bytes, &history); err != nil {
		return nil, errors.WithMessage(err, "Unmarshal history failed")
	}
	return &history, nil
}

func (r *RosterContract) ModifyHistory(ctx contractapi.TransactionContextInterface, historyId, recordId, content, comment string) error {
	history, err := r.GetHistory(ctx, historyId)
	if err != nil {
		return errors.WithMessage(err, "Get History fail")
	}
	rid, err := strconv.Atoi(recordId)
	if err != nil {
		return errors.WithMessage(err, "Parse record id fail")
	}
	if rid >= len(history.Records) {
		return errors.WithMessage(err, "Record out of range")
	}
	history.Records[rid] = datatype.Record{
		Id:      recordId,
		Time:    time.Now(),
		Content: content,
		Comment: comment,
	}
	newBytes, err := json.Marshal(&history)
	if err != nil {
		return errors.WithMessage(err, "Marshal new history failed")
	}
	return ctx.GetStub().PutState(historyId, newBytes)

}
