package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
	"net/url"
	"time"
)

var applications ApplicationGroup
var applications_config ApplicationGroupConfig

type ServiceError struct {
	Msg string
}

func (e *ServiceError) Error() string {
	var time = time.Now()
	return fmt.Sprintf("at %v, %s",
		time, e.Msg)
}

func Error(msg string) error{
	return  &ServiceError{msg}
}


func WsServer(ws *websocket.Conn) {
	var err error
	var appId,uid string
	if appId,uid,err = acceptClientToken(ws);err != nil{
		errMsg := err.Error()
		websocket.Message.Send(ws, errMsg)
		ws.Close()
	}
	config := applications_config[appId]
	for {
		var receiveMsg string

		if err = websocket.Message.Receive(ws, &receiveMsg); err != nil {
			applications.removeConn(appId,uid,ws)
			break
		}
		fmt.Println(receiveMsg)
		if config.MessageTransferApi != ""{
				go http.PostForm(config.MessageTransferApi,url.Values{"uid": {uid}, "message" : {receiveMsg}})
		}
	}
}

func StaticServer(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "demo/demo.html")
	// staticHandler := http.FileServer(http.Dir("./"))
	// staticHandler.ServeHTTP(w, req)
	return
}

func main() {

	var err error

	http.Handle("/", websocket.Handler(WsServer))
	http.Handle("/api/create-channel", &ApiServer{ApiName : "create-channel"})//create a ChannelService
	http.Handle("/api/push", &ApiServer{ApiName : "push"})
	http.Handle("/api/broadcast", &ApiServer{ApiName : "broadcast"})
	http.Handle("/api/get-channel", &ApiServer{ApiName : "get-channel"})
	http.Handle("/api/close-channel", &ApiServer{ApiName : "close-channel"})//close a specific ChannelService
	http.Handle("/api/app-status", &ApiServer{ApiName : "app-status"})//online num and live connection num

	http.HandleFunc("/demo", StaticServer)

	fmt.Println("listen on port 8002")
	//TODO read application info from db or file
	//TODO offer a init commad to reload application info file
	applications = make(ApplicationGroup)
	applications_config = make(ApplicationGroupConfig)

	if err = initServer(); err != nil {
		panic(err.Error())
	}
	fmt.Println(applications)

	if err = http.ListenAndServe(":8002", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
