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
	log.Println("Waiting 30 seconds before starting!")
	for i:=1; i>0;i--  {
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


	go startSimulation("http://188.226.166.33:8080")
	time.Sleep(time.Second)
	go startSimulation("http://188.226.166.33:8081")
	time.Sleep(time.Second)
	go startSimulation("http://188.226.166.33:8082")
	time.Sleep(time.Second)
	go startSimulation("http://188.226.166.33:8083")
	time.Sleep(time.Second)
	go startSimulation("http://188.226.166.33:8084")
	time.Sleep(time.Second)
	go startSimulation("http://188.226.166.33:8085")
	time.Sleep(time.Second)
	startSimulation("http://188.226.166.33:8086")



}
