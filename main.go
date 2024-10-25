package main

import (
	"fmt"
	"math"
	"unsafe"
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

	//i, _ := strconv.ParseInt("1000101", 2, 8)
	//fmt.Println(int8(i))
	//e, _ := strconv.Atoi("69")
	//fmt.Println(e)

	// type byte = uint8
	//var b byte
	//b = 0
	//fmt.Printf("%08b = %d\n", b, b)
	//
	//b = 255
	//fmt.Printf("%08b = %d\n", b, b)

	fmt.Println("int: ", math.MinInt, math.MaxInt)
	fmt.Println("int8: ", math.MinInt8, math.MaxInt8)
	fmt.Println("uint32 :", 0, math.MaxUint32)
	fmt.Println("uint64 :", 0, uint64(math.MaxUint64))
	fmt.Println("float32:", math.SmallestNonzeroFloat32, math.MaxFloat32)
	fmt.Println("float64:", math.SmallestNonzeroFloat64, math.MaxFloat64)
	// memory costs
	fmt.Println("int    :", unsafe.Sizeof(int(1)), "bytes")
	fmt.Println("int8   :", unsafe.Sizeof(int8(1)), "bytes")
	fmt.Println("int16  :", unsafe.Sizeof(int16(1)), "bytes")
	fmt.Println("int32  :", unsafe.Sizeof(int32(1)), "bytes")
	fmt.Println("int64  :", unsafe.Sizeof(int64(1)), "bytes")

	fmt.Println("uint   :", unsafe.Sizeof(uint(1)), "bytes")
	fmt.Println("uint8  :", unsafe.Sizeof(uint8(1)), "bytes")
	fmt.Println("uint16 :", unsafe.Sizeof(uint16(1)), "bytes")
	fmt.Println("uint32 :", unsafe.Sizeof(uint32(1)), "bytes")
	fmt.Println("uint64 :", unsafe.Sizeof(uint64(1)), "bytes")

	fmt.Println("float32:", unsafe.Sizeof(float32(1)), "bytes")
	fmt.Println("float64:", unsafe.Sizeof(float64(1)), "bytes")
}
