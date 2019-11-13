package main

import (
	"fmt"
)

func main(){
	data := 0
	ret := &data
	Multiply(12,25,ret)

	fmt.Println("the ret is:",*ret)
}

func Multiply(a,b int,ret *int){
	*ret = a * b
}