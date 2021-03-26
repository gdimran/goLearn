package main

import "fmt"

func Solve(arr []int) int {
	N := len(arr)
	var primes []int
	b := make([]bool, N)
	for i := 2; i < N; i++ {
		if b[i] == true {
			continue
		}
		primes = append(primes, i)
		for k := i * i; k < N; k += i {
			b[k] = true
		}
	}
	sum := 0
	for _, p := range primes {
		//fmt.Println(p)
		sum += p
	}
	fmt.Println(sum)
	return sum

}

func main() {
	Solve([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15})
}
