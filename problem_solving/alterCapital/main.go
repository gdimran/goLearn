package main

import (
	"fmt"
	"strings"
)

func Capitalize(st string) []string {
	s := strings.Split(st, "")
	textMap := make(map[int]string)

	for i, data := range s {
		textMap[i] = data
	}

	for key, value := range textMap {
		if key%2 == 0 {
			strings.ToUpper(value)
		}
	}

	fmt.Println(s, len(s))
	fmt.Println(s, textMap)

	return s

}

// func Capitalize(st string) []string {
// 	s := strings.Split(st, "")
// 	fmt.Println(s)
// 	for k := range s {
// 		s[k] = strings.TrimSpace(s[k])
// 	}

// 	for i := 0; i < len(s); i++ {
// 		if i%2 == 0 {
// 			s[i] = strings.ToUpper(s[i])
// 		} else {
// 			s[i] = strings.ToLower(s[i])
// 		}
// 	}

// 	fmt.Println(s)
// 	// altText := strings.Join(s, "!")
// 	// fmt.Println(altText)
// 	// var result []string
// 	// result = append(result, st, altText)
// 	// fmt.Println(result)
// 	//return result
// 	return s
//}

func main() {

	Capitalize("abcdef")
}
