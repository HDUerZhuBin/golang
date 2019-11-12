package main 

import(
	"fmt"
	"strings"
)

func main(){
	var str string = "this is just demo of golang!"
	var split string = " "

	var split_ret []string = strings.Split(str,split)
	var len_split int = len(split_ret)
	fmt.Printf("the len of split_ret:%d\n",len_split)
	
	for i:=0;i<len_split;i++{
		fmt.Println("current_element:",split_ret[i])
	}

}