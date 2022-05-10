package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"unsafe"
)

func main() {
	var cities []string

	//nil means the value hasn't been initialized yet
	// a nil slice is empty slice with zero capacity
	fmt.Println(cities == nil)
	fmt.Printf("cities %#v\n", cities)
	fmt.Println(len(cities))

	// cities[0] = "London"
	numbers := []int{2, 3, 4, 5}
	fmt.Println(numbers)

	nums := make([]int, 2)
	fmt.Printf("%#v\n", nums)

	type names []string
	friends := names{"Dan", "Maria"}

	fmt.Println(friends)
	myFriend := friends[0]

	fmt.Println(myFriend)
	friends[0] = "Gap"
	fmt.Println(friends[0])

	for i, v := range numbers {
		fmt.Printf("Index %v, Value %v\n", i, v)
	}

	var n []int
	n = numbers

	fmt.Println(n)

	var nn []int
	fmt.Println(nn == nil)

	m := []int{}
	fmt.Println(m == nil)

	a, b := []int{1, 2, 3}, []int{1, 5, 3}
	//fmt.Println(a == b)

	var eq bool = true

	a = nil
	for i, v := range a {
		if v != b[i] {
			eq = false
			break
		}
	}

	if len(a) != len(b) {
		eq = false
	}

	if eq {
		fmt.Println("a and b slice are equal")
	} else {
		fmt.Println("a and b slice are not equal")
	}

	numbers1 := []int{2, 3}
	numbers1 = append(numbers1, 10)
	fmt.Println(numbers1)

	numbers1 = append(numbers1, 20, 30, 40)
	fmt.Println(numbers1)

	nnn := []int{100, 200}
	numbers1 = append(numbers1, nnn...)
	fmt.Println(numbers1)

	src := []int{10, 20, 30}
	dst := make([]int, 0)
	mn := copy(dst, src)
	fmt.Println(src, dst, mn)

	arr := [5]int{1, 2, 3, 4, 5}
	// arr[start:stop]

	brr := arr[1:3]
	fmt.Printf("%v %T\n", brr, brr)

	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := s1[1:3]
	fmt.Println(s2)

	fmt.Println(s1[2:])
	fmt.Println(s1[:3])
	fmt.Println(s1[:])

	//fmt.Println(s1[:100])
	s1 = append(s1[:4], 100)
	fmt.Println(s1)

	s1 = append(s1[:4], 200)
	fmt.Println(s1)

	ss1 := []int{10, 20, 30, 40, 50}

	s3, s4 := ss1[0:2], ss1[1:3]

	s3[1] = 600
	fmt.Println(ss1, s4)

	arr1 := [4]int{10, 20, 30, 40}
	slice1, slice2 := arr1[0:2], arr1[1:3]

	arr1[1] = 2

	fmt.Println(arr1, slice1, slice2)

	newFunc()
}

func newFunc() {
	cars := []string{"Ford", "Honda", "Range Rover", "Audi"}
	newCars := []string{}

	newCars = append(newCars, cars[0:2]...)

	cars[0] = "Nissan"

	fmt.Println(cars, newCars)

	s1 := []int{10, 20, 30, 40, 50}
	newSlice := s1[0:3]

	fmt.Println(len(newSlice), cap(newSlice))

	newSlice = s1[2:5]
	fmt.Println(len(newSlice), cap(newSlice))

	fmt.Printf("%p\n", &s1)

	fmt.Printf("%p %p\n", &s1, &newSlice)

	newSlice[0] = 1000

	fmt.Println(s1)

	a := [5]int{1, 2, 3, 4, 5}
	s := []int{1, 2, 3, 4, 5}

	fmt.Printf("array size in bytes %d \n", unsafe.Sizeof(a))
	fmt.Printf("slice size in bytes %d \n", unsafe.Sizeof(s))

	newFunc2()
}

func newFunc2() {
	var nums []int
	fmt.Printf("%#v\n", nums)

	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums))

	nums = append(nums, 1, 2)
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums))

	nums = append(nums, 3)
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums))

	nums = append(nums, 4)
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums))

	nums = append(nums, 100)
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums))

	nums = append(nums[0:4], 200, 300, 400, 500, 600)
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums))
	fmt.Println(nums)

	letters := []string{"A", "B", "C", "D", "E", "F"}

	letters = append(letters[:1], "X", "Y")
	fmt.Println(letters)
	fmt.Printf("Length: %d, Capacity: %d \n", len(letters), cap(letters))

	//fmt.Println(letters[4])
	fmt.Println(letters[3:6])

	newNums := []float64{1., 2., 3.}

	newNums = append(newNums, 10.1)
	newNums = append(newNums, 4.1, 5.5, 6.6)
	fmt.Println(newNums)

	nn := []float64{10., 12.}

	newNums = append(newNums, nn...)
	fmt.Println(newNums)

	if len(os.Args) < 2 { //if not run with arguments
		log.Fatal("Please run the program with arguments (2-10 numbers)!")

	}

	//taking the arguments in a new slice
	args := os.Args[1:]

	// declaring and initializing sum and product of type float64
	sum, product := 0., 1.

	if len(args) < 2 || len(args) > 10 {
		fmt.Println("Please enter between 2 and 10 numbers!")
	} else {

		for _, v := range args {
			num, err := strconv.ParseFloat(v, 64)
			if err != nil {
				continue //if it can't convert string to float64, continue with the next number
			}
			sum += num
			product *= num
		}
	}

	fmt.Printf("Sum: %v, Product: %v\n", sum, product)

	nums2 := []int{5, -1, 9, 10, 1100, 6, -1, 6}

	summ := 0
	for _, v := range nums2[2 : len(nums2)-2] {
		fmt.Println(v)
		summ += v
	}

	fmt.Println(summ)

	years := []int{2000, 2001, 2002, 2003, 2004, 2005, 2006, 2007, 2008, 2009, 2010}

	newYears := append(years[:3], years[len(years)-3:]...)

	fmt.Println(newYears)

}
