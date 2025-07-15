package main

import "fmt"

// add returns the sum of two integers.
func add(a, b int) int {
    return a + b
}

// subtract returns the difference of two integers.
func subtract(a, b int) int {
	return a - b
}

// multiply returns the product of two integers.
func multiply(a, b int) int {
	return a * b
}

// divide performs division and returns an error if dividing by zero.
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

// main function demonstrates basic operations and control structures.
func main() {
    // Variable declarations using short assignment
    x := 5
    y := 3

    // Call add function
    sum := add(x, y)

	// Call subtract function
	diff := subtract(x, y)

	// Call multiply function
	prod := multiply(x, y)

	// Call divide function
	quotient, err := divide(x, y)

    // Print result
    fmt.Println("Sum:", sum)
	fmt.Println("Difference:", diff)
	fmt.Println("Product:", prod)
	if err != nil {	
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Quotient:", quotient)
	}

    // Conditional check
    if sum > 5 {
        fmt.Println("Sum is greater than 5")
    }

    // Simple for loop
    for i := 0; i < 3; i++ {
        fmt.Println("Loop iteration:", i)
    }
}
