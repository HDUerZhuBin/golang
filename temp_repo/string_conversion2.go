package main

import (
	"fmt"
	"strconv"
)

func main(){
	var origin string = "ABC"
	var newS string

	fmt.Println("the size of ints is:",strconv.IntSize)
	an , err := strconv.Atoi(origin)
	if err != nil {
		fmt.Println("origin %s is not a integer! - exiting with error",origin)
		return 
	}
	fmt.Println("the integer is %d",an)
	an = an + 5
	newS = strconv.Itoa(an)
	fmt.Println("the new string is:%s",newS)
}