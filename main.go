package main

import (
	"fmt"
)

func main() {
	fmt.Print(`

	Welcome to GREET App

	Give a name to greet.

	`)

	fmt.Print("Enter a name: ")

	var name string
	fmt.Scanln(&name)

	fmt.Printf(`
	
	Hello, %s!

	`, name)
}
