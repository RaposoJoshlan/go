package main

import (
	"fmt"
)

func main() {
	title1, author1, year1 := "The comedy", "Josk", 1320
	title2, author2, year2 := "Macbeth", "Shakespeare", 1606

	fmt.Println("Book 1:", title1, author1, year1)
	fmt.Println("Book 2:", title2, author2, year2)

	type Book struct {
		title  string
		author string
		year   int
	}

	type Book1 struct {
		title, author string
		year          int
	}

	myBook := Book{"The comedy", "Josk", 2030}
	fmt.Printf(" %#v\n", myBook)

	bestBook := Book{year: 1945, title: "Animal Farm", author: "George"}
	_ = bestBook

	aBook := Book{title: "My title"}
	fmt.Printf("%#v\n", aBook)

	lastBook := Book{title: "ANA"}
	fmt.Println(lastBook.title)

	fmt.Printf(" %#v\n", lastBook)
	lastBook.author = "LEO"
	lastBook.year = 1878

	fmt.Printf("%+v\n", lastBook)

	aaBook := Book{title: "ANA", author: "LEO", year: 1878}

	fmt.Println(aaBook == lastBook)

	mybook := aaBook
	mybook.year = 2020

	fmt.Println(mybook, aaBook)

	dian := struct {
		firstName, lastName string
		age                 int
	}{
		firstName: "diana",
		lastName:  "Lol",
		age:       30,
	}

	fmt.Printf("%#v\n", dian)
	fmt.Printf("dian age %d\n", dian.age)

	type bbok struct {
		string
		float64
		bool
	}

	b1 := bbok{"12345 Age", 10.2, false}
	fmt.Printf("%#v\n", b1)

	fmt.Println(b1.string)

	type Employee struct {
		name   string
		salary int
		bool
	}

	e := Employee{"John", 40_000, false}
	fmt.Printf("%#v\n", e)
	e.bool = true
	fmt.Printf("%#v\n", e)
	//new func
	embeddedStruct()

}

func embeddedStruct() {
	type Contact struct {
		email, address string
		phone          int
	}

	type Employee struct {
		name        string
		salary      int
		contactInfo Contact
	}

	John := Employee{
		name:   "Joshlan",
		salary: 30_000,
		contactInfo: Contact{
			email:   "joshlan@.com",
			address: "123 street",
			phone:   1234557,
		},
	}
	fmt.Printf("%+v\n", John)

	fmt.Printf("Employee email: %s\n", John.contactInfo.email)

	John.contactInfo.email = "josh@comp.com"
	fmt.Printf("Employee email: %s\n", John.contactInfo.email)

	myContact := Contact{email: "Domain.com", address: "123 street", phone: 12345677}

	fmt.Println(myContact)
}
