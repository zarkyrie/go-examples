package main

import (
	"fmt"
)

func main() {
	//stopCh := make(chan struct{})
	//fmt.Println("main 1")
	//go Show(stopCh)
	//time.Sleep(2 * time.Second)
	//fmt.Println("main 2")
	//close(stopCh)
	//time.Sleep(2 * time.Second)

	//go func() {
	//	defer func() {
	//		if err := recover(); err != nil {
	//			fmt.Println(err)
	//		}
	//	}()
	//	panic("hllow")
	//}()
	//
	//if err := recover(); err != nil {
	//	fmt.Println(err)
	//}
	//time.Sleep(1 * time.Second)
	//fmt.Println("main")

	aStruct := AStruct{name: "hello"}
	aStruct.Show()
	fmt.Print(aStruct.name)
}

//func Show(stopCh <-chan struct{}) {
//	fmt.Println("step 1")
//	<-stopCh
//	fmt.Println("step 2")
//}

type AStruct struct {
	name string
}

func (receiver AStruct) Show() {
	fmt.Println("A struct")
	receiver.name = "dsdfa"
}

func (receiver AStruct) ShowV2() {
	fmt.Println("A struct V2")
}

type BStruct struct {
	AStruct
}

func (receiver BStruct) Show() {
	fmt.Println("B struct")
}
