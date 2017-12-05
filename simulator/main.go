package simulator

import (
	"fmt"
	"time"
	//"strconv"
	"net/http"
	"io/ioutil"
)
var url string = ""
var bodyByte byte
/**
Join
This call will initialize a proposition for the nodes owner to join the queue.
Route: /join
Method type: GET
Example of return:
Proposition has been made, and you will join the queue
or
Proposition has been denied, you are allready in the queue**/
func simulateJoinQueue() string{
	// TODO Make http request
	// Make a get request
	rs, err := http.Get(url + "/join")
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
	return bodyString
}
/**
Status
This call will return the status: Queue status (In queue: with number, out of queue), and more
Route: /status
Method type: GET
Example of return:
Owner: <ownername>, Queue Status: In queue at position (5)**/
func simulateCheckStatusInQueue() string{
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

	var bodyString = string(bodyBytes)
	fmt.Println(string(bodyByte))
	return bodyString
}
/**
Unjoin
This call will initialize the owner of the node to leave the queue
Route: /join
Method type: DELETE
Example of return:
Proposition has been made, and you will leave the queue
or
Proposition has been denied, you are not in the queue **/
func simulateRemoveFromQueue() bool{
	// TODO Make http request
	return true
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
func simulateAnswerToOffer() bool{
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
func simulateMoveOut() string{
	// TODO Make http request
	// Make a get request
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
	return bodyString
}
func startSimulation(){



	time.Sleep(time.Millisecond* 1000)

	simulateJoinQueue()

	time.Sleep(time.Millisecond * 1000)

	simulateCheckStatusInQueue()

	time.Sleep(time.Millisecond * 1000)

	//simulateRemoveFromQueue()

	//time.Sleep(time.Millisecond * 1000)

	//simulateAnswerToOffer()

	//time.Sleep(time.Millisecond * 1000)

	//simulateMoveOut()

}
func main(){

	startSimulation()

}