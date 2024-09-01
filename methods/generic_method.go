// method_template.go
/*
Method Template in Go
----------------------

This file provides a template for defining methods on structs in Go.

Usage:
    1. Define a method with a receiver argument.
    2. Implement the desired functionality inside the method.

Example:
    type Calculator struct{}

    func (c *Calculator) Add(a, b int) int {
        return a + b
    }
*/

package main

import (
    "fmt"
)

// Define a struct to associate methods with it
type Calculator struct{}

// Define a method on the Calculator struct
func (c *Calculator) Add(a, b int) int {
    return a + b
}

func main() {
    calc := Calculator{}
    result := calc.Add(5, 7)
    fmt.Printf("The sum is: %d\n", result)
}
