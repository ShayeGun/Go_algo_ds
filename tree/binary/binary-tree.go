package binaryTree

type Node struct {
	root *int
	lNode *Node
	rNode *Node
}

func (n *Node) insertNode (root int) int {
	if n.root == nil {
		n.root = &root
		n.lNode = &Node{}
		n.rNode = &Node{}
	}
	if root > *n.root {
		n.rNode.insertNode(root)
	}
	if root < *n.root {
		n.lNode.insertNode(root)
	}
	return root
}

func Printy(){
	node := Node{}
	node.insertNode(6)
	node.insertNode(7)
	node.insertNode(0)
	node.insertNode(2)
}