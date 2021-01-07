package main

import (
	"fmt"
	"time"
)


func main() {
	ch := make(chan int)

	go func() {
		for {
			select {
			case <-ch: fmt.Println(0)
			case <-ch: fmt.Println(1)
			}
		}
	}()

	go func() {
		for {
			ch <- 0
		}
	}()

	time.Sleep(2 * 1e9)
}



