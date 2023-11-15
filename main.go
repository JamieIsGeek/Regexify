package main

import (
	"fmt"
	"regexp"
	"unicode"
)

func generateRegex(word string) string {
	var regexExp string

	for _, char := range word {
		var toAdd string

		toAdd += "["
		if isLetter(char) {
			toAdd += string(char)
			toAdd += string(unicode.ToUpper(char))

			if getNumber(char) != "400" {
				toAdd += getNumber(unicode.ToLower(char))
			}

		} else {
			toAdd += string(char)
			if getLetter(char) != "z" {
				toAdd += getLetter(char)
			}
		}

		toAdd += "]"
		regexExp += toAdd
	}

	return regexExp
}

func isLetter(char rune) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}

func getNumber(char rune) string {
	switch char {
	case 'a':
		return "4"
	case 'e':
		return "3"
	case 'i':
		return "1"
	case 'l':
		return "1"
	case 'o':
		return "0"
	case 't':
		return "7"
	case 'g':
		return "9"
	case 's':
		return "5"
	default:
		return "400"
	}
}

func getLetter(char rune) string {
	switch char {
	case 1:
		return "lL"
	case 3:
		return "eE"
	case 4:
		return "aA"
	case 5:
		return "sS"
	case 7:
		return "tT"
	case 9:
		return "gG"
	case 0:
		return "oO"
	default:
		return "z"
	}
}

func main() {
	word := "test"
	regexPattern := generateRegex(word)

	// Use the generated regex pattern
	matchString := "7357"
	match := regexp.MustCompile(regexPattern).MatchString(matchString)

	fmt.Printf("Original word: %s\n", word)
	fmt.Printf("Generated regex: %s\n", regexPattern)
	fmt.Printf("Does '%s' match the regex? %t\n", matchString, match)
}
