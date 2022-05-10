package main

import "fmt"

//Methods with pointer receiver

type Car struct {
	brand string
	price int
}

func changeCar(c Car, newBrand string, newPrice int) {
	c.price = newPrice
	c.brand = newBrand
}

func (c Car) changeCar1(newBrand string, newPrice int) {
	c.brand = newBrand
	c.price = newPrice
}

func (c *Car) changeCar2(newBrand string, newPrice int) {
	(*c).brand = newBrand
	(*c).price = newPrice
}

type distance *int

// func (d *distance) m1() {
// 	fmt.Println("Just a message")
// }

func main() {
	myCar := Car{brand: "Audi", price: 40_000}

	changeCar(myCar, "Renault", 20_000)
	fmt.Println(myCar)

	myCar.changeCar1("Opel", 21_000)
	fmt.Println(myCar)

	myCar.changeCar2("Seat", 25_000)
	fmt.Println(myCar)

	var yourCar *Car
	yourCar = &myCar
	fmt.Println(*yourCar)

	//Valid ways

	yourCar.changeCar2("VW", 45_000)
	fmt.Println(*yourCar)

	(*yourCar).changeCar2("ANOTHER VW", 45_000)
	fmt.Println(*yourCar)

	fmt.Println(myCar)
}
