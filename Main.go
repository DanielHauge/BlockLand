package main

import (
	"log"
	"net/http"
	"os"
)


// Args= 1: host, 2: Username

func main() {

	Name = os.Args[2]
	log.Println("Initializing P2P Server")
	go server()
	if os.Args[1]!="127.0.0.1"{go introduceMyself(os.Args[1]); log.Println("Connecting to peers")} else { log.Println("I'm the first peer on the network, therefor will not introduce myself") }



	log.Println("Initiazling API")

	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
	log.Println("DONE [x]")


}
