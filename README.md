# go-myBlockchain

##实现简单的区块链

##目的通过新的现实需求来掌握Go语言的知识点

#组成部分

- 链式结构
- 实现一个简单的http server, 对外暴露访问接口
- 效果展示

#步骤

- 创建block(创世区块)
- 创建blockchain(链式结构)
- 创建http server

#创建 Block
- 新建工程demochain
- 创建block文件
- 创建block结构体与相关函数


#创建http Server
- 访问：http://localhost:8888/blockchain/get
- 发送数据：http://localhost:8888/blockchain/write?data=Send 1 BTC To lzc
