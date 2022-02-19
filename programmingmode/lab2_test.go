package programmingmode

import (
	"fmt"
	"testing"
)

func add(i int) func() int {
	return func() int {
		i = i + 1
		return i
	}
}

func TestErrHandle(t *testing.T) {
	fn := add(1)
	fmt.Println(fn())
	fmt.Println(fn())
	fmt.Println(fn())
}
