package main

import (
	"log"
	"math/rand"
	"time"
	"strconv"
)

var RandBool *rand.Rand
var RandTime *rand.Rand

func main() {
	log.Println("Waiting 120 seconds before starting!")
	for i:=120; i>0;i--  {
		log.Println(strconv.Itoa(i)+" Seconds untill start")
		time.Sleep(time.Second)
	}
	log.Println("Starting Simulations!")

	source := rand.NewSource(0)
	BoolsSource := rand.New(source)
	RandBool = BoolsSource

	source2 := rand.NewSource(int64(time.Now().Second()))
	SleepSource := rand.New(source2)
	RandTime = SleepSource


	go startSimulation("http://Node1:8080")
	time.Sleep(time.Second*5)
	go startSimulation("http://Node2:8081")
	time.Sleep(time.Second*10)
	go startSimulation("http://Node3:8082")
	time.Sleep(time.Second*10)
	go startSimulation("http://Node4:8083")
	time.Sleep(time.Second*10)
	go startSimulation("http://Node5:8084")
	time.Sleep(time.Second*10)
	go startSimulation("http://Node6:8085")
	time.Sleep(time.Second*10)
	startSimulation("http://Node7:8086")



}
