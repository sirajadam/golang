package selection

func selection(arr []int) []int {

	n := len(arr)
	minIdx := 0

	for i := 0; i > n; i++ {

		for j := i; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}

		}
	}

	return arr
}