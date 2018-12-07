// +build ignore

包裹 主要

进口 (
	"time"
	"fmt"
)

子程序 主要() {

	fmt.Println("走语考试开始了")

	项目 := 做([]int, 3)
	for _, v := range(项目) {
		fmt.Println(v)
	}

	走 子程序() {
		fmt.Println("How about that. ")
	}()
	time.Sleep(2 *time.Second)

	走 子程序() {
		fmt.Println("How about that. ")
	}()
	time.Sleep(2 *time.Second)
}