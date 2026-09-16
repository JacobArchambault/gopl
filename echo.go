package main

import (
	"fmt"
	"os"
)

// Echo2 prints its command-line arguments.
func main() {
	for i, arg := range os.Args {
		fmt.Println(i, arg)	
	}
}
