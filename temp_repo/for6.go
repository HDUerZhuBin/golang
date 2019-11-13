package main

import (
	"fmt"
)

func main(){

LABEL1:
	for i:=0;i<5;i++{
		for j:=0;j<5;j++{
			if 4==j {
				continue LABEL1
			}
			fmt.Println("i is :",i," and j is:",j)
		}
	}
}