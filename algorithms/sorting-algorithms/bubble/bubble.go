package bubble

func bubble(arr []int) []int {

	n := len(arr)

	unsortedLength := n - 1
	for i := 0; i < n; i++ {

		for j := 0; j < unsortedLength; j++ {
			if arr[j+1] < arr[j] {
				tmp := arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = tmp
			}
		}
		unsortedLength--
	}
	return arr
}
