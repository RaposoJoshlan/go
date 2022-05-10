package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func myName() {
	fmt.Println("Joshlan")
}

func f2(a int, b int) {
	fmt.Println("a + b =", a+b)
}

func f3(a, b, c int, d, e float64, s string) {
	fmt.Println(a, b, c, d, e, s)
}

func f4(a float64) float64 {
	return math.Pow(a, a)
}

func f5(a, b int) (int, int) {
	return a + b, a * b
}

func sum(a, b int) (s int) {
	fmt.Println(s)

	s = a + b

	return
}

// Variadic function
func ff1(a ...int) {
	fmt.Printf("%T\n", a)
	fmt.Printf("%#v\n", a)
}

func ff2(a ...int) {
	a[0] = 50
}

func sumAndProd(a ...float64) (float64, float64) {
	sum := 0.
	product := 1.

	for _, v := range a {
		sum += v
		product *= v

	}

	return sum, product
}

func personInfo(age int, names ...string) string {
	fullName := strings.Join(names, "_")

	returnString := fmt.Sprintf("Age: %d, Full name: %s\n", age, fullName)

	return returnString
}

func foo() {
	fmt.Println("This is foo()")
}

func bar() {
	fmt.Println("This is bar()")
}

func fooBar() {
	fmt.Println("This is foobar()")
}

func increment(i int) func() int {
	return func() int {
		i++
		return i
	}
}

func main() {
	myName()
	f2(10, 20)
	f3(1, 2, 3, 4.0, 5.0, "D")

	p := f4(5.1)
	fmt.Println(p)

	a, b := f5(8, 9)
	fmt.Println(a, b)

	mySum := sum(4, 8)
	fmt.Println(mySum)

	//Variadic function
	ff1(1, 2, 3, 4, 5, 6, 7)
	ff1()

	nums := []int{1, 2}
	nums = append(nums, 3, 4, 5)

	ff1(nums...)

	ff2(nums...)
	fmt.Println("Nums:", nums)

	ss, pp := sumAndProd(2.0, 5., 10., 11., 13., 14.)

	fmt.Println(ss, pp)

	info := personInfo(30, "Wolf", "Dog", "Cat")
	fmt.Println(info)

	defer foo()
	bar()

	fmt.Println("Just a string after defering foo and calling bar")

	defer fooBar()

	file, err := os.Open("functions.go")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	func(msg string) {
		fmt.Println(msg)
	}("I'm an anonymous func")

	aa := increment(10)

	fmt.Printf("%T\n", aa)
	aa()
	fmt.Println(aa())
	fmt.Println(aa())
	fmt.Println(aa())

}
