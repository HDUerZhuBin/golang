package main

import (
	"fmt"
)

func main(){
	callback(1,Add)
}

func Add(a int,b int){
	fmt.Println("the sum is:",a+b)
}

func callback(y int,f func(int,int)){
	f(y,2)
}