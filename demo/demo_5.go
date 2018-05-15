package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		for {
			c:= <-ch
			fmt.Println(c)
		}
	}()

	ch <- 0
	ch <- 0
}
