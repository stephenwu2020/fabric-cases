小明的咖啡馆生意红火，目前为止共有5家分店，分别由小明、小红、小朋、小龙、小虎打理。处理咖啡订单的是一个分布式系统，有5个节点同时运行，使用Raft共识机制。

让我们来看看如何使用Raft共识搭建一个分布式系统。

(源码位于[fabric-cases](https://github.com/stephenwu2020/fabric-cases)下的coffee-raft目录)

首先，为咖啡馆设计数据结构：
```
type CoffeeNode struct {
	ID       string
	Bind     string
	Dir      string
	RaftNode *raft.Raft
	FSM      *CoffeeFSM
}
```
- ID唯一标识咖啡馆
- Bind是端口
- Dir是数据存储位置
- RaftNode是实现Raft的节点
- FSM处理用户数据


5家分店，5个咖啡馆结构:
```
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
```

程序启动，默认小明的咖啡馆是唯一分店，其raft节点是唯一节点，同时也是系统的Leader:
```
func main() {
	cluster := NewCoffeeCluster()

	// Clean all files
	os.RemoveAll(cluster.RootDir)
	time.Sleep(1 * time.Second)

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
```
运行程序:
```
$ go run .
2020-08-31T03:19:44.774+0800 [INFO]  raft: initial configuration: index=0 servers=[]
2020-08-31T03:19:44.774+0800 [INFO]  raft: entering follower state: follower="Node at :11000 [Follower]" leader=
2020-08-31T03:19:46.731+0800 [WARN]  raft: heartbeat timeout reached, starting election: last-leader=
2020-08-31T03:19:46.731+0800 [INFO]  raft: entering candidate state: node="Node at :11000 [Candidate]" term=2
2020-08-31T03:19:46.747+0800 [INFO]  raft: election won: tally=1
2020-08-31T03:19:46.747+0800 [INFO]  raft: entering leader state: leader="Node at :11000 [Leader]"

* What's your request?
  - list: list all raft nodes
  - boot: bootstrap a new raft nodes
  - transfer: leader ship transfer, vote for new candidate
  - set:  set random value
  - get:  get value
  - down: bring down leader node
  - quit: quit

Input: 
```
从日志中看出，raft节点启动，端口号11000的节点当选为Leader。
输入list指令:
```
Input: list
Output:
   Node at :11000 [Leader]
```
当前只有一个raft节点

输入boot指令，启动小红的raft节点:
```
Input: boot
2020-08-31T03:41:24.596+0800 [INFO]  raft: initial configuration: index=0 servers=[]
2020-08-31T03:41:24.596+0800 [INFO]  raft: updating configuration: command=AddStaging server-id=hong server-addr=:11001 servers="[{Suffrage:Voter ID:ming Address::11000} {Suffrage:Voter ID:hong Address::11001}]"
2020-08-31T03:41:24.596+0800 [INFO]  raft: entering follower state: follower="Node at :11001 [Follower]" leader=
2020-08-31T03:41:24.603+0800 [INFO]  raft: added peer, starting replication: peer=hong
2020-08-31T03:41:24.611+0800 [WARN]  raft: failed to get previous log: previous-index=3 last-index=0 error="log not found"
2020-08-31T03:41:24.611+0800 [WARN]  raft: appendEntries rejected, sending older logs: peer="{Voter hong :11001}" next=1
2020-08-31T03:41:24.617+0800 [INFO]  raft: pipelining replication: peer="{Voter hong :11001}"
Output:
  Boot success
```
再次输入list查看节点信息:
```
Input: list
Output:
   Node at :11000 [Leader]
   Node at :11001 [Follower]
```
此刻网络中有一个Leader，一个Follower。现在假设Leader节点宕机，输入down:
```
Input: down
2020-08-31T03:44:38.250+0800 [INFO]  raft: updating configuration: command=RemoveServer server-id=ming server-addr= servers="[{Suffrage:Voter ID:hong Address::11001}]"
2020-08-31T03:44:38.263+0800 [INFO]  raft: removed ourself, shutting down
2020-08-31T03:44:38.263+0800 [INFO]  raft: aborting pipeline replication: peer="{Voter hong :11001}"
2020-08-31T03:44:40.963+0800 [WARN]  raft: heartbeat timeout reached, starting election: last-leader=:11000
2020-08-31T03:44:40.963+0800 [INFO]  raft: entering candidate state: node="Node at :11001 [Candidate]" term=3
2020-08-31T03:44:40.987+0800 [INFO]  raft: election won: tally=1
2020-08-31T03:44:40.987+0800 [INFO]  raft: entering leader state: leader="Node at :11001 [Leader]"
Output:
  Bring down leader success
```
从日志看出，新的Leader是11001端口的节点，输入list查看:
```
Input: list
Output:
   Node at :11000 [Shutdown]
   Node at :11001 [Leader]
```
11000由Leader变为Shutdown, 11001由Follower变为Leader。

现在输入set指令，设置一个随机数:
```
Input: set
Output:
  Set value success
```
输入get，读取刚才设置的值:
```
Input: get
Output:
  Value is: 71
```
客户端的请求，都由Leader处理:
```
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
```

这就是Raft共识的运行机制。总结有以下几点:

1. raft节点启动，推选出Leader
2. 客户端的请求，全部由Leader处理，同时由Leader广播给Follower
3. Leader意外宕机，Follower收不到Leader的心跳，则进入Candidate候选人状态，通知其他节点给自己投票
4. 获得一定数量票的候选人成为Leader，接管工作，系统正常运作