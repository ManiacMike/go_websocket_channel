package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
)

var applications *[]ChannelServiceGroup


func WsServer(ws *websocket.Conn) {
	var err error

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

func main() {

	http.Handle("/", websocket.Handler(WsServer))
	http.Handle("/api/create", &ApiServer{"create"})//create a ChannelService
	http.Handle("/api/push", &ApiServer{"push"})
	http.Handle("/api/broadcast", &ApiServer{"broadcast"})
	http.Handle("/api/check", &ApiServer{"check"})
	http.Handle("/api/close", &ApiServer{"close"})//close a specific ChannelService
	http.Handle("/api/app-status", &ApiServer{"status"})//online num and live connection num

	fmt.Println("listen on port 8002")
	//TODO read application info from db or file
	//TODO offer a init commad to reload application info file

	if err := http.ListenAndServe(":8002", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
