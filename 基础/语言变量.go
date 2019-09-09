package main
import "fmt"
func main() {
	var a string = "Runoob"
	fmt.Println(a)

	var b, c int = 1, 2 // 次声明多个变量
	fmt.Println(b, c)


	//变量声明
	//1. 指定变量类型，如果没有初始化，则变量默认为零值。
	//var v_name v_type
	//v_name = value

	//数值类型（包括complex64/128）为 0
	//布尔类型为 false
	//字符串为 ""（空字符串）
	//以下几种类型为 nil：
	//var a *int
	//var a []int
	//var a map[string] int
	//var a chan int
	//var a func(string) int
	//var a error // error 是接口
	var i int
	var f float64
	var q bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, q, s) // 0 0 false ""

	//第二种，根据值自行判定变量类型。
	var username = "donghao"
	fmt.Printf(username)

	//第三种，省略 var, 注意 := 左侧如果没有声明新的变量，就产生编译错误，格式：
	age:=20
	fmt.Printf("%d", age)


	// 这种因式分解关键字的写法一般用于声明全局变量
	//var (
	//	vname1 v_type1
	//	vname2 v_type2
	//)
}