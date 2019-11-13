package main

import (
	"fmt"
)

func main(){
	var p *int = nil
	*p = 0

	fmt.Println("the content of p:",p)
	fmt.Println("the ptr's content:",*p)
}