package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println(sum)
	fmt.Println(recurse(0))
}

func recurse(start int) int {
	if start == 10 {
		return 0
	} else {
		return start + recurse(start+1)
	}
}
