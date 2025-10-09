package main

import "fmt"

func getTotalX(a []int32, b []int32) int32 {
	var min int = int(b[0])
	var factor int64 = int64(a[0])

	// still find min (kept for your structure, though unused later)
	for i := 0; i < len(b); i++ {
		if int(b[i]) < min {
			min = int(b[i])
		}
	}

	// gcd helper
	gcd := func(x, y int64) int64 {
		for y != 0 {
			x, y = y, x%y
		}
		return x
	}

	// find lcm of a
	factor = int64(a[0])
	for i := 1; i < len(a); i++ {
		factor = factor * int64(a[i]) / gcd(factor, int64(a[i]))
	}

	// find gcd of b
	var g int64 = int64(b[0])
	for i := 1; i < len(b); i++ {
		g = gcd(g, int64(b[i]))
	}

	// find count
	count_length := 0
	count_final := 0
	j := 1
	multi := j * int(factor)
	for i := 0; multi <= int(g) && i < len(b); {
		if int(b[i])%multi == 0 {
			count_length++
			i++
			if count_length == len(b) {
				count_final++
				count_length = 0
				i = 0
				j++
				multi = j * int(factor)
			}
		} else if int(b[i])%multi != 0 {
			count_length = 0
			i = 0
			j++
			multi = j * int(factor)
			continue
		}
	}

	fmt.Println(factor)
	return int32(count_final)
}

func main() {
	fmt.Println(getTotalX([]int32{2 ,4},[]int32{16 ,32 ,96}))
	fmt.Println(getTotalX([]int32{3 ,4},[]int32{24 ,48}))
	fmt.Println(getTotalX([]int32{2,3 ,6},[]int32{42 ,84}))
	fmt.Println(getTotalX([]int32{3,9,6},[]int32{36 ,72}))


}