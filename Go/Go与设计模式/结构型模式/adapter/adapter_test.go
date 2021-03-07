package adapter

import (
	"fmt"
	"testing"
)

func TestAdapteeImpl_SpecificRequest(t *testing.T) {
	adaptee := NewAdaptee()
	adapter := NewAdapter(adaptee)
	fmt.Println(adapter.Request())
}
