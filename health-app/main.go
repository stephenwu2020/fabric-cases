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
	end := start.Add(time.Hour * 7)
	_, err = sdk.ChannelExecute(
		"AddSleepRecord",
		tobytes(datatype.SleepAtNight),
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

	// Report
	report(records)
}

func tobytes(v interface{}) []byte {
	res, _ := json.Marshal(v)
	return res
}

func isHealth(record *datatype.HealthRecord) bool {
	health := false
	diff := record.End - record.Start
	hour := int64(time.Hour.Seconds())
	switch record.Type {
	case datatype.SleepAtNoon:
		if diff >= hour/2 && diff <= hour {
			health = true
		}
	case datatype.SleepAtNight:
		if diff >= hour*6 && diff <= hour*8 {
			health = true
		}
	}
	return health
}

func report(records []datatype.HealthRecord) {
	total := len(records)
	if total == 0 {
		fmt.Println("No record, No Report!")
		return
	}

	healthNum := 0
	for _, r := range records {
		if isHealth(&r) {
			healthNum++
		}
	}
	ratio := float64(healthNum) / float64(total)
	fmt.Printf("Health ratio: %0.2f %%\n", ratio)
}
