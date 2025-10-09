package main

import "fmt"

func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
    var count int32 = 0
    for i := 0 ; i < int(n) ; i++{
        for j := 0 ; j < i ; j++{
            if (ar[i] + ar[j] )%k ==0{
				fmt.Printf("(%v + %v) / %v = %v\n" ,ar[i],ar[j],k, (ar[i] + ar[j] )%k)
                count++
                
            }
        }
    }
    return count
}
func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
    // Step 1: count remainders
    freq := make([]int32, k)
    for _, val := range ar {
        r := val % k
        freq[r]++
    }

    var count int32 = 0

    // Step 2: handle remainder 0 (pairs among themselves)
    count += freq[0] * (freq[0] - 1) / 2

    // Step 3: handle pairs of remainders (r, k-r)
    for r := int32(1); r <= k/2; r++ {
        if r != k-r { // normal case
            count += freq[r] * freq[k-r]
        }
    }

    // Step 4: special case when k is even → remainder k/2 pairs among themselves
    if k%2 == 0 {
        count += freq[k/2] * (freq[k/2] - 1) / 2
    }

    return count
}

func main() {
	fmt.Println(divisibleSumPairs(6,3,[]int32{1, 3, 2, 6, 1, 2}))
}