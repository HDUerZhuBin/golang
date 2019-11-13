package main

import(
	"fmt"
)

func main(){
	b()
}

func trace(s string){
	fmt.Println("entering:",s)
}
func untrace(s string){
	fmt.Println("leaving:",s)
}

func a(){
	trace("a")
	defer untrace("a")
	fmt.Println("in a")
	fmt.Println("before return")
	return  
}

func b(){
	trace("b")
	defer untrace("b")
	fmt.Println("in b")
	fmt.Println("before return")
	a()
	return 
}