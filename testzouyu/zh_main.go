// +build ignore

包裹 主要

进口 (
	"time"
	"fmt"
)

函数 主要() {

	循环 i:=0; i < 10; i++ {
		fmt.Printf("第%v个\n", i+1)
	}

	fmt.Println("走语考试开始了")

	项目 := 初始化([]整数, 3)
	循环 _, v := 范围(项目) {
		fmt.Println(v)
	}

	走 函数() {
		fmt.Println("时间过得很慢")
	}()
	time.Sleep(2 *time.Second)

	走 函数() {
		fmt.Println("时间过得很慢")
	}()
	time.Sleep(2 *time.Second)
}