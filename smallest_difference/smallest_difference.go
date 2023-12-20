package golang_dsa

import (
	"math"
	"sort"
)

func SmallestDifference(arr1, arr2 []int) []int {
	sort.Ints(arr1)
	sort.Ints(arr2)
	i := 0
	j := 0
	difference := math.MaxInt32
	var ans []int
	for i < len(arr1) && j < len(arr2) {
		num_1 := arr1[i]
		num_2 := arr2[j]
		if num_1 > num_2 {
			test := num_1 - num_2
			if test < difference {
				difference = test
				ans = []int{num_1, num_2}
			}
			j++
		} else {
			test := num_2 - num_1
			if test < difference {
				difference = test
				ans = []int{num_1, num_2}
			}
			i++
		}
	}
	return ans
}
