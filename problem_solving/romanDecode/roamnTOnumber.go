package main

import "fmt"

var roman = map[string]int{
	"I":  1,
	"IV": 4,
	"V":  5,
	"IX": 9,
	"X":  10,
	"XL": 40,
	"L":  50,
	"XC": 90,
	"C":  100,
	"CD": 400,
	"D":  500,
	"CM": 900,
	"M":  1000,
}

func romanToInt(s string) int {
	if s == "" {
		return 0
	}
	num, lastint, total := 0, 0, 0
	for i := 0; i < len(s); i++ {
		char := s[len(s)-(i+1) : len(s)-i]
		var b int = int(roman[char])
		//num = roman[char]
		num = b
		if num < lastint {
			total = total - num
		} else {
			total = total + num
		}
		lastint = num
	}

	return total
}

func main() {
	result := romanToInt("XXI")
	fmt.Println(result)
}
