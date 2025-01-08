// author: @UchihaSusie
// reverse-linkedlist.go
package main

import "fmt"

// Node definition
type Node struct {
	data int
	next *Node
}

// LinkedList definition
type LinkedList struct {
	head *Node
}

// insert new node at the head of the linked list
func (ll *LinkedList) Push(data int) {
	node := &Node{data: data}
	node.next = ll.head
	ll.head = node
}

// print the linked list
func (ll *LinkedList) PrintList() {
	temp := ll.head
	for temp != nil {
		fmt.Print(temp.data, " ")
		temp = temp.next
	}
}

// reverse the linked list
func (ll *LinkedList) Reverse() {
	var dummy *Node
	cur := ll.head
	for cur != nil {
		nextTemp := cur.next
		cur.next = dummy
		dummy = cur
		cur = nextTemp
	}
	ll.head = dummy
}

func main() {
	ll := &LinkedList{}
	ll.Push(1)
	ll.Push(2)
	ll.Push(3)
	ll.Push(4)
	ll.Push(5)
	fmt.Println("Before Reverse")
	ll.PrintList()
	ll.Reverse() // reverse the linked list
	fmt.Println("\nAfter Reversing")
	ll.PrintList()
}