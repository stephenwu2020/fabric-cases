# Peer节点启动时做了什么
Peer节点的程序结构与Orderer稍有不同。首先，peer是一个cli应用:
```
func main() {
	// For environment variables.
	viper.SetEnvPrefix(common.CmdRoot)
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	// Define command-line flags that are valid for all peer commands and
	// subcommands.
	mainFlags := mainCmd.PersistentFlags()

	mainFlags.String("logging-level", "", "Legacy logging level flag")
	viper.BindPFlag("logging_level", mainFlags.Lookup("logging-level"))
	mainFlags.MarkHidden("logging-level")

	cryptoProvider := factory.GetDefault()

	mainCmd.AddCommand(version.Cmd())
	mainCmd.AddCommand(node.Cmd())
	mainCmd.AddCommand(chaincode.Cmd(nil, cryptoProvider))
	mainCmd.AddCommand(channel.Cmd(nil))
	mainCmd.AddCommand(lifecycle.Cmd(cryptoProvider))
	mainCmd.AddCommand(snapshot.Cmd(cryptoProvider))

	// On failure Cobra prints the usage message and error string, so we only
	// need to exit with a non-0 status
	if mainCmd.Execute() != nil {
		os.Exit(1)
	}
}
```
Peer节点的启动是peer程序中的node指令下的一条指令:
```
func Cmd() *cobra.Command {
	nodeCmd.AddCommand(startCmd())
	nodeCmd.AddCommand(resetCmd())
	nodeCmd.AddCommand(rollbackCmd())
	nodeCmd.AddCommand(pauseCmd())
	nodeCmd.AddCommand(resumeCmd())
	nodeCmd.AddCommand(rebuildDBsCmd())
	nodeCmd.AddCommand(upgradeDBsCmd())
	return nodeCmd
}

var nodeStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the node.",
	Long:  `Starts a node that interacts with the network.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			return fmt.Errorf("trailing args detected")
		}
		// Parsing of the command line is done so silence cmd usage
		cmd.SilenceUsage = true
		return serve(args)
	},
}
```
由于代码位于internal这个包之中，外部程序无法引用，无法通过代码调用，我们只好用命令行的方式启动节点。

在我们在fabric-cases根目录执行make，以下bin程序被下载到bin目录之下:
```
$ ls -l ../bin/
total 366200
-rwxr-xr-x  1 stephen  staff  18156272 Aug 10 16:31 configtxgen
-rwxr-xr-x  1 stephen  staff  16878512 Aug 10 16:31 configtxlator
-rwxr-xr-x  1 stephen  staff  13140952 Aug 10 16:31 cryptogen
-rwxr-xr-x  1 stephen  staff  17300752 Aug 10 16:31 discover
-rwxr-xr-x  1 stephen  staff  19387296 Aug 10 16:31 fabric-ca-client
-rwxr-xr-x  1 stephen  staff  22989504 Aug 10 16:31 fabric-ca-server
-rwxr-xr-x  1 stephen  staff  12171616 Aug 10 16:31 idemixgen
-rwxr-xr-x  1 stephen  staff  29051248 Aug 10 16:31 orderer
-rwxr-xr-x  1 stephen  staff  38403256 Aug 10 16:31 peer
```
当中的peer程序，就是peer源码编译生成的。我们可以直接调用:
```
../bin/peer 
Usage:
  peer [command]

Available Commands:
  chaincode   Operate a chaincode: install|instantiate|invoke|package|query|signpackage|upgrade|list.
  channel     Operate a channel: create|fetch|join|list|update|signconfigtx|getinfo.
  help        Help about any command
  lifecycle   Perform _lifecycle operations
  node        Operate a peer node: start|reset|rollback|pause|resume|rebuild-dbs|upgrade-dbs.
  version     Print fabric peer version.

Flags:
  -h, --help   help for peer

Use "peer [command] --help" for more information about a command.
```
要启动peer节点，可以调用 ../bin/peer node start，或者，我们也可以调用本目录下make指令:
```
$ make peer
../bin/peer node start
2020-09-03 15:16:30.891 HKT [main] InitCmd -> ERRO 001 Fatal error when initializing core config : Could not find config file. Please make sure that FABRIC_CFG_PATH is set to a path which contains core.yaml
make: *** [peer] Error 1
```
不出意料，程序报错了，提示找不到core.yaml。

阅读源码，找到core.yaml位于fabric/sampleconfig/configtx.yaml，复制到本目录，同时修改msg的路径:
```
# Path on the file system where peer will find MSP local configurations
mspConfigPath: crypto/peerOrganizations/example.com/users/Admin@example.com/msp/
```
然后我们再启动:
```
make peer
../bin/peer node start
2020-09-03 15:24:29.259 HKT [nodeCmd] serve -> INFO 001 Starting peer:
 Version: 2.2.0
 Commit SHA: 5ea85bc54
 Go version: go1.14.4
 OS/Arch: darwin/amd64
 Chaincode:
  Base Docker Label: org.hyperledger.fabric
  Docker Namespace: hyperledger
