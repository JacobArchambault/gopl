package main

import (
	"fmt"
	"os"
	"strings"
)

// Echo2 prints its command-line arguments.
func main() {
	fmt.Println(strings.Join(os.Args[0:], " "))
}
