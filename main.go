package main

import "fmt"

func main() {
	fmt.Println("hello world")
}

type Handle struct {
}

func (Handle) Unpatch() {

}

func Patch(f interface{}, f2 interface{}) Handle {
	return Handle{}
}
