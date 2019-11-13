package main
import(
	"fmt"
)

func main(){
	fmt.Println("multiply 2 * 5 * 6 = ",Multiply3Nums(2,5,6))
}

func Multiply3Nums(a,b,c int)int{
	return a * b * c
}