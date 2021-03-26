package main

import (
	"fmt"
	"strconv"
)

func DigitalRoot(n int) int {
	s := fmt.Sprint(n)

	r := 0

	for _, v := range s {
		l, _ := strconv.Atoi(string(v))

		r += l
	}
	if len(fmt.Sprint(r)) > 1 {
		return DigitalRoot(r)
	}
	return r
}
func DigitalRootB(n int) int {
	return (n-1)%9 + 1
}
func main() {
	fmt.Println(DigitalRoot(493193))
	fmt.Println(DigitalRootB(493193))
}
