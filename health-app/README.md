# 小明的健康记录
今天是周末，阳光明媚。

铃铃铃......电话铃声响起，小明迷迷糊糊的接过电话。

"小明，你到了没？大伙都到了，在等你呢..."

糟糕，睡过头了，小明忘了今天与哥们的聚会!

最近睡过头的事情常有发生。加班、通宵、酗酒、少运动，小明感觉糟糕透了，身体似乎渐渐失控。于是，小明想开发一款健康软件，借助Hyperledger Fabric的技术，追踪个人健康情况，及时发现问题。

## 开发环境设置
- chaincode源码位于[fabric-cases](https://github.com/stephenwu2020/fabric-cases)的chaincode/health目录
- 网络启动脚本位于[fabric-cases](https://github.com/stephenwu2020/fabric-cases)的health-app目录下,Makefile
- 应用访问程序也位于[fabric-cases](https://github.com/stephenwu2020/fabric-cases)的health-app目录下

## 最简单的chaincode
首先定义Health结构体，嵌套contractapi.Contract
```
type Health struct {
	contractapi.Contract
}
```
最简单的Smart Contract就写好了，在main函数启动它:
```
func main() {
	health := new(Health)
	chaincode, err := contractapi.NewChaincode(health)
	if err != nil {
		log.Fatal("Create chaincode failed", err)
	}
	if err := chaincode.Start(); err != nil {
		log.Fatal("Start chaincode failed", err)
	}
}
```
这个chaincode已经可以在网络中安装了!

## 介绍函数
小明写的第一个函数，是介绍这款应用的功能:
```
func (h Health) Intro(ctx contractapi.TransactionContextInterface) (*HealthIntro, error) {
	intro := &HealthIntro{
		Name:     "Health",
		Function: "Record health data, analyse health situation.",
		Version:  "0.0.1",
		Author:   "Ming",
	}
	return intro, nil
}
```
供应用程序调用的方法，需要接收类型contractapi.TransactionContextInterface的参数。在health-app应用程序中，调用该函数:
```
func main() {
	if err := sdk.Init(); err != nil {
		panic(err)
	}
	bytes, err := sdk.ChannelQuery("Intro")
	if err != nil {
		panic(err)
	}
	var intro datatype.HealthIntro

	if err := json.Unmarshal(bytes, &intro); err != nil {
		panic(err)
	}
	fmt.Printf("Health Intro: %+v\n", intro)
}
```
调用成功，看到返回的结果是:
```
$ go run .
Health Intro: {Name:Health Function:Record health data, analyse health situation. Version:0.0.1 Author:Ming}
```
好了，第一个函数完成！

## 随眠记录
睡眠的质量对人的影响是非常巨大的。良好的睡眠，使人精力充沛、充满活力。小明决定将睡眠纳入健康考核。

首先，定义记录睡眠状况的结构体:
```
const (
	SleepAtNoon = iota
	SleepAtNight
)

type HealthRecord struct {
	Type  int
	Start int64
	End   int64
}
```
睡眠的类型分为两种：午睡和晚睡。另外，记录开始时间和结束时间

chaincode定义添加记录AddSleepRecord的方法:
```
func (h Health) AddSleepRecord(ctx contractapi.TransactionContextInterface, sleepT int, start, end int64) error {
	record := datatype.HealthRecord{
		Type:  sleepT,
		Start: start,
		End:   end,
	}
	records, err := h.GetRecords(ctx)
	if err != nil {
		return errors.WithMessage(err, "Get records failed")
	}
	records = append(records, record)
	bytes, err := json.Marshal(&records)
	if err != nil {
		return errors.WithMessage(err, "Marshal recrds failed")
	}
	return ctx.GetStub().PutState(recordkey, bytes)
}
```
方法将新的记录添加到旧的记录列表里，ctx提供了访问状态数据库的接口:
```
ctx.GetStub().GetState()
ctx.GetStub().PutState()
```
在应用程序中调用:
```
	// Add record
	start := time.Now()
	end := start.Add(time.Hour)
	_, err = sdk.ChannelExecute(
		"AddSleepRecord",
		tobytes(datatype.SleepAtNoon),
		tobytes(start.Unix()),
		tobytes(end.Unix()),
	)
	if err != nil {
		panic(err)
	}

	// Get Records
	bytes, err = sdk.ChannelQuery("GetRecords")
	if err != nil {
		panic(err)
	}
	var records []datatype.HealthRecord
	if err := json.Unmarshal(bytes, &records); err != nil {
		panic(err)
	}
	for _, r := range records {
		fmt.Printf("%+v\n", r)
	}
```
输出结果为:
```
$ go run .
Health Intro: {Name:Health Function:Record health data, analyse health situation. Version:0.0.1 Author:Ming}
{Type:0 Start:1598600548 End:1598604148}
{Type:0 Start:1598600831 End:1598604431}
```
到此，睡眠记录添加成功，查询成功