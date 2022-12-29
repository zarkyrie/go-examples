package main

import "fmt"

func main() {
	func1()
}

func func1() {
	defer func() {
		if res := recover(); res != nil {
			fmt.Println(res)
		}
	}()

	func2()
}

func func2() {
	panic("haha")
}
