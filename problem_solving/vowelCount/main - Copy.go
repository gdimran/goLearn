package main

import (
    "os"
    "bufio"
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

func isConsonant(x string) bool {
    consonants := [21]string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "q", "r", "s", "t", "v", "w", "x", "y", "z"}
    consonantLookupTable := make(map[string]bool)
    for _, v := range consonants {
        consonantLookupTable[v] = true
    }

    return consonantLookupTable[x]

}

func formatNumbers(vowelCount int, consonantCount int) {
    fmt.Print("Consonants: ")

    for i := 0; i < consonantCount; i++ {
        fmt.Print("=")
    }

    fmt.Printf(" %d\n", consonantCount)

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

func countConsonants(input string) int {
    consonantCount := 0

    for _, x:= range input {
        var s string
        s = fmt.Sprintf("%c", x)
        if isConsonant(s) {
            consonantCount += 1
        }
    }

    return consonantCount
}

func takeInput() string {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter text: ")
    text, _ := reader.ReadString('\n')

    return text
}

func main() {
    input := takeInput()

    vowelCount := countVowels(input)
    consonantCount := countConsonants(input)

    formatNumbers(vowelCount, consonantCount)
}