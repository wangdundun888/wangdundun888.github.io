package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	data := []byte("hello world!")
	//data.txt出现在wangdundun888.github.io的目录下
	//通过阅读源码可知,如果文件存在,则截断文件但不改变文件权限
	//否认则创建文件并按照参数perm写入权限
	//那么问题来了,第一,可否通过绝对路径或者相对路径创建文件
	//像 name = "D://data.txt"的形式
	//第二,在不读取文件的前提下,如果文件存在,能否保留原内容
	err := ioutil.WriteFile("D://data.txt", data, 0644)
	if err != nil {
		panic(err)
	}
	read, err := ioutil.ReadFile("data.txt")
	if err == nil {
		fmt.Println(string(read))
	}

	//同样是出现在wangdundun888.github.io的目录下
	file1, _ := os.Create("data1.txt")
	defer file1.Close()

	bytes, _ := file1.Write(data)
	fmt.Printf("wrote %d bytes to data1.txt\n", bytes)

	file2, _ := os.Open("data1.txt")
	defer file2.Close()

	read2 := make([]byte, len(data))
	bytes, _ = file2.Read(read2)
	fmt.Printf("read %d bytes from data1.txt\n", bytes)
	fmt.Println(string(read2))
}
