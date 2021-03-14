package main

import (
	"fmt"
	"sort"
	"strconv"
)

func getRoman(number int) string {
	//create a denary_number:roman_symbol map
	romanMap := map[int]string{
		1: "I", 4: "IV", 5: "V", 9: "IX", 10: "X", 40: "XL", 50: "L", 90: "XC", 100: "C", 400: "CD", 500: "D", 900: "CM", 1000: "M",
	}
	//Create a slice of slices
	rows := len(romanMap)
	matrix := make([][]string, rows)
	//create a slice of the map keys
	var key_slice []int

	//range of a map returns key, value pair
	//value is not need so use blank identifier
	for k, _ := range romanMap {
		key_slice = append(key_slice, k)
	}

	//sort the slice in place, heighest number first (decending)
	sort.Sort(sort.Reverse(sort.IntSlice(key_slice)))

	//create a slice of romanMap content slices
	row := 0
	// range of a slice returns index, value pair
	// index not needed so use blank identifier _
	for _, key := range key_slice {
		//convert int key to string key
		skey := strconv.Itoa(key)
		matrix[row] = []string{skey, romanMap[key]}
		row++
	}

	result := ""
	for _, item := range matrix {
		den, err := strconv.Atoi(item[0])
		if err != nil {
			panic(err)
		}
		sys := item[1]
		for number >= den {
			result += sys
			number -= den
		}
	}
	return result
}

func vlaue(r string) int {
	if r == 'I' {
		return 1
	}
	if r == 'V' {
		return 5
	}

	if r == 'X' {
		return 10
	}

	if r == 'L' {
		return 50
	}

	if r == 'C' {
		return 100
	}

	if r == 'D' {
		return 500
	}

	if r == 'M' {
		return 1000
	}

	return -1
}
func romanTONum(romanN string) int {
	//create a denary_number:roman_symbol map
	// roman_Map := map[string]int{
	// 	"I": 1, "IV": 4, "V": 5, "IX": 9, "X": 10, "XL": 40, "L": 50, "XC": 90, "C": 100, "CD": 400, "D": 500, "CM": 900, "M": 1000,
	// }

	// rows := len(roman_Map)
	// matrix := make([][]string, rows)

	// var str_slice []int

	// for t, _ := range roman_Map {
	// 	str_slice = append(str_slice, t)
	// }
	// res := 0

	res := 0
	for i := 0; i < romanN.len(); i++ {
		s1 := value(romanN[i])
		if i+1 < romanN.len() {
			s2 := value(romanN[i+1])
			if s1 >= s2 {
				res = res + s1
			} else {
				res = res + s2 - s1
				i++
			}
		} else {
			res = res + s1
		}
	}
	return res
}

func main() {
	fmt.Println("Convert a denary number (base 10) to a roman numeral:")
	// for readability keep the number below 5000
	number := 21
	roman := getRoman(number)
	fmt.Printf("denary %4v same as roman %v\n", number, roman)

	romanN := "CXXIII"
	numRoman := romanTONum(romanN)
	fmt.Printf("denary %4v same as roman %v\n", roman, numRoman)

	// a few more conversions ...
	numbers := []int{4, 12, 516, 2123}
	for _, number := range numbers {
		roman := getRoman(number)
		fmt.Printf("denary %4v same as roman %v\n", number, roman)
	}
}
