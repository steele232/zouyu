package zouyu

import (
	"fmt"
	"io/ioutil"
	"log"
	"testing"
)

func TestSearchAndReplace(t *testing.T) {
	fmt.Println("Happy Testing!")
	file1 := input1
	fnOut1 := ConvertFile(file1)
	if fnOut1 != output1 {
		t.Fail()
	}
}

// writeTwoOutputsForComparison is
// a small utility for writing to two different
// files so that you can use the 'diff' command on
// the two files and see what the difference is.
// It should be helpful for figuring out why your tests
// aren't passing.
func writeTwoOutputsForComparison(output1, output2 string) {

	// Let's write to a file to see the diff.
	bytes := []byte(output1)
	err := ioutil.WriteFile("input.txt", bytes, 0644)
	if err != nil {
		log.Fatal(err)
	}
	bytes = []byte(output2)
	err = ioutil.WriteFile("output.txt", bytes, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

var input1 = `// +build ignore

包裹 主要

进口 (
	"fmt"
	"time"
)

函数 主要() {

	循环 i := 0; i < 10; i++ {
		fmt.Printf("第%v个\n", i+1)
	}

	fmt.Println("走语考试开始了")

	项目 := 初始化([]整数, 3)
	循环 _, v := 范围 项目 {
		fmt.Println(v)
	}

	走 函数() {
		fmt.Println("时间过得很慢")
	}()
	time.Sleep(2 * time.Second)

	走 函数() {
		fmt.Println("时间过得很慢")
	}()
	time.Sleep(2 * time.Second)
}`

var output1 = `package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 0; i < 10; i++ {
		fmt.Printf("第%v个\n", i+1)
	}

	fmt.Println("走语考试开始了")

	项目 := make([]int, 3)
	for _, v := range 项目 {
		fmt.Println(v)
	}

	go func() {
		fmt.Println("时间过得很慢")
	}()
	time.Sleep(2 * time.Second)

	go func() {
		fmt.Println("时间过得很慢")
	}()
	time.Sleep(2 * time.Second)
}`
