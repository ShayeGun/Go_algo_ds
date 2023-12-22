package greedy

import (
	"fmt"
	"sort"
)

// maybe not the best structs to be used for huffman algo ?o_O?
type HuffmanTableElements struct {
	char      rune
	frequency int
	code      string
}

type HuffmanNode struct {
	char      rune
	frequency int
	left      *HuffmanNode
	right     *HuffmanNode
}

func createSortedElementsSlice(str string) []HuffmanTableElements {
	table := make(map[rune]int)
	result := make([]HuffmanTableElements, 0, len(str))

	// count each letter in string
	for _, i := range str {
		if _, ok := table[i]; ok {
			table[i] += 1
		} else {
			table[i] = 1
		}
	}

	for k, v := range table {
		result = append(result, HuffmanTableElements{char: k, frequency: v})
	}

	// first sorting is based on frequency and then if frequencies were equal sorting is based on char value
	sort.SliceStable(result, func(i, j int) bool { return int(result[i].char) < int(result[j].char) })
	sort.SliceStable(result, func(i, j int) bool { return result[i].frequency < result[j].frequency })

	// create binary code for each letter (slice must be sorted)
	var width int = 0
	for power := 1; power < len(result); power = power * 2 {
		width++
	}
	for i := range result {
		result[i].code = fmt.Sprintf("%0*b", width, i)
	}

	return result
}

func createHuffmanTree(table []HuffmanTableElements) *HuffmanNode {
	// create a copy of the slice (SLICE MUST BE SORTED)
	var copyTree []HuffmanNode
	for i := range table {
		copyTree = append(copyTree, HuffmanNode{frequency: table[i].frequency, char: table[i].char})
	}

	for len(copyTree) != 1 {
		var tempParentTree []HuffmanNode

		/*
		* because we assume the always will be even number of elements in slice (i, i+1)
		* we append slice with an empty node if the number of nodes are odd
		* so wouldn't get an error
		 */
		if len(copyTree)%2 != 0 {
			copyTree = append(copyTree, HuffmanNode{})
		}

		/*
		* this tree creation is a bottom to top approach in each iteration we combine 2
		* nodes into single node and push it into new slice and continue this iterations
		* until there is only one node which is the root node
		 */
		for i := 0; i < len(copyTree); i = i + 2 {
			tempParentTree = append(
				tempParentTree,
				HuffmanNode{
					frequency: copyTree[i].frequency + copyTree[i+1].frequency,
					left:      &copyTree[i],
					right:     &copyTree[i+1]})
		}
		copyTree = tempParentTree
	}

	return &copyTree[0]
}

func printHuffmanTree(node *HuffmanNode, depth int, isRight bool, prefix string) {
	if node == nil {
		return
	}

	if depth > 0 {
		var connector string
		if isRight {
			connector = " └── "
		} else {
			connector = " ├── "
		}
		fmt.Printf("%s%s", prefix, connector)
	}

	fmt.Printf("%c%d\n", node.char, node.frequency)

	// Adjust the spacing for child nodes
	newPrefix := prefix
	if depth > 0 {
		if isRight {
			newPrefix += "     "
		} else {
			newPrefix += " │   "
		}
	}

	printHuffmanTree(node.left, depth+1, false, newPrefix)
	printHuffmanTree(node.right, depth+1, true, newPrefix)
}

func decodeHuffman(str string, start *HuffmanNode) string {
	if len(str) == 0 {
		return "ERROR"
	}

	if start.left == nil && start.right == nil {
		return "ERROR"
	}

	temp := start
	result := ""

	/*
	* because this is greedy algorithm in each iteration we assume that our traverse is
	* correct and if bit is 0 we go to left child if its 1 we traverse to right repeat
	* till we are at leaf node and print the letter then reset pointer to the root node
	 */
	for len(str) > 0 {
		bit := str[0]
		str = str[1:]
		if bit == '0' {
			temp = temp.left
		} else if bit == '1' {
			temp = temp.right
		} else {
			return "ERROR"
		}

		if temp.left == nil && temp.right == nil {
			result += string(temp.char)
			temp = start
		}
	}

	return result
}

func Printy() {
	table := createSortedElementsSlice("salamB")
	printHuffmanTree(createHuffmanTree(table), 0, false, "")
	fmt.Println(createSortedElementsSlice("salamB"))
	fmt.Println(decodeHuffman("000010", createHuffmanTree(table)))
}
