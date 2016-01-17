# go_websocket_channel
go web socket gateway server work with http server

# 服务流程

* 用户使用http登陆app服务器，app服务端调用channel注册该用户uid的channel

* http返回成功后，客户端带着uid和token用websocket连接channel服务器，长连接注册成功

* 接下去的向app服务器的http请求中可以用单人或者广播的方式进行推送

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
