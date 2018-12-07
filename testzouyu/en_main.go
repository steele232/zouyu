package main

import (
	"time"
	"fmt"
)

func main() {

	fmt.Println("走语考试开始了")

	项目 := make([]int, 3)
	for _, v := range(项目) {
		fmt.Println(v)
	}

	go func() {
		fmt.Println("How about that. ")
	}()
	time.Sleep(2 *time.Second)

	go func() {
		fmt.Println("How about that. ")
	}()
	time.Sleep(2 *time.Second)
}