package lib

import (
	"fmt"
)

func main() {
	m, err := NewConcurrentMap(5, newDefaultPairRedistributor(0.2, 2))

	if err != nil {
		panic(struct{}{})
	}

	m.Put("a", "abc")

	fmt.Println(m.Len())

	fmt.Println(m.Get("b"))

	fmt.Println(m.Get("a"))

}
