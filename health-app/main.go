package main

import (
	"encoding/json"
	"fmt"

	"github.com/stephenwu2020/fabric-cases/chaincode/health/datatype"
	"github.com/stephenwu2020/fabric-cases/sdk"
)

func main() {
	// Init
	if err := sdk.Init(); err != nil {
		panic(err)
	}

	// Intro
	bytes, err := sdk.ChannelQuery("Intro")
	if err != nil {
		panic(err)
	}
	var intro datatype.HealthIntro

	if err := json.Unmarshal(bytes, &intro); err != nil {
		panic(err)
	}
	fmt.Printf("Health Intro: %+v\n", intro)

	// Add record
	bytes, err := sdk.ChannelExecute(
		"AddSleepRecord",
		tobytes(datatype.SleepAtNoon),
	)
}

func tobytes(v interface{}) []byte {
	res, _ := json.Marshal(v)
	return res
}
