package main

import (
	"fmt"
	"strings"
	"log"
	"bytes"
	"encoding/json"
)

type QueueStatus struct {
	Queue []string `json:"queue"`
	Queuenumber int `json:"queuenumber"`
}

func DeserializeQueue(obj []byte) QueueStatus{

	var qs QueueStatus
	if err := json.Unmarshal(obj, &qs); err != nil {
		log.Println(err.Error())
	}
	return qs

}



func (qs QueueStatus) Serialize() []byte {
	var result bytes.Buffer
	enc := json.NewEncoder(&result)
	err := enc.Encode(qs)
	if err != nil {log.Println("Something went wrong with encoding block to json byte array")}

	return result.Bytes()
}

func CreateQueueStatus(queue []string, queuenumber int) (QS QueueStatus){
	QS = QueueStatus{Queue:queue, Queuenumber:queuenumber}
	return
}


func ConstructQueue() []string{

	var result []string
	var chain []string
	bcreader := BlockChain.GetReader()
	for {
		block := bcreader.Next()
		chain = append(chain, fmt.Sprintf("%s", block.Data))

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}


	for _, a := range reverseArray(chain){
		if strings.ContainsAny(a, ":"){
			stringSet := strings.Split(a, ":")
			if stringSet[1] == "join"{
				log.Println("It was a join!")
					if !IsUserAllreadyInQueue(result, a) {
						log.Println("The user didn't exist therefor adding")
						result = append(result, stringSet[0])
					}else {log.Println("User did exist therefor will not add")}

				} else if stringSet[1] == "leave"{
					log.Println("it was a leave")
				result = LeaveQueue(result, stringSet[0])
			}

		} else {log.Println("Do nothing, it was genesis block")}
	}

	return result
}

func IsUserAllreadyInQueue(array []string, user string) bool{
	result := false
	for _, b := range array{
		if b == user{
			result = true
		}
	}
	return result
}

func reverseArray(array []string) []string{
	temp := []string{}

	for i := len(array)-1; i >= 0; i-- {
		temp = append(temp, array[i])
	}

	return temp
}

func LeaveQueue(array []string, user string) []string{
	log.Println("I want to remove: "+user)
	temp := []string{}
	for i, u := range array{
		log.Println("Found user: "+u)
		if u != user{
			log.Println("Was not the one")
			temp = append(temp, array[i])
		}else {log.Println("Uhh found user -> will not add to temp")}

	}
	return temp
}