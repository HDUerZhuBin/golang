package main
import (
	"fmt"
	"time"
)

const LIM = 41
var fibs [LIM]uint64

func main(){
	var result uint64 = 0
	start := time.Now()
	for i:=0;i<LIM;i++{
		result = fibonacci(i)
		fmt.Println("fibonacci(",i,") is :",result)
	}
	end := time.Now()
	duration := end.Sub(start)
	fmt.Println("the duration is:",duration)
}

func fibonacci(n int)(res uint64){
	if fibs[n] !=0 {
		res = fibs[n]
		return res
	}
	if n <= 1 {
		res = 1
	}else{
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	fibs[n] = res
	return res
}