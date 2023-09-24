package searchTree

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

func (n *Node) isFullBinaryTree() bool {
	if n == nil {
		return true
	}
	if n.lChild == nil && n.rChild == nil{
		return true
	}
	if n.lChild != nil && n.rChild != nil{
		return (n.lChild.isFullBinaryTree() && n.rChild.isFullBinaryTree())
	}
	return false
}

func (n *Node) isPrefect(depth int, lvl int)  bool {
	if n == nil {
		return  true
	}
	if n.lChild == nil && n.rChild == nil{
		return  depth == lvl
	}
	if n.lChild != nil && n.rChild != nil{
		return n.lChild.isPrefect(depth, lvl + 1) && n.rChild.isPrefect(depth, lvl + 1)
	}
	return  false
}

func (n *Node) isPrefectBinaryTree() bool {
	depth := n.depthFullBinaryTree()
	return n.isPrefect(depth, 0)
}

func (n *Node) depthFullBinaryTree() int {
	depth := 0
	nc := n
	for nc != nil {
		depth += 1
		nc = nc.lChild
	}

	return depth - 1
}

func (n *Node) isCompleteBinaryTree(index int, numNodes int) bool {
	if n == nil {
		return true
	}
	// for complete binary tree we consider that the tree is complete until it shows that it is not -->
	// --> we can realize this by comparing it to an array indexes cuz we can show complete binary trees in array form n, n*2+1, n*2+2
	if index > numNodes {
		return false
	}
	return n.lChild.isCompleteBinaryTree(index * 2 + 1, numNodes) && n.rChild.isCompleteBinaryTree(index * 2 + 2, numNodes)
}

func (n *Node) nodeCount() int {
	if n == nil {
		return 0
	}
	return (1 + n.lChild.nodeCount() + n.rChild.nodeCount()) 
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
	fmt.Println(node.isCompleteBinaryTree(0,node.nodeCount()))
}