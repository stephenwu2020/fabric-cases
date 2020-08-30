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

	/*
			log.Println("Start Hong's raft node")
			raftNodeOfHong, err := createRaftNode(&nodes[1], false)
			_ = raftNodeOfHong
			if err != nil {
				log.Fatal(err)
			}

			log.Println("Join Hong's raft node")
			//f := raftNodeOfHong.AddVoter(raft.ServerID(nodes[0].ID), raft.ServerAddress(nodes[0].Bind), 0, 0)
			f := raftNodeOfMing.AddVoter(raft.ServerID(nodes[1].ID), raft.ServerAddress(nodes[1].Bind), 0, 0)
			if f.Error() != nil {
				log.Fatal(f.Error())
			}
		time.Sleep(5 * time.Second)

	*/
	// Interact with user
	ReadInput(cluster)

	terminate := make(chan os.Signal, 1)
	signal.Notify(terminate, os.Interrupt)
	<-terminate
	log.Println("Exit.")
}
