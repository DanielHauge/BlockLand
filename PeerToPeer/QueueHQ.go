package main

import (
	"fmt"
	"strings"
	"log"
	"bytes"
	"encoding/json"
)

type QueueStatus struct {
	queue []string `json:"queue"`
	queuenumber int `json:"queuenumber"`
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
	QS = QueueStatus{queue:queue, queuenumber:queuenumber}
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

				for _, b := range result{
					if !IsUserAllreadyInQueue(result, b) {
						result = append(result, stringSet[0])
					}
				}
				} else if stringSet[1] == "leave"{
				result = LeaveQueue(result, a)
			}

		} else {log.Println("Do nothing, it was genesis block")}
	}

	return result
}

func IsUserAllreadyInQueue(array []string, user string) bool{
	result := false
	for _, b := range array{
		if strings.Contains(b, user){
			result = true
		}
	}
	return result
}

func reverseArray(array []string) []string{
	if len(array) == 0{
		return array
	}
	return append(reverseArray(array[1:]), array[0])
}

func LeaveQueue(array []string, user string) []string{
	for i, u := range array{
		if u == user{
			return append(array[:i], array[i+1:]...)
		}
	}
	return array
}