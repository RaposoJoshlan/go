package main

import "fmt"

// IN GO Programming language there is only 1 loop and that is for loop

// There is no while loop in go

func main() {
	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}

	j := 10

	for j >= 0 {
		fmt.Print(j)
		j--
	}

	fmt.Printf("\n")

	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println(i)
	}

	count := 0
	for i := 0; true; i++ {
		if i%13 == 0 {
			fmt.Printf("%d is divisible by 13 \n", i)
			count++
		}
		if count == 10 {
			break
		}
	}
	fmt.Println("Just a message after the for loop")
}
