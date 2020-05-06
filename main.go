package main

import (
	"fmt"
	arithmetic "github.com/eddypastika/go-arithmetic"
)

func main() {
	a, b := 20, 20

	m := arithmetic.Multiply(a, b)
	fmt.Println("Multiply result: ", m)

	add := arithmetic.Add(a, b)
	fmt.Println("Add result: ", add)

	s := arithmetic.Subtract(a, b)
	fmt.Println("Subtract result: ", s)

}
