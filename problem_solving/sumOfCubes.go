package main

import "fmt"

func sumOfcubes(n int) int {
	var sum int
	for i := 1; i <= n; i++ {
		sum += (i * i * i)
	}
	fmt.Print("Sum : ", sum)
	return sum
}

func main() {
	var n int
	fmt.Print("Enter a positive integer : ")
	fmt.Scan(&n)
	sumOfcubes(n)

}
