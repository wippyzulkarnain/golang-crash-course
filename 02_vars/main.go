package main

import "fmt"

func main() {
	//shorthand
	// name := "Brad"
	size := 1.3

	name, email := "Brad", "Brad@gmail.com"

	var age int32 = 37
	var isCool = true

	fmt.Println(name, age, isCool, size, email)
	fmt.Printf("%T\n", isCool)
}
