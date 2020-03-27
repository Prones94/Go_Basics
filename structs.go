package main  

import ("fmt")

type car struct { //initializes the class(or struct) of car with parameters
	gasPedal uint16 // min value max 65535
	brakePedal uint16
	steeringWheel int16 // -32k -> +32k
	topSpeedKmh float64
}

func main (){
	newCar := car{gasPedal:2213, // Makes a new car 
				brakePedal:0, 
				steeringWheel: 12572, 
				topSpeedKmh: 225.0}
	fmt.Println(newCar.gasPedal) //This will print the gas Pedal information using the (.) method
}
