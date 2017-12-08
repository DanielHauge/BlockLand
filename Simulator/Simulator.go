package main

import (
	"time"
	"log"
	"math/rand"
	"strconv"
)

type Simulation struct {
	url string
	username string
	status string
	questatus QueueStatus
}


type QueueStatus struct {
	Queue []string `json:"queue"`
	Queuenumber int `json:"queuenumber"`
}

/*
func CreateSimulationStatus(url string, username string) Simulation {

	result := Simulation{}
	result.url = url
	result.username = username
	result.status = "Initizalied - Unknown status"
	result.questatus = QueueStatus{}

	return result
}
*/

func startSimulation(url string){

	/*
	Get request to node about initial information: Username, status
	 */
	inque := false
	SimulationsStatus := simulateCheckStatusInQueue(url)
	if SimulationsStatus.Queuenumber!=0{
		inque = true
	}

	for {

		SimulationsStatus = simulateCheckStatusInQueue(url)
		time.Sleep(time.Duration(2)*time.Second)
		if (SimulationsStatus.Queuenumber == 0&&!inque){
			/*
			WAS NOT IN QUEUE!
			 */

			if RandBool.Intn(2) != 0{
				/*
				50 % Chance to join queue
				 */
				 inque = true
				simulateJoinQueue(url)
				log.Println("I decided to join queue")

			}else {
				/*
				50 % chance to not join queue
				 */

				log.Println("I decided to not join queue yet - btw I'm: "+url)

			}

		} else if SimulationsStatus.Queuenumber != 0 &&inque{
			/*
			WAS IN QUEUE
			 */

			if RandBool.Intn(2) != 0{

				/*
				50 % Chance
				 */
				if RandBool.Intn(2) != 0{
					/*
					25 % Chance to Unjoin queuue
					 */
					simulateRemoveFromQueue(url)
					inque = false
					log.Println("OOOOOOHHH i decided to remove myself from the queue!")

				}else {

					/*
					25% to do	Nothing
					 */
					log.Println("I was in queue, but decided not to leave!")

				}



			}else {


				/*
				50 % to do nothing
				 */
				log.Println("I was in queue, but decided not to leave!")

			}


		} else {
			log.Println("oohh what happens?")
			log.Println(inque)
			log.Println(SimulationsStatus.Queuenumber)
		}





		rand.Seed(time.Now().UnixNano())
		timeInt := random(50000,130000)
		log.Println("Time to sleep! for "+strconv.Itoa(timeInt)+" Miliseconds! - Btw i'm: "+url)
		time.Sleep(time.Duration(timeInt)*time.Millisecond)

	}

}

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}