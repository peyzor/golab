package main

import "github.com/peyzor/golab/car"

func main() {
	//Builder Pattern usage
	assembly := car.NewBuilder().Color(car.GreenColor)

	familyCar := assembly.Wheels(car.SportsWheels).TopSpeed(50 * car.MPH).Build()
	_ = familyCar.Drive()

	sportsCar := assembly.Wheels(car.SteelWheels).TopSpeed(150 * car.MPH).Build()
	_ = sportsCar.Drive()
}
