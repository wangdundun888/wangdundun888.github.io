package observer

import "testing"

func TestObserver(t *testing.T) {
	h := NewHost()
	thing := "月光宝盒"
	price := 99
	c1 := NewCustomer("张三")
	c2 := NewCustomer("李四")
	c3 := NewCustomer("李五")
	h.attach(c1)
	h.attach(c2)
	h.attach(c3)
	h.begin(thing, price)
}
