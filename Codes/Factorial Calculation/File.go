package main
import "fmt"

func factorial(number int) int{
    factorial := 1
    for i := 1; i <= number; i++{
        factorial *= i
    }
    return factorial
}

func main() {
  fmt.Println("Factorial calculation")
  var find int
  fmt.Println("Enter the number to find the factorial: ")
  fmt.Scan(&find)
  fmt.Println("The factorial of number is ", factorial(find))
}
