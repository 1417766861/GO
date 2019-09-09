package main

//如果忽略 [] 中的数字不设置数组大小，Go 语言会根据元素的个数来设置数组的大小：
//
//var balance = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
//该实例与上面的实例是一样的，虽然没有设置数组的大小。
//
//balance[4] = 50.0
import "fmt"

func main() {
	var n [10]int /* n 是一个长度为 10 的数组 */
	var i,j int

	/* 为数组 n 初始化元素 */
	for i = 0; i < 10; i++ {
		n[i] = i + 100 /* 设置元素为 i + 100 */
	}

	/* 输出每个数组元素的值 */
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j] )
	}
	//b :=n[9]
	var b int = n[9]
	fmt.Printf( "%d", b)
}