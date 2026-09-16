package main

import (
	"fmt"
	"os"
)

// Echo2 prints its command-line arguments.
func main() {
	fmt.Println(os.Args[1:])
}
