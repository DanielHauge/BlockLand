package main

import (
	"net"
	"sync"
	"log"
	"os"
	"encoding/json"
)

const PORT = ":2017"

var (

	PeerIPs = make(map[string]string)
	Connections = make(map[string]net.Conn)
	Name string
	testing bool = true
	mutex = new(sync.Mutex)

)


func server(){
	if testing {log.Println("server")}
	tcpAddr, err := net.ResolveTCPAddr("tcp4", PORT)
	if err != nil{ panic(err); log.Println(err.Error())}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil{ panic(err); log.Println(err.Error())}

	for {
		conn, err := listener.Accept()
		if err != nil{
			continue
		}
		if testing {log.Println("Accepted Connection")}
		go receive(conn)
	}
}


func receive(conn net.Conn){

	if testing {log.Println("receive")}
	defer conn.Close()
	dec := json.NewDecoder(conn)
	msg:= new(Message)
	for {
		if err := dec.Decode(msg); err != nil { return }

		switch msg.Kind{
		case "CONNECT":
			if testing {log.Println("KIND = CONNECT")}
			if !handleConnect(*msg, conn){return}
		case "YELL":
			if testing {log.Println("KIND = YELL")}
			log.Println("i just yelled something: " + msg.MSG)

		case "DISCONNECT":
			if testing {log.Println("KIND = DISCONNECT")}
			disconnect(*msg)
			return
		case "HEARTBEAT":
			log.Println("HEARTBEAT")
		case "LIST":
			if testing {log.Println("KIND = LIST")}
			connectToPeers(*msg)
			return
		case "ADD":
			if testing {log.Println("KIND = ADD")}
			addPeer(*msg)

		}
	}
}


//introduces peer to the chat
func introduceMyself(IP string){
	if testing {log.Println("introduceMyself")}
	conn:=createConnection(IP)
	enc:= json.NewEncoder(conn)
	introMessage:= createMessage("CONNECT", Name , getMyIP(), "", make([]string, 0), make([]string, 0))
	enc.Encode(introMessage)
	go receive(conn)
}


//adds a peer to everyone list
func addPeer(msg Message){
	mutex.Lock()
	PeerIPs[msg.Username]=msg.IP
	conn:=createConnection(msg.IP)
	Connections[msg.Username]=conn
	mutex.Unlock()
	log.Println(msg.Username+" just joined")
}


func connectToPeers(msg Message) {
	for index, ip := range msg.IPs {
		conn:=createConnection(ip)
		mutex.Lock()
		PeerIPs[msg.Usernames[index]]=ip
		Connections[msg.Usernames[index]]=conn
		mutex.Unlock()
	}
	addMessage := createMessage("ADD", Name, getMyIP(), "", make([]string, 0), make([]string, 0))
	addMessage.send_all()
}


//creates a new connection, given the IP address, and returns it
func createConnection(IP string) (conn net.Conn){
	service:= IP+PORT
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	if err != nil {log.Println(err.Error())}
	conn, err = net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {log.Println(err.Error())}
	return
}



func disconnect(msg Message){
	mutex.Lock()
	delete(PeerIPs, msg.Username)
	delete(Connections, msg.Username)
	mutex.Unlock()
	log.Println(msg.Username + " left")
}



func handleConnect(msg Message, conn net.Conn) bool{
	if testing {log.Println("handleConnect")}
	Users, IPS := getFromMap(PeerIPs)
	Users = append(Users, Name)
	IPS = append(IPS, getMyIP())
	response := createMessage("LIST", "", "", "", Users, IPS)
	if alreadyAUser(msg.Username){
		response.MSG="Username already taken, choose another one that is not in the list"
		response.send_all()
		return false
	}
	mutex.Lock()
	PeerIPs[msg.Username]=msg.IP
	Connections[msg.Username]=conn
	mutex.Unlock()
	log.Println(Connections)
	response.sendPrivate(msg.Username)
	return true

}


//checks to see if a userName is already in the list
func alreadyAUser(user string) bool {
	for userName,_:= range PeerIPs {
		if userName == user {return true}
	}
	return false
}




func getMyIP() string {
	name, err := os.Hostname()
	if err != nil {log.Println(err.Error())}
	addr, err := net.ResolveIPAddr("ip", name)
	if err != nil {log.Println(err.Error())}
	return addr.String()

}


func getFromMap(mappa map[string]string) ([]string, []string){
	var keys []string
	var values []string
	for key,value := range mappa{
		keys = append(keys,key)
		values = append(values,value)
	}
	return keys,values
}