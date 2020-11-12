package main

import "fmt"

func main() {
	//创建映射的两种方式
	//方式1:make
	m := make(map[string]int)
	m["bob"] = 18
	fmt.Println(m["bob"])
	//方式2:字面量 -> 常用
	m1 := map[string]int{"Alice":15,"Bob":16}
	fmt.Println(m1["Bob"])
	/*
	映射的键可以是任何值,注意是 值 .这个值可以是内置类型,也可以是结构类型,只要这个值可以使用==运算符做比较

	切片,函数以及包含切片的结构类型的这些类型由于具有引用语义,不能作为映射的键,
	m := make(map[[]string]int 是错误的
	但切片可以作为映射的值
	思考:为什么引用语义的类型不能作为映射的键?
		猜测是因为引用类型可能会被回收然后导致空键问题
	 */

}