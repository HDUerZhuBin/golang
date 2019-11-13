package main

import (
	"fmt"
)

func main(){
	var str string = "go is a beautiful language!"
	fmt.Println("the len of str is:",len(str))
	for ix:=0;ix<len(str);ix++{
		fmt.Println("character on position ",ix," is: ",str[ix])
	}
	str2 := "英文语言"
	fmt.Println("the len of the str2 is :",len(str2))
	for i:=0;i<len(str2);i++{
		fmt.Println("current ",i," element is:",str2[i])
	}
}