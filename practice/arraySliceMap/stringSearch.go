package main

import (
	"fmt"
	"strings"
)

func ToCamelCase(s string) string {
	x := s
	var sb strings.Builder
	if strings.Contains(x, "_") {
		split := strings.Split(x, "_")
		var changeCase string
		fmt.Println(len(split))
		var sb strings.Builder
		for i := 0; i < len(split); i++ {
			changeCase = strings.Title(split[i])
			sb.WriteString(changeCase)
		}
		fmt.Println(sb.String())
		return sb.String()
		//return sb
	} else if strings.Contains(x, "-") {
		split := strings.Split(x, "-")

		var changeCase string
		//fmt.Println(len(split))
		var sb strings.Builder
		for i := 0; i < len(split); i++ {
			changeCase = strings.Title(split[i])
			sb.WriteString(changeCase)
		}
		fmt.Println(sb.String())
		//return string(sb
		return sb.String()
	}
	return sb.String()
}

func main() {
	x := "The-Stealth-Warrior-Change"
	ToCamelCase(x)
}
