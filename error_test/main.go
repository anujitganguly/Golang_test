package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Starting Program...")
	os.Exit(0)
	fmt.Println("This line will never print unless we manually insert any error")
}

// try to play around with os.Exit(1) and commenting out and all
