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

	pn2 := person{
		name:    "Ujjal",
		number:  8801933661557,
		message: []string{"Hi this is Ujjal,", "I love programming"},
	}
	fmt.Println(pn2)

	xp := []person{pn, pn2}

	fmt.Println(xp)

	mp := map[string]person{"Me": pn, "Friend": pn2}
	fmt.Println(mp)
}
