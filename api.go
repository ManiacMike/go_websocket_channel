package main

import (
	"fmt"
	"net/http"
	// "golang.org/x/net/websocket"
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
	case "check":
		this.Check(w,r)
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


func (this *ApiServer ) Create(w http.ResponseWriter, r *http.Request){
	uid := r.PostFormValue("uid");
	if uid == ""{
		fmt.Fprint(w, "Invalid uid")
		return
	}
	token := GenerateId();
	channelService := ChannelService{Uid:uid,Token:token}
	appId := this.AppId
	applications[appId].Services[uid] = channelService
  fmt.Fprint(w, "create channel success on " + this.AppId)
}

func (this *ApiServer ) Push(w http.ResponseWriter, r *http.Request){
  fmt.Fprint(w, "push api")
}

func (this *ApiServer ) Broadcast(w http.ResponseWriter, r *http.Request){
  fmt.Fprint(w, "Broadcast api")
}

func (this *ApiServer ) Check(w http.ResponseWriter, r *http.Request){
  fmt.Fprint(w, "Check api")
}

func (this *ApiServer ) Close(w http.ResponseWriter, r *http.Request){
  fmt.Fprint(w, "Close api")
}

func (this *ApiServer ) Status(w http.ResponseWriter, r *http.Request){
  fmt.Fprint(w, "Status api")
}
