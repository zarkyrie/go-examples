package programmingmode

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestSlice(t *testing.T) {
	path := []byte("AAAA/BBBBBBBBB")
	sepIndex := bytes.IndexByte(path, '/')
	//limit capacity
	dir1 := path[:sepIndex:sepIndex]
	dir2 := path[sepIndex+1:]
	fmt.Println("dir1 =>", string(dir1)) //prints: dir1 => AAAA
	fmt.Println("dir2 =>", string(dir2)) //prints: dir2 => BBBBBBBBB
	dir1 = append(dir1, "suffix"...)
	fmt.Println("dir1 =>", string(dir1)) //prints: dir1 => AAAAsuffix
	fmt.Println("dir2 =>", string(dir2)) //prints: dir2 => uffixBBBB
}

//深度比较
func TestDeepCmp(t *testing.T) {
	type data struct {
	}

	v1 := data{}
	v2 := data{}
	fmt.Println("v1 == v2:", reflect.DeepEqual(v1, v2))
	//prints: v1 == v2: true

	m1 := map[string]string{"one": "a", "two": "b"}
	m2 := map[string]string{"two": "b", "one": "a"}
	fmt.Println("m1 == m2:", reflect.DeepEqual(m1, m2))
	//prints: m1 == m2: true

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	fmt.Println("s1 == s2:", reflect.DeepEqual(s1, s2))
	//prints: s1 == s2: true
}

type iValid interface {
	name() string
	toStr() string
}

type demo struct {
}

func (s *demo) name() string {
	return "hello"
}

func (s *demo) toStr() string {
	return "demo"
}

//强制验证接口
func TestValidInterface(t *testing.T) {
	var _ iValid = (*demo)(nil)
}

//拼接字符串，可能比+拼接性能更好
func TestConcatStr(t *testing.T) {
	a := "hello "
	b := "world"

	var str strings.Builder
	str.WriteString(a)
	str.WriteString(b)
	fmt.Println(str.String())
}
