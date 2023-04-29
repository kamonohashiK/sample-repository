package main

import "fmt"

func main() {
	fmt.Println(Hello())
}

func Hello() string {
	return "Hello World!"
}

func NewHello() string {
	return "Hello New World!"
}
