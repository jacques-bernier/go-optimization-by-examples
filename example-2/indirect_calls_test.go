package main

import "testing"

type add struct {
	a int64
	b int64
}

func (a *add) Add() int64 {
	return a.a + a.b
}

type adder interface {
	Add() int64
}

type dummy struct{ m int64 }

func (d *dummy) Direct(a *add) int64 {
	return a.Add() + d.m
}

func (d *dummy) Indirect(a adder) int64 {
	return a.Add() + d.m
}

var ret int64

func Benchmark_direct(b *testing.B) {
	b.ReportAllocs()

	d := &dummy{
		m: 123,
	}

	a := &add{
		a: 321,
		b: 123,
	}

	for i := 0; i < b.N; i++ {
		ret = d.Direct(a)
	}
}

func Benchmark_indirect(b *testing.B) {
	b.ReportAllocs()

	d := &dummy{
		m: 123,
	}

	a := &add{
		a: 321,
		b: 123,
	}

	for i := 0; i < b.N; i++ {
		ret = d.Indirect(a)
	}
}
