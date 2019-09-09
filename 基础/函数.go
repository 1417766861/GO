package main

import (
	"fmt"
)

func main() {
	//函数声明告诉了编译器函数的名称，返回类型，和参数。

	//函数定义
	//Go 语言函数定义格式如下：
	//func function_name( [parameter list] ) [return_types] {
	//函数体
	//}
	//parameter list：参数列表，参数就像一个占位符，当函数被调用时，你可以将值传递给参数，这个值被称为实际参数。
	// 参数列表指定的是参数类型、顺序、及参数个数。参数是可选的，也就是说函数也可以不包含参数。
	//return_types：返回类型，函数返回一列值。return_types 是该列值的数据类型。有些功能不需要返回值，
	// 这种情况下 return_types 不是必须的

	//函数返回多个值
	//Go 函数可以返回多个值，例如：
	//
	//实例
	//package main
	//
	//import "fmt"
	//
	//func swap(x, y string) (string, string) {
	//	return y, x
	//}
	//
	//func main() {
	//	a, b := swap("Google", "Runoob")
	//	fmt.Println(a, b)
	//}

	///* 定义局部变量 */
	//var a int = 100
	//var b int = 200
	//var ret int
	//
	///* 调用函数并返回最大值 */
	//ret = max(a, b)
	//
	//fmt.Printf( "最大值是 : %d\n", ret )


	var a int = 100
	var b int= 200

	fmt.Printf("交换前，a 的值 : %d\n", a )
	fmt.Printf("交换前地址，a  : %d\n",  &a )
	fmt.Printf("交换前地址，b  : %d\n",  &b )
	fmt.Printf("交换前，b 的值 : %d\n", b )

	/* 调用 swap() 函数
	 * &a 指向 a 指针，a 变量的地址
	 * &b 指向 b 指针，b 变量的地址
	 */
	swap(&a, &b)

	fmt.Printf("交换后，a 的值 : %d\n", a )
	fmt.Printf("交换后，b 的值 : %d\n", b )
	fmt.Printf("交换前地址，a  : %d\n",  &a )
	fmt.Printf("交换前地址，b  : %d\n",  &b )
	//引用传递是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。

	//语言函数作为实参
	/* 声明函数变量 */
	//getSquareRoot := func(x float64) float64 {
	//	return math.Sqrt(x)
	//}
	///* 使用函数 */
	//fmt.Println(getSquareRoot(9))


	//// 声明一个函数类型
	//type cb func(int) int
	//
	//func main() {
	//	testCallBack(1, callBack)
	//	testCallBack(2, func(x int) int {
	//		fmt.Printf("我是回调，x：%d\n", x)
	//		return x
	//	})
	//}
	//
	//func testCallBack(x int, f cb) {
	//	f(x)
	//}
	//
	//func callBack(x int) int {
	//	fmt.Printf("我是回调，x：%d\n", x)
	//	return x
	//}
}

func max(num1, num2 int) int {
	/* 声明局部变量 */
	var result int

	if (num1 > num2) {
		result = num1
	} else {
		result = num2
	}
	return result
}


func swap(x *int, y *int) {
	var temp int
	temp = *x    /* 保存 x 地址上的值 */
	*x = *y      /* 将 y 值赋给 x */
	*y = temp    /* 将 temp 值赋给 y */
}