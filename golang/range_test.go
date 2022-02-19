package main

import (
	"fmt"
	"testing"
)

func TestRange(t *testing.T) {
	arr := []int{1, 2, 3}
	arrCopy := make([]int, len(arr))
	for i, v := range arr {
		arrCopy[i] = v
	}

	fmt.Println(arrCopy)
}
