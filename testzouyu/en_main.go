package main

import (
	"time"
	"fmt"
)

func main() {

	for i:=0; i < 10; i++ {
		fmt.Printf("第%v个\n", i+1)
	}

	fmt.Println("走语考试开始了")

	项目 := make([]int, 3)
	for _, v := range(项目) {
		fmt.Println(v)
	}

	go func() {
		fmt.Println("时间过得很慢")
	}()
	time.Sleep(2 *time.Second)

	go func() {
		fmt.Println("时间过得很慢")
	}()
	time.Sleep(2 *time.Second)
}