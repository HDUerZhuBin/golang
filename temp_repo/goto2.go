package main

import (
	"fmt"
)

func main(){
	a := 1 
	goto TARGET
	b := 9

TARGET:
	b += a
	fmt.Println("a is: ",a,"**** b is: ",b)
}