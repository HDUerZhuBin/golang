package main

import (
	"fmt"
	"strings"
)

func main(){
	str := "the quick brown fox jumps over the lazy dog"
	sl := strings.Fields(str)

	fmt.Println("the filed's ret:",sl)

	len := len(sl)
	fmt.Println("the len of sl:",len)

	for key,val:=range sl{
		fmt.Println("key:,val:",key,val)
	}

	str2 := "go1|the abc of go|25"
	sl2 := strings.Split(str2,"|")
	for _,value := range sl2{
		fmt.Println("current element:",value)
	}


	str3 := strings.Join(sl2,"***")
	fmt.Println("join-ret:%s",str3)

}