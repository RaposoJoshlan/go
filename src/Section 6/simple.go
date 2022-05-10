package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	i, err := strconv.Atoi("45a")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}

	if i, err := strconv.Atoi("2z0"); err == nil {
		fmt.Println("no error. i is ", i)
	} else {
		fmt.Println(err)
	}

	if args := os.Args; len(args) != 2 {
		fmt.Println("One argument is required")
	} else if km, err := strconv.Atoi(args[1]); err != nil {
		fmt.Println("the argument must be an integer", err)
	} else {
		fmt.Printf("%d km in mile mile is %v\n", km, float64(km)*0.621)
		fmt.Printf("%T\n", args[1])
	}
}
