package main

import "fmt"

func SquaresInRect(lng int, wdth int) []int {
	if lng == wdth {
		return nil
	}

	result := []int{}

	for {
		if lng <= 0 || wdth <= 0 {
			break
		}
		if lng < wdth {
			result = append(result, lng)
			wdth -= lng
		} else {
			result = append(result, wdth)
			lng -= wdth
		}
	}
	return result
}

func main() {
	fmt.Println(SquaresInRect(3, 5))
	fmt.Println(SquaresInRect(15, 8))
}
