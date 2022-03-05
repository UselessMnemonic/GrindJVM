package main

type Object struct {
	lock  uintptr
	class *Class
}
