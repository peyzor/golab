package main

import "fmt"

func main() {
	//Builder Pattern usage
	//assembly := car.NewBuilder().Color(car.GreenColor)
	//
	//familyCar := assembly.Wheels(car.SportsWheels).TopSpeed(50 * car.MPH).Build()
	//_ = familyCar.Drive()
	//
	//sportsCar := assembly.Wheels(car.SteelWheels).TopSpeed(150 * car.MPH).Build()
	//_ = sportsCar.Drive()

	// Arg indexing starts from 1
	var planet = "venus"
	var distance = 261
	fmt.Printf("%v is %v km away from us. think! %[2]v away is %[1]v", planet, distance)
}
