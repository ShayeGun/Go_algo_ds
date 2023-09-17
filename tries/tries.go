package tries

const AlphabetLetters = 26

type node struct {
	children [AlphabetLetters]*node
	isEnd bool
}

type Tries struct {
	start *node
}

func InitTrie() *Tries {
	result := Tries{ &node{} }
	return &result
}

func (t *Tries) insertT(word string) {
	currentNode := t.start

	for _,v:= range word{
		i := v - 'a'
		if currentNode.children[i] == nil{
			currentNode.children[i] = &node{}
		}
		currentNode = currentNode.children[i]
	}

	currentNode.isEnd = true
}

func (t *Tries) searchT(word string) bool{
	currentNode := t.start

	for _,v:= range word{
		i := v - 'a'
		if currentNode.children[i] == nil{
			return false
		}
		currentNode = currentNode.children[i]
	}
	if currentNode.isEnd == true {
		return true
	}
	return false
}
