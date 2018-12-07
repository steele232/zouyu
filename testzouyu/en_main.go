package main

import (
	"time"
	"fmt"
)

function main() {

	fmt.Println("ZouYu test started")

	go function() {
		fmt.Println("How about that. ")
	}()
	time.Sleep(2 *time.Second)

	go function() {
		fmt.Println("How about that. ")
	}()
	time.Sleep(2 *time.Second)
}