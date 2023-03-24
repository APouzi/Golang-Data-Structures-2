package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

type Tree struct {
	root *Node
}

func (t *Tree) insert(n int) {
	if t.root == nil {
		t.root = &Node{data: n}
	} else {
		t.root.insertion(n)
	}
}

func (t *Tree) deletion(n int) {
	if t.root == nil {
		return
	} else {
		t.root.deletion(n)
	}
}

func (n *Node) insertion(num int) *Node {
	if n == nil {
		return &Node{data: num}
	}
	if num >= n.data {
		n.right = n.right.insertion(num)
	}
	if num < n.data {
		n.left = n.left.insertion(num)
	}
	// This will never hit because it will always hit "nil" and that will be us setting the new return.
	return n
}

func (n *Node) deletion(num int) *Node {
	// If the given tree is empty, this means we want to return nil, can't do anything
	if n == nil {
		return nil
	}
	// traverse down the tree
	if num > n.data {
		n.right = n.right.deletion(num)
	} else if num < n.data {
		n.left = n.left.deletion(num)
	} else if num == n.data {
		if n.right == nil {
			return n.left
		} else if n.left == nil {
			return n.right
		} else if n.right != nil && n.left != nil {
			smallest := leftSubTreeSmallest(n.left)
			n.data = smallest
			n.left.deletion(num)
		}
	}
	return n
}

func leftSubTreeSmallest(node *Node) int {
	curr := node
	for curr.left != nil {
		curr = curr.left
	}
	return curr.data
}
func (n *Node) insert(num int) {
	if n.data >= num {
		if n.left == nil {
			new := &Node{data: num}
			n.left = new
			fmt.Println("inserted", num)
		} else {
			n.left.insert(num)
		}
	} else {
		if n.right == nil {
			new := &Node{data: num}
			n.right = new
		} else {
			n.right.insert(num)
		}
	}
}

func maxPathSum(n *Node) {

}

func invert(n *Node) {
	temp := n.right
	n.right = n.left
	n.left = temp
	if n.left != nil {
		invert(n.left)
	}
	if n.right != nil {
		invert(n.right)
	}

}

func display(n *Node) {
	if n != nil {
		fmt.Println(n.data)
		display(n.left)
		display(n.right)

	} else {
		return
	}
}

func main() {
	myTree := &Tree{}
	myTree.insert(1)
	myTree.insert(2)
	myTree.insert(3)
	myTree.insert(-1)
	myTree.insert(-2)
	myTree.insert(-3)
	myTree.deletion(-3)
	display(myTree.root)
	invert(myTree.root)
	fmt.Println("Inverted Tree")
	display(myTree.root)
}
