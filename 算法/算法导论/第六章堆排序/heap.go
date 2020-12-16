package main

/*
	二叉堆是一个数组,被看成一个近似的完全二叉树,除最底层外,该树是完全满的,而且是从左到右填充
	表示堆的数组A包括两个属性,A.length表示数组元素的个数,A.heap-size表示有多少个堆元素存储在数组中
	A[1..A.length](此处意指下标从1开始)可能存有数据,但只有A[1..A.heap-size]为堆的有效数据,且0<=A.heap-size<=A.length,
	A[1]为树的根节点.

	对于任意给定的一个结点的下标i,
		父结点:	parent(i) return i/2 //(向下取整)
		左孩子结点: left(i) return i*2
		右孩子结点: right(i) return i*2+1

	最大堆: 对于所有的结点i,均满足 A[parent(i)]>=A[i]
	最小堆: 对于所有的结点i,均满足 A[parent(i)]<=A[i]

*/
import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func getRandomSlice(size int) []int {
	if size < 0 || size > math.MaxInt32/2 {
		return nil
	}
	s := make([]int, size)
	rand.Seed(time.Now().UnixNano())
	for i, _ := range s {
		s[i] = rand.Intn(1000)
	}
	return s
}

var heapsize int

func parent(i int) int { return i / 2 }
func left(i int) int   { return i * 2 }
func right(i int) int  { return i*2 + 1 }

//维护最大堆性质
//假定根节点为left(i)和right(i)的二叉树都是最大堆,此为调用该函数的前提
//但A[i]可能小于其孩子,违反最大堆性质,故调用该函数维护堆
//维护之后,A[i]下沉到左子树或右子树,可能会继续违反最大堆性质,对其递归调用本函数
//时间复杂度为O(lgn),对于一个高度为h的结点来说,时间复杂度为O(h)
func maxHeapify(A []int, i int) {
	l := left(i)
	r := right(i)
	largest := 0
	if l <= heapsize && A[l] > A[i] {
		largest = l
	} else {
		largest = i
	}
	if r <= heapsize && A[r] > A[largest] {
		largest = r
	}
	if largest != i {
		A[i], A[largest] = A[largest], A[i]
		maxHeapify(A, largest)
	}

}

//创建堆
//调用maxHeapify
//从len(A)/2开始,都是叶子结点,本身就是一个最大堆,符合maxHeapify的调用条件
//时间复杂度为O(n)
func buildHeap(A []int) {
	heapsize = len(A) - 1
	for i := len(A) / 2; i >= 1; i-- {
		maxHeapify(A, i)
	}
}

//最大元素总在根节点A[1]处,把它与A[n]交换,然后A[n]是已经排好序的,通过heapsize减1屏蔽A[n]
//此时,除了根节点,依然是最大堆,调用maxHeapify(A,1)维护性质
//循环以上步骤,完成堆排序
//原址排序,时间复杂度为O(nlgn)
func sortHeap(A []int) {
	buildHeap(A)
	for i := len(A) - 1; i >= 2; i-- {
		A[1], A[i] = A[i], A[1]
		heapsize -= 1
		maxHeapify(A, 1)
	}
}

func main() {
	A := getRandomSlice(10)
	fmt.Println(A)
	sortHeap(A)
	fmt.Println(A)
}
