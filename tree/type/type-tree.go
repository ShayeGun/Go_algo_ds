package typeTree

import (
	"fmt"
)

type Node struct{
	root int
	lChild *Node
	rChild *Node
}

func (n *Node) isFullTree() bool {
	if n == nil {
		return true
	}
	if n.lChild == nil && n.rChild == nil{
		return true
	}
	if n.lChild != nil && n.rChild != nil{
		return (n.lChild.isFullTree() && n.rChild.isFullTree())
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

func (n *Node) isPrefectTree() bool {
	depth := n.depthFullTree()
	return n.isPrefect(depth, 0)
}

func (n *Node) depthFullTree() int {
	depth := 0
	nc := n
	for nc != nil {
		depth += 1
		nc = nc.lChild
	}

	return depth - 1
}

func (n *Node) isCompleteTree(index int, numNodes int) bool {
	if n == nil {
		return true
	}
	// for complete tree we consider that the tree is complete until it shows that it is not -->
	// --> we can realize this by comparing it to an array indexes cuz we can show complete trees in array form n, n*2+1, n*2+2
	if index > numNodes {
		return false
	}
	return n.lChild.isCompleteTree(index * 2 + 1, numNodes) && n.rChild.isCompleteTree(index * 2 + 2, numNodes)
}

func (n *Node) nodeCount() int {
	if n == nil {
		return 0
	}
	return (1 + n.lChild.nodeCount() + n.rChild.nodeCount()) 
}

func (n *Node) isBalancedTree(height int) (bool, int) {
	if n == nil {
		return true, height
	}
	lStatus,lHeight  := n.lChild.isBalancedTree(height + 1)
	rStatus,rHeight := n.rChild.isBalancedTree(height + 1)
	var df int

	if lStatus && rStatus {
		df = lHeight - rHeight
		if df == -1 || df == 0 || df == 1{
			height = max(lHeight, rHeight)
			return true, height
		}
	}
	
	return false, -1;
	
}


func Printy(){
	node := Node{root:1, lChild: &Node{root:2}, rChild: &Node{root:3}}
	// TEST for perfect tree --> returns true
	// node.lChild = &Node{5, &Node{root:6}, &Node{root:7}}
	// node.rChild = &Node{5, &Node{root:6}, &Node{root:7}}
	// node.lChild.lChild = &Node{5, &Node{root:6}, &Node{root:7}}
	// node.lChild.rChild= &Node{5, &Node{root:6}, &Node{root:7}}


	// TEST for full tree
	// node.lChild.lChild = &Node{4,&Node{root:5}, &Node{root:6}}
	// node.rChild.rChild = &Node{7,&Node{root:8}, &Node{root:9}}
	fmt.Println(node.isBalancedTree(0))
}