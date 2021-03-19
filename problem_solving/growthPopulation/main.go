package main

import "fmt"

func NbYear(p0 int, percent float64, aug int, p int) int {
	var y int
	for y = 0; p0 < p; y++ {
		p0 = int(float64(p0)*(1.0+percent*0.01)) + aug
	}
	return y
}

func main() {
	fmt.Println(NbYear(1500, 5, 100, 5000))
}
