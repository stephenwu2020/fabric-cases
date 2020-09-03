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

我们的工程位于[fabric-cases](https://github.com/stephenwu2020/fabric-cases)下的[code-analyse]()目录下