/****************************************************************************************
闭包函数(closure function):
函数使用了上下文环境中的变量，但是没有显式的传入该参数，而是通过将上下文环境通过闭包的形式进行封装
*****************************************************************************************/

package main
import (
	"fmt"
)

func main(){
	p2 := Add2()
	fmt.Println("call add2 for 3 gives:",p2(3))
	TwoAdder := Adder(2)
	fmt.Println("the result is:",TwoAdder(3))
}

// 闭包函数返回的是一个函数名称（地址），需要对返回值（函数名称）进行传参显性调用
func Add2() func(b int) int{ //这里是将传入的参数b进行了隐藏，该函数通常是匿名函数
	return func(b int) int{
		return b + 2
	}
}

func Adder(a int) func(b int) int{
	return func(b int)int{
		return a + b
	}
}