package main

import (
	"fmt"
	"os"
)

// password_strength_analyzer - Analyze password security
func password_strength_analyzer(path string) {
	fmt.Println("========================================")
	fmt.Println("  Password-Strength-Analyzer")
	fmt.Println("  Analyze password security")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Target:", path)
	fmt.Println("Processing...")
	fmt.Println("Done!")
}

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	password_strength_analyzer(path)
}
