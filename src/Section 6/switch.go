package main

import (
	"fmt"
	"time"
)

func main() {
	language := "GoLang"

	switch language {
	case "Python":
		fmt.Println("You are learning python") // go lang adds break statements to the end of switch statements automatically. no need to add break

	case "Go", "GoLang":
		fmt.Println("Good. Go for go")

	default:
		fmt.Println("Any programming language is good start ")
	}

	n := 5

	switch true {
	case n%2 == 0:
		fmt.Println("Even number")

	case n%2 != 0:
		fmt.Println("Odd number")
	}

	hour := time.Now().Hour()
	//fmt.Println(hour)

	switch true {
	case hour < 12:
		fmt.Println("Good morning")

	case hour < 17:
		fmt.Println("Good Afternoon")

	default:
		fmt.Println("Good Evening")
	}
}
