package main
import "fmt"

func main(){
	var str string = "go is a beautiful language!"
	len_str := len(str)
	fmt.Println("the len of the str is:",len_str)

	for idx,val := range str{
		fmt.Println("the idx is:",idx," and the value is:",val)
	}

	str2 := "vivo,照亮你的美"
	len_str2 := len(str2)
	fmt.Println("the len of str2 is:",len_str2)
	for key,value := range str2{
		fmt.Println("the key is:",key," and the value is:",value)
	}
	fmt.Println("********************")
	fmt.Println("super sb")
}