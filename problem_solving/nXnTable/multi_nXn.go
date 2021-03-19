//Program to display Multiplication Table of any number
package main

import "fmt"

func MultiplicationTable(size int) [][]int {
	var sum [100][100]int
	var row, col int
	fmt.Print("Enter number of rows: ")
	fmt.Scanln(&row)
	fmt.Print("Enter number of cols: ")
	fmt.Scanln(&col)

	fmt.Println()
	fmt.Println("========== Matrix1 =============")
	fmt.Println()
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Printf("Enter the element for Matrix1 %d %d :", i+1, j+1)
			fmt.Scanln(&matrix1[i][j])
		}
	}
}

func main() {
	size := 3
	MultiplicationTable(size)
}
