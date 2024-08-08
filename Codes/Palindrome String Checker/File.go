package main
import (
    "fmt"
    "strings"
    )

func checkPalindrome(a string) bool{
    a = strings.ToLower(a)
    reverse := ""
    
    for i := len(a) - 1; i >= 0; i--{
        reverse += string(a[i])
    }
    
    return reverse == a
    
}

func main() {
    var String string
    
    fmt.Println("Enter the string to check : ")
    fmt.Scan(&String)
    
    if checkPalindrome(String) == true{
        fmt.Println("Palindrome String")
    } else {
        fmt.Println("Not Palindrome string")
    }
    
}
