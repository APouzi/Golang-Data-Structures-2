package main

import "fmt"

type Tree struct {
	root *Node
}

type Node struct {
	data  int
	right *Node
	left  *Node
}

func (t *Tree) insert(num int) {
	if t.root == nil {
		t.root = &Node{data: num}
	} else {
		t.root.insertion(num)
	}
}

// My version
func (n *Node) insertMy(num int) {
	if num <= n.data {
		if n.left == nil {
			new := &Node{data: num}
			n.left = new
		} else {
			n.left.insertMy(num)
		}
	} else if num >= n.data {
		if n.right == nil {
			new := &Node{data: num}
			n.right = new
		} else {
			n.right.insertMy(num)
		}
	}
}

func (n *Node) insertion(val int) *Node {
	if n == nil {
		return &Node{data: val}
	}

	if val >= n.data {
		n.right = n.right.insertion(val)
	}
	if val < n.data {
		n.left = n.left.insertion(val)
	}
	return n
}

func (n *Node) deletion2(val int) *Node {
	if n == nil {
		return nil
	}
	if val > n.data {
		n.right = n.right.deletion2(val)
	} else if val < n.data {
		n.left = n.left.deletion2(val)
	} else {
		if n.right == nil {
			return n.left
		} else if n.left == nil {
			return n.right
		} else {
			replacedWith := biggestFromLeftSubtree(n.left)
			n.data = replacedWith.data
			n.left.deletion2(replacedWith.data)
		}
	}
	// never hit.
	return n

}

func biggestFromLeftSubtree(n *Node) *Node {
	curr := n
	for curr.right != nil {
		curr = curr.right
	}
	return curr
}

// My old version
func (n *Node) deletion(num int) {
	if num < n.data {
		if n.left.data == num {
			maxNum := findMax(n.left)
			fmt.Println("max from found searched item is:", maxNum)
		} else {
			n.left.deletion(num)
		}
	} else {
		if n.right.data == num {
			maxNum := findMax(n.right)
			fmt.Println("max from found searched item is:", maxNum)
		} else {
			n.right.deletion(num)
		}
	}
}

var count int = 0

func findMax(t *Node) int {
	if t.data > count {
		count = t.data
	}
	if t.left != nil || t.right != nil {
		if t.left != nil {
			findMax(t.left)
		}
		if t.right != nil {
			findMax(t.right)
		}
	} else {
		return count
	}
	return count
}

func PreOrder(n Node) {
	fmt.Println(n.data)
	if n.left != nil {
		PostOrder(n.left)
	}
	if n.right != nil {
		PostOrder(n.right)
	}

}

func InOrder(n *Node) {
	if n.left != nil {
		InOrder(n.left)
	}
	fmt.Println(n.data)
	if n.right != nil {
		InOrder(n.right)
	}
}

func PostOrder(n *Node) {
	if n.left != nil {
		PostOrder(n.left)
	}
	if n.right != nil {
		PostOrder(n.right)
	}
	fmt.Println(n.data)
}

func main() {
	tree := Tree{}
	tree.insert(7)
	tree.insert(2)
	tree.insert(3)
	tree.insert(4)
	tree.insert(5)
	tree.insert(6)
	// tree.insert(7)
	tree.insert(8)
	tree.insert(9)
	tree.insert(100)
	tree.insert(11)
	tree.insert(12)
	tree.insert(111)
	InOrder(tree.root)
	tree.root.deletion2(111)
	fmt.Println("deletion")
	InOrder(tree.root)

	fmt.Println("the max is", findMax(tree.root))
	// PostOrder(tree.root)
	// PreOrder(tree.root)
}
