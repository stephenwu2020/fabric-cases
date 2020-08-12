# 从test-network说起
相信大部分开发者接触Hyperledger Fabric，都会从官方教程fabric-samples中的例子入手。首先运行test-network下的测试网络，再运行各种教材案例。当我们开始进行开发chaincode的时候，问题就来了。

假设我们写了一段代码: 函数A...，然后进行测试，清除网络，重新启动网络；随后我们又写了一段代码：函数B...，清除网络，重新启动网络；发现有一个bug，清除网络，重新启动网络......

重启网络，test-network大致做了以下工作:
1. 关闭，删除container，删除docker网络
2. 清空证书文件
3. 重新生成证书文件
4. 创建docker网络，创建container
5. 创建channel，加入channel
6. 打包，安装，实例化chaincode

时间足够喝一杯咖啡。

咖啡不能多喝，所以需要配置专门开发chaincode的网络，具有以下特点：
1. 将节点的数据写入至本地文件，重启时读取本地文件，恢复链的状态
2. 重启不清除已经生成的证书文件
3. 修改chaincode之后，用peer chaincode upgrade升级，而不是重新安装
4. 只需要一个orderer节点，一个peer节点，简化网络架构

# 解释本网络
## 目录结构解释
首先运行Makefile，下载fabric的工具，拉取fabric镜像:
```
make
```
脚本运行成功之后，生成了两个目录：
1. bin: 工具bin文件
2. config: 配置文件

其他目录分别有：
1. chaincode: 链码文件
2. pkg:链码打包后存放的目录
3. productions: peer和orderer数据存放目录
4. scripts: 网络部署，channel设置，chaincode安装等工具脚本
5. builder.sh: 网络部署的cli工具

## 如何解决test-network存在的问题
第一，节点的数据存放在本地，容器重启后，数据还在，其中主要的设置在docker-compose.yaml文件中：
```
 peer0.org1.develop.com:
    ......
    volumes:
        - ./productions/peer:/var/hyperledger/production
    .......
```

```
  orderer.develop.com:
    ......
    volumes:
        - ./productions/orderer:/var/hyperledger/production/orderer
    ......
```

第二，builder.sh细化了网络的启动逻辑，查看帮助:
```
#$ ./builder.sh 
Usage: 
  ./builder.sh <cmd>
cmd: 
  - new
  - destroy
  - up
  - down
  - upgrade
  - network
  - channel
  - chaincode

```
指令解释:
- new: 创建全新网络，会清空之前全部网络的数据，证书，链的数据等等；
- destroy: 销毁网络，清空证书，链的数据等等；
- up: 启动网络，只是启动docker container
- down: 关闭网络，只是删除docker container
- upgrade: 更新链码
- network, channel, chaincode 分别对应网络，channel, chaincode的更细的功能，请查看帮助

第三，提供upgrade指令

chaincode修改后，运行upgrade，重新部署chaincode

# 使用示例

1 如果没有下载工具，请先执行:
```
make
```

2 新建全新的网络:
```
#$ ./builder.sh new
Removing network develop_basic
WARNING: Network develop_basic not found.
org1.develop.com
2020-08-11 09:55:23.165 HKT [common.tools.configtxgen] main -> INFO 001 Loading configuration
2020-08-11 09:55:23.199 HKT [common.tools.configtxgen.localconfig] completeInitialization -> INFO 002 orderer type: etcdraft
2020-08-11 09:55:23.200 HKT [common.tools.configtxgen.localconfig] completeInitial
......
c1
Blockchain info: {"height":2,"currentBlockHash":"jn8Npe0AfpjLiVYA3+EUb7Drvd7F7xIMB4q9KZcTaJA=","previousBlockHash":"0jPkgU/EY/4ZHft2BMBk2yprOFAGyqPwDwkbp3nMJLc="}
pkg already exist
Installed chaincodes on peer:
Package ID: abstore_1:f038023086151baad502e5b82165d3630ac2c482ea502244a8f52cc1f25c5d22, Label: abstore_1
PackageID is abstore_1:f038023086151baad502e5b82165d3630ac2c482ea502244a8f52cc1f25c5d22
{
	"approvals": {
		"Org1MSP": true
	}
}
Committed chaincode definitions on channel 'c1':
Name: mycc, Version: 1.0.0, Sequence: 1, Endorsement Plugin: escc, Validation Plugin: vscc
100
```

3 关闭网络
```
#$ ./builder.sh down
Stopping cli                    ... done
Stopping orderer.develop.com    ... done
Stopping peer0.org1.develop.com ... done
Removing cli                    ... done
Removing orderer.develop.com    ... done
Removing peer0.org1.develop.com ... done
Removing network develop_basic

```
4 启动网络
```
#$ ./builder.sh up
Creating network "develop_basic" with the default driver
Creating peer0.org1.develop.com ... done
Creating orderer.develop.com    ... done
Creating cli                    ... done

```

5 查询chaincode
```
#$ ./builder.sh chaincode queryCommit
Committed chaincode definitions on channel 'c1':
Name: mycc, Version: 1.0.0, Sequence: 1, Endorsement Plugin: escc, Validation Plugin: vscc
```
```
#$ ./builder.sh chaincode query
100

```

6 修改chaincode，注意./chaincode/abstore/go/abstore.go的55行，初始化a的值时，额外增加10
```
	fmt.Printf("Aval = %d, Bval = %d\n", Aval, Bval)

	Aval = Aval + 10

	// Write the state to the ledger
	err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
```

7 运行upgrade指令，更新链码
```
#$ ./builder.sh chaincode upgrade
upgrade chaincode...
options: version:1.0.1, sequence:2
fetch go dependency
/opt/gopath/src/github.com/hyperledger/chaincode/abstore/go /opt/gopath/src/github.com/hyperledger/fabric/peer
/opt/gopath/src/github.com/hyperledger/fabric/peer
package chaincode
Installed chaincodes on peer:
Package ID: abstore_1:5496f0b8dabffeb1e498582b78b676c62aeb3cf7ec263d68281320dc6da7f7cd, Label: abstore_1
Package ID: abstore_1:f038023086151baad502e5b82165d3630ac2c482ea502244a8f52cc1f25c5d22, Label: abstore_1
PackageID is abstore_1:f038023086151baad502e5b82165d3630ac2c482ea502244a8f52cc1f25c5d22
110
chaincode upgraded.

```
8 再次查询链码
```
#$ ./builder.sh chaincode query
110

#$ ./builder.sh chaincode queryCommit
Committed chaincode definitions on channel 'c1':
Name: mycc, Version: 1.0.1, Sequence: 2, Endorsement Plugin: escc, Validation Plugin: vscc
```

可以看到, a的值变为了110， chaincode版本已经是1.0.1了

9 销毁网络
```
#$ ./builder.sh destroy
```

## 后话
down命令只是停止、删除container，container的数据都存放在production目录，下次启动时，读取数据，读取证书文件，启动，恢复状态，时间大大的缩短了。另一点，chaincode的升级步骤被抽离出来了，为日常开发提供了便利。