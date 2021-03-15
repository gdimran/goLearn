package main

import (
	"fmt"
	"regexp"
	"strings"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")
var matchAllDasher = regexp.MustCompile("(_)")

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}${2}")
	snake = matchAllDasher.ReplaceAllString(snake, "")
	return strings.ToUpper(snake)
}

func main() {
	fmt.Println(ToSnakeCase("i_got_intern_at_geeks_for_geeks"))
}

// func toCamleCase(str string) string {
// // 	snake

// }
