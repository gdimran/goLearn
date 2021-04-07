package main

import "fmt"

type person struct {
	name    string
	number  int
	message []string
}

func main() {
	pn := person{
		name:    "Imran Hossain",
		number:  8801921213329,
		message: []string{"Hi this is Imran", "I love programming"},
	}
	fmt.Println(pn)

}
