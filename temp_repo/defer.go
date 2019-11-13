package main
import (
	"fmt"
)

func main(){
	function1()
}

func function1(){
	fmt.Println("in function1 at the top")
	function2()
	fmt.Println("in function1 at the bottom")
}

func function2(){
	fmt.Println("function2:deffered until the end of the calling funciton")
}