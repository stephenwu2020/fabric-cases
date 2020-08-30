package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func help() {
	fmt.Println("\n*What's your request?")
	fmt.Println("  - list: list all raft nodes")
	fmt.Println("  - quit: quit")
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
		fmt.Println()
		switch input {
		case "list":
			cluster.ListRaftNodes()
		case "quit":
			fmt.Println("Bye!")
			os.Exit(0)
		default:
			fmt.Println("No such service, guy!")
		}
		time.Sleep(3 * time.Second)
	}
}
