package main

import(
	"fmt"
	"strings"
)

func main(){
	var str string = "hello,how is it going,hugo?"
	var manyG string = "ggggggggggggg"

	fmt.Printf("number of h's in \"%s\" is:",str)
	fmt.Printf("%d\n",strings.Count(str,"h"))

	fmt.Printf("number of double g's in \"%s\" is:",manyG)
	fmt.Printf("%d\n",strings.Count(manyG,"gg"))
}