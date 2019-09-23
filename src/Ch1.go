package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go sendData(ch)
	getDataByRange(ch)
	printmy("hello")
	// getDataByLoop(ch)
}
func printmy(word ...string) {
	fmt.Printf("printmy %v\n", word)
}
func sendData(ch chan string) {
	defer close(ch)
	ch <- "北京"
	fmt.Printf("写入北京\n")
	time.Sleep(1e9)
	ch <- "上海"
	fmt.Printf("写入上海\n")
	time.Sleep(1e9)
	ch <- "山东"
	fmt.Printf("写入山东\n")
	time.Sleep(1e9)
	// close(ch)
}

func getDataByRange(ch chan string) {
	for input := range ch {
		fmt.Printf("收到的数据：%s\n", input)
	}
}

func getDataByLoop(ch chan string) {
	for {
		fmt.Printf("正在接收\n")
		// 若前面已经别读取干净，则不能继续进行
		input, open := <-ch
		if !open {
			break
		}
		fmt.Printf("for循环收到的数据：%s\n", input)
	}
}
