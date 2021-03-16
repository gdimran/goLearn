package main

import (
	"fmt"
	"unicode/utf8"
)

func Reverse(s string) string {
	size := len(s)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}

	textRev := string(buf)
	fmt.Println(textRev)
	return textRev
}
func main() {

	input := "Imran Hossain"
	Reverse(input)
}
