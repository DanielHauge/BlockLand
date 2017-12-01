package main

import (
	"encoding/json"
	"log"
)


type Message struct {
	Kind string
	Username string
	IP string
	MSG string
	Usernames []string
	IPs []string
}

func createMessage(Kind string, Username string, IP string, MSG string, Usernames []string, IPs []string) (msg *Message) {
	msg = new(Message)
	msg.Kind = Kind
	msg.Username = Username
	msg.IP = IP
	msg.MSG = MSG
	msg.Usernames = Usernames
	msg.IPs = IPs
	return
}

//sends message to all peers
func (msg *Message) send_all(){
	if testing {log.Println("send_all")}
	if testing {log.Println(Connections)}
	for _,peerConnection := range Connections{
		enc := json.NewEncoder(peerConnection)
		enc.Encode(msg)
	}
}

//sends message to a peer
func (msg *Message) sendPrivate(receiver string){
	if testing {log.Println("sendPrivate")}
	if alreadyAUser(receiver){
		peerConnection:=Connections[receiver]
		enc:=json.NewEncoder(peerConnection)
		enc.Encode(msg)
	}else{
		log.Println("Was not a real user")
	}
}