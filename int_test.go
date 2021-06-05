package types

import (
	"fmt"
	"sort"
	"testing"
)

func TestAdd(t *testing.T) {
	fmt.Println(INT_MAX)
	fmt.Println(INT_MIN)
	bucketCntBits := 3
	bucketCnt := 1 << bucketCntBits
	fmt.Println(bucketCnt)
}

func TestUniqueInt(t *testing.T) {
	a := []int{2, 1, 2, 5, 6, 3, 4, 5, 2, 3, 9}
	b := UniqueInts(a)
	sort.Ints(b)
	fmt.Println(b)
}

func BenchmarkUniqueInt(t *testing.B) {
	a := []int{2, 1, 2, 5, 6, 3, 4, 5, 2, 3, 9, 3, 13, 3426, 24, 3, 2, 7, 23, 796, 35, 53, 35, 870, 53, 35, 37, 346, 3, 1, 23, 34, 23, 513, 513, 66, 124, 124, 124}
	b := UniqueInts(a)
	sort.Ints(b)
	fmt.Println(b)
}
