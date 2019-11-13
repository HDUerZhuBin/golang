package main
import (
	"fmt"
)

func main(){
	slice := []int{1,5,7,8,2,4,6,8,5}
	min_ret := min(slice...)
	fmt.Println("the min_ret:",min_ret)

}

func min(s ...int)(int){
	len_params := len(s)
	if len_params == 0{
		return 0
	}

	min := s[0]
	for _,val := range s{
		if val < min{
			min = val
		}
	}
	return min
}