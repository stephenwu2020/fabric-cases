package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/hashicorp/raft"
	boltdb "github.com/hashicorp/raft-boltdb"
	"github.com/pkg/errors"
)

var key = "coffeenum"

type CoffeeNode struct {
	ID       string
	Bind     string
	Dir      string
	RaftNode *raft.Raft
	FSM      *CoffeeFSM
}

type CoffeeCluster struct {
	RootDir     string
	CoffeeNodes []*CoffeeNode
}

func NewCoffeeCluster() *CoffeeCluster {
	nodes := []*CoffeeNode{
		{ID: "ming", Bind: ":11000", Dir: "./nodes/ming"},
		{ID: "hong", Bind: ":11001", Dir: "./nodes/hong"},
		{ID: "peng", Bind: ":11002", Dir: "./nodes/peng"},
		{ID: "long", Bind: ":11003", Dir: "./nodes/long"},
		{ID: "hu", Bind: ":11004", Dir: "./nodes/hu"},
	}
	return &CoffeeCluster{
		RootDir:     "./nodes",
		CoffeeNodes: nodes,
	}

}

func (cf *CoffeeCluster) CreateRaftNode(coffeeNode *CoffeeNode, runCluster bool) (*raft.Raft, error) {
	config := raft.DefaultConfig()
	config.LogLevel = "info"
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
	coffeeNode.FSM = fsm
	coffeeNode.RaftNode = raftNode

	if runCluster {
		configuration := raft.Configuration{
			Servers: []raft.Server{{ID: config.LocalID, Address: transport.LocalAddr()}},
		}
		cluster := raftNode.BootstrapCluster(configuration)
		if cluster.Error() != nil {
			return nil, errors.WithMessage(cluster.Error(), "Create cluster failed")
		}
	}
	return raftNode, nil
}

func (cf *CoffeeCluster) ListRaftNodes() {
	for _, node := range cf.CoffeeNodes {
		if node.RaftNode != nil {
			fmt.Printf("%s\n", node.RaftNode)
		}
	}
}

func (cf *CoffeeCluster) GetLeader() *CoffeeNode {
	var target *CoffeeNode
	for _, rn := range cf.CoffeeNodes {
		if rn.RaftNode != nil && rn.RaftNode.State() == raft.Leader {
			target = rn
			break
		}
	}
	return target
}

func (cf *CoffeeCluster) BootCaffeeNode() error {
	var target *CoffeeNode
	for _, node := range cf.CoffeeNodes {
		if node.RaftNode == nil {
			target = node
			break
		}
	}
	if target == nil {
		return errors.New("Reach max limit")
	}

	raftNote, err := cf.CreateRaftNode(target, false)
	target.RaftNode = raftNote
	if err != nil {
		return errors.WithMessage(err, "Create raft node failed")
	}
	leader := cf.GetLeader()
	if leader == nil {
		return errors.New("Leader not found")
	}
	f := leader.RaftNode.AddVoter(raft.ServerID(target.ID), raft.ServerAddress(target.Bind), 0, 0)
	if f.Error() != nil {
		return errors.WithMessage(f.Error(), "Add voter fail")
	}
	return nil
}

func (cf *CoffeeCluster) Set() error {
	cm := &command{
		Op:    "set",
		Key:   key,
		Value: strconv.Itoa(rand.Intn(100)),
	}
	bytes, err := json.Marshal(cm)
	if err != nil {
		return err
	}
	leader := cf.GetLeader()
	if leader == nil {
		return errors.New("Leader not found")
	}
	f := leader.RaftNode.Apply(bytes, 10*time.Second)
	return f.Error()
}

func (cf *CoffeeCluster) Get() (string, error) {
	leader := cf.GetLeader()
	if leader == nil {
		return "", errors.New("Leader not found")
	}
	val, ok := leader.FSM.m[key]
	if !ok {
		return "", errors.New("Key not founds")
	}
	return val, nil
}

func (cf *CoffeeCluster) Transfer() error {
	leader := cf.GetLeader()
	if leader == nil {
		return errors.New("Leader not found")
	}
	f := leader.RaftNode.LeadershipTransfer()
	return f.Error()
}

func (cf *CoffeeCluster) LeaderDown() error {
	leader := cf.GetLeader()
	if leader == nil {
		return errors.New("Key not founds")
	}
	f := leader.RaftNode.RemoveServer(raft.ServerID(leader.ID), 0, 0)
	return f.Error()
}
