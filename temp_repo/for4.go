package main

import(
	// "fmt"
)

func main(){
	for i:=0;i<3;i++{
		for j:=0;i<10;j++{
			if j>5{
				break
			}
			print(j)
		}
		print("       ");
	}
}