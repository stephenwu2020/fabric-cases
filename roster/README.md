# 花名册
使用Hyperledger Fabric技术搭建的花名册应用。

## 拉取仓库
1. 仓库地址: [fabric-cases](https://github.com/stephenwu2020/fabric-cases)
2. 首先，阅读devnet的介绍，确保devnet能够正常启动、关闭

## 启动网络
```
cd roster
make new
```

## 运行 roster cli
查看帮助手册:
```
$ go run .
Roster records person info on Hyperledger Fabric blockchain

Usage:
  roster [flags]
  roster [command]

Available Commands:
  group       Mark person with group tag
  help        Help about any command
  history     Person history operations
  person      Add person, delete person and modify person
  version     Print the version number of roster

Flags:
  -h, --help   help for roster

Use "roster [command] --help" for more information about a command.
```
添加人物:
```
$ go run . person add -n "小明"
Add person with name: 小明
Add person success.
```
搜索人物:
```
$ go run . person search -n "小明"
Search person by name: 小明
(0) 小明: 
  Id		: person_1000
  Name		: 小明
  Age		: 0
  Gender	: 0
  Birth		: 2020-08-25
  BirthPlace	: 
  GroupTags	: []
  HistroyId	: history_1000
```
添加记录:
```
$ go run . history add --id person_1000 -c "上班" -C "勤奋工作好榜样"
Add record into history: 上班 勤奋工作好榜样
Add history record success.
```
列出记录:
```
$ go run . history show --id person_1000
Show history with person id: person_1000
(0) Record: 
  Id		: 0
  Content	: 上班
  Comment	: 勤奋工作好榜样
  Time		: 2020-08-25
```

更多功能请查看帮助手册.

## 网络操作
- 新建: make new
- 销毁: make destroy
- 启动: make up
- 关闭: make down
- 更新链码: make upgradeCC

更多操作细节请查看devnet的介绍.