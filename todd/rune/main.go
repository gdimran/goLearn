package main

import "fmt"

func main() {
	x := "ABCDEF"

	fmt.Println(x)
	fmt.Printf("%T\n", x)

	xr := []rune(x)

	fmt.Println(xr)

	for i, v := range x {
		fmt.Print(i, v)
		fmt.Printf("%d- %b - %0x\n", v, v, v)
	}
}
