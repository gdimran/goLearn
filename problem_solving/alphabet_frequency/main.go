package main

import (
	"fmt"
)


func solve(slice []string) []int {
	// Your code here and happy coding!
	
	alphabets := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	var positionArr []int
	
	for i := 0; i < len(slice); i++ {
		var incriment int
		for j := 0; j < len(slice[i]); j++ {
			for n := 0; n < len(alphabets); n++ {
				if slice[i] == alphabets[j] {
					incriment++
				}
				
			}
		}
		positionArr = append(positionArr, incriment)
	}

fmt.Println(positionArr)
return positionArr
}
func main() {

	textSlice := []string{"abode", "ABc", "xyzD"}

	solve(textSlice)

}