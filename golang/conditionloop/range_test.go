package conditionloop

import (
	"fmt"
	"testing"
)

type Foo struct {
	str string
}

func TestRange1(t *testing.T) {
	list := []Foo{{"a"}, {"b"}, {"c"}}
	cp := make([]*Foo, len(list))
	for i, v := range list {
		cp[i] = &v
	}

	for i := 0; i < len(cp); i++ {
		fmt.Println(*cp[i])
	}
}

func TestRange2(t *testing.T) {
	numbers2 := [...]int{1, 2, 3, 4, 5, 6}
	maxIndex2 := len(numbers2) - 1
	for i, e := range numbers2 {
		if i == maxIndex2 {
			numbers2[0] += e
		} else {
			numbers2[i+1] += e
		}
	}
	fmt.Println(numbers2)
}

func TestRange3(t *testing.T) {
	numbers2 := []int{1, 2, 3, 4, 5, 6}
	maxIndex2 := len(numbers2) - 1
	for i, e := range numbers2 {
		if i == maxIndex2 {
			numbers2[0] += e
		} else {
			numbers2[i+1] += e
		}
	}
	fmt.Println(numbers2)
}

func TestRange4(t *testing.T) {
	list := []Foo{{"a"}, {"b"}, {"c"}}
	for i := range list {
		fmt.Println(i)
	}
}
