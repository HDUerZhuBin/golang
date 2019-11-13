package main

import(
	"fmt"
)

func main(){
	var s string = "this is a string"
	var ptr *string = &s 
	fmt.Println("the address of s:%p",&s)
	fmt.Println("the content of ptr:",ptr)

	fmt.Println("before the change:%s",s)
	*ptr = "change"
	fmt.Println("after the change:%s",s)

}