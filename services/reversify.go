package services

// Reversify reverses each word by characters
type Reversify struct{}

// Transform text by reversing it by characters
func (service Reversify) Transform(text string) string {
	return MapWords(text, reversify)
}

func reversify(word string) string {
	chars := []rune(word)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}
