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
	for i:=0;i<10;i++{
		if i==11{
			fmt.Println(i)
			return
		}
	}
	fmt.Println(1111111)
}
