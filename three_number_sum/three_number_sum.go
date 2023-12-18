package golang_dsa

import (
	"sort"
)

func ThreeNumberSum(array []int, target int) [][]int {
	solution := [][]int{}
	sort.Ints(array)
	for i := 0; i < len(array)-1; i++ {
		lo := i + 1
		hi := len(array) - 1
		num := array[i]
		for lo < hi {
			test := num + array[lo] + array[hi]
			if test == target {
				arr := []int{num, array[lo], array[hi]}
				solution = append(solution, arr)
				lo++
				hi--
			} else if test < target {
				lo++
			} else {
				hi--
			}
		}
	}
	return solution
}
