## SDK Package
Package提供了访问devnet的方法。

## 如何使用
1. 启动devenet
   ```
   cd ../devnet
   ./builder.sh new
   ```
2. 初始化
   ```
   err := Init()
   ```
3. 读数据调用query
   ```
   bytes, err := ChannelQuery("query", []byte("a"))
   ```
4. 写数据调用execute
   ```
   _, err = ChannelExecute("invoke", []byte("a"), []byte("b"), []byte("10"))
   ```

详情查看test:
```
func TestSDK(t *testing.T) {
	err := Init()
	if err != nil {
		t.Error("Init faild", err)
	}
	amount1, err := getA()
	if err != nil {
		t.Error(err)
	}

	if _, err = ChannelExecute("invoke", []byte("a"), []byte("b"), []byte("10")); err != nil {
		t.Error("Invoke faild", err)
	}

	amount2, err := getA()
	if err != nil {
		t.Error(err)
	}
	println("Amount of A:", amount1)
	println("Amount of A:", amount2)
	if amount1-amount2 != 10 {
		t.Error(errors.New("A subtract 10 faild"))
	}
}
```

## 说明
其一，sdk用的官方fabric-sdk-go的修改版本: [stephenwu2020/fabric-sdk-go](https://github.com/stephenwu2020/fabric-sdk-go)，在go.mod里：
   ```
   replace github.com/hyperledger/fabric-sdk-go => github.com/stephenwu2020/fabric-sdk-go v1.0.0-beta3-modify
   ```
  只是为了添加两个方便的方法，接收byte数组

    ```
    func (c *Contract) EvaluateTransactionWithBytes(name string, args ...[]byte) ([]byte, error) {
      txn, err := c.CreateTransaction(name)

      if err != nil {
        return nil, err
      }

      return txn.EvaluateWithBytes(args...)

    }
    func (c *Contract) SubmitTransactionWithBytes(name string, args ...[]byte) ([]byte, error) {
      txn, err := c.CreateTransaction(name)

      if err != nil {
        return nil, err
      }

      return txn.SubmitWithBytes(args...)
    }
    ```
其二，sdk的初始化wallet用的是内存版本，由于开发阶段网络频繁创建，生成本地的钱包也要删除，使用内存版本免去这一烦恼:
```
wallet = gateway.NewInMemoryWallet()
```