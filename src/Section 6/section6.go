package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	price, inStock := 100, true

	if price > 80 {
		fmt.Println("too expensive")
	}

	_ = inStock

	if price < 100 {
		fmt.Println("Buy it")
	} else if price == 100 {
		fmt.Println("edge")
	} else {
		fmt.Println("expensive")
	}

	cmdLineArgs()
}

func cmdLineArgs() {
	// fmt.Println("os.Args", os.Args)
	// fmt.Println("Path:", os.Args[0])
	// fmt.Println("Path:", os.Args[1])
	// fmt.Println("Path:", os.Args[2])
	// fmt.Println("Path:", len(os.Args))

	var result, err = strconv.ParseFloat(os.Args[1], 64)

	fmt.Printf("%T\n", os.Args[1])
	fmt.Printf("%T\n", result)

	_ = err
}
