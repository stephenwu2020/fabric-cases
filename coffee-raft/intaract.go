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
	fmt.Println("  - boot: bootstrap a new raft nodes")
	fmt.Println("  - set:  set random value")
	fmt.Println("  - get:  get value")
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
		case "boot":
			fmt.Println("Boostrap a new raft node...")
			if err := cluster.BootCaffeeNode(); err != nil {
				log.Println("Bootstrap raft node failed", err)
			}
		case "set":
			if err := cluster.Set(); err != nil {
				log.Println(err)
			} else {
				log.Println("Value has been set")
			}
		case "get":
			val, err := cluster.Get()
			if err != nil {
				log.Println(err)
			} else {
				log.Println("Value is", val)
			}
		default:
			fmt.Println("No such service, guy!")
		}
		time.Sleep(3 * time.Second)
	}
}
