package main

import "fmt"

type tax float64

var x tax

func main() {
	x = 85.50
	fmt.Printf("%T\n", x)
	fmt.Println(x)

	y := float64(x)
	fmt.Printf("%T\n", y)
	fmt.Println(y)
}
