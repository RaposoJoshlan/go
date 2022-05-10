package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var age int = 30

	fmt.Println("AGE :", age)

	var name = "Hello_World"

	//fmt.Println("Name :", name)

	_ = name //To ignore unused variable

	s := "Learning goLang"

	fmt.Println("S :", s)

	car, cost := "Audi", 50000

	fmt.Println(car, cost)

	var opened = false

	closed := true

	_, _ = opened, closed

	var (
		salary    float64
		firstName string
		gender    bool
	)

	fmt.Println(salary, firstName, gender)

	var a, b, c int
	fmt.Println(a, b, c)

	var i, j int
	i, j = 5, 7

	j, i = i, j //swaping variables

	fmt.Println(i, j)

	sum := 5 + 2.3

	fmt.Println(sum)

	var cities = []string{"a", "b", "c"}
	fmt.Println(cities)

	var cities2 = [3]int{1, 2, 3}
	fmt.Println(cities2)

	balances := map[string]int{
		"USD": 1,
		"GBP": 2,
	}

	fmt.Printf("%T\n", balances)

	type Person struct {
		name string
		age  int
	}

	var you Person
	you.age = 10
	you.name = "Donkey"

	fmt.Println(you.age, you.name)

	codEx5()
	codEx6(86400, 365)

}

func codingEx() {
	var a int = 10
	var b float64 = 20
	var c string = "j"

	fmt.Println(a, b, c)

	x := -20
	y := -15.5
	z := "Gopher"

	fmt.Println(x, y, z)

	var (
		aa int
		bb float64
		cc string
	)

	fmt.Println(aa, bb, cc)

	xx, yy, zz := -20, -15.5, "Gopher"
	_, _, _ = xx, yy, zz

}

func codEx3() {
	var a float64 = 7.1

	x, y := true, 3.7

	a, x = 5.5, false

	_, _, _ = a, x, y
}

const version = "3.1"

func codEx4() {
	name := "Golang"
	fmt.Println(name, version)
}

func codEx5() {
	//const daysWeek int = 7
	//const lightSpeed float64 = 299792458
	//const pi float64 = 3.14159

	const (
		daysWeek   = 7
		lightSpeed = 299792458
		pi         = 3.14159
	)

	fmt.Println(daysWeek, lightSpeed, pi)

}

func codEx6(secPerDay int, daysYear int) {
	var secYear = secPerDay * daysYear
	fmt.Println("Number of Seconds in a Year: ", secYear)
}

func codEx7() {

	const x int = 10

	// declaring a constant of type slice int ([]int): Cannot declare a slice constant
	var m = []int{1, 3, 4, 5, 6, 8}
	_ = m

}

func codEx8() {
	const a = 7
	const b float64 = 5.6
	const c = a * b

	x := 8
	_ = x

	var noIPv6 = math.Pow(2, 128)
	_ = noIPv6

	const (
		//iota starts from zero
		jan = iota + 2
		Jun = iota + 6
		Jul //automatically incremented by one
		Aug
	)
}

func codEx10() {
	x, y, z := 10, 15.5, "Gophers"
	score := []int{10, 20, 30}

	fmt.Printf("x is %d y is %f z is %s \n", x, y, z)
	fmt.Printf("score is %#v\n", score)
	fmt.Printf("z is %q\n", z)

	fmt.Printf(" y type is %T z type is %T", y, score)

	const xx float64 = 1.422349587101

	fmt.Printf("x is %.4f", xx)

}

func codEx11() {
	shape := "circle"
	radius := 3.2
	const pi float64 = 3.14159
	circumference := 2 * pi * radius

	fmt.Printf("Shape: %q\n", shape)
	fmt.Printf("Circle's circumference with radius %f is %f\n", radius, circumference)
	_ = shape
}

func codEx12() {
	var i int = 3
	var f float64 = 3.2

	_, _ = i, f

	fi := float64(i)
	fii := int(f)

	fmt.Printf("i1's type: %T, i1's value: %d\n", fi, fii)

	var s1, s2 = "3.14", "5"

	s := strconv.Itoa(i)
	fmt.Printf("s is %T s value is %q\n", s, s)

	stringToInt, err := strconv.Atoi(s2)

	if err == nil {
		fmt.Printf("i type is %T, i value is %v\n", stringToInt, stringToInt)
	} else {
		fmt.Println("Cannot convert string to int")
	}

	floatToString := fmt.Sprintf("%f", f)
	fmt.Printf("f is %T", floatToString)

	if stringToFloat, err := strconv.ParseFloat(s1, 64); err == nil {
		fmt.Printf("s1 is %T", stringToFloat)
	} else {
		fmt.Println("Cannot convert from string to float")
	}

	x, y := 4, 5.1

	z := x * int(y)
	fmt.Println(z)

	a := 5
	b := 6.2 * float64(a)
	fmt.Println(b)

}

func sunLightToErath() {
	const (
		distance = 149_600_000 * 1000

		speed = 299_792_458
	)

	const time = distance / speed

	_ = time
}

func gfgf() {
	x, y := 0.1, 5
	var z float64

	// Write the correct logical operator (&&, || , !)
	// inside the expression so that result1 will be false and result2 will be true.

	result1 := x < z || int(x) != int(z)
	fmt.Println(result1)

	result2 := y == 1*5 && int(z) == 0
	fmt.Println(result2)
}

type duration int

func hdh() {
	var hr duration

	fmt.Printf("value of %v type of hr is %T", hr, hr)

	hr = 3600

	fmt.Printf("value of %v", hr)

}

type dur int

func ghgh() {
	var hour duration = 3600
	minute := 60
	fmt.Println(hour != duration(minute))
}

type mile float64
type kilometer float64

const m2km = 1.609

func hjk() {
	var miletoberlientoparis mile = 655.3
	var kmBerlinToParis kilometer

	kmBerlinToParis = kilometer(miletoberlientoparis * m2km)

	fmt.Printf("type of this is %T", kmBerlinToParis)
}
