package main

import (
	"fmt"
	"net/http"
	"log"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
	log.Println("Peers IP's: ",PeerIPs)
	log.Println("Connections: ",Connections)
	msg := createMessage("YELL", Name, getMyIP(), "HEEEEELOOOOO WELCOME!", make([]string, 0), make([]string, 0))
	msg.send_all()
}


func Join(w http.ResponseWriter, r *http.Request){

	msg := createMessage("JOIN", Name, getMyIP(), Name, make([]string, 0), make([]string, 0))
	msg.send_all()
}

func UnJoin(w http.ResponseWriter, r *http.Request){

	msg := createMessage("UNJOIN", Name, getMyIP(), Name, make([]string, 0), make([]string, 0))
	msg.send_all()
}

func Status(w http.ResponseWriter, r *http.Request){

}

func AnswerOffer(w http.ResponseWriter, r *http.Request){

}

func MoveOut(w http.ResponseWriter, r *http.Request){

}
