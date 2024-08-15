package main
import ("fmt" // Library that is used for input & Output Functions
        "reflect" // Library that is printing the types of variable
        "unsafe" // Library that is printing the size of variable
        )

func main() {
  var varing int16 = 10
  fmt.Println("The data types of the varing variable is ", reflect.TypeOf(varing))
  fmt.Println("The data size of the varing variable is ", unsafe.Sizeof(varing), " bytes = ", (unsafe.Sizeof(varing) * 8), "Bits")
}
