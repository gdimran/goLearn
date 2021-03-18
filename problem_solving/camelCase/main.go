package main

import (
	"fmt"
	"strings"

)

// var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
// var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")
// var matchAllDasher = regexp.MustCompile("(_)")

// func ToSnakeCase(str string) string {
// 	snake := matchFirstCap.ReplaceAllString(str, "${1}${2}")
// 	snake = matchAllCap.ReplaceAllString(snake, "${1}${2}")
// 	snake = matchAllDasher.ReplaceAllString(snake, "")
// 	return strings.ToUpper(snake)
// }
func ToCamelCase(s string) string {

	

    i := strings.Index(s, "-")
    fmt.Println("Index: ", i)
    if i > -1 {
        //chars := s[:i]
        arefun := s[i+1:]
       // fmt.Println(chars)
        fmt.Println(arefun)
    } else {
        fmt.Println("Index not found")
        fmt.Println(s)
    }

	return s
}
func main() {

	str := "the-Stealth-Warrior"
	ToCamelCase(str)
}

// func toCamleCase(str string) string {
// // 	snake

// }
