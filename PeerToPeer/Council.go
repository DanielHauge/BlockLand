package main

import (
	"time"
	"log"
	"github.com/boltdb/bolt"
	"os"
	"strconv"
)

type Discussion struct {
	Data string
	owner string
}


func createDiscussion(owner string, data string) (dis Discussion) {

	dis.owner = owner
	dis.Data = data
	return
}

var (

	DiscussionInSession bool
	DiscussionSpeaker string
	DiscussionTopic string
	DiscussionSpeakerPort string
	DiscussionQueue chan Discussion
	DiscussionAgreement map[string]bool
	DiscussionParticipants []string
	DiscussionBlock *ProofOfWork
)

func PutDiscussionInQueue(dis Discussion){
	DiscussionQueue<- dis
}

func StartDiscussion(dis Discussion) {
	if (DiscussionInSession){


		Proposal := createMessage("DISCUSSION-TO-QUEUE", Name, getMyIP(), "Here's my list of participants", make([]string, 0), make([]string, 0))

		Proposal.send_all()
		DiscussionQueue <- dis


	} else {


		/*
		PARTICIPANTS CHECKING -> Speaker will count who is here and who is not here, any disagreement will result in a veto
		 */
		DiscussionSpeaker = dis.owner
		DiscussionTopic = "Participants Voting"
		DiscussionAgreement = map[string]bool{}
		DiscussionParticipants = []string{}
		DiscussionInSession = true
		Users, IPS := getFromMap(PeerIPs)
		Proposal := createMessage("SESSION-PROPOSAL", Name, getMyIP(), os.Args[3], Users, IPS)
		Proposal.send_all()

		VotingTime := make(chan bool)

		ticker := time.NewTicker(10 * time.Second)
		go func(ticker *time.Ticker) {
			for {
				select {
				case <-ticker.C:
					VotingTime<-true

				}
			}
		}(ticker)

		<-VotingTime
		if CountCouncilMembersDecision() == 0 {

			/*
			Council agreed upon participants and will now begin discussing
			PROPOSITION TIME -> Speaker will present a block for addition, everyone have to agree. Any disagreement will result in a veto!
			 */

			log.Println("Discussion succesfully started, everyone agrees on participants")
			log.Println("Will start to mine and work and then send to others for approval")
			DiscussionTopic = "Mining Block"
			DiscussionAgreement = map[string]bool{}


			var lastHash []byte

			err := BlockChain.db.View(func(tx *bolt.Tx) error{
				b:= tx.Bucket([]byte("blocksBucket"))
				lastHash = b.Get([]byte("1"))

				return nil
			})
			if err != nil{
				log.Println("Something went wrong at join to get last hash")
			}

			newBlock := NewBlock(dis.Data, lastHash)
			DiscussionBlock = GenerateProofOfWork(newBlock)

			Proposal := createMessage("JOIN-PROPOSITION", Name, getMyIP(), "Here is my Block", Users, IPS)
			Proposal.Block = newBlock.Serialize()
			Proposal.send_all()
			DiscussionTopic = "Discussing legality of block"

			VotingTime := make(chan bool)

			ticker := time.NewTicker(10 * time.Second)
			go func(ticker *time.Ticker) {
				for {
					select {
					case <-ticker.C:
						VotingTime<-true

					}
				}
			}(ticker)

			<-VotingTime

			if CountCouncilMembersDecision() == 0 {

				/*
				PROPOSITION HAS PASSED -> Everyone agrees -> Fire off to add to the chain!
				 */
				join := createMessage("END", Name, getMyIP(), "Everyone agrees", Users, IPS)
				join.send_all()
				EndDiscussion()

			} else {
				/*
				PROPOSITIION HAS BEEN DENIED! -> Everyone needs to agree
				Logic here to solve disagreement
				 */
				abort := createMessage("ABORT", Name, getMyIP(), "Everyone does not agree", Users, IPS)
				abort.send_all()
				AbourtDiscussion()

				log.Println("UUUUH SOMETHING WENT WRONG!!!!!!, council members didn't agree. Disagreement: "+strconv.Itoa(CountCouncilMembersDecision()))

			}

		} else {

			/*
			Participants VETO! -> Council will not continue untill everyone agrees!
			 */
			log.Println("Everyone didn't agree on participants, starting healthchecks on connections")


			abort := createMessage("ABORT", Name, getMyIP(), "Everyone does not agree", Users, IPS)
			abort.send_all()
			AbourtDiscussion()







		}
	}
}


func EndDiscussion(){
	BlockChain.AddKnownGoodBlock(DiscussionBlock)
	DiscussionBlock = nil
	DiscussionSpeaker = ""
	DiscussionSpeakerPort = ""
	DiscussionAgreement = map[string]bool{}
	DiscussionParticipants = make([]string,0)
	DiscussionInSession = false
	if len(DiscussionQueue) > 0{
		NewDiscussion := <-DiscussionQueue
		if NewDiscussion.owner == Name{
			StartDiscussion(NewDiscussion)
		}
	}

}

func AbourtDiscussion(){
	DiscussionBlock = nil
	DiscussionSpeaker = ""
	DiscussionSpeakerPort = ""
	DiscussionAgreement = map[string]bool{}
	DiscussionParticipants = make([]string,0)
	DiscussionInSession = false
	if len(DiscussionQueue) > 0{
		NewDiscussion := <-DiscussionQueue
		if NewDiscussion.owner == Name{
			StartDiscussion(NewDiscussion)
		}
	}
}








