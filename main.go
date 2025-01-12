package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
}

func cleanInput(text string) []string {
	var stringSlice []string
	output := strings.ToLower(text)
	stringSlice = strings.Fields(output)
	return stringSlice
}
