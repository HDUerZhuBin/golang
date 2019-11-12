package main

import (
	"fmt"
	"strconv"
)

func main(){
	var origin string = "555"
	var an int
	var newS string

	fmt.Println("the size of ints is:",strconv.IntSize)

	an,_ = strconv.Atoi(origin)
	fmt.Println("the integer is:\n",an)

	an = an + 5
	newS = strconv.Itoa(an)
	fmt.Println("the new string is:\n",newS)
}