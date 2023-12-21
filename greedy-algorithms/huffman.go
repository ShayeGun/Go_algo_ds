package greedy

import (
	"fmt"
	"sort"
)

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

func createTable(str string) []HuffmanTableElements {
	table := make(map[rune]int)
	result := make([]HuffmanTableElements, 0, len(str))

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
	// create a base array tree from table
	var copyTree []HuffmanNode
	for i := range table {
		copyTree = append(copyTree, HuffmanNode{frequency: table[i].frequency, char: table[i].char})
	}

	for len(copyTree) != 1 {
		var tempParentTree []HuffmanNode

		if len(copyTree)%2 != 0 {
			copyTree = append(copyTree, HuffmanNode{})
		}

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

// TODO: fix createTable for odd number of elements

func Printy() {
	table := createTable("salamB")
	printHuffmanTree(createHuffmanTree(table), 0, false, "")
	fmt.Println(createTable("salamB"))
	fmt.Println(decodeHuffman("000010", createHuffmanTree(table)))
}
