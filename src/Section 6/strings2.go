package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s1 := "Golang"

	fmt.Println(len(s1))

	name := "Codruța"
	fmt.Println(len(name))

	fmt.Println(len("%^&*"))

	n := utf8.RuneCountInString(name)
	m := utf8.RuneCountInString("")
	fmt.Println(n, m)

	//Slicing strings
	SlicingStrings()
}

//Slicing strings
func SlicingStrings() {
	s1 := "I love Golang"
	fmt.Println(s1[2:5])

	s2 := "你好"
	fmt.Println(s2[0:2])

	rs := []rune(s2)
	fmt.Printf("%T\n", rs)

	fmt.Println(string(rs[0:2]))

	s := 'U'
	fmt.Printf("%T\n", s)

}
