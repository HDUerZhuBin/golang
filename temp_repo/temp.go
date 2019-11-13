package main
import(
	"fmt"
)

func main(){
	f()
}

func a(){
	i := 0
	defer fmt.Println(i)
	i++
	return
}

func f(){
	for i:=0;i<5;i++{
		defer fmt.Println("current i:",i)
	}
}