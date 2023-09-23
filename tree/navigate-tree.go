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

func (n *Node) isPrefectBinaryTree(lvl int) (int, bool) {
	if n == nil {
		return lvl, true
	}
	if n.lChild == nil && n.rChild == nil{
		return lvl, true
	}
	if n.lChild != nil && n.rChild != nil{
		lvl += 1
		lvlL, lChild := n.lChild.isPrefectBinaryTree(lvl)
		lvlR, rChild := n.rChild.isPrefectBinaryTree(lvl)
		// check if two children are same (empty || full) at the same level of the tree
		if (lChild && rChild && lvlL == lvlR){
			return lvl, true
		}
	}
	return lvl, false
}

func Printy(){
	node := Node{2, &Node{root:3}, &Node{root:4}}
	// TEST for perfect tree --> returns true
	// node.lChild = &Node{5, &Node{root:6}, &Node{root:7}}
	// node.rChild = &Node{5, &Node{root:6}, &Node{root:7}}

	// TEST for full tree
	// node.lChild.lChild = &Node{root:5,rChild: &Node{root:6}}
	// node.rChild.rChild = &Node{7,&Node{root:8}, &Node{root:9}}
	fmt.Println(node.isPrefectBinaryTree(0))
}