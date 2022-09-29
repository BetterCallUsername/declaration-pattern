package main

import "fmt"

func declaration(targetFunc func()) {
	fmt.Println("Function called")

	targetFunc()

	fmt.Println("After called function")
}

func sayHello() {
	fmt.Println("Hello world")
}

func main() {
	declaration(sayHello)
}
