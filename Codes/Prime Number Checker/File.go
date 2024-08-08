package main
import "fmt"

func checkerPrime(num_1 int) bool{
    var checker bool = false
    
    if num_1 == 1{
        return false
    } else if num_1 == 2{
        return true
    } else if num_1 == 0{
        return false
    }
    
    for i := 2; i <= num_1 / 2; i++{
        if num_1 % i == 0{
            return false
            break;
        } else{
            checker = true
        }
    }
    
    return checker
}


func main() {
    var number int
    
    fmt.Println("Enter the number : ")
    fmt.Scan(&number)
    
    if checkerPrime(number) == true{
        fmt.Println("Prime Number")
    } else{
        fmt.Println("Not prime number")
    }
    
}
