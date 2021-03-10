package LFU

import (
	"fmt"
	"strconv"
	"testing"
)

var mlc *myLFUCache = NewMLC(3)

func TestMyLFUCache_Set(t *testing.T) {
	for i := 1; i < 5; i++ {
		mlc.Set(strconv.Itoa(i), strconv.Itoa(i))
		fmt.Print("now lfuCache : ")
		printKV(mlc)
	}
}

func TestMyLFUCache_Get(t *testing.T) {
	for i := 1; i < 4; i++ {
		value := mlc.Get(strconv.Itoa(i))
		fmt.Print("now lfuCache : ")
		printKV(mlc)
		fmt.Println("get value, key :", i, "value :", value)
	}
	mlc.Set("5", "5")
	fmt.Println("now put k-v :5-5")
	fmt.Print("now lfuCache : ")
	printKV(mlc)
}

func printKV(mlc *myLFUCache) {
	for k, index := range mlc.KeyMap {
		fmt.Print("( ", k, ",", mlc.Data.value[index], " )-", mlc.Data.cnt[index])
	}
	defer fmt.Println()
}
