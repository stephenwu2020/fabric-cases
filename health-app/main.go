package main

import (
	"encoding/json"
	"fmt"
	"time"

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
	start := time.Now()
	end := start.Add(time.Hour)
	_, err = sdk.ChannelExecute(
		"AddSleepRecord",
		tobytes(datatype.SleepAtNoon),
		tobytes(start.Unix()),
		tobytes(end.Unix()),
	)
	if err != nil {
		panic(err)
	}

	// Get Records
	bytes, err = sdk.ChannelQuery("GetRecords")
	if err != nil {
		panic(err)
	}
	var records []datatype.HealthRecord
	if err := json.Unmarshal(bytes, &records); err != nil {
		panic(err)
	}
	for _, r := range records {
		fmt.Printf("%+v\n", r)
	}

}

func tobytes(v interface{}) []byte {
	res, _ := json.Marshal(v)
	return res
}
