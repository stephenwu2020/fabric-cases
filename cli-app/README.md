# 使用fabric-go-sdk开发简单的app

## 目标
Fabric的官方案例提供了许多简单的chaincode，我们选择abstore作为本次示例，结合fabric-sdk-go，开发一款简单的cli应用。

## 技术概要
1. 使用[devnet](https://github.com/stephenwu2020/fabric-cases/tree/master/devnet)提供的网络，默认部署了abstore
2. sdk选用go的版本
3. cli的框架选用[cobra](github.com/spf13/cobra)

## abstore的基本功能
- Init，初始化用户余额，设置用户a,b各有100块
- invoke，a对b转账x块
- delete，清空用户的数据
- query，查询用户余额

## 启动网络
- 启动执行:
  ```
	make new
	```
- 销毁执行:
  ```
	make destroy
	```
- 更详细的用法，查看builder.sh提供的命令:
  ```
	cd devnet
	./builder.sh
	```
   
## cli简要说明
运行，查看帮助：
```
$ cd cli-app
$ go run .
Simple abstore cli app

Usage:
  hf-simple-app [flags]
  hf-simple-app [command]

Available Commands:
  help        Help about any command
  query       Query someone's balance
  transfer    transfer someone's balance to another one

Flags:
  -h, --help   help for hf-simple-app

Use "hf-simple-app [command] --help" for more information about a command.
```
Cli只提供了两个命令，query和transfer

query的用法(查询a的余额):
```
$ go run . query -n a
Balance is: 100
```

transfer的用法(a给b转10块):
```
$ go run . transfer -f a -t b -v 10
Transfer success!
$ go run . query -n a
Balance is: 90
```

## 关于fabric-go-sdk的使用
devenet使用fabric-ca的形式启动，同时创建了用于连接fabric的配置文件:organizations/peerOrganizations/org1.develop.com/connection-org1.yaml

cli/sdk/sdk.go文件中可以看到，sdk启动时，读取该文件，创建钱包，创建channel网络的对象，创建对应chaincode对象：
```
func Init() {
	os.Setenv("DISCOVERY_AS_LOCALHOST", "true")

	logging.SetLevel("fabsdk/core", logging.ERROR)

	wallet, err = gateway.NewFileSystemWallet("wallet")
	if err != nil {
		fmt.Printf("Failed to create wallet: %s\n", err)
		os.Exit(1)
	}

	if !wallet.Exists("appUser") {
		err = populateWallet(wallet)
		if err != nil {
			fmt.Printf("Failed to populate wallet contents: %s\n", err)
			os.Exit(1)
		}
	}

	ccpPath := filepath.Join(
		"..",
		"fixtures",
		"organizations",
		"peerOrganizations",
		"org1.develop.com",
		"connection-org1.yaml",
	)

	gw, err := gateway.Connect(
		gateway.WithConfig(config.FromFile(filepath.Clean(ccpPath))),
		gateway.WithIdentity(wallet, "appUser"),
	)
	if err != nil {
		fmt.Printf("Failed to connect to gateway: %s\n", err)
		os.Exit(1)
	}

	network, err = gw.GetNetwork(channel)
	if err != nil {
		fmt.Printf("Failed to get network: %s\n", err)
		os.Exit(1)
	}

	contract = network.GetContract(chaincode)
}
```

与fabric进行交互时，有两种类型的方法：
1. 只读:
   ```
   func ChannelQuery(fcn string, args ...string) ([]byte, error) {
      result, err := contract.EvaluateTransaction(fcn, args...)
      return result, err
   }
   ```
2. 读写:
   ```
   func ChannelExecute(fcn string, args ...string) ([]byte, error) {
      result, err := contract.SubmitTransaction(fcn, args...)
      return result, err
   }
   ```

以上就是fabric-sdk-go最基本的使用方法
