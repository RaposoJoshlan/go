package main

import (
	"fmt"
	f "fmt"
)

// File scope
//Permitted in Go

const m = 10 // Package scope

func main() { // Package scope
	x := 10 // block or local scope
	_ = x

	f.Println(m, x)

	const m = 10 //Local scoped and can be declared again

	f.Println(m)

	f1()
}

func f1() {
	const m = 9
	fmt.Printf("done from f1 %v\n", m)

	a := 10
	_ = a
}
