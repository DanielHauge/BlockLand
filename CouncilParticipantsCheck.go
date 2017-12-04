package main

import "log"

func ShareMyFellowPeeps(speaker string, speakersUsernames []string, speakerport string){

	Users, IPS := getFromMap(PeerIPs)
	DiscussionInSession = true
	DiscussionSpeakerPort = speakerport
	if (CrossCheckList(Users, speakersUsernames) && len(Users) == len(speakersUsernames)){

		HeadCountMessage := createMessage("SESSION-HEAD-COUNT", Name, getMyIP(), "Yes", Users, IPS)
		HeadCountMessage.sendPrivate(speaker)
	} else {
		HeadCountMessage := createMessage("SESSION-HEAD-COUNT", Name, getMyIP(), "No", Users, IPS)
		HeadCountMessage.sendPrivate(speaker)
	}

}

func CrossCheckList(list1 []string, list2 []string) bool{
	result := true

	for _, b := range list1{
		if !DoesUserExist(b, list2){
			result = false
		}
	}


	return result
}
func DoesUserExist(name string, strings []string) bool {
	result := false;

	for _, b := range strings{
		if b == name {
			result = true
			log.Println("Oooohh found a user, it exists!")
		}
	}


	return result
}

func HeadCount(who string, answer string, users []string){
	if answer == "Yes" {
		log.Println("HeadCount Status: Yes!")
		Users, _ := getFromMap(PeerIPs)
		if CrossCheckList(users, Users){
			log.Println("Yes CrossCheck was correct")
			DiscussionAgreement[who] = true
			DiscussionParticipants = append(DiscussionParticipants, who)
		} else {
			log.Println("HeadCount Status: NO!")
			DiscussionAgreement[who] = false
		}

	} else {
		log.Println("HeadCount Status: NO!")
		DiscussionAgreement[who] = false
	}
}

func CountCouncilMembersDecision() int{
	result := len(DiscussionAgreement)

	for _, b := range DiscussionAgreement{
		if b{
			result--
		}
	}


	return result
}