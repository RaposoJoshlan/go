package main

import "fmt"

func main() {

	outer := 19
	_ = outer

	people := [5]string{"Helen", "Mark", "mar", "ma", "kram"}

	friends := [2]string{"Mark", "Mary"}

outer:
	for i, name := range people {
		for _, friend := range friends {
			if name == friend {
				fmt.Printf("Found a friend %q at index %d\n", friend, i)
				break outer
			}
		}
	}

	fmt.Println("Next instruction after the break")

	i := 0

loop:
	if i < 5 {
		fmt.Println(i)
		i++
		goto loop
	}

	// 	goto todo
	// 	x := 5
	// todo:
	// 	fmt.Println("Something here")

}
