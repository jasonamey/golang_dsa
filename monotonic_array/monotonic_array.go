package golang_dsa

func MonotnicArray(array []int) bool {
	increasing := true
	decreasing := true
	for i := 1; i < len(array); i++ {
		if array[i] < array[i-1] {
			increasing = false
		}
		if array[i] > array[i-1] {
			decreasing = false
		}
	}
	return increasing || decreasing
}
