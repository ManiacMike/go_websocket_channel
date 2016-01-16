package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
	"time"
)

var applications ChannelServer
var applications_config ChannelServerConfig

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
	if err = acceptClientToken(ws);err != nil{
		ws.Close()
	}

	for {
		var receiveMsg string

		if err = websocket.Message.Receive(ws, &receiveMsg); err != nil {
			//fmt.Println("Can't receive,user ", uid, " lost connection")
			//CurrentUsers.Remove(uid)
			break
		}
		fmt.Println(receiveMsg)
		if err := websocket.Message.Send(ws, receiveMsg); err != nil {
			//fmt.Println("Can't send user ", user.uid, " lost connection")
			//CurrentUsers.Remove(user.uid)
			break
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
	http.Handle("/api/create", &ApiServer{ApiName : "create"})//create a ChannelService
	http.Handle("/api/push", &ApiServer{ApiName : "push"})
	http.Handle("/api/broadcast", &ApiServer{ApiName : "broadcast"})
	http.Handle("/api/get-channel", &ApiServer{ApiName : "get-channel"})
	http.Handle("/api/close", &ApiServer{ApiName : "close"})//close a specific ChannelService
	http.Handle("/api/app-status", &ApiServer{ApiName : "status"})//online num and live connection num

	http.HandleFunc("/demo", StaticServer)

	fmt.Println("listen on port 8002")
	//TODO read application info from db or file
	//TODO offer a init commad to reload application info file
	applications = make(ChannelServer)
	applications_config = make(ChannelServerConfig)

	if err = initServer(); err != nil {
		panic(err.Error())
	}
	fmt.Println(applications)

	if err = http.ListenAndServe(":8002", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
