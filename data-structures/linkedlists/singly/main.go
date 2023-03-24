package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head   *Node
	length int
}

func (l *LinkedList) insert(n *Node) {
	replace := l.head
	l.head = n
	l.head.next = replace
	l.length++

}

func (l *LinkedList) delete(search int) {
	curr := l.head
	var prev *Node = nil
	for curr != nil {
		if curr.data == search {
			prev.next = curr.next
		}
		prev = curr
		curr = curr.next
	}

}

func (l *LinkedList) insertAfter(n *Node, target int) {
	curr := l.head
	for curr.next != nil {
		if curr.data == target {
			saveNext := curr.next
			curr.next = n
			n.next = saveNext
			return
		}
		curr = curr.next
	}
}

func (l *LinkedList) pop() int {
	returnThis := l.head
	l.head = l.head.next
	return returnThis.data
}

func (ll *LinkedList) insertList(llInsert *LinkedList, target int) {
	curr := ll.head
	curr2 := llInsert.head
	for curr.next != nil {
		if curr.data == target {

			save := curr.next
			curr.next = curr2
			for curr2.next != nil {
				curr2 = curr2.next
			}
			curr2.next = save
			return
		}
		curr = curr.next
	}
}

func (ll *LinkedList) reverseList() {
	curr := ll.head
	var prev *Node
	fmt.Println(curr.data)
	for curr != nil {
		next := curr.next
		curr.next = prev
		if curr.next != nil {
			fmt.Println("inside loop:", curr.data, "next", curr.next.data)
		}
		prev = curr
		curr = next
	}
	ll.head = prev
}

func recursiveList(n *Node) *Node {
	if n == nil || n.next == nil {
		return n
	}
	p := recursiveList(n.next)
	n.next.next = n
	n.next = nil

	return p
}

//made this sucka, declare a next, prev as current and the next as nil. Then do the loop for continually moving this around
func reverseList(n *Node, ll *LinkedList) {
	next := n.next
	prev := n
	n.next = nil
	for n != nil {
		n = next
		next = n.next
		n.next = prev
		prev = n
		n = next
	}
	ll.head = prev
}

// Notice we don't send the pointer, this is because we don't want to send in a changed "mylist" object that is going to be altered and displayed once.
func Display(l LinkedList) {
	if l.head != nil {
		fmt.Println("first line in display", l.head.data)
	}
	for l.head != nil {
		fmt.Println(l.head.data)
		l.head = l.head.next
	}
}

func main() {

	mylist := LinkedList{}
	node1 := &Node{data: 48}
	node2 := &Node{data: 49}
	node3 := &Node{data: 47}
	node4 := &Node{data: 59}
	node5 := &Node{data: 69}

	mylist.insert(node1)
	mylist.insert(node2)
	mylist.insert(node3)
	Display(mylist)

	// mylist.delete(49)
	// Display(mylist)

	mylist.insertAfter(node4, 49)
	mylist.insert(node5)
	Display(mylist)

	fmt.Println("Reversed")
	reverseList2(mylist.head, nil, &mylist)
	Display(mylist)

	// fmt.Println("pop")
	// mylist.pop()
	// Display(mylist)

	// mylist2 := LinkedList{}
	// node1a := &Node{data: 480}
	// node2a := &Node{data: 490}
	// node3a := &Node{data: 470}
	// mylist2.insert(node1a)
	// mylist2.insert(node2a)
	// mylist2.insert(node3a)
	// mylist.insertList(&mylist2, 59)
	// Display(mylist)

}

// func variabicFunc(num1 int, words ...string) {
// 	var y = 10
// 	var yPointer = &y
// 	yPointer = 10
// 	fmt.Println("derefrenced", *yPointer)
// }
