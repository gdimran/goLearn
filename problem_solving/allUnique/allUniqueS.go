package main

func HasUniqueChar(str string) bool {
	//var x bool
	for i := 0; i < len(str)-1; i++ {
		for j := i + 1; j < len(str); j++ {
			if str[i] == str[j] {
				return false
			}
		}

	}
	return true
}

func main() {
	str := "abcdE"

	HasUniqueChar(str)
}
