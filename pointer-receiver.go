package main

import "fmt"

const usixteenbitmax float64 = 65535
const kmhMultiple float64 = 1.60934

type car struct { //initializes the class(or struct) of car with parameters
	gasPedal uint16 // min value max 65535
	brakePedal uint16
	steeringWheel int16 // -32k -> +32k
	topSpeedKmh float64
}

func (c car) kmh() float64{ //c is the variable for the car, which we can use to reference its parameters
	return float64(c.gasPedal) * (c.topSpeedKmh/usixteenbitmax)
}

func (c car) mph() float64{ //c is the variable for the car, which we can use to reference its parameters
	return float64(c.gasPedal) * (c.topSpeedKmh/usixteenbitmax/kmhMultiple)
}

func (c *car) newTopSpeed(newspeed float64) {
	c.topSpeedKmh = newspeed
}

func newerTopSpeed(c car, speed float64) car {
	c.topSpeedKmh = speed
	return c
}

func main (){
	newCar := car{gasPedal:2213, // Makes a new car 
				brakePedal:0, 
				steeringWheel: 12572, 
				topSpeedKmh: 225.0}

	fmt.Println(newCar.gasPedal) //This will print the gas Pedal information using the (.) method
	fmt.Println(newCar.kmh())
	fmt.Println(newCar.mph())
	// newCar.newTopSpeed(550)
	newCar = newerTopSpeed(newCar, 500)
	fmt.Println(newCar.kmh())
	fmt.Println(newCar.mph())
}

