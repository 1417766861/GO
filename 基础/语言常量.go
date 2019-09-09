package main

import (
	"fmt"
)

func main() {
	//显式类型定义：
	//const a string = "abc"
	//隐式类型定义：
	//const b = "abc"
	//fmt.Println(a, b)

	//常量可以用len(), cap(), unsafe.Sizeof()函数计算表达式的值
	//c := "abc"
	//d := len(a)
	//e := unsafe.Sizeof(a)
	//println(c, d, e)

	//iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，
	// const 中每新增一行常量声明将使
	// iota 计数一次(iota 可理解为 const 语句块中的行索引)

	const (
		a = iota   //0
		b          //1
		c          //2
		d = "ha"   //独立值，iota += 1
		e          //"ha"   iota += 1
		f = 101    //iota +=1
		g          //100  iota +=1
		h = iota   //7,恢复计数
		i          //8
	)
	fmt.Println(a,b,c,d,e,f,g,h,i)
}
