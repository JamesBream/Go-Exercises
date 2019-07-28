// Package twofer returns a string "One for `name`, one for me"
package twofer

import (
	"fmt"
)

// ShareWith builds a string with the name param passed in
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
