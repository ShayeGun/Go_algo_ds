package avlTree

import (
	"fmt"
)

type Node struct{
	root int
	lChild *Node
	rChild *Node
	height int
}

func (n *Node) getHeight() int{
	if n == nil {
		return 0
	}

	return n.height
}

func (n *Node) rightRotation() *Node{
	x := n.lChild
	temp := x.rChild
	x.rChild = n
	n.lChild = temp
	n.height = max(n.lChild.getHeight(), n.rChild.getHeight())
	x.height = max(x.lChild.getHeight(), x.rChild.getHeight()) + 1
	*n = *x
	return x
}

func (n *Node) leftRotation() *Node{
	x := n.rChild
	temp := x.lChild
	x.lChild = n
	n.rChild = temp 
	n.height = max(n.lChild.getHeight(), n.rChild.getHeight())
	x.height = max(x.lChild.getHeight(), x.rChild.getHeight()) + 1
	*n = *x
	return x
}

func (n *Node) getBalance() int{
	if n == nil {
		return 0 
	}
	return n.lChild.getHeight() - n.rChild.getHeight()
}

func (n *Node) insertNode(value int) *Node {
	if n.root == 0 {
		*n = Node{root: value, height: 1}
		return n
	}
	if value < n.root {
		if n.lChild == nil {
			n.lChild = &Node{}
		}
		n.lChild.insertNode(value)

	} else if value > n.root {
		if n.rChild == nil {
			n.rChild = &Node{}
		}
		n.rChild.insertNode(value)

	} else {
		return n
	}
	// Update the balance factor of each node
    // And, balance the tree
	n.height = max(n.lChild.getHeight(), n.rChild.getHeight()) + 1

	// bf = balance factor
	bf := n.getBalance()
	if bf > 1 {
		if value < n.lChild.root {
			return  n.rightRotation()
		} else if value > n.lChild.root {
			n.lChild.leftRotation()
			return n.rightRotation()
		}
	}

	if bf < -1 {
		if value > n.rChild.root {
			return  n.leftRotation()
		} else if value < n.rChild.root {
			n.rChild.rightRotation()
			n.leftRotation()
			return n
		}
	}

	return n
}

// TODO: THERE IS A PROBLEM WITH HEIGHT IT'S WORKING BUT MY GUTS SAYS IT HAS SOME PROBLEMS :| 
func Printy(){
	node := &Node{}
	node.insertNode(4)
	node.insertNode(5)
	node.insertNode(3)
	node.insertNode(2)
	node.insertNode(1)

	fmt.Println(node.lChild)
}