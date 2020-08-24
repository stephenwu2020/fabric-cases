package main

import (
	"github.com/stephenwu2020/fabric-cases/roster/cmd"
	"github.com/stephenwu2020/fabric-cases/roster/sdk"
)

func main() {
	sdk.Init()
	cmd.Execute()
}
