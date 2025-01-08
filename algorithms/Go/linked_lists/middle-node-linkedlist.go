// author: @UchihaSusie
// algorithms/Go/linked_lists/middle_node_linkedlist.go

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
    newNode := &Node{data: data}
    newNode.next = ll.head
    ll.head = newNode
}

// return the data of the middle node
func (ll *LinkedList) MiddleElement() *Node {
    if ll.head == nil {
        return nil
    }
    slow := ll.head
    fast := ll.head
    for fast != nil && fast.next != nil {
        slow = slow.next
        fast = fast.next.next
    }
    return slow
}

func main() {
    ob := &LinkedList{}
	ob.Push(7)
	ob.Push(6)
    ob.Push(5)
    ob.Push(4)
    ob.Push(3)
    ob.Push(2)
    ob.Push(1)
    middleNode := ob.MiddleElement()
    if middleNode != nil {
        fmt.Println(middleNode.data) // expected output: 4
    }
}