package main

import ("fmt")

func add(x,y float64) float64 {
	return x+y
}

func multiple(a,b string)(string, string){
	return a,b
}

func main(){
	// Defining a Variable
	// var num1, num2 float64 = 5.6, 9.5
	num1, num2 := 5.6, 9.5 // Shorthand if variables will always stay the same type using :=

	w1, w2 := "Hey", "there"

	fmt.Println(add(num1, num2))
	fmt.Println(multiple(w1,w2))

	//Convert a type

	var a int= 62
	var b float64 = float64(a)

	x := a // x wil be type int
}	