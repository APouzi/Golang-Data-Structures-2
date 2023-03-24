package main

import "fmt"

type Node struct {
	data string
	next *Node
}

type LinkedList struct {
	head   *Node
	length int
}

func (l *LinkedList) insert(n *Node) {
	replace := l.head
	l.head = n
	n.next = replace
}

func (l *LinkedList) insert2(n string) {
	curr := l.head
	new := &Node{data: n}
	new.next = curr
	l.head = new
}

func (l *LinkedList) jointTwo(lInsert *LinkedList, val *string) {
	curr := l.head
	for curr.next != nil {
		curr = curr.next

	}
	fmt.Println(curr.data)
	curr.next = lInsert.head
}

func (l *LinkedList) insertList(lInsert *LinkedList, val string) {
	curr := l.head
	for curr != nil {
		if curr.data == val {
			hold := curr.next
			curr2 := lInsert.head
			curr.next = curr2
			for curr2.next != nil {
				curr2 = curr2.next
			}
			curr2.next = hold
			val = "changed"
		}
		curr = curr.next
	}
}

func (l *LinkedList) insertAfter(n *Node, target string) {
	curr := l.head
	for curr.next != nil {
		if curr.next.data == target {
			save := curr.next
			n.next = curr.next.next
			curr.next = save
			save.next = n
			return
		}
		curr = curr.next
	}
}

func (l *LinkedList) delete(search string) {
	curr := l.head
	var prev *Node

	if curr.data == search {
		fmt.Println("found in first")
		replace := curr.next
		l.head = replace
		return
	}
	for curr != nil {
		if curr.data == search {
			prev.next = curr.next
			curr.next = nil
		}
		prev = curr
		curr = curr.next
	}
}

func display(l LinkedList) {
	for l.head != nil {
		fmt.Println(l.head.data)
		l.head = l.head.next
	}
}

func recursiveList(n *Node, prev *Node, ll *LinkedList) {
	if n == nil || n.next == nil {
		ll.head = prev
		return
	}
	next := n.next
	n.next.next = n
	prev = n
	recursiveList(next, prev, ll)
	n.next = prev
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

//testing these commit
func main() {
	myList := LinkedList{}
	node1 := &Node{data: "first"}
	node2 := &Node{data: "second"}
	node3 := &Node{data: "third"}

	// node4 := &Node{data: "first"}
	// node5 := &Node{data: "second"}
	// node6 := &Node{data: "third"}
	node4 := &Node{data: "InsertAfter"}
	myList.insert(node1)
	myList.insert(node2)
	myList.insert(node3)
	myList.insertAfter(node4, "second")
	display(myList)
	fmt.Println("-----------------")

	// reverseRecursive(myList.head, nil, &myList)
	recursiveList(myList.head, nil, &myList)
	display(myList)
	// myList2 := LinkedList{}
	// myList2.insert(node4)
	// myList2.insert(node5)
	// myList2.insert(node6)
	// val := "second"
	// myList.insertList(&myList2, val)
	// display(myList)

}
