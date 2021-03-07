package proxy

type Subject interface {
	Do() string
}

type RealSubject struct{}

func (RealSubject) Do() string {
	return "real"
}

type Proxy struct {
	real RealSubject
}

//感觉和外观模式有点像,只不过外观模式是直接调用目标方法
//而代理模式是需要做一些自己的工作
func (p Proxy) Do() string {
	res := ""

	res += "pre:"

	res += p.real.Do()

	res += ":after"

	return res
}
