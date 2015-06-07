package main

import "fmt"

func main() {
	defer fmt.Println("hold off");
	v := 200
	for i := 100; i <= v; i *= 2 {
		fmt.Println(i)
	}
}
