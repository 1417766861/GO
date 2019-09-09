package main

import "fmt"
//动态数组

//使用make()函数来创建切片:
//var slice1 []type = make([]type, len)
//也可以简写为
//slice1 := make([]type, len)

//也可以指定容量，其中capacity为可选参数。
//make([]T, length, capacity)

func main() {
	s := []string{"1", "2", "3", "4", "5", "6", "7"}
	b := s[2:4]
	fmt.Printf("s is：%s\n", s)
	fmt.Printf("b is：%s\n", b)
	//将arr中从下标startIndex到endIndex-1 下的元素创建为一个新的切片
	c := s[3:]
	fmt.Printf("c is：%s\n", c)
	d := s[:len(s)-1]
	fmt.Printf("d is：%s\n", d)

	//e :=make([]int,1,2)
	//fmt.Printf("e is：%s\n", e)

	//cap() 可以测量切片最长可以达到多少。

	numbers := make([]int,3,5)

	printSlice(numbers)


	//空(nil)切片
	//一个切片在未初始化之前默认为 nil，长度为 0，实例如下：
	//func main() {
	//	var numbers []int
	//	printSlice(numbers)
	//	if(numbers == nil){
	//		fmt.Printf("切片是空的")
	//	}
	//}
	//func printSlice(x []int){
	//	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}

func printSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}