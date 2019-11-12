package main

import(
	"fmt"
	"strings"
)

func main(){
	var originS string = "hi there! "
	var newS string 

	newS = strings.Repeat(originS,3)
	fmt.Printf("the new repeated string is:%s\n",newS)
}