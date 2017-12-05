package main

import (
	"fmt"
	"net/http"
	"log"
	"strconv"
 	"github.com/shurcooL/github_flavored_markdown"
	"bytes"
	"gopkg.in/russross/blackfriday.v2"
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

	var buffer bytes.Buffer
	buffer.WriteString("# BlockLand Status!\n\n")
	buffer.WriteString("## Username: **"+Name+"**\n\n")
	buffer.WriteString("----------------------\n\n")
	buffer.WriteString("----------------------\n\n")
	buffer.WriteString("## Network\n\n")
	buffer.WriteString("| User              | IP                  | TCP Connection          |\n")
	buffer.WriteString("| ----------------- |:-------------------:| -----------------------:|\n")
	for a, b := range PeerIPs{
		constring := fmt.Sprintf("%v", Connections[a])
		buffer.WriteString("| "+a+" | "+b+" | "+constring+" |\n")
	}
	buffer.WriteString("\n## Discussion\n\n")
	buffer.WriteString("\n##### Discussion queue: "+strconv.Itoa(len(DiscussionQueue))+"\n\n")

	if DiscussionInSession {
		buffer.WriteString("### Discussion in session - Topic: "+DiscussionTopic+"\n\n")
		buffer.WriteString("###### Speaker: "+DiscussionSpeaker+"\n\n")

		if Name == DiscussionSpeaker{

			buffer.WriteString("##### Participants\n\n")
			for _, b := range DiscussionParticipants{
				buffer.WriteString(b+", ")
			}
			buffer.WriteString("\n\n##### Current Agreements\n\n")
			buffer.WriteString("| User              | Vote                |\n")
			buffer.WriteString("| ----------------- |:-------------------:|\n")
			for a, b := range DiscussionAgreement{
				buffer.WriteString("| "+a+" | "+strconv.FormatBool(b)+" |\n")
			}
			buffer.WriteString("\n")

		}else {
			buffer.WriteString("Go to ```"+PeerIPs[DiscussionSpeaker]+DiscussionSpeakerPort+"/status``` for more info\n\n")
		}


	} else {
		buffer.WriteString("No discussion is in session\n\n")
	}

	markdown := blackfriday.Run([]byte(buffer.String()))
	w.Write(markdown)

}

func AnswerOffer(w http.ResponseWriter, r *http.Request){

}

func MoveOut(w http.ResponseWriter, r *http.Request){

}

func GetChain(w http.ResponseWriter, r *http.Request){

	var buffer bytes.Buffer

	buffer.WriteString("# Full Blockchain\n\n")
	buffer.WriteString("User: "+Name+"\n\n")
	buffer.WriteString("| Index | Nounce | Data              | Hash                         | Prev. hash                      | Proof of Work |\n")
	buffer.WriteString("|:-----:|:-------------:|:--------------------------- |:--------------------------------------------------- |:----------------------------------------------------------- |:-------------:|\n")
	bcreader := BlockChain.GetReader()
	test := 0
	for {
		block := bcreader.Next()
		test ++
		pow := GenerateProofOfWork(block)
		buffer.WriteString("| "+fmt.Sprintf("%x", test)+" | "+strconv.Itoa(block.Nounce)+" | "+fmt.Sprintf("%s", block.Data) + " | "+ fmt.Sprintf("%x", block.Hash)+" | "+fmt.Sprintf("%x", block.PrevBlockHash)+" | "+fmt.Sprintf("%s", strconv.FormatBool(pow.InspectGemCarat()))+" |\n")

		/*
		fmt.Fprintf(w,"Number: %x\n", test)
		fmt.Fprintf(w, "Nounce: %v\n", block.Nounce)
		fmt.Fprintf(w, "Data: %s\n", block.Data)
		fmt.Fprintf(w,"Hash: %x\n", block.Hash)
		fmt.Fprintf(w,"Prev. Hash: %x\n", block.PrevBlockHash)
		fmt.Fprintf(w, "PoW: %s\n", strconv.FormatBool(pow.InspectGemCarat()))

		fmt.Fprintln(w)
		*/

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}


	markdown := blackfriday.Run([]byte(buffer.String()))
	rendered := github_flavored_markdown.Markdown(markdown)
	w.Write(rendered)
}
