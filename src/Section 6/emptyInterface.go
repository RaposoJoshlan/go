package main

import "fmt"

type emptyInterface interface {
}

type Person struct {
	info emptyInterface
}

type Money float64

func (m Money) print() {
	fmt.Printf("%.2f\n", m)
}

func (m Money) printStr() string {
	return fmt.Sprintf("%.2f\n", m)
}

type Book struct {
	price float64
	title string
}

func (b Book) vat() float64 {
	return b.price * 0.09
}

func (b *Book) discount() {
	(*b).price = b.price * 0.09
}

type vehicle interface {
	License() string
	Name() string
}

type car struct {
	licenceNo string
	brand     string
}

func (c car) License() string {
	return "12345"
}

func (c car) Name() string {
	return "Audi"
}

func main() {
	var empty interface{}

	empty = 5
	fmt.Println(empty)

	empty = "Go"
	fmt.Println(empty)

	empty = []int{1, 2, 3}
	fmt.Println(empty)

	fmt.Println(len(empty.([]int)))

	you := Person{}

	you.info = "Your name"
	you.info = 40
	you.info = []int{}
	fmt.Println(you.info)

	// cod exe

	var salary Money = 4.4
	salary.print()

	book := Book{price: 9.9, title: "My head is small"}

	vat := book.vat()
	fmt.Println(vat)

	book.discount()
	fmt.Println(book)

	var v vehicle = car{licenceNo: "987654321", brand: "Ford"}
	v.License()
	v.Name()

	fmt.Println(v.License(), v.Name())

}
