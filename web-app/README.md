# Fabric Web应用
针对chaincode fabcar设计的web应用

## 程序架构
1. Fabric
2. Chaincode Fabcar
3. SDK选取fabric-sdk-go
4. Web server，用go语言开发，web框架gin
5. 前端的App采取Vue

## 启动
1. 启动网络
    ```
    cd web-app
    make
    ```
2. 启动 Web Server
    ```
    cd web-app/server
    go run .
    ```
3. 访问 Web App: localhost:8000

## 功能
1. 打开首页，查询所有的car的信息
2. 点击+按钮，添加car

## 其他
1. 更新chaincode
   ```
   make upgradeCC
   ```
2. 销毁网络
   ```
   make destroy
   ```