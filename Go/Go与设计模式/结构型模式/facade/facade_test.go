package facade

import "testing"

func TestNewSwitch(t *testing.T) {
	s := NewSwitch()
	for _, e := range s.electrical {
		e.Run()
	}
}
