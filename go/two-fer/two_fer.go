package twofer

import "strings"

// ShareWith should have a comment documenting it.
func ShareWith(name string) string {
	var s strings.Builder

	s.WriteString("One for ")

	if len(name) == 0 {
		s.WriteString("you, ")
	} else {
		s.WriteString(name + ", ")
	}

	s.WriteString("one for me.")

	return s.String()
}
