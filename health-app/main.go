package main

import (
	"encoding/json"
	"fmt"

	"github.com/stephenwu2020/fabric-cases/chaincode/health/datatype"
	"github.com/stephenwu2020/fabric-cases/sdk"
)

func main() {
	if err := sdk.Init(); err != nil {
		panic(err)
	}
	bytes, err := sdk.ChannelQuery("Intro")
	if err != nil {
		panic(err)
	}
	var intro datatype.HealthIntro

	if err := json.Unmarshal(bytes, &intro); err != nil {
		panic(err)
	}
	fmt.Printf("Health Intro: %+v\n", intro)
}
