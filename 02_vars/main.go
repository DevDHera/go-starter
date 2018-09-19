package main

import "fmt"

func main() {

	// Using Var and Const
	// var name = "Devin"
	var age int32 = 23
	const flag = true
	// flag = false
	var size float32 = 1.3

	// Shorthand Method
	name := "Devin"
	// size := 1.3

	school, institute := "Maliyadeva College", "NIBM"

	fmt.Println(name, age, flag, school, institute)
	fmt.Printf("%T\n", size)
}
