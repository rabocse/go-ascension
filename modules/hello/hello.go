//Tutorial: Getting Started > Create a module > "Call your code from another module" from golang.org/doc

package main

import (
	"fmt"
	"github.com/rabocse/go-lab/ascension/modules/greetings" // In production, this is where the package would be.
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
