package binaryTree

import (
	"fmt"
)

type Node struct{
	root int
	lChild *Node
	rChild *Node
}

func (n *Node) inOrder () {
	if n == nil {
		return
	}
	n.lChild.inOrder()
	fmt.Printf("%d ---> ",n.root)
	n.rChild.inOrder()
}

func (n *Node) preOrder () {
	if n == nil {
		return
	}
	fmt.Printf("%d ---> ",n.root)
	n.lChild.preOrder()
	n.rChild.preOrder()
}

func (n *Node) postOrder () {
	if n == nil {
		return
	}
	n.lChild.postOrder()
	n.rChild.postOrder()
	fmt.Printf("%d ---> ",n.root)
}

func Printy(){
	node := Node{2, &Node{root:3}, &Node{root:4}}
	// TEST for perfect tree --> returns true
	// node.lChild = &Node{5, &Node{root:6}, &Node{root:7}}
	// node.rChild = &Node{5, &Node{root:6}, &Node{root:7}}
	// node.lChild.lChild = &Node{5, &Node{root:6}, &Node{root:7}}
	// node.lChild.rChild= &Node{5, &Node{root:6}, &Node{root:7}}


	// TEST for full tree
	// node.lChild.lChild = &Node{root:5,rChild: &Node{root:6}}
	// node.rChild.rChild = &Node{7,&Node{root:8}, &Node{root:9}}
	node.inOrder()
}