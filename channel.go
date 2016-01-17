package main

import (
	"fmt"
	//"errors"
	"github.com/larspensjo/config"
	"golang.org/x/net/websocket"
	"net/http"
	"net/url"
	"strconv"
)

const MAX_CLIENT_CONN, DEFAULT_CLIENT_CONN = 5, 1

//const TOKEN_METHOD_COOKIE,TOKEN_METHOD_GET = 1,2

//refer to one user with multiple client connections
type ChannelService struct {
	Uid   string
	Token string
	Conns []*websocket.Conn
}

type ApplicationConfig struct {
	AppId     string
	AppSecret string
	//TokenMethod int
	MaxClientConn      int
	GetConnectApi      string
	LoseConnectApi     string
	MessageTransferApi string
}

//refer to one application with multiple channel services
type Application struct {
	Services map[string]ChannelService
	Config   ApplicationConfig
}

type ApplicationGroup map[string]Application
type ApplicationGroupConfig map[string]ApplicationConfig

func initServer() error {
	configMap := make(map[string]ApplicationConfig)

	cfg, err := config.ReadDefault("config.ini")
	if err != nil {
		return Error("unable to open config file or wrong fomart")
	}
	sections := cfg.Sections()
	if len(sections) == 0 {
		return Error("no app config")
	}

	for _, section := range sections {
		if section != "DEFAULT" {
			sectionData, _ := cfg.SectionOptions(section)
			tmp := make(map[string]string)
			for _, key := range sectionData {
				value, err := cfg.String(section, key)
				if err == nil {
					tmp[key] = value
				}
			}
			maxClientConn, _ := strconv.Atoi(tmp["MaxClientConn"])
			configMap[section] = ApplicationConfig{tmp["AppId"], tmp["AppSecret"], maxClientConn, tmp["GetConnectApi"], tmp["LoseConnectApi"], tmp["MessageTransferApi"]}
		}
	}
	fmt.Println(configMap)

	valid_config := make(map[string]ApplicationConfig)

	for appid, appconfig := range configMap {
		// if appconfig.TokenMethod != TOKEN_METHOD_GET && appconfig.TokenMethod != TOKEN_METHOD_COOKIE{
		//   return Error("invalid TokenMethod appid: " + appid )
		// }
		if appconfig.MaxClientConn < 1 || appconfig.MaxClientConn > MAX_CLIENT_CONN {
			return Error("invalid MaxClientConn appid: " + appid)
		}
		channelGroup := make(map[string]ChannelService)
		app := Application{channelGroup, appconfig}

		applications[appid] = app
		valid_config[appid] = appconfig
	}
	applications_config = valid_config
	return nil
}

func acceptClientToken(ws *websocket.Conn) (string, string, error) {
	appId := ws.Request().FormValue("app_id")
	if appId == "" {
		return "", "", Error("app_id missing")
	}
	config, ok := applications_config[appId]
	if ok == false {
		return "", "", Error("app_id invalid")
	}
	uid := ws.Request().FormValue("uid")
	if uid == "" {
		return "", "", Error("uid missing")
	}
	channelService, ok := applications[appId].Services[uid]
	if ok == false {
		return "", "", Error("uid invalid")
	}
	token := ws.Request().FormValue("token")
	if token == "" {
		return "", "", Error("token missing")
	}
	if token != channelService.Token {
		return "", "", Error("invalid token")
	}
	conns := channelService.Conns
	if len(conns) > (config.MaxClientConn - 1) {
		//close the first conn
		conns[0].Close()
		conns = conns[1:]
	}
	conns = append(conns, ws)
	if len(conns) == 1 && config.GetConnectApi != "" {
		go http.PostForm(config.GetConnectApi, url.Values{"uid": {uid}})
	}
	channelService.Conns = conns
	applications[appId].Services[uid] = channelService
	return appId, uid, nil
}

//remove lost conns
func (this *ApplicationGroup) removeConn(appId, uid string, ws *websocket.Conn) error {
	cs, ok := (*this)[appId].Services[uid]
	if ok == false {
		return nil
	}
	config := applications_config[appId]
	for i, conn := range cs.Conns {
		if ws == conn {
			cs.Conns = append((cs.Conns)[:i], (cs.Conns)[i+1:]...)
			break
		}
	}
	(*this)[appId].Services[uid] = cs
	if len(cs.Conns) == 0 && config.LoseConnectApi != "" {
		go http.PostForm(config.LoseConnectApi, url.Values{"uid": {uid}})
	}
	return nil
}

func (this *ApplicationGroup) removeChannel(appId, uid string) error {
	cs := (*this)[appId].Services[uid]
	for _, conn := range cs.Conns {
		conn.Close()
	}
	delete((*this)[appId].Services, uid)
	return nil
}
