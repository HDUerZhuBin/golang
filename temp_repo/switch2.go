package main

import "fmt"

func main(){
	var num1 int = 7

	switch {
	case num1<0:
		fmt.Println("number is negative")
	case num1>0 && num1<10:
		fmt.Println("number is between 0 and 10")
	default:
		fmt.Println("number is 10 or greater")
	}
}