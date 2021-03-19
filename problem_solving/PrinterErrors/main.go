package main

import (
	"fmt"
)

func PrinterError(s string) string {
	lable := map[rune]bool{}
	lable['a'] = true
	lable['b'] = true
	lable['c'] = true
	lable['d'] = true
	lable['e'] = true
	lable['f'] = true
	lable['g'] = true
	lable['h'] = true
	lable['i'] = true
	lable['j'] = true
	lable['k'] = true
	lable['l'] = true
	lable['m'] = true
	errors := 0

	for _, v := range s {
		if !lable[v] {
			errors++
		}
	}
	return fmt.Sprintf("%d/%d", errors, len(s))
}

func main() {
	//s := "aaaxbbbbyyhwawiwjjjwwm"
	s := "aaabbbbhaijjjm"
	fmt.Println(PrinterError(s))
}
