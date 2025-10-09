package main

import (
    "fmt"
)
func SquareRoot(number int) int {
    if number == 0 {
        return 0
    }
    if number ==1{
        return 1
    }
    for i := number/2 ; i * i > 0; i--{
        if i * i <= number{
            return i
        }
    }
    return -1
}
func main() {
    fmt.Println(SquareRoot(9))
    fmt.Println(SquareRoot(16))
    fmt.Println(SquareRoot(25))
    fmt.Println(SquareRoot(26))
    fmt.Println(SquareRoot(0))
    fmt.Println(SquareRoot(-1))
    fmt.Println(SquareRoot(1))
    fmt.Println(SquareRoot(2))

}