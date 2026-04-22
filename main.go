package main

import "fmt"

func greet(name string) string {
	return fmt.Sprintf("hello %s", name)
}

func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(greet("world"))
	fmt.Println(add(2, 3))
}
