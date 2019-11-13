package main

import (
	"fmt"
	"strings"
)

func main(){
	addBmp := MakeAddSuffix(".bmp")
	addJpeg := MakeAddSuffix(".jpeg")

	bmp_ret := addBmp("file")
	jpeg_ret := addJpeg("file")

	fmt.Println("bmp_ret:",bmp_ret)
	fmt.Println("jpeg_ret:",jpeg_ret)
}

// 函数名为MakeAddSuffix,传入参数为string类型的suffix，返回值为匿名函数（可引用地址变量），该匿名函数的传入参数是string类型，返回值类型为string
func MakeAddSuffix(suffix string) func(string) string{
	// 返回匿名函数，并对匿名函数进行定义
	return func(name string) string{
		if !strings.HasSuffix(name,suffix){
			return name + suffix
		}
		return name
	}
}