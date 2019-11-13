package main

import(
	"fmt"
)

func main(){
	var min,max int
	min,max = MinMax(78,65)
	fmt.Println("the min is:",min,", and the max is:",max)
}

func MinMax(a int,b int)(int,int){
	var min,max int
	if a<b{
		min = a 
		max = b
	}else{
		min = b
		max = a
	}
	return min,max
}