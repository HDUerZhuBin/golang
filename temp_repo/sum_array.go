package main
import "fmt"

func main(){
	arr := [...]int{3,6,8}
	sum := Sum(&arr)
	fmt.Println("sum:",sum)
}

func Sum(a *[3]int) int{
	sum := 0
	for _,value := range a{
		sum += value
	}

	return sum
}