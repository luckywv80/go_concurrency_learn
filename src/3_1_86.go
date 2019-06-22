package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	var stdoutBuff bytes.Buffer
	defer stdoutBuff.WriteTo(os.Stdout)
	inStream := make(chan int, 4)
	go func() {
		defer close(inStream)
		defer fmt.Fprintf(&stdoutBuff, "Product Done.\n")
		for i := 0; i < 8; i++ {
			fmt.Fprintf(&stdoutBuff, "Sending:%d\n", i)
			inStream <- i
		}
	}()

	for integer := range inStream {
		fmt.Fprintf(&stdoutBuff, "Received %v.\n", integer)
	}
}
