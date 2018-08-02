package main

import (
	"fmt"

	"github.com/kevincowan/go-examples/mathsy"
)

func main() {

	x := mathsy.Sum(6, 7)
	y := mathsy.Multiply(6, 7)

	fmt.Printf("x = %d\n", x)
	fmt.Printf("y = %d\n", y)
}
