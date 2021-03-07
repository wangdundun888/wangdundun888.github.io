package proxy

import "testing"

func TestProxy_Do(t *testing.T) {
	p := &Proxy{
		real: RealSubject{},
	}
	result := p.Do()

	if result != "pre:real:after" {
		t.Fatal("expext pre:real:after ,no ", result)
	}
}
