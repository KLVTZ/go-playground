package main

import "fmt"

func main() {
	i := 4
	i = 8 << uint(i)
	fmt.Println(i)
}

