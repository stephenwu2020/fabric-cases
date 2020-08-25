package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/stephenwu2020/fabric-cases/insurance/server/sdk"
)

func main() {
	fmt.Println("hi")
	if err := sdk.Init(); err != nil {
		log.Error("SDK init fail", err)
	}
}
