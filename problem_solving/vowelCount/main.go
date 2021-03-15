package main

import (
   // "os"
   // "bufio"
    "fmt"
)

func isVowel(x string) bool {
    vowels := [5]string{"a", "e", "i", "o", "u",}
    vowelLookupTable := make(map[string]bool)
    for _, v := range vowels {
        vowelLookupTable[v] = true
    }

    return vowelLookupTable[x]
}


func formatNumbers(vowelCount int) {
  

    fmt.Print("Vowels:     ")

    for i := 0; i < vowelCount; i++ {
        fmt.Print("=")
    }

    fmt.Printf(" %d\n", vowelCount)
}

func countVowels(input string) int {
    vowelCount := 0

    for _, x:= range input {
        var s string
        s = fmt.Sprintf("%c", x)
        if isVowel(s) {
            vowelCount += 1
        }
    }

    return vowelCount
}



func main() {
    //input := takeInput()

    vowelCount := countVowels("imranHossain")
    //consonantCount := countConsonants(input)

    formatNumbers(vowelCount)
}