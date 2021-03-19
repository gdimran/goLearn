package main

import "fmt"

func MaximumSubarraySum(numbers []int) int {
	max_local := 0
	max := 0
	for _, x := range numbers {
		max_local += x
		if max_local < 0 {
			max_local = 0
		} else {
			if max_local > max {
				max = max_local
			}
		}
	}
	fmt.Println(max)
	return max
}

func main() {
	var numbers = []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	MaximumSubarraySum(numbers)
}
