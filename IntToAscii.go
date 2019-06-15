// Package inttoascii provide two options to convert a string int an ascii char.
package inttoascii

import (
	"fmt"
	"strconv"
)

// IntegerToASCII convert a int in a string using sprintf.
func IntegerToASCII(x int) string {
	return fmt.Sprintf("%d", x)
}

// IntegerToASCII2 convert a int in a string using Itoa.
func IntegerToASCII2(x int) string {
	return strconv.Itoa(x)
}
