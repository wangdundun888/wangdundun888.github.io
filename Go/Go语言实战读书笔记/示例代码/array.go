package main

import "fmt"

func main() {
	//声明与初始化
	//方式1:var
	var a [3]int
	fmt.Println(a)
	//方式2:字面量
	a = [3]int{10, 20, 30}         //使用字面量声明,{}里元素的个数不可超过声明的长度
	a1 := [...]int{10, 20, 30, 40} //可以对长度使用省略号,让编译器根据{}元素个数确认长度
	fmt.Println(a1)
	a2 := [...]int{2: 2} //使用索引声明
	fmt.Println(a2)
	//访问与复制
	a[0] = 100
	fmt.Println(a)
	fmt.Println(a[1:3], a[:3], a[2:], a[2])

	if a == a2 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
	a = a2 //当类型与长度相同时可以直接赋值
	//a = a1[0:3]  是错误的,a1[0:3]返回的是[]int ,切片类型
	if a == a2 { //可以直接比较两个数组是否相等,索引及索引下边对应的值
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
	//迭代数组
	//方式1:下标访问
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i], " ")
	}
	//方式2:range
	for index, value := range a {
		fmt.Println(index, " ", value)
	}
}
