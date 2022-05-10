package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

//Spawning go routine

func f1(wg *sync.WaitGroup) {
	fmt.Println("f1 starts")

	for i := 0; i < 3; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}

	fmt.Println("f1 exits")
	wg.Done()
	// (*wg).Done()
}

func f2() {
	fmt.Println("f2 starts")

	for i := 0; i < 8; i++ {
		fmt.Println(i)
	}

	fmt.Println("f2 exits")

}

func main() {

	var wg sync.WaitGroup

	wg.Add(1)

	go f1(&wg)
	fmt.Println("No of goroutines after go f1()", runtime.NumGoroutine())

	f2()

	// time.Sleep(time.Second * 2)
	wg.Wait()
	fmt.Println("Main execution stopped")

}
