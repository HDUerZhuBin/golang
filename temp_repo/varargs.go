package main
import "fmt"

func main(){
	slice := []string{"dtie","hie","jkii"}
	print1(slice...)
	fmt.Println("*********")
	print2(slice)
}

func print1(s ...string){
	for _,str := range s{
		fmt.Println("the str ele:",str)
	}
}

func print2(s []string){
	len_s := len(s)
	for i:=0;i<len_s;i++{
		fmt.Println("in method2,the ele is:%s",s[i])
	}
}