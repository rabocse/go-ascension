//Tutorial: Getting Started > "Call code in an external package" from golang.org/doc
package main

import (
	"fmt"
	"rsc.io/quote"
)

func main() {

	fmt.Println(quote.Go())
}
