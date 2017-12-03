package main

import (
	"net"
	"encoding/json"
	"log"
)


func receive(conn net.Conn){

	defer conn.Close()
	dec := json.NewDecoder(conn)
	msg := new(Message)



	for {

		if err := dec.Decode(msg); err != nil { return }

		if testing {log.Print("\nReceieved Message: \n Kind: "+msg.Kind+ "\n MSG: "+ msg.MSG+"\n Username: "+ msg.Username+"\n IP: "+ msg.IP+ "\n Number of Peers: ", len(msg.Usernames), "\n Number of Connections: ",len(msg.IPs), "\n\n")}

		switch msg.Kind{

		case "CONNECT":
			if !handleConnect(*msg, conn){return}

		case "YELL":
			log.Println("i just yelled something: " + msg.MSG)

		case "END":
			EndDiscussion()

		case "HEALTH":

		case "JOIN-PROPOSITION":
			EvaluateProposition(Deserialize(msg.Block), msg.Username)

		case "JOIN-PROPOSITION-ANSWER":
			CountVotes(msg.Username, msg.MSG)


		case "SESSION-HEAD-COUNT":
			HeadCount(msg.Username, msg.MSG, msg.Usernames)

		case "SESSION-PROPOSAL":
			ShareMyFellowPeeps(msg.Username, msg.Usernames)

		case "DISCUSSION-TO-QUEUE":
			DiscussionQueue <- createDiscussion(msg.Username, msg.MSG)

		case "DISCONNECT":
			disconnect(*msg)
			return

		case "LIST":
			connectToPeers(*msg)
			return

		case "ADD":
			addPeer(*msg)

		}
	}

}
