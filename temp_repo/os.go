package main

import (
	"runtime"
	"fmt"
)

func main(){
	goos := runtime.GOOS 
	if "linux"==goos{
		fmt.Println("the os is linux")
	}else{
		fmt.Println("the os is windows")
	}
}