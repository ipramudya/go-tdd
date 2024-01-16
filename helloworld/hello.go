package main

import "fmt"

const HiPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}

	return HiPrefix + name
}

func main() {
	fmt.Println(Hello("John"))
}
