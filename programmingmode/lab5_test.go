package programmingmode

import (
	"testing"
)

func MapStrToStr(arr []string, fn func(string) string) []string {
	var newArray []string
	for _, it := range arr {
		newArray = append(newArray, fn(it))
	}
	return newArray
}

func MapStrToInt(arr []string, fn func(s string) int) []int {
	var newArray []int
	for _, it := range arr {
		newArray = append(newArray, fn(it))
	}
	return newArray
}

func TestGeneric(t *testing.T) {
}
