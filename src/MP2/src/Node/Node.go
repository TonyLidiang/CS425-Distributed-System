package node

import (
	"MP2/src/node"
	"log"
	"net"
)

type Node struct {
	MemList    []string
	Sender     node.Sender
	Listener   node.Listener
	Introducer node.Introducer
	Updater    node.Updater
}

func CreateIntroducerNode() Node{
	newSender     := NewSender()
	newListener   := NewListener()
	newIntroducer := NewIntroducer()
	newUpdater    := NewUpdater()
	
}

func CreateNonIntroducerNode() Node{

}

func RunNonIntroducerNode(node) {

}

func RunIntroducerNode(node) {

}

//Called from main.go when the command is "JOIN\n"
//Create new node and run the node until LEAVE or crash
func RunNode(isIntroducer bool) {
	var node Node
	if(!isIntroducer){
		node = CreateNonIntroducerNode()
		RunNonIntroducerNode(node)
	} else {
		node = CreateIntroducerNode()
		RunIntroducerNode(node)
	}
}

//Called from main.go when the command is "LEAVE\n"
//Delete the Node
func StopNode() {
}
