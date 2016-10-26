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

// Add adds a number to the end of the linked list
func (el *Node) Add(number int) *Node {
	node := &Node{nil, number}
	el.next = node
	return node
}

func (el Node) String() string {
	var str string
	current := &el
	for current != nil {
		str += strconv.Itoa(current.value)
		current = current.next
	}
	return str
}

// RemoveNode removes the first node with the searchvalue
func RemoveNode(head **Node, searchValue int) *Node {
	doublePointer := head
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
	RemoveNode(&list, 2)
	fmt.Println(list)
}
