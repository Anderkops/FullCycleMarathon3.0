package main

import (
	"fmt"
	"os"
)

func mainReturnWithCode() int {
	// do stuff, defer functions, etc.
	return 0 // a suitable exit code
}

func main() {
	fmt.Println("Hello Full Cycle!")
	os.Exit(mainReturnWithCode())
}
