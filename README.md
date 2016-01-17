# go_websocket_channel
go web socket gateway server work with http server

# 服务流程

*用户使用http登陆app服务器，app服务端调用channel注册该用户uid的channel

*http返回成功后，客户端带着uid和token用websocket连接channel服务器，长连接注册成功

*接下去的向app服务器的http请求中可以用单人或者广播的方式进行推送
