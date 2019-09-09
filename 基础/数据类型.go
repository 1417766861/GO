package main

import (
	"fmt"
)

func main(){
	// package 基础
	// bool 类型
	var a bool = true

	//	数字类型
	//  整型 int 和浮点型 float32、float64 、复数
	var b int = 10
	var c float32 = 10000
	var d float64 = 999
	var e float32 = 3.14
	// 字符串
	var f string = "123"
	fmt.Println(a,b,c,d,e,f)


	//派生类型:
	//	包括：
	//	(a) 指针类型（Pointer）
	//Go 语言的取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址。
	fmt.Printf("变量的地址: %x\n", &a  )

	var g int = 20
	var ip *int
	ip = &g
	fmt.Printf("a 变量的地址是: %x\n", &g  )
	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip )
	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip )
	//空指针 值为：nil
	//if(ptr != nil)     /* ptr 不是空指针 */
	//if(ptr == nil)    /* ptr 是空指针 */


	//	(b) 数组类型
	//	(c) 结构化类型(struct)
	//	(d) Channel 类型
	//	(e) 函数类型
	//	(f) 切片类型
	//	(g) 接口类型（interface）
	//	(h) Map 类型

}



