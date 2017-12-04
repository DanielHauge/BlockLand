package main

import (
	"fmt"
	"net/http"
	"log"
	"strconv"

)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
	log.Println("Peers IP's: ",PeerIPs)
	log.Println("Connections: ",Connections)
	msg := createMessage("YELL", Name, getMyIP(), "HEEEEELOOOOO WELCOME!", make([]string, 0), make([]string, 0))
	msg.send_all()
}

/*
func Join(w http.ResponseWriter, r *http.Request){



	msg := createMessage("JOIN", Name, getMyIP(), Name, make([]string, 0), make([]string, 0))
	msg.send_all()
	BlockChain.AddBlock("New Block Hello")
	fmt.Fprintf(w, "Data: %s\n", "Discussion Started")
}
*/

func Join(w http.ResponseWriter, r *http.Request){

	dis := createDiscussion(Name, Name+":join")
	go StartDiscussion(dis)
	fmt.Fprintf(w, "Data: %s\n", "Discussion Started")
}

func UnJoin(w http.ResponseWriter, r *http.Request){

	msg := createMessage("UNJOIN", Name, getMyIP(), Name, make([]string, 0), make([]string, 0))
	msg.send_all()
}

func Status(w http.ResponseWriter, r *http.Request){

	fmt.Fprintln(w, "My Username: "+Name+"\n\n")
	fmt.Fprintln(w, "Network IPs")
	for a, b := range PeerIPs{
		fmt.Fprintf(w, "User: %v\n", a)
		fmt.Fprintf(w, "IP: %v\n", b)
		fmt.Fprintln(w)
	}

	fmt.Fprintln(w)
	fmt.Fprintln(w, "Network TCP connections")
	for a, b := range Connections{
		fmt.Fprintf(w, "User: %v\n", a)
		fmt.Fprintf(w, "Connection: %v\n", b)
		fmt.Fprintln(w)

	}

	if DiscussionInSession {
		fmt.Fprintln(w,"\n\nDiscussion is in session")
		fmt.Fprintln(w, "\nSpeaker: "+DiscussionSpeaker)

		if Name == DiscussionSpeaker{

			fmt.Fprintln(w, "\n\nParticipants: \n")
			for _, b := range DiscussionParticipants{
				fmt.Fprintf(w, "%s\n", b)
			}
			fmt.Fprintln(w, "\n\nCurrentAgreements")
			for a, b := range DiscussionAgreement{
				fmt.Fprintf(w, "Seat: %s - Vote: %s\n", a,strconv.FormatBool(b))
			}

		}else {
			fmt.Fprintln(w, "Go to "+PeerIPs[DiscussionSpeaker]+DiscussionSpeakerPort+"/status for more info")
		}


	} else {
		fmt.Fprintln(w, "No discussion is in session")
	}

}

func AnswerOffer(w http.ResponseWriter, r *http.Request){

}

func MoveOut(w http.ResponseWriter, r *http.Request){

}

func GetChain(w http.ResponseWriter, r *http.Request){

	bcreader := BlockChain.GetReader()
	test := 0
	for {
		block := bcreader.Next()
		test ++
		fmt.Fprintf(w,"Number: %x\n", test)
		fmt.Fprintf(w, "Nounce: %v\n", block.Nounce)
		fmt.Fprintf(w, "Data: %s\n", block.Data)
		fmt.Fprintf(w,"Hash: %x\n", block.Hash)
		fmt.Fprintf(w,"Prev. Hash: %x\n", block.PrevBlockHash)
		pow := GenerateProofOfWork(block)
		fmt.Fprintf(w, "PoW: %s\n", strconv.FormatBool(pow.InspectGemCarat()))
		fmt.Fprintln(w)

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}
