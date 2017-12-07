package main

import (
	"log"
	"math/rand"
	"os"
)

var RandBool *rand.Rand
var RandTime *rand.Rand

func main() {
	log.Println("Starting Simulations!")

	source := rand.NewSource(0)
	RandBools := rand.New(source)
	RandBool = RandBools


	go startSimulation("188.226.166.33:8080")
	go startSimulation("188.226.166.33:8081")
	go startSimulation("188.226.166.33:8082")
	go startSimulation("188.226.166.33:8083")
	go startSimulation("188.226.166.33:8084")
	go startSimulation("188.226.166.33:8085")
	startSimulation("188.226.166.33:8086")



}
