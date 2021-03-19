package main

import "fmt"

func Solve(arr []int) int {
	//var sieve []int
	var primes []int
	total := 0
	for i := 2; i < len(arr); i++ {
		count := 0
		for j := 2; j < arr[i]; j++ {
			if arr[i]%j == 0 {
				count = 1
				break
			}
		}
		if count == 0 {
			primes[total] = arr[i]
			total++
		}
	}
	fmt.Println("prime numbers: ")
	for i := 0; i < total; i++ {
		fmt.Println(primes[i])
	}
	return 0
	//sum:=0

}

func main() {
	Solve([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15})
}
