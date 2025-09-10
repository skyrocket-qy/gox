package main

import "sync"

type Question struct {
	A int
	B int
	C string
}

type Answer struct {
	Res int
}

type WhoAnswer struct {
	Name  string
	Mutex sync.Mutex
}
