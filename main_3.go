package main

import "fmt"

// Function with Parameter Example

func familyName(fname string) {
	fmt.Println("Hello", fname, "Refsnes")
}

func main() {
	familyName("Liam")
	familyName("Jenny")
	familyName("Anja")
}
