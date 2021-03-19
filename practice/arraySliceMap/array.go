package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var x [10]float64
	x[0] = 98
	x[1] = 89
	x[2] = 92
	x[3] = 70
	x[4] = 88
	x[5] = 99
	x[6] = 91
	x[7] = 93

	var total float64 = 0

	for i := 0; i < len(x); i++ {
		total += x[i]
	}
	fmt.Println(total / float64(len(x)))

	//slice

	var numbers []int
	numbers = make([]int, 5)
	matrix := make([][]int, 3*3)
	for i := 0; i < 5; i++ {
		numbers[i] = rand.Intn(100)
	}

	//insert matrix data
	for i := 0; i < 3; i++ {
		matrix[i] = make([]int, 3)
		for j := 0; j < 3; j++ {
			matrix[i][j] = rand.Intn(100)
			fmt.Println(matrix[i][j])
		}
	}

	num := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	for i, min := range num {
		if min < i {
			fmt.Println("minmum:", min)
		}
	}
	fmt.Println("list:", num)

}
