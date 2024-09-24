package twofer

import "strings"

// ShareWith takes a name as a string parameter and returns a string that
// represents the phrase "One for you, one for me." or "One for name, one for me."
// If the name parameter is an empty string, it will use "you" instead.
func ShareWith(name string) string {
	var sb strings.Builder

	sb.WriteString("One for ")

	if name == "" {
		sb.WriteString("you, ")
	} else {
		sb.WriteString(name + ", ")
	}

	sb.WriteString("one for me.")

	return sb.String()
}
