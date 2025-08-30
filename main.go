package main

import(
	"fmt"
	"strings"
)

func cleanInput(text string) []string {
	s := strings.TrimSpace(text)
	return strings.Fields(s)
}

func main() {
	fmt.Println("Hello, World!")
}