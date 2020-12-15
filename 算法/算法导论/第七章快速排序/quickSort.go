package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	快速排序用了分治思想,对于一个数组A[p..r],进行一下三步排序:
	第一步:
		分解,数组A[p..r]被划分为两个(可能为空)的子数组A[p,q-1]和A[q+1,r],
		使得子数组A[q..p-1]每一个元素都小于A[q],A[p+1..r]都大于A[q]
	第二步:
		解决,通过递归调用快速排序,对子数组的子数组A[p,q-1]和A[q+1,r]进行排序
	第三步:
		合并,由于是原址排序,此时A[p..r]已经有序

	循环不变量: 在循环中不变的量,可以为证明正确性提供有用的性质.
			  需要证明在"初始化-过程-终止"三个过程中,定好的量的不变性
	循环不变量可为循环提供一个证明循环正确的有力的依据.
	以快排的partition函数中的从尾选划分值的循环为例:
		x := A[r]  //从尾选划分值
		i := p-1
		for j:=p;j<r;j++{
		if A[j] <= x {
			i = i + 1
			A[j],A[i] = A[i],A[j]
		}
	}
		对于任意下标k,不变量有:
			1.若 p <= k <= i,则A[k] <= x
			2.若 i+1 <= k <= j-1,则A[k] > x
			3.如 k = r,则 A[k] = x
			证明:
				1.初始化
					开始时,i = p-1,j = p,所以p与i,i+1与j-1之间都不存在值,所以不变量1,2满足
				2.过程中保持
					j在向前行走的过程中,有A[j] <= x和A[j] > x两种划分,先来讨论A[j] > x,此时j值加1,此时A[j-1]依然满足不变量2,
					且其他也不变.当A[j] <= x,i值加1,此时A[i] > x,交换A[i]和A[j],则满足不变量1,然后j值加1,此时A[j-1] > x,满足不变量2,
					所以过程中总能保持不变量1 2
				3.终止
					当j=r时,循环终止,A[p..r]被划分为3个子集,一个是所有元素小于等于x,另一个是所有元素大于x,最后一个是只有一个元素,即x
			最后交换A[i+1]与A[r],因为A[i+1] > A[r],所以交换后,在A[i+1]的左边全是小于等于它的元素,右边全是大于它的元素,而它本身是已经被排序
			好的,然后返回它的下标i+1,继续递归调用排序左右两边的元素.



*/
//获得一个有n个元素的切片,值小于1000
func getRandomSlice(size int) []int {
	if size < 0 {
		return nil
	}
	s := make([]int, size)
	rand.Seed(time.Now().UnixNano())
	for i, _ := range s {
		s[i] = rand.Intn(1000)
	}
	return s
}

func quickSort(A []int, p, r int) {
	if p < r {
		q := partition(A, p, r)
		quickSort(A, p, q-1)
		quickSort(A, p+1, r)
	}
}
func partition(A []int, p, r int) int {
	//x := A[r]  //从尾选划分值
	//i := p-1
	//for j:=p;j<r;j++{
	//	if A[j] <= x {
	//		i = i + 1
	//		A[j],A[i] = A[i],A[j]
	//	}
	//}
	//A[i+1],A[r] = A[r],A[i+1]
	//return i+1

	//注意两者之间细微的区别

	x := A[p] //从头开始选划分值
	i := p
	for j := p + 1; j < r+1; j++ {
		if A[j] <= x {
			i = i + 1
			A[j], A[i] = A[i], A[j]
		}
	}
	A[i], A[p] = A[p], A[i]
	return i
}
func main() {
	s := getRandomSlice(10)
	s1 := getRandomSlice(10)
	fmt.Println(s, "\n", s1)
	quickSort(s, 0, len(s)-1)
	//冒泡排序
	for i := 0; i < len(s1); i++ {
		for j := len(s1) - 1; j > i; j-- {
			if s1[j] < s1[j-1] {
				s1[j], s1[j-1] = s1[j-1], s1[j]
			}
		}
	}
	fmt.Println(s, "\n", s1)
}
