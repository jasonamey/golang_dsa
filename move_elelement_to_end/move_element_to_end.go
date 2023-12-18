package golang_dsa

// array : [2, 1, 2, 2, 2, 3, 4, 2], to move : 2
// answer: [1, 3, 4, 2, 2, 2, 2, 2]

func MoveElementToEnd(arr []int, toMove int) []int {
	i := 0
	j := len(arr) - 1
	for i < j {
		if arr[i] == toMove && arr[j] != toMove {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		}
		for i < len(arr) && arr[i] != toMove {
			i++
		}
		for j > -1 && arr[j] == toMove {
			j--
		}
	}
	return arr
}
