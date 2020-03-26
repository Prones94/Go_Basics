package main

import "fmt"

func main() {
	x := 15
	a := &x // This will point to the memory address
	fmt.Println(a)
	fmt.Println(*a) // This will read through the memory address
	*a = 5 // Reading through the memory address and changing the data inside 
}