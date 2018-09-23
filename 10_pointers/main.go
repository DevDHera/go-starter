package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n", b)

	// Reading value assigned to pointer
	fmt.Println(*b)
	fmt.Println(*&a)

	// Change val with pointer
	*b = 10
	fmt.Println(a)
}
