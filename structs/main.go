package structs

import "fmt"

// Person defines a simple struct with Name and Age fields.
type Person struct {
    Name string
    Age  int
}

// Greet is a method on Person that prints a greeting.
func (p Person) Greet() {
    fmt.Printf("Hello, I am %s and I'm %d years old.\n", p.Name, p.Age)
}

func main() {
    // Create a new Person instance
    p := Person{Name: "Karma", Age: 17}

    // Call the Greet method on Person
    p.Greet()
}