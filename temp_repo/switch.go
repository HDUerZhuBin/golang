package main
import (
	"fmt"
)

func main(){
	var num1 int = 99


	switch num1{
	case 98,99:
		fmt.Println("equal to 98 or 99")
	case 100:
		fmt.Println("equal to 100")
	default:
		fmt.Println("it's not equal to 98~100 in integer")
	}
}