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
	fmt.Println("  - transfer: leader ship transfer, vote for new candidate")
	fmt.Println("  - set:  set random value")
	fmt.Println("  - get:  get value")
	fmt.Println("  - down: leader down")
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
		case "transfer":
			if err := cluster.Transfer(); err == nil {
				log.Println("transfer success")
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
		case "down":
			if err := cluster.LeaderDown(); err != nil {
				log.Println(err)
			} else {
				log.Println("Leader has been shutdown")
			}
		default:
			fmt.Println("No such service, guy!")
		}
		time.Sleep(3 * time.Second)
	}
}
