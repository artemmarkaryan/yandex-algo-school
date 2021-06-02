package main

import (
	"fmt"
)

func Calculate(a, b, c int) bool {
	return (a + b) > c && (a + c) > b && (c + b) > a
}

func main() {
	var a, b, c int
	_, _ = fmt.Scanln(&a)
	_, _ = fmt.Scanln(&b)
	_, _ = fmt.Scanln(&c)

	if Calculate(a, b, c) {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
