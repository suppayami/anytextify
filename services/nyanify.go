package services

import (
	"math/rand"
	"strings"
	"time"
)

// Nyanify transform given text to nyannn~
type Nyanify struct{}

// Transform text to nyannn~
func (service Nyanify) Transform(text string) string {
	return MapWords(text, nyanify)
}

func nyanify(word string) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	length := len(word)

	if length == 1 {
		return ToProperCase("n", word)
	}

	if length == 2 {
		return ToProperCase("ny", word)
	}

	if length == 3 {
		return ToProperCase("nya", word)
	}

	if length == 4 {
		return ToProperCase("nyan", word)
	}

	left := length - len("nyan")
	y := r.Intn(left)
	a := r.Intn(left - y)

	newWord := "n" +
		"y" + strings.Repeat("y", y) +
		"a" + strings.Repeat("a", a) +
		"n" + strings.Repeat("n", left-y-a)

	return ToProperCase(newWord, word)
}
