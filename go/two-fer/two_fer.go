// Package twofer implements the ShareWith method.
package twofer

import "fmt"

// ShareWith returns formatted string with the name supplied
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
