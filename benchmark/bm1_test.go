package main

import (
	"sync"
	"testing"
)

type Input struct {
	a int64
	b int64
}

type Result struct {
	sumA int64
	sumB int64
}

type ResultOpt struct {
	sumA int64
	_    [56]byte
	sumB int64
}

func SumCommon(inputs []Input) Result {
	var wg sync.WaitGroup
	ret := Result{}
	wg.Add(2)
	go func() {
		for i := 0; i < len(inputs); i++ {
			ret.sumA = ret.sumA + inputs[i].a
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < len(inputs); i++ {
			ret.sumB = ret.sumB + inputs[i].b
		}
		wg.Done()
	}()
	wg.Wait()
	return ret
}

func SumOpt(inputs []Input) ResultOpt {
	var wg sync.WaitGroup
	ret := ResultOpt{}
	wg.Add(2)
	go func() {
		for i := 0; i < len(inputs); i++ {
			ret.sumA = ret.sumA + inputs[i].a
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < len(inputs); i++ {
			ret.sumB = ret.sumB + inputs[i].b
		}
		wg.Done()
	}()
	wg.Wait()
	return ret
}

func CreateInputs(n int) []Input {
	inputs := make([]Input, n)
	for i := 0; i < n; i++ {
		inputs[i].a = int64(i)
		inputs[i].b = int64(i * 2)
	}
	return inputs
}

func BenchmarkSum(b *testing.B) {
	inputs := CreateInputs(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SumCommon(inputs)
	}
}

func BenchmarkSumOpt(b *testing.B) {
	inputs := CreateInputs(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SumOpt(inputs)
	}
}
