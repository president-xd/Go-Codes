package main

import "fmt"

// Function to sum two integers and return the result
func sum(a int, b int) int {
    return a + b
}

// Function to subtract two integers and return the result
func subtract(a int, b int) int {
    return a - b
}

// Function to divide two integers and return the result as float32
func divide(a int, b int) float32 {
    if b == 0 {
        fmt.Println("Error: Division by zero is not allowed.")
        return 0
    }
    return float32(a) / float32(b)
}

// Function to multiply two integers and return the result
func multiply(a int, b int) int {
    return a * b
}

func main() {
    var num1, num2 int
    var choice int
    
    fmt.Println("Calculator --By president")
    fmt.Println("-------------------------------")
    fmt.Println("Press 1: Addition")
    fmt.Println("Press 2: Subtraction")
    fmt.Println("Press 3: Multiplication")
    fmt.Println("Press 4: Division")
    fmt.Println("-------------------------------")
    
    fmt.Print("Enter your choice: ")
    fmt.Scan(&choice)
    
    fmt.Print("Enter first number: ")
    fmt.Scan(&num1)
    
    fmt.Print("Enter second number: ")
    fmt.Scan(&num2)
    
    switch choice {
    case 1:
        fmt.Println("Sum of both is", sum(num1, num2))
    case 2:
        fmt.Println("Subtraction of both is", subtract(num1, num2))
    case 3:
        fmt.Println("Multiplication of both is", multiply(num1, num2))
    case 4:
        fmt.Println("Division of both is", divide(num1, num2))
    default:
        fmt.Println("Invalid choice")
    }
}
