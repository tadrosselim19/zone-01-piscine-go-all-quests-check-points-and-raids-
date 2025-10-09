package main

import "fmt"

// func validated_string(s string)bool{
    
// }
func Unzipstring(s string) string {
    if s == ""{
        return "Invalid Input"
    }
    final := ""
    for i := 0 ; i < len(s) ; i++{
        if i % 2 == 0 && !(s[i] >='0' && s[i] <= '9') {
            return "Invalid Input"
        }
        if i+1 <len(s) && (i+1) % 2 != 0 && (!(s[i+1] <= 'Z' && s[i+1] >= 'A') && !(s[i+1] <= 'z' && s[i+1] >= 'a')){
            return "Invalid Input"
        }
        if s[i] >= '0' && s[i] <= '9'{
            for j := '0' ; j < rune((s[i])) ; j++{
                final += string(rune(s[i+1]))
            }
            i++
        }else{
            final += string(rune(s[i]))
        }
    }
    return final
}
func main() {
    fmt.Println(Unzipstring("2f"))
    fmt.Println(Unzipstring("2O5u2H0e"))
    fmt.Println(Unzipstring("3p6i1W"))
    fmt.Println(Unzipstring("6H8a"))
    fmt.Println(Unzipstring("2p4;7g"))
    fmt.Println(Unzipstring("2a 6p8f"))
    fmt.Println(Unzipstring("2t4dD"))
    fmt.Println(Unzipstring("82t4D"))
    fmt.Println(Unzipstring(""))

}