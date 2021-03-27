package main

import "fmt"

var store map[float64]float64

func factorial(n float64) float64 {
	if n <= 1 {
		return 1
	}
	if store[n] == 0 {
		store[n] = n * factorial(n-1)
	}
	return store[n]
}

func Going(n int) float64 {
	res, inter := 1.0, 1.0

	for i := n; i >= 2; i-- {
		inter = inter * (1.0 / float64(i))
		res += inter
	}
	return float64(int64(res*1e6)) / 1e6
}
func main() {
	fmt.Println(113)
}
