package main

import (
	"fmt"

	"github.com/Mahmud139/vatcalculator"
)

func main() {
	test := vatcalculator.ExclusiveTax(200, 7.5)
	fmt.Println(test)
}
