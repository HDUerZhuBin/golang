package main

import (
	"fmt"
)

func main(){
	fmt.Println("in main before calling greeting!")
	greeting()
	fmt.Println("in main after calling greeting!")

}

func greeting(){
	fmt.Println("in calling function:this is greeting function!")
}