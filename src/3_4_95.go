package main

import (
	"fmt"
	"time"
)

func main() {
	var c <-chan int
	// c := make(chan int)
	// close(c)
	select {
	case <-c:
		fmt.Print("close")
	case <-time.After(5 * time.Second):
		fmt.Printf("Time out!")
	}
}
