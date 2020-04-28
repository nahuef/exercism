// Package twofer returns a string with a message with a name,
// or defaults to "you".
package twofer

import (
	"fmt"
)

// ShareWith returns a string with a phrase.
func ShareWith(name string) string {
	if len(name) == 0 {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
