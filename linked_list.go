package main

import (
	"fmt"
	"strconv"
)

// List is a single linked list.
type List struct {
	head *Node
	len  int
}

// Node is an element of a single linked list.
type Node struct {
	next  *Node
	value int
}

// New returns an initialized list.
func New() *List { return new(List).Init() }

// Init initializes or clears list l.
func (l *List) Init() *List {
	l.head = new(Node)
	l.len = 0
	return l
}

// Add adds a new node holding given value number to the end of the linked list
func (l List) Add(value int) {
	lastNode := l.head
	for lastNode.next != nil {
		lastNode = lastNode.next
	}
	next := &Node{nil, value}
	lastNode.next = next
}

func (l List) String() string {
	var str string
	current := l.head
	for current != nil {
		str += strconv.Itoa(current.value)
		current = current.next
	}
	return str
}

// Insert insert a node at given index position
func (l *List) Insert(value, index int) {
	nodeToInsert := new(Node)
	nodeToInsert.value = value
	var prev *Node
	next := l.head
	for ; index != 0; index-- {
		prev = next
		next = next.next
	}
	if prev == nil {
		l.head = nodeToInsert
	} else {
		prev.next = nodeToInsert
	}
	nodeToInsert.next = next
}

// RemoveNode removes the first node in the list holding the given searchValue
func (l *List) RemoveNode(searchValue int) *Node {
	doublePointer := &(l.head)
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
func (l List) Search(searchValue int) *Node {
	current := l.head
	for current.next != nil {
		if current.value == searchValue {
			return current
		}
		current = current.next
	}
	return nil
}

func main() {

	list := New()
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.Add(5)

	fmt.Println(list.Search(3).value)

	fmt.Println(list)
	list.RemoveNode(0)
	fmt.Println(list)
	list.Insert(9, 0)
	fmt.Println(list)
}
