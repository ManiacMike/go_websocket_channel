package main

import (
	"fmt"
	"net/http"
	// "golang.org/x/net/websocket"
)

//handle api request from api
type ApiServer struct{
	ApiName string
}

func (this *ApiServer) ServeHTTP(w http.ResponseWriter, r *http.Request){
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

func (this *ApiServer ) Create(w http.ResponseWriter, r *http.Request){
  fmt.Fprint(w, "create api")
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
