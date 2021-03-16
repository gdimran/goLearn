package main

import (
	"fmt"
	"strings"
)
func arrayToString(a []uint, delim string) string {
    return strings.Trim(strings.Replace(fmt.Sprint(a), ",", "delim", -1), "[]")
}

func CreatePhoneNumber(numbers [10]uint) string {
	h := numbers[0:3]
	head:= arrayToString(h, ",")
	first:= strings.ReplaceAll(head, " ", "")
	//fmt.Println(head)
	m := numbers[3:6]
	mid:= arrayToString(m, ",")
	middle:= strings.ReplaceAll(mid, " ", "")
	//fmt.Println(mid)
	l := numbers[6:]
	tail:= arrayToString(l, ",")
	last:=strings.ReplaceAll(tail, " ", "")
	//fmt.Println(tail)
	 
	num_formate:= "(" + first + ") "  + middle + "-" +last
	//fmt.Println(num_formate)
	//PhoneNumber := strings.ReplaceAll(num_formate, " ", "")
	fmt.Println(num_formate)
	return num_formate
}
func main() {

	numbers := [10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	CreatePhoneNumber(numbers)
}