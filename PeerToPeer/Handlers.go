package main

import (
	"fmt"
	"net/http"
	"log"
	"strconv"
 	"github.com/shurcooL/github_flavored_markdown"
	"bytes"
	"gopkg.in/russross/blackfriday.v2"
	"encoding/json"
	"strings"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
	log.Println("Peers IP's: ",PeerIPs)
	log.Println("Connections: ",Connections)
	msg := createMessage("YELL", Name, getMyIP(), "HEEEEELOOOOO WELCOME!", make([]string, 0), make([]string, 0))
	msg.send_all()
}

func Join(w http.ResponseWriter, r *http.Request){

	dis := createDiscussion(Name, Name+":join")
	go StartDiscussion(dis)
	fmt.Fprintf(w, "Data: %s\n", "Discussion Started")

	/*TODO retrieve queue information */
	w.WriteHeader(http.StatusOK)
	obj := "Ok"
	msgs, err := json.Marshal(obj); if err != nil{fmt.Println(err) }
	fmt.Fprint(w, string(msgs))
}



func UnJoin(w http.ResponseWriter, r *http.Request){

	dis := createDiscussion(Name, Name+":leave")
	go StartDiscussion(dis)
	fmt.Fprintf(w, "Data: %s\n", "Discussion Started")
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
	/*TODO retrieve queue information */

	w.WriteHeader(http.StatusOK)
	queueNo := helpGetStatusInQueue()
	obj := QueueStatus{queueNo, "something"}
	msgs, err := json.Marshal(obj); if err != nil{fmt.Println(err) }
	fmt.Fprint(w, string(msgs))
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

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}


	markdown := blackfriday.Run([]byte(buffer.String()))
	rendered := github_flavored_markdown.Markdown(markdown)
	w.Write(rendered)
}

func helpGetStatusInQueue() int {
	bcreader := BlockChain.GetReader()

	test := 0
	a := make([]string, 0)
	for {
		block := bcreader.Next()
		test ++
		a = append(a, fmt.Sprintf("%s", block.Data))
		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
	queue := make([]string, test+1)
	for i := 0; i < len(a) ; i++ {
		if(i == 0){
			continue;
		}
		c := strings.Split(a[i], ":")
		if( len(c) > 2){
			name := c[0]
			command := c[1]
			_, inarray  := contains(a, name);
			if command == "join" && !inarray {
				queue = append(queue, name)
			} else if command == "unjoin" && inarray  {
				queue = remove( queue , name )
			}
		}
	}
	status , _ :=contains(queue, Name)
	return status
}
func contains(arr []string, e string) (int, bool) {
	for i, a := range arr {
		if a == e {
			return i, true
		}
	}
	return -1, false
}
func remove(slice []string, s string) []string {
	index, inarray  := contains(slice, s)
	if( inarray ) {
		return append(slice[:index], slice[index+1:]...)
	} else {
		return slice
	}
}
func GetQueue(w http.ResponseWriter, r *http.Request){

	var buffer bytes.Buffer
	queue := ConstructQueue()
	mypos := "\n\n### You are not in queue\n\n"
	buffer.WriteString("# Block-Land Queue\n\n")
	buffer.WriteString("User: "+Name+"\n\n")
	buffer.WriteString("| Queue Number | User |\n")
	buffer.WriteString("|:------------:| ----:|\n")

	for i, u := range queue{
		buffer.WriteString("| "+strconv.Itoa(i+1)+" | "+u)
		if (u == Name){
			mypos = "\n\n### You are at position: "+strconv.Itoa(i+1)+"\n\n"
		}
	}
	buffer.WriteString(mypos)

	markdown := blackfriday.Run([]byte(buffer.String()))
	rendered := github_flavored_markdown.Markdown(markdown)
	w.Write(rendered)
}

func GetSimulationData(w http.ResponseWriter, r *http.Request){

}
