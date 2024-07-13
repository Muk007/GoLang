package main

import (
	"fmt"
)

func main() {
	var a = 5
	fmt.Println(a)
	// Converting a variable into into address
	address := &a
	fmt.Println(address)

	// Converting a address into its value
	value := *address
	fmt.Print(value)
}
