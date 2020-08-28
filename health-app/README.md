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
	AtNoon = iota
	AtNight
)

type SleepType int

type Sleep struct {
	Type  SleepType
	Start time.Time
	End   time.Time
}
```
睡眠的类型分为两种：午睡和晚睡。另外，记录开始时间和结束时间


