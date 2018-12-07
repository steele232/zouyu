// +build ignore

包裹 主要

进口 (
	"time"
	"fmt"
)

子程序 主要() {

	fmt.Println("ZouYu test started")

	走 子程序() {
		fmt.Println("How about that. ")
	}()
	time.Sleep(2 *time.Second)

	走 子程序() {
		fmt.Println("How about that. ")
	}()
	time.Sleep(2 *time.Second)
}