package concurrent

import (
	"fmt"
	"strings"
	"sync"
	"testing"
)

func TestWaitGroup(t *testing.T) {
	letter, number := make(chan bool), make(chan bool)
	wait := sync.WaitGroup{}

	go func() {
		i := 1
		for {
			select {
			case <-number:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				letter <- true
			}
		}
	}()
	wait.Add(1)
	go func(wait *sync.WaitGroup) {
		i := 'A'
		for {
			select {
			case <-letter:
				if i >= 'Z' {
					wait.Done()
					return
				}

				fmt.Print(string(i))
				i++
				fmt.Print(string(i))
				i++
				number <- true
			}

		}
	}(&wait)
	number <- true
	wait.Wait()
}

func TestName(t *testing.T) {
	f := func(s string) string {
		fmt.Println(s)
		return "hello " + s
	}

	invokeFunc(f)
}

func invokeFunc(f func(string) string) {
	str := f("wold")
	fmt.Println(str)
}

func TestCountChar(t *testing.T) {
	fmt.Println(strings.Count("梁结恒", ""))
}

func TestTextBlocks(t *testing.T) {
	str := `hello 
world
!`
	fmt.Println(str)
}
