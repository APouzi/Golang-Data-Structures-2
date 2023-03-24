package main

import "fmt"

type Node struct {
	data string
	prev *Node
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) insert(str string) {
	if l.head == nil {
		new := Node{data: str}
		l.head = &new
	} else {
		new := Node{data: str}
		replace := l.head
		new.next = replace
		replace.prev = &new
		l.head = &new

	}
}

func (l *LinkedList) insertMid(str string, index string) {
	curr := l.head
	for curr != nil {
		if curr.data == index {
			new := &Node{data: str}
			next := curr.next
			curr.next = new
			new.prev = curr
			new.next = next
			next.prev = new
		}
		curr = curr.next
	}
}

func display(l *LinkedList) {
	curr := l.head
	for curr != nil {
		fmt.Println(curr.data)
		if curr.next != nil {
			curr = curr.next
		} else {
			break
		}
	}
	for curr != nil {
		fmt.Println(curr.data)
		if curr.prev != nil {
			curr = curr.prev
		} else {
			break
		}
	}
}

func main() {
	myList := LinkedList{}
	myList.insert("first")
	myList.insert("second")
	myList.insert("thirds")
	myList.insertMid("insert", "second")
	display(&myList)
}
