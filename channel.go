package main

import (
	//"fmt"
	"golang.org/x/net/websocket"
)


const MAX_CLIENT_CONN,DEFAULT_CLIENT_CONN = 5,1

const TOKEN_METHOD_COOKIE,TOKEN_METHOD_GET = 1,2

//refer to one user with multiple client connections
type ChannelService struct{
	uid string
	token string
	con []*websocket.Conn
}

type ChannelServiceGroupConfig struct{
	AppId string
	AppSecret string
	TokenMethod int
	MaxClientConn int
	GetConnectApi string
	LoseConnectApi string
}

//refer to one application with multiple channel services
type ChannelServiceGroup struct{
	Services []ChannelService
	Config ChannelServiceGroupConfig
}
