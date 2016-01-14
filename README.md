# go_websocket_channel
go web socket gateway server work with http server

# 服务流程

1. 用户使用http登陆，服务端调用channel注册 uid的channel

php代码：
$channelService = ChannelService::getInstance();
$channelService -> setUid($uid); //注册允许建立连接的uid


2. http返回成功后，客户端带着uid或者token写入cookie用websocket连接channel服务器，长连接注册成功

js代码
var channelService = new ChannelService(token); //模版渲染写入或者ajax回调
var ws = channelService.open(); //建立web socket对象并返回
ws.onopen = function() {
};
ws.onclose = function() {
};


3. 接下去的web http请求中可以用单人或者广播的方式进行推送
php代码：
$channelService = ChannelService::getInstance();
$msg = '{"message":"test"}';
$channelService -> push($targetUid, $msg); //单人推送
$channelService -> broadcast($msg); //广播到所有连接
