package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	str := "hello"
	data := []byte(str)
	fmt.Println(str)
	MD5 := md5.Sum(data)
	fmt.Printf("%s\n", data)
	fmt.Printf("%v", MD5)
}
