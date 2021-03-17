package main

import (
	"fmt"
	"strings"
)

var text = "FUN@WITH@GO@LANG@"

// ANSWER 1:
func findTheSymbol(original string, desiredOne string) []int {
	var foundIndexes []int
	for i := 0; i < len(original); i++ {
		if desiredOne == string(original[i]) {
			foundIndexes = append(foundIndexes, i)
		}
	}
	return foundIndexes
}

// ANSWER 2: using Slice
func usingSlice(original string, desiredOne string) []string {
	return strings.Split(text, "@")
}

func main() {
	// Found Indexes
	ind := findTheSymbol(text, "@")
	// checking length of array/slice
	if len(ind) > 0 {
		var startIndex int = 0
		fmt.Println("The Index/s: ", ind)
		for i := 0; i < len(ind); i++ {
			fmt.Println("part = ", i, string(text[startIndex:ind[i]]))
			startIndex = ind[i] + 1
		}
	} else {
		fmt.Println("Unable to find: ", ind)
	}

	// use the slice
	fmt.Println(strings.Join(usingSlice(text, "@"), ", "))
}
