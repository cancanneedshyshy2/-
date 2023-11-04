package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	num1 := 76
	num2 := 30
	sum := add(num1, num2)
	fmt.Println("两个数的和为：", sum)
}

