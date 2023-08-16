package main

import (
	"fmt"
)



func main() {

	myPython := programmingSkills{}

	myPython.programName = "Python"

	fmt.Println(myPython.programName)


	myMsg := messageToSend{}
	myMsg.sender.name = "ravi kabir"
	myMsg.recipient.name = "pachakat"
	myMsg.sender.number = 1231412
	//myMsg.recipient.number = 1353513

	canI := canSendMessage(myMsg)

	fmt.Println(canI)

}


type programmingSkills struct {
	programName string
	frameworkName string
	numberOfProjects int
	interest enjoyment

}

type enjoyment struct {
	complexity int
	fun int
	useful int
	community int
}

type messageToSend struct{
	message string
	sender user
	recipient user
}

type user struct {
	name string
	number int
}

func canSendMessage(mToSend messageToSend) bool {
	if mToSend.sender.name == "" {
		return false
	}
	if mToSend.sender.number == 0 {
		return false
	}
	if mToSend.recipient.name == "" {
		return false
	}
	if mToSend.recipient.number == 0 {
		return false
	}
	return true



}

