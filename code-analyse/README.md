# Orderer启动时做了什么？

日常开发，部署时，Orderer被编译成bin程序，运行在Docker容器中。或许有人会问：Orderer可以运行在主机中吗？如何启动呢？启动的时候做了什么？

## 本地启动

阅读Orderer源码后，发现，启动的入口位于 fabric/orderer/common/server/main.go中:
```
// Main is the entry point of orderer process
func Main() {
  fullCmd := kingpin.MustParse(app.Parse(os.Args[1:]))
  ...
}
```
注意，这个Main函数的package并不是main，因此不可以直接启动。我们可以创建新的GO项目，调用此Main函数，启动Orderer。

我们的工程位于[fabric-cases](https://github.com/stephenwu2020/fabric-cases)下的[code-analyse](https://github.com/stephenwu2020/fabric-cases/tree/master/code-analyse)目录下。

启动之前，请确保fabic-case的根目录下，已经下载bin文件，否则，请查看根目录的README，执行:
```
make
```

然后进入本项目的目录:
```
cd code-analyse
```

编写Main函数: (main.go)
```
package main

import (
	"github.com/hyperledger/fabric/orderer/common/server"
)

func main() {
	server.Main()
}
```
Main函数非常简单，直接调用orderer的Main函数。不出意外的话，程序会奔溃，因为配置文件并未设置！！！

## 配置

Orderer的源码中，很多参数都设置了默认值，比如配置文件order.yaml，channel id，msp路径等等。在源码中，发现默认的配置在fabric/orderer/common/server/testdata/下，本工程目录下的3个配置文件就是从那里复制过来的:
```
$ ls -l *.yaml
-rw-r--r--  1 stephen  staff   6113 Sep  3 08:58 configtx.yaml
-rw-r--r--  1 stephen  staff   3589 Sep  3 08:58 examplecom-config.yaml
-rw-r--r--  1 stephen  staff  14066 Sep  3 08:58 orderer.yaml
```

在网络的部署过程中，头两个步骤分别是生成认证文件、生成创世块，我们把这个过程写如Makefile的crypto指令中，执行指令生成crypto文件夹和genesisblock:
```
$ make crypto
Clean materials...
Done
Generate materials...
example.com
2020-09-03 10:40:56.867 HKT [common.tools.configtxgen] main -> INFO 001 Loading configuration
2020-09-03 10:40:56.988 HKT [common.tools.configtxgen.localconfig] completeInitialization -> INFO 002 orderer type: solo
2020-09-03 10:40:56.988 HKT [common.tools.configtxgen.localconfig] Load -> INFO 003 Loaded configuration: /Users/stephen/Develop/Blockchain/fabric-cases/code-analyse/configtx.yaml
2020-09-03 10:40:56.997 HKT [common.tools.configtxgen] doOutputBlock -> INFO 004 Generating genesis block
2020-09-03 10:40:56.998 HKT [common.tools.configtxgen] doOutputBlock -> INFO 005 Writing genesis block
Done!
```
然后我们启动程序:
```
$ go run .
2020-09-03 10:41:58.889 HKT [localconfig] completeInitialization -> WARN 001 General.GenesisMethod should be replaced by General.BootstrapMethod
2020-09-03 10:41:58.889 HKT [localconfig] completeInitialization -> WARN 002 General.GenesisFile should be replaced by General.BootstrapFile
2020-09-03 10:41:58.889 HKT [localconfig] completeInitialization -> INFO 003 FileLedger.Prefix unset, setting to hyperledger-fabric-ordererledger
2020-09-03 10:41:58.889 HKT [localconfig] completeInitialization -> INFO 004 Kafka.Version unset, setting to 0.10.2.0
2020-09-03 10:41:58.889 HKT [orderer.common.server] prettyPrintStruct -> INFO 005 Orderer config values:
        General.ListenAddress = "127.0.0.1"
        General.ListenPort = 7050
        General.TLS.Enabled = false
        General.TLS.PrivateKey = "/Users/stephen/Develop/Blockchain/fabric-cases/code-analyse/crypto/ordererOrganizations/example.com/orderers/127.0.0.1.example.com/tls/server.key"
        General.TLS.Certificate = "/Users/stephen/Develop/Blockchain/fabric-cases/code-analyse/crypto/ordererOrganizations/example.com/orderers/127.0.0.1.example.com/tls/server.crt"
        General.TLS.RootCAs = [/Users/stephen/Develop/Blockchain/fabric-cases/code-analyse/crypto/ordererOrganizations/example.com/orderers/127.0.0.1.example.com/tls/ca.crt]
        General.TLS.ClientAuthRequired = false
        General.TLS.ClientRootCAs = []
        General.Cluster.ListenAddress = ""
        General.Cluster.ListenPort = 0
        General.Cluster.ServerCertificate = ""
        General.Cluster.ServerPrivateKey = ""
        General.Cluster.ClientCertificate = ""
        General.Cluster.ClientPrivateKey = ""
        General.Cluster.RootCAs = []
        General.Cluster.DialTimeout = 5s
        General.Cluster.RPCTimeout = 7s
        General.Cluster.ReplicationBufferSize = 20971520
        General.Cluster.ReplicationPullTimeout = 5s
        General.Cluster.ReplicationRetryTimeout = 5s
        General.Cluster.ReplicationBackgroundRefreshInterval = 5m0s
        General.Cluster.ReplicationMaxRetries = 12
        General.Cluster.SendBufferSize = 10
        General.Cluster.CertExpirationWarningThreshold = 168h0m0s
        General.Cluster.TLSHandshakeTimeShift = 0s
        General.Keepalive.ServerMinInterval = 1m0s
        General.Keepalive.ServerInterval = 2h0m0s
        General.Keepalive.ServerTimeout = 20s
        General.ConnectionTimeout = 0s
        General.GenesisMethod = "file"
        General.GenesisFile = "genesisblock"
        General.BootstrapMethod = "file"
        General.BootstrapFile = "/Users/stephen/Develop/Blockchain/fabric-cases/code-analyse/genesisblock"
        General.Profile.Enabled = false
        General.Profile.Address = "0.0.0.0:6060"
        General.LocalMSPDir = "/Users/stephen/Develop/Blockchain/fabric-cases/code-analyse/crypto/ordererOrganizations/example.com/orderers/127.0.0.1.example.com/msp"
        General.LocalMSPID = "SampleOrg"
        General.BCCSP.ProviderName = "SW"
        General.BCCSP.SwOpts.SecLevel = 256
        General.BCCSP.SwOpts.HashFamily = "SHA2"
        General.BCCSP.SwOpts.Ephemeral = true
        General.BCCSP.SwOpts.FileKeystore.KeyStorePath = ""
        General.BCCSP.SwOpts.DummyKeystore =
        General.BCCSP.SwOpts.InmemKeystore =
        General.Authentication.TimeWindow = 15m0s
        General.Authentication.NoExpirationChecks = false
        FileLedger.Location = "/var/hyperledger/production/orderer"
        FileLedger.Prefix = "hyperledger-fabric-ordererledger"
        Kafka.Retry.ShortInterval = 5s
        Kafka.Retry.ShortTotal = 10m0s
        Kafka.Retry.LongInterval = 5m0s
        Kafka.Retry.LongTotal = 12h0m0s
        Kafka.Retry.NetworkTimeouts.DialTimeout = 10s
        Kafka.Retry.NetworkTimeouts.ReadTimeout = 10s
        Kafka.Retry.NetworkTimeouts.WriteTimeout = 10s
        Kafka.Retry.Metadata.RetryMax = 3
        Kafka.Retry.Metadata.RetryBackoff = 250ms
        Kafka.Retry.Producer.RetryMax = 3
        Kafka.Retry.Producer.RetryBackoff = 100ms
        Kafka.Retry.Consumer.RetryBackoff = 2s
        Kafka.Verbose = false
        Kafka.Version = 0.10.2.0
        Kafka.TLS.Enabled = false
        Kafka.TLS.PrivateKey = ""
        Kafka.TLS.Certificate = ""
        Kafka.TLS.RootCAs = []
        Kafka.TLS.ClientAuthRequired = false
        Kafka.TLS.ClientRootCAs = []
        Kafka.SASLPlain.Enabled = false
        Kafka.SASLPlain.User = ""
        Kafka.SASLPlain.Password = ""
        Kafka.Topic.ReplicationFactor = 3
        Debug.BroadcastTraceDir = ""
        Debug.DeliverTraceDir = ""
        Consensus = map[snapdir:/var/hyperledger/production/orderer/etcdraft/snapshot waldir:/var/hyperledger/production/orderer/etcdraft/wal]
        Operations.ListenAddress = "127.0.0.1:8443"
        Operations.TLS.Enabled = false
        Operations.TLS.PrivateKey = ""
        Operations.TLS.Certificate = ""
        Operations.TLS.RootCAs = []
        Operations.TLS.ClientAuthRequired = false
        Operations.TLS.ClientRootCAs = []
        Metrics.Provider = "disabled"
        Metrics.Statsd.Network = "udp"
        Metrics.Statsd.Address = "127.0.0.1:8125"
        Metrics.Statsd.WriteInterval = 30s
        Metrics.Statsd.Prefix = ""
        ChannelParticipation.Enabled = false
        ChannelParticipation.RemoveStorage = false
2020-09-03 10:41:58.933 HKT [orderer.common.server] Main -> INFO 006 Not bootstrapping the system channel because of existing channels
2020-09-03 10:41:58.940 HKT [orderer.common.server] selectClusterBootBlock -> INFO 007 Cluster boot block is bootstrap (genesis) block; Blocks Header.Number system-channel=0, bootstrap=0
2020-09-03 10:41:58.945 HKT [orderer.common.server] Main -> INFO 008 Starting with system channel: testchannelid, consensus type: solo
2020-09-03 10:41:58.957 HKT [orderer.consensus.solo] HandleChain -> WARN 009 Use of the Solo orderer is deprecated and remains only for use in test environments but may be removed in the future.
2020-09-03 10:41:58.957 HKT [orderer.commmon.multichannel] Initialize -> INFO 00a Starting system channel 'testchannelid' with genesis block hash baf6470148beed52854637f803db5473458c5c7de06cdc714944d13e1a9d12d5 and orderer type solo
2020-09-03 10:41:58.957 HKT [orderer.common.server] Main -> INFO 00b Starting orderer:
 Version: latest
 Commit SHA: development build
 Go version: go1.14.3
 OS/Arch: darwin/amd64
2020-09-03 10:41:58.957 HKT [orderer.common.server] Main -> INFO 00c Beginning to serve requests
```
Orderer启动成功啦！！！

## 配置初始化

Orderer的Main函数中，有几个比较重要的步骤，第一个是配置初始化:
```
conf, err := localconfig.Load()
```
函数会加载当前目录下的orderer.yaml文件:
```
config := viper.New()
coreconfig.InitViper(config, "orderer")
config.SetEnvPrefix(Prefix)
```
怎么知道是当前目录下呢？InitViper函数里，若FABRIC_CFG_PATH未设置，则将当前目录添加进搜索路径：
```
func InitViper(v *viper.Viper, configName string) error {
	var altPath = os.Getenv("FABRIC_CFG_PATH")
	if altPath != "" {
		// If the user has overridden the path with an envvar, its the only path
		// we will consider

		if !dirExists(altPath) {
			return fmt.Errorf("FABRIC_CFG_PATH %s does not exist", altPath)
		}

		AddConfigPath(v, altPath)
	} else {
		// If we get here, we should use the default paths in priority order:
		//
		// *) CWD
		// *) /etc/hyperledger/fabric

		// CWD
		AddConfigPath(v, "./")

		// And finally, the official path
		if dirExists(OfficialPath) {
			AddConfigPath(v, OfficialPath)
		}
	}

	// Now set the configuration file.
	if v != nil {
		v.SetConfigName(configName)
	} else {
		viper.SetConfigName(configName)
	}

	return nil
}
```
我们并没有设置该环境变量的值，因此会找当前目录下的orderer.yaml.

查看orderer.yaml，才发现很多选项很熟悉，对应了docker容器的环境变量，例如，监听端口:
```
# Listen port: The port on which to bind to listen.
ListenPort: 7050
```
TLS证书:
```
# TLS: TLS settings for the GRPC server.
TLS:
    Enabled: false
    # PrivateKey governs the file location of the private key of the TLS certificate.
    PrivateKey: crypto/ordererOrganizations/example.com/orderers/127.0.0.1.example.com/tls/server.key
    # Certificate governs the file location of the server TLS certificate.
    Certificate: crypto/ordererOrganizations/example.com/orderers/127.0.0.1.example.com/tls/server.crt
    RootCAs:
      - crypto/ordererOrganizations/example.com/orderers/127.0.0.1.example.com/tls/ca.crt
    ClientAuthRequired: false
    ClientRootCAs:
```
Genesis Block的位置、msp的路径:
```
GenesisFile: genesisblock

# LocalMSPDir is where to find the private crypto material needed by the
# orderer. It is set relative here as a default for dev environments but
# should be changed to the real location in production.
LocalMSPDir: crypto/ordererOrganizations/example.com/orderers/127.0.0.1.example.com/msp

# LocalMSPID is the identity to register the local MSP material with the MSP
# manager. IMPORTANT: The local MSP ID of an orderer needs to match the MSP
# ID of one of the organizations defined in the orderer system channel's
# /Channel/Orderer configuration. The sample organization defined in the
# sample configuration provided has an MSP ID of "SampleOrg".
LocalMSPID: SampleOrg
```

## GenesisBlock，Channel
随后是读取、解释创世块:
```
	var bootstrapBlock *cb.Block
	if conf.General.BootstrapMethod == "file" {
		bootstrapBlock = file.New(conf.General.BootstrapFile).GenesisBlock()
		if err := onboarding.ValidateBootstrapBlock(bootstrapBlock, cryptoProvider); err != nil {
			logger.Panicf("Failed validating bootstrap block: %v", err)
		}

		// Are we bootstrapping with a genesis block (i.e. bootstrap block number = 0)?
		// If yes, generate the system channel with a genesis block.
		if len(lf.ChannelIDs()) == 0 && bootstrapBlock.Header.Number == 0 {
			logger.Info("Bootstrapping the system channel")
			initializeBootstrapChannel(bootstrapBlock, lf)
		} else if len(lf.ChannelIDs()) > 0 {
			logger.Info("Not bootstrapping the system channel because of existing channels")
		} else {
			logger.Infof("Not bootstrapping the system channel because the bootstrap block number is %d (>0), replication is needed", bootstrapBlock.Header.Number)
		}
	} else if conf.General.BootstrapMethod != "none" {
		logger.Panicf("Unknown bootstrap method: %s", conf.General.BootstrapMethod)
	}
```
解释、创建系统channel:
```
// select the highest numbered block among the bootstrap block and the last config block if the system channel.
	sysChanConfigBlock := extractSystemChannel(lf, cryptoProvider)
	clusterBootBlock := selectClusterBootBlock(bootstrapBlock, sysChanConfigBlock)

	// determine whether the orderer is of cluster type
	var isClusterType bool
	if clusterBootBlock == nil {
		logger.Infof("Starting without a system channel")
		isClusterType = true
	} else {
		sysChanID, err := protoutil.GetChannelIDFromBlock(clusterBootBlock)
		if err != nil {
			logger.Panicf("Failed getting channel ID from clusterBootBlock: %s", err)
		}

		consensusTypeName := consensusType(clusterBootBlock, cryptoProvider)
		logger.Infof("Starting with system channel: %s, consensus type: %s", sysChanID, consensusTypeName)
		_, isClusterType = clusterTypes[consensusTypeName]
	}
```

## 系统管理服务
Orderer.yaml配置中的Operations字段，用于启动系统管理服务:
```
opsSystem := newOperationsSystem(conf.Operations, conf.Metrics)
if err = opsSystem.Start(); err != nil {
		logger.Panicf("failed to start operations subsystem: %s", err)
}
```
此服务监听8443端口:
```
Operations:
    # host and port for the operations server
    ListenAddress: 127.0.0.1:8443
```
分析源码，这个一个http服务，提供的API有:
```
s.mux.Handle("/metrics", s.handlerChain(promhttp.Handler(), s.options.TLS.Enabled))
s.mux.Handle("/logspec", s.handlerChain(httpadmin.NewSpecHandler(), s.options.TLS.Enabled))
s.mux.Handle("/healthz", s.handlerChain(s.healthHandler, false))
s.mux.Handle("/version", s.handlerChain(versionInfo, false))
```
大多是反馈系统状态的请求。我们尝试调用这些API：
```
curl localhost:8443/version
{"CommitSHA":"development build","Version":"latest"}

curl localhost:8443/healthz
{"status":"OK","time":"2020-09-03T11:17:48.613373+08:00"}
```

## Orderer服务
最后，启动orderer服务:
```
grpcServer := initializeGrpcServer(conf, serverConfig)
logger.Info("Beginning to serve requests")
if err := grpcServer.Start(); err != nil {
  logger.Fatalf("Atomic Broadcast gRPC server has terminated while serving requests due to: %v", err)
}
```
这是一个google grpc服务:
```
// Start starts the underlying grpc.Server
func (gServer *GRPCServer) Start() error {
	// if health check is enabled, set the health status for all registered services
	if gServer.healthServer != nil {
		for name := range gServer.server.GetServiceInfo() {
			gServer.healthServer.SetServingStatus(
				name,
				healthpb.HealthCheckResponse_SERVING,
			)
		}

		gServer.healthServer.SetServingStatus(
			"",
			healthpb.HealthCheckResponse_SERVING,
		)
	}
	return gServer.server.Serve(gServer.listener)
}
```
它监听的是7050的端口:
```
# Listen port: The port on which to bind to listen.
ListenPort: 7050
```
当grpc服务接收到信息，它将交给ab.AtomicBroadcastServer处理:
```
server := NewServer(
		manager,
		metricsProvider,
		&conf.Debug,
		conf.General.Authentication.TimeWindow,
		mutualTLS,
		conf.General.Authentication.NoExpirationChecks,
	)
ab.RegisterAtomicBroadcastServer(grpcServer.Server(), server)
```
它处理两大类的消息：
1. 接收客户端发起的交易，打包
2. 广播打包好的区块
```
// Broadcast receives a stream of messages from a client for ordering
func (s *server) Broadcast(srv ab.AtomicBroadcast_BroadcastServer) error {
	logger.Debugf("Starting new Broadcast handler")
	defer func() {
		if r := recover(); r != nil {
			logger.Criticalf("Broadcast client triggered panic: %s\n%s", r, debug.Stack())
		}
		logger.Debugf("Closing Broadcast stream")
	}()
	return s.bh.Handle(&broadcastMsgTracer{
		AtomicBroadcast_BroadcastServer: srv,
		msgTracer: msgTracer{
			debug:    s.debug,
			function: "Broadcast",
		},
	})
}

// Deliver sends a stream of blocks to a client after ordering
func (s *server) Deliver(srv ab.AtomicBroadcast_DeliverServer) error {
	logger.Debugf("Starting new Deliver handler")
	defer func() {
		if r := recover(); r != nil {
			logger.Criticalf("Deliver client triggered panic: %s\n%s", r, debug.Stack())
		}
		logger.Debugf("Closing Deliver stream")
	}()

	policyChecker := func(env *cb.Envelope, channelID string) error {
		chain := s.GetChain(channelID)
		if chain == nil {
			return errors.Errorf("channel %s not found", channelID)
		}
		// In maintenance mode, we typically require the signature of /Channel/Orderer/Readers.
		// This will block Deliver requests from peers (which normally satisfy /Channel/Readers).
		sf := msgprocessor.NewSigFilter(policies.ChannelReaders, policies.ChannelOrdererReaders, chain)
		return sf.Apply(env)
	}
	deliverServer := &deliver.Server{
		PolicyChecker: deliver.PolicyCheckerFunc(policyChecker),
		Receiver: &deliverMsgTracer{
			Receiver: srv,
			msgTracer: msgTracer{
				debug:    s.debug,
				function: "Deliver",
			},
		},
		ResponseSender: &responseSender{
			AtomicBroadcast_DeliverServer: srv,
		},
	}
	return s.dh.Handle(srv.Context(), deliverServer)
}
```

## 监听
最后，检查系统的监听服务，系统管理服务和grpc服务都在:
```
$ make listener
order-ana 7673 stephen    9u     IPv4 0xc1d0e97275765439        0t0                 TCP 127.0.0.1:7050 (LISTEN)
order-ana 7673 stephen   17u     IPv4 0xc1d0e97275764079        0t0                 TCP 127.0.0.1:8443 (LISTEN)
```