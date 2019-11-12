package main

import (
	"fmt"
)

func main(){
	var i1 int = 5
	fmt.Println("an integer:%d,its location in memory:%p",i1,&i1)

	var int_ptr *int = &i1
	fmt.Println("the value of memory location %p is %d",int_ptr,*int_ptr)
}