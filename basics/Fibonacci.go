package main

func fibonacci(x int) int {
	if x <= 1 {
		return x
	}
	return fibonacci(x-1) + fibonacci(x-2)
}
