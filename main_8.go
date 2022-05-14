// Multiple Return Values

package main

import "fmt"

func myFunction(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + " World!"
	return
}

func main() {
	fmt.Println(myFunction(5, "Hello"))

	a, b := myFunction(5, "Hello")
	fmt.Println(a, b)

	_, xb := myFunction(5, "Hello")
	fmt.Println(xb)
}
