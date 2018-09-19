package main

import (
	"fmt"
	"strconv"
)

func greeting(name string) string {
	return "Hello " + name
}

func getSum(num1, num2 int) string {
	return "Sum of " + strconv.Itoa(num1) + " and " + strconv.Itoa(num2) + " is " + strconv.Itoa(num1+num2)
}

func main() {
	fmt.Println(greeting("Devin"))
	fmt.Println(getSum(12, 12))
}
