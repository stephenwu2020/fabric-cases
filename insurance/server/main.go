package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/stephenwu2020/fabric-cases/insurance/server/sdk"
)

func main() {
	fmt.Println("hi")
	if err := sdk.Init(); err != nil {
		log.Fatal("SDK init fail", err)
	}
	args := [][]byte{}
	rsp, err := sdk.ChannelQuery("Invoke", args)
	if err != nil {
		log.Fatal("Query fail", err)
	}
	log.Info(rsp.Payload)
}
