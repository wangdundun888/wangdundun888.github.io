package main

import "fmt"

//内存卡放入读卡器插入电脑读取内容

type cache struct {
	content string
}

func (c *cache) getContent() string {
	return c.content
}

type readUsb struct {
	cache
}

type computer struct{}

func (c *computer) read(u readUsb) {
	fmt.Println(u.content)
}

func main() {
	c := cache{"内存卡的内容"}
	ru := readUsb{c}
	com := &computer{}
	com.read(ru)
}
