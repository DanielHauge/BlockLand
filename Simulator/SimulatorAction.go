package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"bytes"
)

/**
Join
This call will initialize a proposition for the nodes owner to join the queue.
Route: /join
Method type: GET
Example of return:
Proposition has been made, and you will join the queue
or
Proposition has been denied, you are allready in the queue**/
func simulateJoinQueue(url string){
	// TODO Make http request
	// Make a get request
	req, err := http.NewRequest("GET", url+"/join", nil)
	// Process response
	if err != nil {
		//panic(err) // More idiomatic way would be to print the error and die unless it's a serious error
		fmt.Println(err);
	}
	defer req.Body.Close()
}
/**
Status
This call will return the status: Queue status (In queue: with number, out of queue), and more
Route: /status
Method type: GET
Example of return:
Owner: <ownername>, Queue Status: In queue at position (5)**/

/**
Unjoin
This call will initialize the owner of the node to leave the queue
Route: /join
Method type: DELETE
Example of return:
Proposition has been made, and you will leave the queue
or
Proposition has been denied, you are not in the queue **/
func simulateRemoveFromQueue(url string){
	// TODO Make http request
	req, err := http.NewRequest("DELETE", url+"/join", nil)
	// Process response
	if err != nil {
		//panic(err) // More idiomatic way would be to print the error and die unless it's a serious error
		fmt.Println(err);
	}
	defer req.Body.Close()

}
/**
Answer
This call is to Confirm answer, the user will have to answer yes or no.
Route: /answer
Method type: POST
Accepts Json & text/plain Syntax:
{"answer": <bool>}
Example:
{"answer": true}
Return of Example:
Succesfully transfered contract to you.
or
You denied the contract, you will hereby be forced to leave the queue. **/
func simulateAnswerToOffer(answer bool) bool{
	// TODO Make http request
	return true
}
/**
Move out
This call will initialize the owner of the node to move out and begin a contract transfer
Route: /moveout
Method type: GET
Example of return:
You have moved out and released the contract for other peers
or
You do not currently have a contract, therfore the proposition has been denied
Get Chain
This call will iterate the whole chain and post this
Route: /chain
Method type: GET**/
func simulateMoveOut(url srting) string{
	// TODO Make http request
/*	// Make a get request
	rs, err := http.Get(url + "/moveout")
	// Process response
	if err != nil {
		//panic(err) // More idiomatic way would be to print the error and die unless it's a serious error
		fmt.Println(err);
	}
	defer rs.Body.Close()

	bodyBytes, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		panic(err)
	}

	var bodyString = string(bodyBytes)
	fmt.Println(string(bodyByte))
	return bodyString*/
	return ""
}

func simulateCheckStatusInQueue(url string) QueueStatus{
	// TODO Make http request
	// TODO Make http request
	// Make a get request
	rs, err := http.Get(url + "/status")
	// Process response
	if err != nil {
		//panic(err) // More idiomatic way would be to print the error and die unless it's a serious error
		fmt.Println(err);
	}
	defer rs.Body.Close()

	bodyBytes, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		panic(err)
	}


	var tempSim QueueStatus
	enc := json.NewDecoder(bytes.NewReader(bodyBytes))
	enc.Decode(&tempSim)
	return tempSim

}

func HandShakeWithNode(url string) Simulation{
	rs, err := http.Get(url + "/status")
	// Process response
	if err != nil {
		//panic(err) // More idiomatic way would be to print the error and die unless it's a serious error
		fmt.Println(err);
	}
	defer rs.Body.Close()

	bodyBytes, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		panic(err)
	}

	var sim Simulation
	enc := json.NewDecoder(bytes.NewReader(bodyBytes))
	enc.Decode(&sim)
	return sim
}
