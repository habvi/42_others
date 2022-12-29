package main

import (
	"fmt"
	"math"
)

func min(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}

func max(x, y int) int {
	return int(math.Max(float64(x), float64(y)))
}

func subSlice(slice []int, begin int, length int, capacity int) []int {
	ret := make([]int, length, max(length, capacity))
	slice_len := len(slice)
	copy_len := min(slice_len, begin+length) - begin
	for i := 0; i < copy_len; i++ {
		ret[i] = slice[i+begin]
	}
	return ret
}

func main() {
	var orig = []int{0, 1, 2, 3, 4, 5}
	var ret []int

	ret = subSlice(orig, 0, 3, 3)
	fmt.Printf("ret = %v, len = %d, cap = %d\n", ret, len(ret), cap(ret))

	ret = subSlice(orig, 2, 7, 10)
	fmt.Printf("ret = %v, len = %d, cap = %d\n", ret, len(ret), cap(ret))

	ret = subSlice(orig, 2, 7, 3)
	fmt.Printf("ret = %v, len = %d, cap = %d\n", ret, len(ret), cap(ret))
}
