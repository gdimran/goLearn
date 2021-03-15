package main

import "fmt"

func rotate(s1 string) string {
	return s1[1:] + s1[:1]
}

func reverse(s1 string) string {
	s2 := ""
	for i := len(s1); i > 0; i-- {
		s2 += string(s1[i-1])
	}
	return s2
}
func Revrot(s string, n int) string {
	if len(s) < n || n == 0 {
		return ""
	}
	result := ""
	for i := 0; i < len(s)/n; i++ {
		subs := s[0+n*i : n*(i+1)]
		y := 0
		for _, x := range subs {
			y += int((x - 48) * (x - 48) * (x - 48))
		}
		if y%2 == 0 {
			result += reverse(subs)
		} else {
			result += rotate(subs)
		}
	}
	fmt.Println(result)
	return result
}

func main() {
	Revrot("Hello Imran", 5)
}
