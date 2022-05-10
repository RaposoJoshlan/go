package main

import "fmt"

func main() {
	s1 := "Hi there! Go"
	fmt.Println(s1)
	fmt.Println("He says: \"Hello\"")

	fmt.Println(`He says: "Hello"`)

	s2 := `I like \n go`
	fmt.Println(s2)

	fmt.Println("Price: 100\nBrand: Nike")
	fmt.Println(`
Price: 100
Brand: Nike`)

	fmt.Println(`C:\Users\Andrie`)
	fmt.Println("C:\\Users\\Andrie")

	var s3 = "I love " + "Go " + "Programming"

	fmt.Println(s3)

	fmt.Println("Element at index 0:", s3[0])

	// s3[5] = 'X'

	fmt.Printf("%s\n", s3)
	fmt.Printf("%q\n", s3)

}
