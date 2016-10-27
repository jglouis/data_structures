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

// Insert insert a node at given index position
func (node *Node) Insert(value, index int) {
	nodeToInsert := new(Node)
	nodeToInsert.value = value
	var prev *Node
	next := node
	for ; index != 0; index-- {
		prev = next
		next = next.next
	}
	if prev == nil {
		node = nodeToInsert
	} else {
		prev.next = nodeToInsert
	}
	nodeToInsert.next = next
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

// Search returns the first Node holding the given searchValue
// Returns nil if none found
func (node Node) Search(searchValue int) *Node {
	current := &node
	for current.next != nil {
		if current.value == searchValue {
			return current
		}
		current = current.next
	}
	return nil
}

func main() {

	list := new(Node)
	list.Add(1).Add(2).Add(3).Add(4).Add(5)

	fmt.Println(list.Search(3).value)

	fmt.Println(list)
	list.RemoveNode(2)
	fmt.Println(list)
	list.Insert(9, 0)
	fmt.Println(list)
}
