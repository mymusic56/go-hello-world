package main

import (
	"fmt"
	"time"
)

func forTest() {
	for true {
		fmt.Println(time.Now())
		fmt.Println(time.Now().Unix())
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		fmt.Println(time.Now().Date())
		fmt.Println("--sleep--")
		time.Sleep(1000000000)
		fmt.Println("--sleep--")
		time.Sleep(1*time.Second)
	}
}

func main() {
	forTest()
}
