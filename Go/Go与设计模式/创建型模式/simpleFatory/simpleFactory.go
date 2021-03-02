package simpleFatory

import "fmt"

type API interface {
	Say(name string)
}

//简单工厂模式
func NewAPI(t int) API {
	if t == 1 {
		return &hi{}
	} else if t == 2 {
		return &hello{}
	}
	return nil
}

type hi struct{}

type hello struct{}

func (h *hi) Say(name string) {
	fmt.Println("hi," + name)
}

func (h *hello) Say(name string) {
	fmt.Println("hello," + name)
}
