package main

import (
	"fmt"
)

func f()(ret int){
	defer func(){
		ret++
		fmt.Println("the ret is:",ret)
	}()
	fmt.Println("before return in f()")
	return 1
}

func main(){
	f()
	fmt.Println("*******************this is just split line********************")
	fmt.Println(f())
}