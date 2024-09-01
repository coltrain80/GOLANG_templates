// class_template.go
/*
Class Template in Go
---------------------

This file provides a template for defining a Go struct with methods. 
Go does not have classes, but you can achieve similar functionality using structs and methods.

Usage:
    1. Define a struct with fields representing its state.
    2. Define methods on the struct to perform operations.

Example:
    type Person struct {
        Name string
        Age  int
    }

    func (p *Person) Greet() string {
        return fmt.Sprintf("Hello, my name is %s and I am %d years old.", p.Name, p.Age)
    }
*/

package main

import (
    "fmt"
)

// Define a struct with fields
type Person struct {
    Name string
    Age  int
}

// Method to greet, defined on the Person struct
func (p *Person) Greet() string {
    return fmt.Sprintf("Hello, my name is %s and I am %d years old.", p.Name, p.Age)
}

func main() {
    person := Person{Name: "John Doe", Age: 30}
    fmt.Println(person.Greet())
}
