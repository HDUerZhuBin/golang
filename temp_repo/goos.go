package main

import(
	"fmt"
	"runtime"
	"os"
)

func main(){
	var goos string = runtime.GOOS 
	fmt.Println("the operation system is: ",goos)
	path := os.Getenv("PATH")
	fmt.Println("Path is :",path)
}