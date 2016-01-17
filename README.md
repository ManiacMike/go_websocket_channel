# go_websocket_channel
go web socket gateway server work with http server

channel服务指代本项目

# 服务流程

* http服务端调用channel服务注册该用户uid的channel

* api返回一个token，客户端用于使用websocket连接channel服务器

* 接下去的http服务端可以利用channel服务用单人或者广播的方式进行推送

# 提供给app服务器的api

* ip:port/api/create-channel 创建一个channel
* ip:port/api/get-channel 获取channel信息
* ip:port/api/push 向一个用户推送
* ip:port/api/broadcast 向所有的app用户推送
* ip:port/api/get-channel 获取channel信息
* ip:port/api/close-channel 关闭channel
* ip:port/api/app-status 获取app的状态

# app配置

* AppId string
*	AppSecret string
*	MaxClientConn int  一个uid的最大连接数
*	GetConnectApi string  建立socket连接回调app服务地址
*	LoseConnectApi string  失去所有socket回调app服务地址
*	MessageTransferApi string  收到socket消息调用app服务地址

# 命令
todo
