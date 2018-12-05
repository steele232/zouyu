package main

import (
	"time"
	"fmt"
)

func main() {

	fmt.Println("ZouYu test started")

	go func() {
		fmt.Println("How about that. ")
	}()
	time.Sleep(2 *time.Second)
}