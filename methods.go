package main // This is to import start with 

import ("fmt")

type car struct {
	gasPedal uint16 // min value max 65535
	brakePedal uint16
	steeringWheel int16 // -32k -> +32k
	topSpeedKmh float64
}

func main (){
	aCar := car{gasPedal:2213, 
				brakePedal:0, 
				steeringWheel: 12572, 
				topSpeedKmh: 225.0}
	fmt.Println(aCar.gasPedal)
}
