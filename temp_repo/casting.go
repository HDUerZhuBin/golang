package main

import (
	"fmt"
)

func main(){
	var n int16 = 34
	var m int32

	m = int32(n)

	fmt.Println("32bit int is:",m)
	fmt.Println("16bit int is:",n)
}