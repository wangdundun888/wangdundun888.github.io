package observer

import "fmt"

type Host struct {
	customer []Person
	thing    string
	price    int
}

func NewHost() *Host {
	return &Host{
		customer: make([]Person, 0),
	}
}

func (h *Host) attach(p Person) {
	h.customer = append(h.customer, p)
}

func (h *Host) notifyAll() {
	for _, c := range h.customer {
		c.update(h)
	}
}

func (h *Host) begin(thing string, price int) {
	h.thing = thing
	h.price = price
	fmt.Println("Sale begin...")
	h.notifyAll()
}

type Person interface {
	update(*Host)
}

type customer struct {
	name string
}

func NewCustomer(name string) *customer {
	return &customer{name}
}

func (c *customer) update(h *Host) {
	fmt.Printf("%s knows %s price : %d\n", c.name, h.thing, h.price)
}
