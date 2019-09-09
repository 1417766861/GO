package main
// package main表示一个可独立执行的程序，每个 Go 应用程序都包含一个名为 main 的包
import "fmt"
// fmt 包实现了格式化 IO（输入/输出）的函数。

func main(){

	// 声明一个变量并初始化
	var a = "RUNOOB"
	fmt.Println(a)

	// 没有初始化就为零值
	var b int
	fmt.Println(b)

	// bool 零值为 false
	var c bool
	fmt.Println(c)
	fmt.Println("hello donghao")



	_,numb,strs := numbers() //只获取函数返回值的后两个
	fmt.Println(numb,strs)

	//一个可以返回多个值的函数

}
func numbers()(int,int,string){
	a , b , c := 1 , 2 , "str"
	return a,b,c
}