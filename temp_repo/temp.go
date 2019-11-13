package main

import(
	"runtime"
	"fmt"
)

var prompt string = "enter a digit,e.g. 3 or %s to quit"

func init(){

	fmt.Println("first in the init func...")

	if runtime.GOOS == "window"{
		fmt.Println(prompt,"ctrl + z,enter")
	}else{
		fmt.Println(prompt,"ctrl + d")
	}
}

func Abs(x int) int {
	if x > 0{
		return x
	}
	return -1 * x

}

func isGreate(x int,y int) bool{
	if x > y {
		return true
	}
	return false
}

func main(){
	fmt.Println("******************************")
	fmt.Println("after the init func...")
}