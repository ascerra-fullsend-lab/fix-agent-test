package main

import "fmt"

const defaultGreeting = "hello stranger"

func greet(name string) string {
	if name == "" {
		return defaultGreeting
	}
	return fmt.Sprintf("hello %s", name)
}

func add(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

func main() {
	fmt.Println(greet(""))
	fmt.Println(add(2, 3))
	fmt.Println(multiply(4, 5))
}
