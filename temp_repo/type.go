package main

import (
	"fmt"
)

type TZ int

func main(){
	var a,b TZ = 3,4
	c := a + b
	fmt.Println("c has the value of :%d",c) 
}