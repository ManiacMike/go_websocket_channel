package main

import (
	"fmt"
	"net/http"
	"golang.org/x/net/websocket"
)

//handle api request from api
type ApiServer struct{
	ApiName string
	AppId string
}

func (this *ApiServer) ServeHTTP(w http.ResponseWriter, r *http.Request){
	if err := this.CheckParams(r);err != nil{
		fmt.Fprint(w, err.Error())
		return
	}

	switch this.ApiName {
	case "create":
		this.Create(w,r)
	case "push":
		this.Push(w,r)
	case "broadcast":
		this.Broadcast(w,r)
	case "get-channel":
		this.GetChannel(w,r)
	case "close":
		this.Close(w,r)
	case "status":
		this.Status(w,r)
	default:
		fmt.Fprint(w, "Invalid api")
	}
}

func (this *ApiServer) CheckParams(r *http.Request) error{
	appId := r.PostFormValue("app_id");
	if appId == ""{
		return Error("app_id missing")
	}
	appSecret := r.PostFormValue("app_secret");
	if appSecret == ""{
		return Error("app_secret missing")
	}
	config, ok := applications_config[appId];
	if ok == false{
		return Error("app_id invalid")
	}
	if config.AppSecret != appSecret{
		return Error("app_secret invalid")
	}
	this.AppId = appId
	return nil
}


func (this *ApiServer ) Create(w http.ResponseWriter, r *http.Request) error{
	uid := r.PostFormValue("uid");
	if uid == ""{
		return Error("uid missing")
	}
	token := GenerateId();
	fmt.Println("token: ",token)
	channelService := ChannelService{Uid:uid,Token:token}
	appId := this.AppId
	applications[appId].Services[uid] = channelService
  this.Success("success",w)
	return nil
}

func (this *ApiServer ) Push(w http.ResponseWriter, r *http.Request) error{
	uid := r.PostFormValue("uid");
	if uid == ""{
		return Error("uid missing")
	}
	appId := this.AppId
	channelService,ok := applications[appId].Services[uid]
	if ok == false{
		return Error("uid invalid")
	}
	msg := r.PostFormValue("message");
	for i,conn := range channelService.Conns{
		if err := websocket.Message.Send(conn, msg); err != nil {
			channelService.Conns = append((channelService.Conns)[:i], (channelService.Conns)[i+1:]...)
			applications[appId].Services[uid] = channelService
		}
	}
  this.Success("success",w)
	return nil
}

func (this *ApiServer ) Broadcast(w http.ResponseWriter, r *http.Request) error{
	appId := this.AppId
	channelServices := applications[appId]
	msg := r.PostFormValue("message");
	for uid,cs := range channelServices.Services{
		for i,conn := range cs.Conns{
			if err := websocket.Message.Send(conn, msg); err != nil {
				cs.Conns = append((cs.Conns)[:i], (cs.Conns)[i+1:]...)
				applications[appId].Services[uid] = cs
			}
		}
	}
	this.Success("success",w)
	return nil
}

func (this *ApiServer ) GetChannel(w http.ResponseWriter, r *http.Request) error{
	uid := r.PostFormValue("uid");
	if uid == ""{
		return Error("uid missing")
	}
	appId := this.AppId
	channelService,ok := applications[appId].Services[uid]
	if ok == false{
		fmt.Fprint(w, "{\"code\":101,\"msg\":\"channel not created\"}")
	}
	msg := fmt.Sprintf("{\"uid\":\"%v\",\"token\":\"%v\",\"connections\":%d}",channelService.Uid,channelService.Token,len(channelService.Token))
	fmt.Fprint(w, msg)
	return nil
}

func (this *ApiServer ) Close(w http.ResponseWriter, r *http.Request) error{
	uid := r.PostFormValue("uid");
	if uid == ""{
		return Error("uid missing")
	}
	appId := this.AppId
	channelService,ok := applications[appId].Services[uid]
	if ok == false{
		return Error("uid invalid")
	}
	for _,conn := range channelService.Conns{
		conn.Close()
	}
	delete(applications[appId].Services, uid)
  this.Success("success",w)
	return nil
}

func (this *ApiServer ) Status(w http.ResponseWriter, r *http.Request) error{
	appId := this.AppId
	application := applications[appId]
	channelNum := len(application.Services)
	returnMsg := fmt.Sprintf("\"code\":0,\"channelNum\":%d",channelNum)
  fmt.Fprint(w, returnMsg)
	return nil
}

func (this *ApiServer ) Success(msg string, w http.ResponseWriter){
	returnMsg := fmt.Sprintf("\"code\":0,\"msg\":%v",msg)
	fmt.Fprint(w, returnMsg)
}
