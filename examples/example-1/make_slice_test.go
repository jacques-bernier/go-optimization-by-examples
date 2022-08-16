package main

import "testing"

type A struct {
	Value int64
}

var a []A

const size = 10000

func Benchmark_make(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		tmp := make([]A, 0)
		for j := int64(0); j < size; j++ {
			tmp = append(tmp, A{Value: j})
		}
		a = tmp
	}
}

func Benchmark_make_capacity(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		tmp := make([]A, 0, size)
		for j := int64(0); j < size; j++ {
			tmp = append(tmp, A{Value: j})
		}
		a = tmp
	}
}

func Benchmark_make_length(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		tmp := make([]A, size)
		for j := int64(0); j < size; j++ {
			tmp[j] = A{Value: j}
		}
		a = tmp
	}
}
