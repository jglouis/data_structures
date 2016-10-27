package main

import (
	"fmt"
	"strconv"
)

// Node is a element of a linked list
type Node struct {
	next  *Node
	value int
}

// Add adds a new node holding given value number to the end of the linked list
func (node *Node) Add(value int) *Node {
	next := &Node{nil, value}
	node.next = next
	return next
}

func (node Node) String() string {
	var str string
	current := &node
	for current != nil {
		str += strconv.Itoa(current.value)
		current = current.next
	}
	return str
}

// RemoveNode removes the first node in the list holding the given searchValue
func (node *Node) RemoveNode(searchValue int) *Node {
	doublePointer := &node
	for *doublePointer != nil && (*doublePointer).value != searchValue {
		doublePointer = &((*doublePointer).next)
	}
	if *doublePointer == nil {
		return nil
	}
	deletedNode := *doublePointer
	*doublePointer = deletedNode.next
	deletedNode.next = nil

	return deletedNode
}

func main() {

	list := new(Node)
	list.Add(1).Add(2).Add(3).Add(4).Add(5)

	fmt.Println(list)
	list.RemoveNode(2)
	fmt.Println(list)
}
