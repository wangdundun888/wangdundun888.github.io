package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}

	var b [5]int

	fmt.Println(a, "\n", b)

	b = a

	fmt.Println(a, "\n", b)

	c := a

	fmt.Println(c)

	a[0] = 2
	b[0] = 3
	c[0] = 4
	//根据实验可知,数组赋值为复制赋值
	fmt.Println(":::", a, "\n", b, "\n", c)

}
