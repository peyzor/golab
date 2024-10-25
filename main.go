package main

import (
	"fmt"
	"strconv"
)

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
	//var planet = "venus"
	//var distance = 261
	//fmt.Printf("%v is %v km away from us. think! %[2]v away is %[1]v\n", planet, distance)

	// keyed items in array initialization
	//count := 0
	//for range [...]int{9: 0} {
	//	count += 1
	//}
	//fmt.Println(count)
	//
	//isVowel := []bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
	//fmt.Println(isVowel)
	//
	//isWeekend := []bool{5: true, 6: true}
	//fmt.Println(isWeekend)

	// Confusion
	//a := [...]int{5, 4: 1, 0, 2: 3, 2, 1: 4}
	//fmt.Println(a)

	//result := 1.0 / 10.0
	//for range [...]int{8: 0} {
	//	result += 1.0 / 10.0
	//}
	//
	//fmt.Printf("%.60f\n", result)

	//fmt.Printf("%b\n", 69)
	//fmt.Printf("%010b\n", 69)
	//fmt.Printf("%010b\n", 70)
	//fmt.Printf("%010b\n", 420)

	i, _ := strconv.ParseInt("1000101", 2, 8)
	fmt.Println(int8(i))
	e, _ := strconv.Atoi("69")
	fmt.Println(e)

}
