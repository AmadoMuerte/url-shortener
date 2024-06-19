package random

import (
	"math/rand"
	"strings"
)

func NewRandomString(stringLength int) string {
	var letters = []string{"q", "w", "e", "r", "t", "y", "u", "i", "v", "c", "x", "z", "a", "b", "j", "l", "p"}
	var resString string

	for i := 0; i < stringLength; i++ {
		letter := letters[rand.Intn(len(letters))]
		isUpperCaseRand := rand.Intn(2)
		if isUpperCaseRand == 1 {
			resString += strings.ToUpper(letter)
			continue
		}

		resString += strings.ToLower(letter)
	}

	return resString
}
