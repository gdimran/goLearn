package main

import "fmt"

func getMaxNsecondMax(nums []int) (int, int) {

	max := nums[0] // assume first value is the smallest

	for _, value := range nums {
		if value > max {
			max = value // found another smaller value, replace previous value in max
		}
	}
	return max
}

func main() {
	nums := []int{1, 5, 87, 45, 8, 8}
	getMax := getMaxNsecondMax(nums)
	fmt.Println("Tow Oldest ages are: ", getMax)
}
