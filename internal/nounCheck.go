package internal

import (
	"unicode"
)

func isPropperNounCandidate(word string) bool {
	if len(word) == 0 {
		return false
	}

	if !unicode.IsUpper(rune(word[0])) {
		return false
	}

	for i, letter := range word {
		if !unicode.IsLetter(letter) {
			return false
		}

		if unicode.IsUpper(letter) && i > 0 {
			return false
		}
	}
	return true
}
