package sdk

import (
	"encoding/json"
	"testing"

	"github.com/pkg/errors"
)

func TestSDK(t *testing.T) {
	err := Init()
	if err != nil {
		t.Error("Init faild", err)
	}
	amount1, err := getA()
	if err != nil {
		t.Error(err)
	}

	if _, err = ChannelExecute("invoke", []byte("a"), []byte("b"), []byte("10")); err != nil {
		t.Error("Invoke faild", err)
	}

	amount2, err := getA()
	if err != nil {
		t.Error(err)
	}
	println("Amount of A:", amount1)
	println("Amount of A:", amount2)
	if amount1-amount2 != 10 {
		t.Error(errors.New("A subtract 10 faild"))
	}
}

func getA() (int, error) {
	bytes, err := ChannelQuery("query", []byte("a"))
	if err != nil {
		return -1, errors.WithMessage(err, "Query failed")
	}
	var amount int
	if err := json.Unmarshal(bytes, &amount); err != nil {
		return -1, errors.WithMessage(err, "Marshal value faild")
	}
	return amount, nil
}
