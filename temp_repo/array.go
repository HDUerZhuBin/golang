package main
import (
	"fmt"
)

func main(){
	arrKeyValue := []string{3: "chris" ,14:"ron"}
	len_arr := len(arrKeyValue)
	fmt.Println("len_arr:",len_arr)
	for key,val := range arrKeyValue{
		fmt.Println(key,":",val)
	}
}