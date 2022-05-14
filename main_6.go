package main

import "fmt"

// Named Return Value

func myFunction(x int, y int) (result int) {
	result = x + y
	return
}

func myFunction2(x int, y int) (result int) {
	result = x + y
	return result
}

func main() {
	fmt.Println(myFunction(1, 2))
}
