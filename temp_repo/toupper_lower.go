package main

import(
	"fmt"
	"strings"
)

func main(){
	var origin string = "Hey, how are you George?"

	var lower , upper string

	fmt.Printf("the original string is:%s\n",origin)
	lower = strings.ToLower(origin)
	upper = strings.ToUpper(origin)
	fmt.Printf("the lowercase string is:%s\n",lower)
	fmt.Printf("the uppercase string is:%s\n",upper)
}