// command_line_script.go
/*
Command Line Script Template in Go
-----------------------------------

This script demonstrates how to handle command line arguments in Go.

Usage:
    1. Run the script with command line arguments (e.g., go run command_line_script.go arg1 arg2).
    2. Access the arguments using os.Args.

Example:
    go run command_line_script.go hello world
*/

package main

import (
    "fmt"
    "os"
)

func main() {
    if len(os.Args) < 3 {
        fmt.Println("Usage: go run command_line_script.go <parameter1> <parameter2>")
        return
    }

    parameter1 := os.Args[1]
    parameter2 := os.Args[2]

    fmt.Printf("Your inputs were %s and %s\n", parameter1, parameter2)
}
