package main

import (
	"fmt"
	"math"
)

type BST struct {
	root *Node
}

type Node struct {
	data  int
	right *Node
	left  *Node
}

func (t *BST) insert(n int) {
	if t.root == nil {
		t.root = &Node{data: n}
	} else {
		t.root.insertion(n)
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

func display(n *Node) {
	if n != nil {

		display(n.left)
		fmt.Println(n.data)
		display(n.right)

	} else {
		return
	}
}

//--------------------------^UTILITY CLASS ^ -----------------------

func invertTree(n *Node) {

	temp := n.right
	n.right = n.left
	n.left = temp
	if n.left != nil {
		invertTree(n.left)
	}
	if n.right != nil {
		invertTree(n.right)
	}

}

func invertTree2(n *Node) *Node {
	if n.right == nil && n.left == nil {
		return n
	}

	if n.right != nil {
		n.right = invertTree2(n.left)
	}

	if n.left != nil {
		n.left = invertTree2(n.right)
	}

	return n
}

func maxPathSum(n *Node, sum int) int {
	if n == nil {
		return sum
	}

	left := maxPathSum(n.left, sum+n.data)
	right := maxPathSum(n.right, sum+n.data)
	return int(math.Max(float64(left), float64(right)))
}

func main() {
	tree := BST{}
	tree.insert(5)
	tree.insert(6)
	tree.insert(7)
	tree.insert(3)
	tree.insert(1)
	tree.insert(4)
	display(tree.root)
	// invertTree2(tree.root)
	// fmt.Println("Inverted")
	display(tree.root)

	fmt.Println("max", maxPathSum(tree.root, 0))

}
