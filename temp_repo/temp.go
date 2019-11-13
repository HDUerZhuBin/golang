package main
import (
	"fmt"
)

func main(){

	arr := [5]int{0,1,2,3,4}
	
	ret := sum(arr[:])
	fmt.Println("the sum_ret:",ret)

}

func sum(a []int)int{
	var sum_ret int
	for _,v := range a{
		sum_ret += v
	}
	return sum_ret
}