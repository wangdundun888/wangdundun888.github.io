package prototype

import "testing"

type Type1 struct {
	name string
}

func (t *Type1) Clone() Cloneable {
	tc := *t
	return &tc
}

type Type2 struct {
	name string
}

func (t *Type2) Clone() Cloneable {
	tc := *t
	return &tc
}

var p *PrototypeManager

func init() {
	p = NewPrototypeManager()
	t1 := &Type1{
		name: "type1",
	}
	p.Set("t1", t1)
}

func TestClone(t *testing.T) {
	t1 := p.Get("t1")
	t2 := t1.Clone()
	if t1 == t2 {
		t.Fatal("clone fail")
	}
}

func TestCloneFromManager(t *testing.T) {
	c := p.Get("t1").Clone()

	t1 := c.(*Type1)
	if t1.name != "type1" {
		t.Fatal("error")
	}

}
