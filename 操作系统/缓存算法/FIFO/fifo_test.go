package FIFO

import (
	"container/list"
	"fmt"
	"strconv"
	"testing"
)

var mfc *myFifoCache

func TestMyFifoCache_Set(t *testing.T) {
	mfc = NewMFC(3)
	for i := 1; i < 5; i++ {
		mfc.Set(strconv.Itoa(i), strconv.Itoa(i))
		fmt.Print("now fifoCache : ")
		printList(mfc.Data)
	}
	mfc.Set(strconv.Itoa(3), strconv.Itoa(1))
	fmt.Print("now fifoCache : ")
	printList(mfc.Data)
}

func TestMyFifoCache_Get(t *testing.T) {
	for i := 1; i < 5; i++ {
		fmt.Println("get key : ", i, " value :", mfc.Get(strconv.Itoa(i)))
	}
}

func printList(list *list.List) {
	l := list
	for e := l.Front(); e != nil; e = e.Next() {
		pair := e.Value.(Pair)
		fmt.Print("(", pair.key, "-", pair.value, ")")
	}
	fmt.Println()
}
