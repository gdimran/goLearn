package main

import (
	"fmt"
)

func DigPow(n, p int) int {
	nums := string(n)
	total := 0
	for i := 0; i < len(nums); i++ {

		fmt.Println(total)
	}
}
func main() {
	DigPow(695, 2)
}
