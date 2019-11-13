package main
import (
	"fmt"
)

func main(){
	temp_ret := 0
	for i:=0;i<=10;i++{
		temp_ret = fibonacci(i)
		fmt.Println("fibonacci(",i,") is : ",temp_ret)
	}
}

func fibonacci(n int)(int){
	var ret int
	if n<=1{
		ret = 1
	}else{
		ret = fibonacci(n-2)+fibonacci(n-1)
	}
	return ret
}