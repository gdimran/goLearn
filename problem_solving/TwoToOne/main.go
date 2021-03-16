package main

import (
	"fmt"
	"strings"
)

func TwoToOne(s1 string, s2 string) string {
	var res string

	s := s1 + s2
	for i := 'a'; i <= 'z'; i++ {
		if strings.Index(s, string(i)) != -1 {
			res += string(i)
		}
	}
	fmt.Println(res)
	return res
}

func main() {
	s1 := "imran"
	s2 := "Hossain"
	TwoToOne(s1, s2)
}
