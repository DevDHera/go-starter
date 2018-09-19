package main

import "fmt"

func main() {
	// Define Map
	emails := make(map[string]string)

	// Assign Key Values
	emails["Devin"] = "aaa@gmail.com"
	emails["Herath"] = "herath@gmail.com"
	emails["Don"] = "don@gmail.com"

	fmt.Println(emails)
	fmt.Println(emails["Devin"])
	fmt.Println(len(emails))
	fmt.Println(emails)

	// Delete from Map
	delete(emails, "Don")
	fmt.Println(emails)

	// Declare Map and add key values
	studentIds := map[int]string{1: "Devin", 2: "Don"}

	fmt.Println(studentIds)
}
