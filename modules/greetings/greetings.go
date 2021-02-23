//Tutorial: Getting Started > Create a module > "Start a module that others can use" from golang.org/doc

package greetings

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v, Welcome !", name)
	return message
}
