package main

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"time"

	"github.com/hashicorp/raft"
	boltdb "github.com/hashicorp/raft-boltdb"
	"github.com/pkg/errors"
)

type CoffeeNode struct {
	ID   string
	Bind string
	Dir  string
}

type CoffeeCluster struct {
	RootDir     string
	CoffeeNodes []CoffeeNode
	RaftNodes   []*raft.Raft
}

func NewCoffeeCluster() *CoffeeCluster {
	nodes := []CoffeeNode{
		{ID: "ming", Bind: ":11000", Dir: "./nodes/ming"},
		{ID: "hong", Bind: ":11001", Dir: "./nodes/hong"},
		{ID: "peng", Bind: ":11002", Dir: "./nodes/peng"},
		{ID: "long", Bind: ":11003", Dir: "./nodes/long"},
		{ID: "hu", Bind: ":11004", Dir: "./nodes/hu"},
	}
	return &CoffeeCluster{
		RootDir:     "./nodes",
		CoffeeNodes: nodes,
		RaftNodes:   []*raft.Raft{},
	}
}

func (cf *CoffeeCluster) CreateRaftNode(coffeeNode *CoffeeNode, runCluster bool) (*raft.Raft, error) {
	config := raft.DefaultConfig()
	config.LocalID = raft.ServerID(coffeeNode.ID)
	addr, err := net.ResolveTCPAddr("tcp", coffeeNode.Bind)
	if err != nil {
		return nil, errors.WithMessage(err, "Resole tcp addr fail")
	}
	transport, err := raft.NewTCPTransport(coffeeNode.Bind, addr, 3, 10*time.Second, os.Stderr)
	if err != nil {
		return nil, errors.WithMessage(err, "Create tcp transport fail")
	}
	snapshots, err := raft.NewFileSnapshotStore(coffeeNode.Dir, 2, os.Stderr)
	if err != nil {
		return nil, errors.WithMessage(err, "Create snapshot fail")
	}
	logStore, err := boltdb.NewBoltStore(filepath.Join(coffeeNode.Dir, "raft.db"))
	if err != nil {
		return nil, errors.WithMessage(err, "Create boltstore fail")
	}
	fsm := NewFSM()
	raftNode, err := raft.NewRaft(config, fsm, logStore, logStore, snapshots, transport)
	if err != nil {
		return nil, errors.WithMessage(err, "Create raft nodefail")
	}

	if runCluster {
		configuration := raft.Configuration{
			Servers: []raft.Server{{ID: config.LocalID, Address: transport.LocalAddr()}},
		}
		cluster := raftNode.BootstrapCluster(configuration)
		if cluster.Error() != nil {
			return nil, errors.WithMessage(cluster.Error(), "Create cluster failed")
		}
	}
	cf.RaftNodes = append(cf.RaftNodes, raftNode)
	return raftNode, nil
}

func (cf *CoffeeCluster) ListRaftNodes() {
	for _, r := range cf.RaftNodes {
		fmt.Printf("%s\n", r)
	}
}
