package main

import "fmt"

// Function of Fibonacci Series
func fibonacci(num1 int, num2 int, nooftimes int) {
    fmt.Print(num1, " ", num2, " ") // Print the initial values once

    for i := 0; i < nooftimes-2; i++ { // Adjust the loop to print the remaining numbers
        sum := num1 + num2
        fmt.Print(sum, " ")
        num1 = num2
        num2 = sum
    }
    fmt.Println() // Print a newline at the end
}

func main() {
    fmt.Println("Fibonacci Series")
    fibonacci(0, 1, 9)
}
