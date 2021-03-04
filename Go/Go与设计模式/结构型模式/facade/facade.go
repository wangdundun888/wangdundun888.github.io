package facade

import "fmt"

type Switch struct {
	electrical []Electrical
}

func NewSwitch() Switch {
	e := Switch{
		electrical: make([]Electrical, 0),
	}
	e.electrical = append(e.electrical, &TV{})

	e.electrical = append(e.electrical, &Light{})

	return e
}

type Electrical interface {
	Run()
}

type TV struct{}

func (tv *TV) Run() {
	fmt.Println("TV is running")
}

type Light struct{}

func (l *Light) Run() {
	fmt.Println("Light is running")
}
