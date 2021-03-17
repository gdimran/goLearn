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

	if strings.Contains(s, "_") {
		i := strings.Index(s, "_")
		if i > -1 {
			//chars := s[i]
			arefun := s[i+1]
			changeCase := strings.ToUpper(arefun)
			fmt.Println(changeCase)
		}
	} else if strings.Contains(s, "-") {
		i := strings.Index(s, "-")
		if i > -1 {
			//chars := x[i]
			arefun := s[i+1]
			changeCase := strings.ToUpper(arefun)
			fmt.Println(changeCase)
		}
	}
	return changeCase
}
func main() {

	str := "the-Stealth-Warrior"
	ToCamelCase(str)
}

// func toCamleCase(str string) string {
// // 	snake

// }
