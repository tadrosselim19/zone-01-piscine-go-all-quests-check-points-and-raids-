package main

import (
    "fmt"
    
)
func Slice(a []string, nbrs... int) []string{
    final := []string{}
    if len(nbrs) == 0 {
        return nil
    }
    count := -1

    // check the arrangment of values in nbrs , counting the array and mange negative number
    for i:= 0 ; i < len(nbrs) ; i++{
        if i+1 < len(nbrs) && nbrs[i] > nbrs[i+1]{
            return nil
        }
        if nbrs[i] * -1 > 0{
            nbrs[i] = len(a)+ nbrs[i]
        }
        count++
    }
    if count == 0{
        nbrs = append(nbrs, len(a))
        count++
    }
    j := 0
    for count >= 0 {
        stared := nbrs[j]
        end := nbrs[j+1]
        for i := stared ; i < end && i <len(a) ;i++{
            final = append(final, a[i])
        }
        count -=2
        j++
    }  
    return final
}
func main(){
    a := []string{"coding", "algorithm", "ascii", "package", "golang"}
    fmt.Printf("%#v\n",Slice(a, 1))
    fmt.Printf("%#v\n",Slice(a, 2 ,4))
    fmt.Printf("%#v\n",Slice(a, -3))
    fmt.Printf("%#v\n",Slice(a, -2, -1))
    fmt.Printf("%#v\n",Slice(a, 2, 0))
    fmt.Printf("%#v\n",Slice(a, 0 ,2,5))
    fmt.Printf("%#v\n",Slice(a, -5 , -2))


}