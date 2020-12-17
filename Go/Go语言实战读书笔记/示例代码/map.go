package main

import "fmt"

func main() {
	//1.创建映射的两种方式
	//方式1:make
	m := make(map[string]int)
	m["bob"] = 18
	fmt.Println("bob: ", m["bob"])
	//方式2:字面量 -> 常用
	m1 := map[string]int{"Alice": 15, "Bob": 16}
	fmt.Println("Bob: ", m1["Bob"])
	/*
		映射的键可以是任何值,注意是 值 .这个值可以是内置类型,也可以是结构类型,只要这个值可以使用==运算符做比较

		切片,函数以及包含切片的结构类型的这些类型由于具有引用语义,不能作为映射的键,会发生编译错误
		m := make(map[[]string]int 是错误的
		但切片可以作为映射的值
		思考:为什么引用语义的类型不能作为映射的键?
			猜测是因为引用类型可能会被回收然后导致空键问题
	*/

	//2.使用映射
	//2.1赋值
	m["Tom"] = 19
	fmt.Println("Tom: ", m["Tom"])
	m["Tom"] = 20
	fmt.Println("Tom: ", m["Tom"])
	//2.2获取映射值的两种方式
	//2.2.1
	value, exists := m["bob"]
	if exists {
		fmt.Println("bob: ", value)
	}
	//2.2.2
	value = m["Tony"]
	//判断这个值是否为零值来确定键是否存在
	if value != 0 {
		fmt.Println("Tony :", value)
	}
	//但如果这个值本身就为零值呢?
	m["Tony"] = 0
	value = m["Tony"]
	if value != 0 {
		fmt.Println("Tony :", value)
	}
	//很明显,根本不会显示.那么,问题来了,其他类型可否用相应的零值来做映射值
	m2 := make(map[int]string)
	m2[1] = "bob"
	m2[2] = ""
	fmt.Println("1 :", m2[1])
	fmt.Println("2 :", m2[2])
	value1 := m2[2]
	if value1 != "" {
		fmt.Println("2 :", m2[2])
		//这里永远不会执行,因为value1本身值就为空值
	}
	m3 := map[int][]int{}
	//m3 := make(map[int][]int)  为什么这种方式定义一个slice类型值会发生编译错误,是无法解析吗
	m3[1] = []int{1, 2, 3}
	m3[2] = nil
	fmt.Println("1 :", m3[1])
	fmt.Println("2 :", m3[2])
	//由此看来,其他类型也可以用相应的零值来做映射值
	//使用range迭代映射
	for key, value := range m {
		fmt.Println(key, " ", value)
	}
	//删除映射的键值对
	delete(m, "bob")
	delete(m, "noExist") //删除一个不存在的键并不会发生异常
	//func(m map[int]string)
	//在函数里传递映射就和传递切片一样,只是传递一个引用值,并不是复制整个映射,所以在函数里对映射操作,原映射也会同步
}
