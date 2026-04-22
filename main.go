package main

import "fmt"

func greet(name string) string {
	if name == "" {
		return "hello stranger"
	}
	return fmt.Sprintf("hello %s", name)
}

func add(a, b int) int {
	result := a + b
	return result
}

func multiply(a, b int) int {
	return a * b
}

func main() {
	result := greet("")
	fmt.Println(result)
	fmt.Println(add(2, 3))
	fmt.Println(multiply(4, 5))
}
