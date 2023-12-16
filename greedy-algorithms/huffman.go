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
	char  rune
	left  *HuffmanNode
	right *HuffmanNode
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

func createHuffmanTree(table []HuffmanTableElements) HuffmanNode {
	tree := HuffmanNode{}

	for _, v := range table {
		binaryCode := v.code
		for i := range binaryCode {
			if binaryCode[i] == '0' {
			} else {
			}
		}
	}

	return tree
}

func Printy() {
	fmt.Println(createTable("salam"))
}
