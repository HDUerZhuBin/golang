package main
import (
	"fmt"
)

var num int = 10
var numx2,numx3 int

func main(){
	numx2,numx3 = getX2AndX3(num)
	PrintValues()
	numx2,numx3 = getX2AndX3_2(num)
	PrintValues()
}

func PrintValues(){
	fmt.Println("num=",num,", 2x num = ",numx2,", 3x num = ",numx3)
}

func getX2AndX3(input int)(int ,int){
	return 2*input,3*input
}

func getX2AndX3_2(input int)(x2 int,x3 int){
	x2 = 2 * input
	x3 = 3 * input

	return 
}