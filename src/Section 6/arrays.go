package main

import (
	"fmt"
	"strings"
)

func main() {
	var numbers [4]int

	fmt.Printf("%v\n", numbers)
	fmt.Printf("%#v\n", numbers)

	var a1 = [4]float64{}
	fmt.Printf("%#v\n", a1)

	var a2 = [3]int{-10, 1, 100}
	fmt.Printf("%#v\n", a2)

	a3 := [4]string{"a", "b", "c", "d"}
	fmt.Printf("%#v\n", a3)

	a4 := [4]string{"x", "y"}
	fmt.Printf("%#v\n", a4)

	//ellipsis operator
	a5 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("%#v\n", a5)
	fmt.Printf("the length of the array is %d\n", len(a5))

	newArr := [3]int{10, 20, 30}
	fmt.Printf("%#v\n", newArr)

	newArr[0] = 4
	newArr[1] = 5
	newArr[2] = 6
	fmt.Printf("%#v\n", newArr)

	for i, v := range newArr {
		fmt.Println("Index:", i, "Value is:", v)
	}

	fmt.Println(strings.Repeat("#", 20))

	for i := 0; i < len(newArr); i++ {
		fmt.Println("Index:", i, "Value is:", newArr[i])
	}

	balances := [2][3]int{
		{5, 6, 7},
		{8, 9, 10},
	}

	fmt.Println(balances)

	m := [2]int{1, 2}
	n := m
	fmt.Println("n equals m is", n == m)

	m[0] = -1
	fmt.Println("n equals m is", n == m)

	grades := [3]int{
		1: 10,
		0: 5,
		2: 7,
	}

	fmt.Println(grades)

	accounts := [3]int{
		2: 50,
	}
	fmt.Println(accounts)

	names := [...]string{
		5: "Dan",
	}

	fmt.Println(names, len(names))

	cities := [...]string{
		5: "Paris",
		1: "new york",
		"london",
	}
	fmt.Printf("%#v\n", cities)

	weekend := [7]bool{
		5: true,
		6: true,
	}

	fmt.Println(weekend)

	nums := [...]int{30, -1, -6, 90, -6}
	i := 0

	for _, v := range nums {
		if v%2 == 0 && v > 0 {
			i++
			fmt.Println(v, i)
		} else {
			continue
		}
	}

	myArray := [3]float64{1.2, 5.6}

	myArray[2] = 6

	a := 10
	myArray[0] = float64(a)

	//myArray[3] = 10.10

	fmt.Println(myArray)

}
