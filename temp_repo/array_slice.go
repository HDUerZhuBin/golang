package main
import "fmt"


func main(){
	var arr1 [6]int
	var slice []int = arr1[2:5]

	for i:=0;i<len(arr1);i++{
		arr1[i] = i
	}
	for i:=0;i<len(slice);i++{
		fmt.Println("slice at %d is %d",i,slice[i])
	}
	
	fmt.Println("the len of the arr1:",len(arr1))
	fmt.Println("the len of the slice:",len(slice))
	fmt.Println("the capacity of the sllice:",cap(slice))

	slice = slice[:4]
	for i:=0;i<len(slice);i++{
		fmt.Println("slice at %d is %d",i,slice[i])
	}

	fmt.Println("the len of slice:",len(slice))
	fmt.Println("the capacity of slice:",cap(slice))
	
}