2020-09-03 15:24:29.260 HKT [peer] getLocalAddress -> INFO 002 Auto-detected peer address: 192.168.1.3:7051
2020-09-03 15:24:29.260 HKT [peer] getLocalAddress -> INFO 003 Host is 0.0.0.0 , falling back to auto-detected address: 192.168.1.3:7051
2020-09-03 15:24:29.267 HKT [nodeCmd] initGrpcSemaphores -> INFO 004 concurrency limit for endorser service is 2500
2020-09-03 15:24:29.267 HKT [nodeCmd] initGrpcSemaphores -> INFO 005 concurrency limit for deliver service is 2500
2020-09-03 15:24:29.293 HKT [ledgermgmt] NewLedgerMgr -> INFO 006 Initializing LedgerMgr
2020-09-03 15:24:29.427 HKT [ledgermgmt] NewLedgerMgr -> INFO 007 Initialized LedgerMgr
2020-09-03 15:24:29.437 HKT [gossip.service] New -> INFO 008 Initialize gossip with endpoint 192.168.1.3:7051
2020-09-03 15:24:29.439 HKT [gossip.gossip] New -> INFO 009 Creating gossip service with self membership of Endpoint: , InternalEndpoint: 192.168.1.3:7051, PKI-ID: bc25609ffe8f936a25a5eab13e2e481309e9e63cf28ac4f1fd09e13eda9f213f, Metadata: 
2020-09-03 15:24:29.440 HKT [gossip.gossip] New -> WARN 00a External endpoint is empty, peer will not be accessible outside of its organization
2020-09-03 15:24:29.440 HKT [gossip.gossip] start -> INFO 00b Gossip instance 192.168.1.3:7051 started
2020-09-03 15:24:29.440 HKT [lifecycle] InitializeLocalChaincodes -> INFO 00c Initialized lifecycle cache with 0 already installed chaincodes
2020-09-03 15:24:29.441 HKT [nodeCmd] computeChaincodeEndpoint -> INFO 00d Entering computeChaincodeEndpoint with peerHostname: 192.168.1.3
2020-09-03 15:24:29.441 HKT [nodeCmd] computeChaincodeEndpoint -> INFO 00e Exit with ccEndpoint: 192.168.1.3:7052
2020-09-03 15:24:29.441 HKT [nodeCmd] createChaincodeServer -> WARN 00f peer.chaincodeListenAddress is not set, using 192.168.1.3:7052
2020-09-03 15:24:29.445 HKT [sccapi] DeploySysCC -> INFO 010 deploying system chaincode 'lscc'
2020-09-03 15:24:29.446 HKT [sccapi] DeploySysCC -> INFO 011 deploying system chaincode 'cscc'
2020-09-03 15:24:29.447 HKT [sccapi] DeploySysCC -> INFO 012 deploying system chaincode 'qscc'
2020-09-03 15:24:29.447 HKT [sccapi] DeploySysCC -> INFO 013 deploying system chaincode '_lifecycle'
2020-09-03 15:24:29.447 HKT [nodeCmd] serve -> INFO 014 Deployed system chaincodes
2020-09-03 15:24:29.447 HKT [discovery] NewService -> INFO 015 Created with config TLS: false, authCacheMaxSize: 1000, authCachePurgeRatio: 0.750000
2020-09-03 15:24:29.448 HKT [nodeCmd] registerDiscoveryService -> INFO 016 Discovery service activated
2020-09-03 15:24:29.448 HKT [nodeCmd] serve -> INFO 017 Starting peer with ID=[jdoe], network ID=[dev], address=[192.168.1.3:7051]
2020-09-03 15:24:29.448 HKT [nodeCmd] serve -> INFO 018 Started peer with ID=[jdoe], network ID=[dev], address=[192.168.1.3:7051]
2020-09-03 15:24:29.448 HKT [kvledger] LoadPreResetHeight -> INFO 019 Loading prereset height from path [/var/hyperledger/production/ledgersData/chains]
2020-09-03 15:24:29.448 HKT [blkstorage] preResetHtFiles -> INFO 01a No active channels passed
```
Peer节点启动成功了!

## 配置初始化
在core.yaml，发现很多选项很熟悉，因为这些变量是docker容器的环境变量:
```
listenAddress: 0.0.0.0:7051

gossip:
  bootstrap: 127.0.0.1:7051

mspConfigPath: crypto/peerOrganizations/example.com/users/Admin@example.com/msp/

...
```
与Orderer的配置不同的是，peer的配置多了chaincode和ledger的配置。因为chaincode的生命周期、ledger都是peer节点管理的。

## 启动逻辑
总体上，peer的启动与orderer的启动，步调一致: 读取配置，创建服务，运行服务。

具体来说，peer节点的启动逻辑更加复杂，主要体现在chaincode服务和ledger的配置，以及policy管理、identity管理、gossip服务、snapshot服务等等

