package main

import (
	"fmt"
)

func Solve(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		rns[i], rns[j] = rns[j], rns[i]
	}

	return string(rns)
}
func main() {
	str := "Geeks"

	// returns the reversed string.
	strRev := Solve(str)
	fmt.Println(str)
	fmt.Println(strRev)
}
