package errors

import (
    "errors"
    "fmt"
)

// divide performs division and returns an error if dividing by zero.
func divide(a, b float64) (float64, error) {
    if b == 0 {
        // Return a new error
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}

func main() {
    // Attempt division by zero
    result, err := divide(10, 0)
    
    // Check for error and handle it
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    // If no error, print result
    fmt.Println("Result:", result)
}