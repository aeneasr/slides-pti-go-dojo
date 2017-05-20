package main

type T struct {}

func (T) Foo(s string) {
	println(s)
}

var fn func(T, string) = T.Foo

