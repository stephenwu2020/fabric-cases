package main

import (
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {
	cluster := NewCoffeeCluster()

	// Clean all files
	os.RemoveAll(cluster.RootDir)
	time.Sleep(3 * time.Second)

	// Start Ming's raft node, and boostrap cluser
	_, err := cluster.CreateRaftNode(cluster.CoffeeNodes[0], true)
	if err != nil {
		log.Fatal(err)
	}

	// Need some time to vote a leader
	time.Sleep(5 * time.Second)

	// Interact with user
	ReadInput(cluster)

	terminate := make(chan os.Signal, 1)
	signal.Notify(terminate, os.Interrupt)
	<-terminate
	log.Println("Exit.")
}
