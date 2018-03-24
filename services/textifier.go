package services

import (
	"regexp"
	"strings"
	"unicode"
)

// Textifier transforms input text
type Textifier interface {
	Transform(text string) string
}

// GetService returns service based on given type
func GetService(serviceType string) Textifier {
	stdType := strings.ToLower(serviceType)
	if stdType == "nyanify" {
		return &Nyanify{}
	}
	if stdType == "reversify" {
		return &Reversify{}
	}
	panic("no service was found with given type: " + serviceType)
}

//ToProperCase Convert potato word to proper case
func ToProperCase(word string, original string) string {
	originalRunes := []rune(original)
	newRunes := []rune(word)

	for i := 0; i < len(newRunes); i++ {
		if unicode.IsUpper(originalRunes[i]) {
			newRunes[i] = unicode.ToUpper(newRunes[i])
		} else if unicode.IsLower(originalRunes[i]) {
			newRunes[i] = unicode.ToLower(newRunes[i])
		}
	}

	return string(newRunes)
}

// MapWords maps word by word with given transform function
func MapWords(text string, transform func(string) string) string {
	var result []string
	words := strings.Fields(text)

	for _, word := range words {
		newWord := word
		tokens := getWords(newWord)
		for _, token := range tokens {
			newWord = strings.Replace(newWord, token, transform(token), -1)
		}
		result = append(result, newWord)
	}

	return strings.Join(result, " ")
}

func getWords(token string) []string {
	r, _ := regexp.Compile("\\w+")
	return r.FindAllString(token, -1)
}
