package main

import "fmt"

// Reference: https://www.w3schools.com/go/

func main() {
	// Go Data Types
	var a bool = true    // Boolean
	var b int = 5        // Integer
	var c float32 = 3.14 // Floating point number
	var d string = "Hi!" // String
	fmt.Printf(`Boolean: %t
Integer: %d
Float: %f
String: %s
`, a, b, c, d)

	// Variable Declaration With Initial Value
	var hello string = "Hello" // type is string
	var world = "World"        // type is inferred
	exclamationMark := "1"     // type is inferred
	fmt.Printf("%s %s %s\n", hello, world, exclamationMark)

	// Variable Declaration Without Initial Value
	var xa string
	var xb int
	var xc bool
	fmt.Printf(`empty String: %s
empty Integer: %d
empty Boolean: %t
`, xa, xb, xc)

	// Value Assignment After Declaration
	var xstudent1 string
	xstudent1 = "John"
	fmt.Println(xstudent1)

	// Go Multiple Variable Declaration
	var ya, yb, yc, yd int = 1, 3, 5, 7
	fmt.Printf("%d %d %d %d\n", ya, yb, yc, yd)
	// If the type keyword is not specified, you can declare different types of variables in the same line
	var za, zb = 6, "Hello"
	zc, zd := 7, "World!"
	fmt.Printf("%d %s %d %s\n", za, zb, zc, zd)

	// Go Variable Declaration in a Block
	// Multiple variable declarations can also be grouped together into a block for greater readibility
	var (
		aa int
		ab int    = 1
		ac string = "hello"
	)
	fmt.Printf("%d %d %s\n", aa, ab, ac)

	// Multi-Word Variable Names
	myVariableName := "John"   // Camel Case
	MyVariableName := "John"   // Pascal Case
	my_variable_name := "John" // Snake Case
	fmt.Printf("%s %s %s\n", myVariableName, MyVariableName, my_variable_name)

	// Go Constants
	const PI = 3.14 // Constant names are usually written in uppercase letters

	// Multiple Constants Declaration
	const (
		A int = 1
		B     = 3.14
		C     = "Hi!"
	)

	// The Printf() function
	/* %v is used to print the value of the arguments
	%T is used to print the type of the arguments */
	fmt.Printf("A has value: %v and type: %T\n", A, A)

	// General Formatting Verbs
	/* %v Prints the value in the default format
	%#v Prints the value in Go-syntax format
	%% Prints the % sign */
	var txt = "Hello World!"
	fmt.Printf(`Prints the value in default format: %v\n
Prints the value in Go-syntax format: %#v
Prints the %% sign: 15.5%%
`, txt, txt)

	// Integer Formatting Verbs
	var bi = 15 // Base 10
	fmt.Printf(`
Base 2: %b
Base 10: %d
Base 10 and always show sign: %+d
Base 8: %o
Base 8, with leading 0o: %O
Base 16, lowercase: %x
Base 16, with leading 0x: %#x
Pad with spaces (width 4, right adjusted): %4d
Pad with spaces (width 4, left justified): %-4d
Pad with zeroes (width 4): %04d
`, bi, bi, bi, bi, bi, bi, bi, bi, bi, bi)

	// String Formatting Verbs
	var atxt = "Hello"
	fmt.Printf(`Prints the value as plain string: %s
Prints the value as double-quoted string: %q
Prints the value as plain string (width 8, right justified): %8s
Prints the value as plain string (width 8, left justified): %-8s
Prints the value as hex dump of bytes value: %x
Prints the value as hex dump with spaces: % x
`, atxt, atxt, atxt, atxt, atxt, atxt)

	// BOolean Formatting Verbs
	var i = true
	var j = false
	fmt.Printf("%v %t\n", i, j)

	// Float Formatting Verbs
	var ai = 3.141
	fmt.Printf(`Scientific notation with 'e' as exponent: %e
Decimal point, no exponent: %f
Default width, precision 2: %.2f
Width 6, precision 2: %6.2f
Exponent as needed, only necessary digits: %g
`, ai, ai, ai, ai, ai)

	/* Go Integer Data Types
	Depends on platform: int32 in 32 bit systems and int64 in 64 bit systems
	int = Signed Integer, can store both positive and negtative values
	uint = Unsigned Integer, can only store non-negative values.

	int8 = 8 bits / 1 byte = -128 to 127
	int16 = 16 bits / 2 byte = -32768 to 32767
	int32, int64, uint8, uint16, uint32, uint64
	*/

	/* Go Float Data Types
	float32, float64
	*/

	// Go Arrays
	/* In Go, arrays have a fixed length.
	The length of the array is either defined by a number or is inferred (means that the compiler decides the length of the array, based on the number of values)
	*/
	var array_name = [1]int{1}          // here length is defined
	var xarray_name = [...]int{1, 2, 3} // here length is inferred

	fmt.Printf(`%v
%v
`, array_name, xarray_name)

	// Change Elements of an Array
	prices := [3]int{10, 20, 30}
	prices[2] = 50

	fmt.Println(prices)

	// Array Initialization
	// If an array or one of its elements has not been initialized in the code, it is assigned the default value of its type
	arr1 := [5]int{}              // not initialized
	arr2 := [5]int{1, 2}          // partially initialized
	arr3 := [5]int{1, 2, 3, 4, 5} // fully initialized

	fmt.Printf(`arr1: %v
arr2: %v
arr3: %v
`, arr1, arr2, arr3)

	// Initialize Only Specific Elements
	arr4 := [5]int{1: 10, 2: 40}

	fmt.Println(arr4) // expected: [ 0 10 40 0 0 ]

	// Find the Length of an Array
	arr5 := [4]string{"Volvo", "BMW", "Ford", "Mazda"}
	arr6 := [...]int{1, 2, 3, 4, 5, 6}

	fmt.Printf(`arr5 len: %v
arr6 len: %v
`, len(arr5), len(arr6))

	// Go Slices
	// unlike arrays, the length of a slice can grow and shrink asw you see fit
	myarray := [...]int{1, 2, 3}
	myslice := []int{1, 2, 3}

	// len() function - returns the length of the slice (the number of elements in the slice)
	// cap() function - returns the capacity of the slice (the number of elements the slice can grow or shrink to)
	fmt.Printf(`len(myarray): %v
cap(myarray): %v
len(myslice): %v
cap(myslice): %v
`, len(myarray), cap(myarray), len(myslice), cap(myslice))

	// Create a Slice From an Array
	xarr1 := [6]int{10, 11, 12, 13, 14, 15}
	xmyslice := xarr1[2:4] // xmyslice := xarr1[start:end]

	fmt.Printf(`xmyslice, expected [ 12 13 ]: %v
length, expected 2: %v
capacity, expected 4: %v
`, xmyslice, len(xmyslice), cap(xmyslice))

	// Create a Slice With The make() Function
	// make([]type, length, capacity)
	myslice1 := make([]int, 5, 10)
	fmt.Printf(`myslice1: %v
length: %v
capacity: %v
`, myslice1, len(myslice1), cap(myslice1))

	// Append Elements to a Slice
	// slice_name = append(slice_name, element1, element2)
	xmyslice1 := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf(`xmyslice1: %v
length: %v
capacity: %v
`, xmyslice1, len(xmyslice1), cap(xmyslice1))

	xmyslice1 = append(xmyslice1, 20, 21)
	fmt.Printf(`xmyslice1: %v
length: %v
capacity: %v
`, xmyslice1, len(xmyslice1), cap(xmyslice1))

	// Append One Slice To Another Slice
	ymyslice1 := []int{1, 2, 3}
	ymyslice2 := []int{4, 5, 6}
	ymyslice3 := append(ymyslice1, ymyslice2...)

	fmt.Printf(`ymyslice3: %v
length: %v
capacity: %v
`, ymyslice3, len(ymyslice3), cap(ymyslice3))

	// Change The Length of a Slice
	barr1 := [6]int{9, 10, 11, 12, 13, 14} // An Array
	bmyslice1 := barr1[1:5]                // Slice array
	fmt.Printf(`bmyslice1: %v
length: %v
capacity: %v
`, bmyslice1, len(bmyslice1), cap(bmyslice1))

	bmyslice1 = barr1[1:3] // Change length by-reslicing the array
	fmt.Printf(`bmyslice1: %v
length: %v
capacity, expected capacity will be the same: %v
`, bmyslice1, len(bmyslice1), cap(bmyslice1))

	bmyslice1 = append(bmyslice1, 20, 21, 22, 23) // Change length by appending items
	fmt.Printf(`bmyslice1: %v
length: %v
capacity, expected capacity doubled or 10: %v
`, bmyslice1, len(bmyslice1), cap(bmyslice1))

	// Memory Efficiency <- disputed, I assume w3schools is wrong.
	/* When using slices, Go loads all the underlying elements into the memory <- understand this carefully to differentiate between the copy() function and the slice[start:end] operator.
	If the array is large and you need only a few elements, it is better to copy those elements using the copy() function
	The copy() function creates a new underlying array with only the required elements for the slice. This will reduce the memory used for the program.
	copy(dest, src)
	*/

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	// Original slice
	fmt.Printf(`numbers: %v
length: %d
capacity: %d
`, numbers, len(numbers), cap(numbers))

	// Create copy with only needed numbers
	neededNumbers := numbers[:len(numbers)-10] // neededNumbers := numbers[0:end] <- this will COPY the reference, any changes to the neededNumbers' slice will affect the numbers' slice aswell.
	numbersCopy := make([]int, len(neededNumbers))
	copy(numbersCopy, neededNumbers) // <- this will COPY the data.

	fmt.Printf(`neededNumbers: %v
length: %d
capacity: %d
`, neededNumbers, len(neededNumbers), cap(neededNumbers))

	fmt.Printf(`numbersCopy: %v
length: %d
capacity: %d
`, numbersCopy, len(numbersCopy), cap(numbersCopy))

	// Go Operators
	p := 34
	q := 20

	result1 := p % q
	fmt.Printf("Modulus: p %% q = %d\n", result1)

	result2 := p / q
	fmt.Printf(`Division: p / q = %d
Expected result: 1 is %t
`, result2, result2 == 1)

	result3 := p & q
	fmt.Printf(`Bitwise AND: p & q = %d
100010
010100
------ &
000000
`, result3)

	// The result of OR is 1 any of the two bits is 1
	result4 := p | q
	fmt.Printf(`Bitwise OR: p | q = %d
100010
010100
------ |
110110
`, result4)

	// The result of XOR is 1 if the two bits are different
	result5 := p ^ q
	fmt.Printf(`Bitwise XOR: p ^ q = %d
100010
010100
------ ^
110110
`, result5)

	result6 := p << 1
	fmt.Printf(`left shift: p << 1 = %d
 100010
------- << 1
1000100
`, result6)

	result7 := p >> 1
	fmt.Printf(`right shift: p >> 1 = %d
100010
------ >> 1
010001
`, result7)

	// To get the bits that are in 3 AND NOT 6 (order matters)
	var result8 = 3 &^ 6
	fmt.Printf(`AND NOT: 3 &^ 6 = %d
011
110
--- &^
001
`, result8)

	// Go if Statement
	if 20 > 18 {
		fmt.Println("20 is greater than 18")
	}

	// The switch Statement
	day := 3
	switch day {
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	default:
		fmt.Println("no case match")
	}

	// Go Multi-case switch Statement
	switch day {
	case 1, 3, 5:
		fmt.Println("Odd weekday")
	case 2, 4:
		fmt.Println("Even weekday")
	}

	// Go for Loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// The continue Statement
	// The continue statement is used to skip one or more iterations in the loop. It then continues with the next iteration in the loop.
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}

	// The break statement
	// The break statement is used to break/terminate the loop execution
	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println(i)
	}

	// Nested Loops
	adj := [2]string{"big", "tasty"}
	fruits := [3]string{"apple", "orange", "banana"}
	for i := 0; i < len(adj); i++ {
		for j := 0; j < len(fruits); j++ {
			fmt.Println(adj[i], fruits[j])
		}
	}

	// The Range Keyword
	// The range keyword is used to more easily iterate over an array, slice or map. It returns both the index and the value.
	xfruits := [3]string{"apple", "orange", "banana"}
	for idx, val := range xfruits {
		fmt.Printf("%v\t%v\n", idx, val)
	}

	// To only show the value or the index, you can omit the other output using an underscore (_)
	yfruits := [3]string{"apple", "orange", "banana"}
	for _, val := range yfruits {
		fmt.Printf("%v\n", val)
	}
}
