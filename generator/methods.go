package generator

import (
	"math/rand"
	"strings"
)

func (generator Generator) Generate() (result string) {
	var alphabet string
	if generator.UseLowerCaseLetters {
		alphabet += lowerCaseLetters
	}
	if generator.UseUpperCaseLetters {
		alphabet += upperCaseLetters
	}
	if generator.UseNumbers {
		alphabet += numbers
	}
	if generator.UseSymbols {
		alphabet += symbols
	}
	if generator.ExcludeSimilarCharacters {
		generator.Exclude += similarCharacters
	}
	for _, letter := range generator.Exclude {
		alphabet = strings.ReplaceAll(alphabet, string(letter), "")
	}

	if len(alphabet) == 0 {
		return result
	}

	randGenrator := rand.New(rand.NewSource(1337))
	for index := uint(0); index < generator.Length; index++ {
		randIndex := randGenrator.Intn(len(alphabet))
		result += string(alphabet[randIndex])
	}
	return result
}
