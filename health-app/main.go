package main

import (
	"encoding/json"
	"log"

	"github.com/stephenwu2020/fabric-cases/chaincode/health/datatype"
	"github.com/stephenwu2020/fabric-cases/sdk"
)

func main() {
	if err := sdk.Init(); err != nil {
		log.Fatal("Init sdk failed", err)
	}
	bytes, err := sdk.ChannelQuery("Intro")
	if err != nil {
		log.Fatal("Call Intro failed", err)
	}
	var intro datatype.HealthIntro
	if err := json.Unmarshal(bytes, &intro); err != nil {
		log.Fatal("Marshal Intro failed", err)
	}
	log.Printf("Health Intro: %+v", intro)
}
