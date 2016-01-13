package main

import (
	//"fmt"
  //"errors"
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
	CreateConnectApi string
	LoseConnectApi string
}

//refer to one application with multiple channel services
type ChannelServiceGroup struct{
	Services []ChannelService
	Config ChannelServiceGroupConfig
}


func initServer() error{
  config := make(map[string]ChannelServiceGroupConfig)
  //test application
  config["test"] = ChannelServiceGroupConfig{"test","test_secret",TOKEN_METHOD_GET,2,"http://localhost","http://localhost"}

  for appid,appconfig := range config{
    if appconfig.TokenMethod != TOKEN_METHOD_GET && appconfig.TokenMethod != TOKEN_METHOD_COOKIE{
      return Error("invalid TokenMethod appid: " + appid )
    }
    if appconfig.MaxClientConn < 1 || appconfig.MaxClientConn > MAX_CLIENT_CONN{
      return Error("invalid MaxClientConn appid: " + appid )
    }
    channelGroup := []ChannelService{}
    app := ChannelServiceGroup{channelGroup, appconfig}
    applications = append(applications,app)
  }
  return nil
}
