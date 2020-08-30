package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func help() {
	fmt.Println("\n* What's your request?")
	fmt.Println("  - list: list all raft nodes")
	fmt.Println("  - boot: bootstrap a new raft nodes")
	fmt.Println("  - transfer: leader ship transfer, vote for new candidate")
	fmt.Println("  - set:  set random value")
	fmt.Println("  - get:  get value")
	fmt.Println("  - down: bring down leader node")
	fmt.Println("  - quit: quit")
}

func output(a ...interface{}) {
	fmt.Println("Output:")
	fmt.Print("  ")
	fmt.Println(a...)
}

func ReadInput(cluster *CoffeeCluster) {
	var input string
	for {
		help()
		fmt.Printf("\nInput: ")
		_, err := fmt.Scanln(&input)
		if err != nil {
			log.Fatal("Read user input fail", err)
		}
		switch input {
		case "list":
			fmt.Println("Output:")
			for _, node := range cluster.CoffeeNodes {
				if node.RaftNode != nil {
					fmt.Println("  ", node.RaftNode)
				}
			}
		case "quit":
			output("Bye")
			os.Exit(0)
		case "boot":
			err := cluster.BootCaffeeNode()
			time.Sleep(3 * time.Second)
			if err != nil {
				output("Boot raft node failed", err)
			}
			output("Boot success")
		case "transfer":
			if err := cluster.Transfer(); err != nil {
				output("Transfer leadership fail", err)
			} else {
				output("Transfer leadership success")
			}
		case "set":
			if err := cluster.Set(); err != nil {
				output("Set value fail", err)
			} else {
				output("Set value success")
			}
		case "get":
			val, err := cluster.Get()
			if err != nil {
				output("Get value fail", err)
			} else {
				output("Value is:", val)
			}
		case "down":
			err := cluster.LeaderDown()
			time.Sleep(3 * time.Second)
			if err != nil {
				output("Bring down leader fail", err)
			} else {
				output("Bring down leader success")
			}
		default:
			output("Not such service, guy!")
		}
		time.Sleep(1 * time.Second)
	}
}
