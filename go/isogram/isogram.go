package isogram

import (
	"unicode"
)

// IsIsogram determines if a word or phrase is an isogram.
// An isogram is a word or phrase without a repeating letter,
// ignoring case and non-letter characters.
func IsIsogram(word string) bool {
	seen := make(map[rune]bool)

	for _, char := range word {
		// Convert to lowercase for case-insensitive comparison
		char = unicode.ToLower(char)

		// Skip non-letter characters
		if !unicode.IsLetter(char) {
			continue
		}

		// If we've seen this letter before, it's not an isogram
		if seen[char] {
			return false
		}

		// Mark this letter as seen
		seen[char] = true
	}

	return true
}
