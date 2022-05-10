package main

import "fmt"

func changeValue(qty int, price float64, name string, sold bool) {
	qty = 3
	price = 500.4
	name = "Mobile Phone"
	sold = false
}

func changeValuesByPointers(qty *int, price *float64, name *string, sold *bool) {
	*qty = 3
	*price = 500.5
	*name = "Mobile Phone"
	*sold = false

}

type Product struct {
	price       float64
	productName string
}

func changeProduct(p Product) {
	p.price = 300.
	p.productName = "Bicycle"
}

func changeProductByPointer(p *Product) {
	(*p).price = 300.
	p.productName = "Bicycle"
}

func changeSlice(s []int) {
	for i := range s {
		s[i]++
	}
}

func changeMap(m map[string]int) {
	m["a"] = 10
	m["b"] = 20
	m["c"] = 30
}

func main() {
	qty, price, name, sold := 5, 300.4, "Laptop", true

	fmt.Println("BEFORE calling changeValue():", qty, price, name, sold)

	changeValue(qty, price, name, sold)

	fmt.Println("AFTER calling changeValue():", qty, price, name, sold)

	changeValuesByPointers(&qty, &price, &name, &sold)
	fmt.Println("AFTER calling changeValuesByPointers():", qty, price, name, sold)

	gift := Product{
		price:       100.,
		productName: "Watch",
	}

	changeProduct(gift)

	fmt.Println(gift)

	fmt.Println("BEFORE calling changeProductByPointer():", gift)

	changeProductByPointer(&gift)
	fmt.Println("AFTER calling changeValue():", gift)

	prices := []int{1, 2, 3}
	changeSlice(prices)
	fmt.Println(prices)

	myMap := map[string]int{"a": 1, "b": 2}
	changeMap(myMap)
	fmt.Println(myMap)

	//When should we pass values and pointers and the answer depends on the variable type

	//Coding exercise
	x := 10.10
	fmt.Printf("Address of x is: %p\n", &x)
	ptr := &x
	fmt.Printf("Address of x is: %p and the type is %T\n", ptr, ptr)
	fmt.Printf("Address of pointer: %p\n", &ptr)
	fmt.Printf("Value of x is:%f\n", *ptr)

	xx, yy := 10, 2
	ptrx, ptry := &xx, &yy

	z := *ptrx / *ptry
	fmt.Println(z)

	xxx, yyy := 5.5, 8.8
	change3(&xxx, &yyy)
	fmt.Println(xxx, yyy)

}

func change3(x *float64, y *float64) {
	*x, *y = *y, *x
}
