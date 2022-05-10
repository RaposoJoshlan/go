package main

import "fmt"

func main() {
	var employees map[string]string

	fmt.Printf("%#v\n", employees)

	fmt.Printf("No of pairs: %d\n", len(employees))

	fmt.Printf("Value for key %q is %q\n", "Dan", employees["Dan"])

	var accounts map[string]float64
	fmt.Printf("%v\n", accounts["0x323"])

	var m1 map[[5]int]string
	_ = m1

	//employees["Dan"] = "Programmer"

	people := map[string]float64{}
	people["John"] = 21.4
	people["Mary"] = 10
	fmt.Println(people)

	map1 := make(map[int]int)
	map1[4] = 8

	fmt.Println(map1)

	balances := map[string]float64{
		"USD": 45.4,
		"EUR": 34.4,
		"INR": 45.5,
	}

	m := map[int]int{1: 10, 2: 19}
	fmt.Println(m)
	fmt.Println(balances)

	balances["USD"] = 500.3
	balances["GBP"] = 100.
	fmt.Println(balances)

	fmt.Println(balances["RON"])

	balances["RON"] = 0

	v, ok := balances["RON"]
	if ok {
		fmt.Println("Ron Balance is:", v)
	} else {
		fmt.Println("RON key doesn't exist in the map")
	}

	for k, v := range balances {
		fmt.Printf("Key: %#v, Value: %#v\n", k, v)
	}

	delete(balances, "USD")
	fmt.Println(balances)

	comparingMaps()

}

func comparingMaps() {
	a := map[string]string{"A": "X"}
	b := map[string]string{"A": "X"}

	// fmt.Println(a == b)
	s1 := fmt.Sprintf("%s", a)
	s2 := fmt.Sprintf("%s", b)

	fmt.Println(s1, s2)

	if s1 == s2 {
		fmt.Println("Maps are equal")
	} else {
		fmt.Println("Maps are not equal")
	}

	mapHeader()

}

func mapHeader() {
	friends := map[string]int{"Dan": 40}
	neighbors := friends

	friends["Dan"] = 50
	fmt.Println(neighbors)

	people := make(map[string]int)

	for k, v := range friends {
		people[k] = v
		fmt.Println("Here")
	}

	people["Anne"] = 18
	fmt.Printf("%#v --------- %#v\n", people, friends)

	delete(friends, "Dan")
	fmt.Println(friends)

	delete(people, "Andr")

	m := map[int]bool{1: true, 2: false, 3: false}
	delete(m, 2)

	for k, v := range m {
		fmt.Printf("Key is: %d, Value is %v\n", k, v)
	}
}
