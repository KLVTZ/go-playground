package main

import (
	"fmt"
	"strings"
)

func swap(x, y string) (string, string) {
	return swapAgain(y, x)
}

func swapAgain(x, y string) (string, string) {
	return capPhrase(y), capPhrase(x)
}

func capPhrase(phrase string) string {
	return strings.Title(phrase)
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
