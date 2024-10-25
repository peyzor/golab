package main

import (
	"fmt"
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

	//fmt.Println("int: ", math.MinInt, math.MaxInt)
	//fmt.Println("int8: ", math.MinInt8, math.MaxInt8)
	//fmt.Println("uint32 :", 0, math.MaxUint32)
	//fmt.Println("uint64 :", 0, uint64(math.MaxUint64))
	//fmt.Println("float32:", math.SmallestNonzeroFloat32, math.MaxFloat32)
	//fmt.Println("float64:", math.SmallestNonzeroFloat64, math.MaxFloat64)
	// memory costs
	//fmt.Println("int    :", unsafe.Sizeof(int(1)), "bytes")
	//fmt.Println("int8   :", unsafe.Sizeof(int8(1)), "bytes")
	//fmt.Println("int16  :", unsafe.Sizeof(int16(1)), "bytes")
	//fmt.Println("int32  :", unsafe.Sizeof(int32(1)), "bytes")
	//fmt.Println("int64  :", unsafe.Sizeof(int64(1)), "bytes")
	//
	//fmt.Println("uint   :", unsafe.Sizeof(uint(1)), "bytes")
	//fmt.Println("uint8  :", unsafe.Sizeof(uint8(1)), "bytes")
	//fmt.Println("uint16 :", unsafe.Sizeof(uint16(1)), "bytes")
	//fmt.Println("uint32 :", unsafe.Sizeof(uint32(1)), "bytes")
	//fmt.Println("uint64 :", unsafe.Sizeof(uint64(1)), "bytes")
	//
	//fmt.Println("float32:", unsafe.Sizeof(float32(1)), "bytes")
	//fmt.Println("float64:", unsafe.Sizeof(float64(1)), "bytes")

	//fmt.Println(int8(math.MaxInt8 + 1)) // overflows
	//var n int8 = math.MaxInt8

	// wrap arounds to its negative maximum
	//fmt.Println("max int8     :", n)   // 127
	//fmt.Println("max int8 + 1 :", n+1) // -128

	// wrap arounds to its positive maximum
	//n = math.MinInt8
	//fmt.Println("min int8     :", n)   // -128
	//fmt.Println("min int8 - 1 :", n-1) // 127

	// wrap arounds to its maximum
	//var un uint8
	//
	//fmt.Println("max uint8    :", un)   // 0
	//fmt.Println("max uint8 - 1:", un-1) // 255

	// wrap around to its minimum
	//un = math.MaxUint8
	//fmt.Println("max uint8 + 1:", un+1) // 0

	// floats goes to infinity when overflowed
	//f := float32(math.MaxFloat32)
	//fmt.Println("max float    :", f*1.2)

	// iota expression
	//const (
	//	monday = iota + 1
	//	tuesday
	//	wednesday
	//	thursday
	//	friday
	//	saturday
	//	sunday
	//)
	//
	//fmt.Println(monday, tuesday, wednesday, thursday, friday, saturday, sunday)

	// blank identifier usage in iota
	//const (
	//	xd = -(5 + iota)
	//	_
	//	xdd
	//	x3d
	//)
	//
	//fmt.Println(xd, xdd, x3d)

	//fmt.Println(strings.Repeat("-", 25))

	//city := os.Args[1]
	//
	//switch city {
	//case "Paris", "Lyon":
	//	fmt.Println("Really")
	//	fmt.Println("France")
	//case "Tokyo":
	//	fmt.Println("Now")
	//	fmt.Println("We're talking")
	//default:
	//	fmt.Println("Where?")
	//}

	//num := os.Args[1]
	//i, _ := strconv.Atoi(num)
	//switch {
	//case i > 100:
	//	fmt.Println("greater than 100")
	//	fallthrough
	//case i > 0:
	//	fmt.Println("positive")
	//case i < -10:
	//	fmt.Println("too small")
	//default:
	//	fmt.Println("xd")
	//}

	//source := rand.NewSource(time.Now().UnixNano())
	//rng := rand.New(source)
	//start := 5
	//end := 10
	//switch x := rng.Intn(end-start+1) + start; {
	//case x > 7:
	//	fmt.Println(x)
	//	fmt.Println("greater than 7")
	//case x <= 7:
	//	fmt.Println(x)
	//	fmt.Println("less than or equals 7")
	//}

	//switch h := time.Now().Hour(); {
	//case h >= 18: // 18 to 23
	//	fmt.Println("good evening")
	//case h >= 12: // 12 to 18
	//	fmt.Println("good afternoon")
	//case h >= 6: // 6 to 12
	//	fmt.Println("good morning")
	//default: // 0 to 5
	//	fmt.Println("good night")
	//}

	//for i := 1; i < 10; i++ {
	//	fmt.Printf("%d ", i)
	//}
	//fmt.Println()
	//
	//j := 1
	//for {
	//	if j >= 10 {
	//		break
	//	}
	//
	//	fmt.Printf("%d ", j)
	//	j++
	//}
	//fmt.Println()
	//
	//k := 1
	//for k < 10 {
	//	fmt.Printf("%d ", k)
	//
	//	k++
	//}
	//fmt.Println()
	//
	//list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//for _, x := range list {
	//	fmt.Printf("%d ", x)
	//}

	//words := strings.Fields("Montag Dienstag Mittwoch Donnerstag Freitag Samstag Sonntag")
	//for x, y := range words {
	//	fmt.Println(x+1, y)
	//}

	//outer:
	//	for i := 1; i < 10; i++ {
	//		for j := i; j < 10; j++ {
	//			if i == 5 && j == 5 {
	//				break outer
	//				// also works with continue
	//			}
	//
	//			fmt.Printf("%d ", j)
	//		}
	//		fmt.Println()
	//	}

	//	i := 1
	//loop:
	//	if i < 10 {
	//		fmt.Printf("%d ", i)
	//		i++
	//		goto loop
	//	}

	// # in printf acts like __repr__ in python
	zero := [0]byte{}
	fmt.Printf("%#v\n", zero)

	var names [3]string
	names[0] = "Peyman"
	names[1] = "Ari"
	names[2] = "Thusnelda"
	fmt.Printf("%#v\n", names)
	fmt.Printf("%v\n", names)

	books := [4]string{
		"Kafka's Revenge",
		"Stay Golden",
	}
	fmt.Printf("%#v\n", books)

	// but slices can only be compared to nil
	x := [3]int{1, 2, 3}
	y := [3]int{1, 2, 3}
	fmt.Printf("%t\n", x == y)

	// can not use [...][...]int it's not possible for the compiler to infer
	structure := [...][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	fmt.Printf("%#v\n", structure)
}
