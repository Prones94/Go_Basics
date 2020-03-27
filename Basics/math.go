package main

import ("fmt"
		// "math" // This will import the math package
		"math/rand" //Importing the random subpackage from the main math package
)

func main() {
	// fmt.Println("The square root of 4 is",math.Sqrt(4))
	fmt.Println("A number from 1-100", rand.Intn(100))
}