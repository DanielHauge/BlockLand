package main

import (
	"time"
	"log"
)

type Simulation struct {
	url string
	username string
	status string
	questatus QueueStatus
}


type QueueStatus struct {
	queuenumber int
	queuelog string

}


func CreateSimulationStatus(url string, username string) Simulation {

	result := Simulation{}
	result.url = url
	result.username = username
	result.status = "Initizalied - Unknown status"
	result.questatus = QueueStatus{}

	return result
}

func startSimulation(url string){

	/*
	Get request to node about initial information: Username, status
	 */

	SimulationsStatus := HandShakeWithNode(url)

	for {

		SimulationsStatus.questatus = simulateCheckStatusInQueue(url)
		time.Sleep(2000)
		if (SimulationsStatus.questatus.queuenumber != 0){
			/*
			WAS NOT IN QUEUE!
			 */

			if RandBool.Intn(2) != 0{
				/*
				50 % Chance to join queue
				 */
				simulateJoinQueue(url)
				log.Println("I decided to join queue")

			}else {
				/*
				50 % chance to not join queue
				 */

				log.Println("I decided to not join queue yet")

			}

		} else {
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

				}else {

					/*
					25% to do	Nothing
					 */


				}



			}else {


				/*
				50 % to do nothing
				 */

			}


		}

		if RandBool.Intn(2) != 0{





		}else {




		}




		timeInt := RandTime.Intn(10000 - 30000)+10000

		time.Sleep(time.Duration(timeInt))

	}

}