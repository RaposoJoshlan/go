package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	bs := []byte("The Go gopher is an iconic mascot!")
	err := ioutil.WriteFile("info.txt", bs, 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", bs)
}
