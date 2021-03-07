package adapter

type Target interface {
	Request() string
}

type Adaptee interface {
	SpecificRequest() string
}

type adapteeImpl struct{}

func NewAdaptee() Adaptee {
	return &adapteeImpl{}
}

func (*adapteeImpl) SpecificRequest() string {
	return "adaptee method"
}

type adapter struct {
	Adaptee
}

func NewAdapter(adaptee Adaptee) Target {
	return &adapter{
		Adaptee: adaptee,
	}
}
func (a *adapter) Request() string {
	return a.SpecificRequest()
}
