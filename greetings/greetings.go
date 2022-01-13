// Tutorial - create a module ✅

package greetings

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

//:= operator is a shortcut for declaring and initializing a variable in one line
