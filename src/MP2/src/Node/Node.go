package node

import (
	"fmt"
	msg "../Helper"
)

var curNode *Node = CreateNewNode()
var UpQryChan chan UpdateQuery = make(chan UpdateQuery)
var MemListChan chan []string = make(chan []string)
var LocalAddress string
var LocalID string

type Node struct {
	MemList []string
	InGroup bool
	Sender
	Listener
	Updater
	Introducer
}

func CreateNewNode() *Node {
	var newNode *Node = new(Node)
	newNode.MemList = []string{}
	newNode.InGroup = true
	// newSender := NewSender()
	// newListener := NewListener()
	// newIntroducer := NewIntroducer()
	// newUpdater := NewUpdater()
	// newNode := Node{
	// 	MemList:  newMemList,
	// 	Sender:   newSender,
	// 	Listener: newListener,
	// 	Updater:  newUpdater,
	// 	InGroup:  false,
	// }
	return newNode
}

//Called from main.go when the command is "JOIN\n"
//Create new node and run Listener,Sender and Updater
//in seperate goroutines
func RunNode(isIntroducer bool) {
	LocalID = msg.CreateID()
	fmt.Println("Node: Local ID is: " + LocalID)
	LocalAddress = msg.GetHostName()
	fmt.Println("Node: Local Address is: " + LocalAddress)

	go curNode.Updater.UpdateMembershipList()

	// curNode = CreateNewNode()
	if !isIntroducer {
		//Firstly, send Join Msg to Introducer
		curNode.Sender.NodeSend(msg.JoinMsg)
		//false for non-intro, true for intro
		go curNode.Listener.NodeListen()
	} else {
		go curNode.Introducer.NodeHandleJoin()
		go curNode.Listener.NodeListen()
	}

	go curNode.Sender.NodeSend(msg.HeartbeatMsg)
}

//Called from main.go when the command is "LEAVE\n"
//Delete the Node
func StopNode() {
	curNode.Sender.NodeSend(msg.LeaveMsg)
}
