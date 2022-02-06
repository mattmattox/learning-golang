package main

import "fmt"

var hello = "Hello!"
var ok = true

func nope() {
	const ok = true
	var hello = "Hello!"
	_ = hello
}

func main() {
	fmt.Println(hello, ok)
}
