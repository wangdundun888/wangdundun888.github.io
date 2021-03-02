package simpleFatory

import (
	"testing"
)

func TestNewAPI(t *testing.T) {
	a := NewAPI(1)
	a.Say("tom")
	b := NewAPI(2)
	b.Say("tom")
}
