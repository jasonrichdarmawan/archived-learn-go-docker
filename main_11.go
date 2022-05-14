// Go Maps

package main

import "fmt"

func main() {
	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}

	fmt.Printf("a\t%v\n", a)
	fmt.Printf("a\t%v\n", b)

	// Creating Maps using make() function
	var xa = make(map[string]string) // The map is empty now
	xa["brand"] = "Ford"
	xa["model"] = "Mustang"
	xa["year"] = "1964"

	xb := make(map[string]int)
	xb["Oslo"] = 1
	xb["Bergen"] = 2
	xb["Trondheim"] = 3
	xb["Stavanger"] = 4

	fmt.Printf("a\t%v\n", a)
	fmt.Printf("b\t%v\n", b)

	// Creating an Empty Map
	// The make() function is the right way to create an empty map. If you make an empty map in a different way and write to it, it will causes a runtime panic
	var ya = make(map[string]string)
	var yb map[string]string

	fmt.Println(ya == nil)
	fmt.Println(yb == nil)

	// Remove Element from Map
	var ia = make(map[string]string)
	ia["brand"] = "Ford"
	ia["model"] = "Mustang"
	ia["year"] = "1964"

	fmt.Println(ia)

	delete(ia, "year")

	fmt.Println(ia)

	// Check For Specific Elements in a Map
	var ja = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964", "day": ""}
	val1, ok1 := ja["brand"] // Checking for existing key and its value
	val2, ok2 := ja["color"] // Checking for non-existing key and its value
	val3, ok3 := ja["day"]   // Checking for existing key and its value
	_, ok4 := ja["model"]    // Only checking for existing key and not its value

	fmt.Println(val1, ok1)
	fmt.Println(val2, ok2)
	fmt.Println(val3, ok3)
	fmt.Println(ok4)

	// Maps Are References
	// Maps are references to hash tables
	// If two map variables refer to the same hash table, changing the content of one variable affect the content of the other.
	var xya = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	xyb := xya

	fmt.Println(xya)
	fmt.Println(xyb)

	xyb["year"] = "1970"
	fmt.Println("After change to b:")
	fmt.Println(xya)
	fmt.Println(xyb)

	// Iterating Over Maps
	// You can use range to iterate over maps
	// Note the order of the elements in the output
	yxa := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}
	for k, v := range yxa {
		fmt.Printf("%v : %v\n", k, v)
	}

	// Iterate Over Maps in a Specific Order
	ija := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}
	var ijb []string // defining the order
	ijb = append(ijb, "one", "two", "three", "four")

	for k, v := range ija { // loop with no order
		fmt.Printf("%v : %v\n", k, v)
	}

	for _, element := range ijb { // loop with the defined order
		fmt.Printf("%v : %v\n", element, ija[element])
	}
}
