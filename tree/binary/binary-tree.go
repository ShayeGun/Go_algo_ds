package binaryTree

import (
	"fmt"
)

type Node struct {
	root  *int
	lNode *Node
	rNode *Node
}

func (n *Node) insertNode(root int) int {
	if n.root == nil {
		n.root = &root
	}
	if root > *n.root {
		if n.rNode == nil {
			n.rNode = &Node{}
		}
		n.rNode.insertNode(root)
	}
	if root < *n.root {
		if n.lNode == nil {
			n.lNode = &Node{}
		}
		n.lNode.insertNode(root)
	}
	return root
}

func (n *Node) searchNode(value int) (bool, *Node) {
	if n.root == nil {
		return false, nil
	}
	if *n.root == value {
		return true, n
	}

	if n.rNode != nil && value > *n.root {
		return n.rNode.searchNode(value)
	}

	if n.lNode != nil && value < *n.root {
		return n.lNode.searchNode(value)
	}

	return false, nil
}

func (n *Node) deleteNode(value int) {
	truth, node := n.searchNode(value)
	if truth {
		// leaf node
		if node.lNode == nil && node.rNode == nil {
			*node = Node{}
			return
		}
		// node with only 1 child node
		if node.lNode != nil && node.rNode == nil {
			node.swapAndDelete(node.lNode)
			return
		}
		if node.rNode != nil && node.lNode == nil {
			node.swapAndDelete(node.rNode)
			return
		}
		// node with 2 child nodes
		childNode := node.inOrderPredecessor()
		node.swapAndDelete(childNode)
	}
}

func (n *Node) swapAndDelete(child *Node) {
	n.root = child.root
	*child = Node{}
}

func (n *Node) inOrderPredecessor() *Node {
	newRoot := n.lNode
	if newRoot != nil {
		for newRoot.rNode != nil {
			newRoot = newRoot.rNode
		}
	}

	return newRoot
}

func Printy() {
	node := Node{}
	node.insertNode(6)
	node.insertNode(7)
	node.insertNode(2)
	node.insertNode(3)
	node.insertNode(1)
	node.insertNode(8)
	node.deleteNode(6)

	fmt.Println(*node.root)
}
