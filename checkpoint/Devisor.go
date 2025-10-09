package main

import (
    "fmt"
    
)
func Divisors(n int) int {
    if n <= 0 { return 0 }

    count := 0
    for i := 1; i*i <= n; i++ {
        if n % i == 0 {
            count += 2       // i and n/i
            if i*i == n {
                count--      // perfect square, don’t double count
            }
        }
    }
    return count
}

func Divisors1(n int) int {
	count := 0
	if n < 0 { 
		return 0 
	}

	for i := 1 ; i <= n ; i++{
		if n % i == 0{
			count++
		}
	}
	return  count
	     
    
}
func main() {
    fmt.Println(Divisors(4))// 4 can be divided by 1 and 2 and 4
    fmt.Println(Divisors(9999999999999))//5 can be divided by 1 and 5
}
