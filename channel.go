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
	Uid string
	Token string
	Conns []*websocket.Conn
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
	Services map[string]ChannelService
	Config ChannelServiceGroupConfig
}

type ChannelServer map[string]ChannelServiceGroup
type ChannelServerConfig map[string]ChannelServiceGroupConfig

func initServer() error{
  config := make(map[string]ChannelServiceGroupConfig)
  //test application
  config["test"] = ChannelServiceGroupConfig{"test","test_secret",TOKEN_METHOD_GET,2,"http://localhost","http://localhost"}

  valid_config := make(map[string]ChannelServiceGroupConfig)

  for appid,appconfig := range config{
    if appconfig.TokenMethod != TOKEN_METHOD_GET && appconfig.TokenMethod != TOKEN_METHOD_COOKIE{
      return Error("invalid TokenMethod appid: " + appid )
    }
    if appconfig.MaxClientConn < 1 || appconfig.MaxClientConn > MAX_CLIENT_CONN{
      return Error("invalid MaxClientConn appid: " + appid )
    }
    channelGroup := make(map[string]ChannelService)
    app := ChannelServiceGroup{channelGroup, appconfig}


    applications[appid] = app
    valid_config[appid] = appconfig
  }
  applications_config = valid_config
  return nil
}
