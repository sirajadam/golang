package insertion

func insertion(arr []int) []int {
	for i := range arr {
		j := i
		for j > 0 && arr[j-1] > arr[j] {
			temp := arr[j]
			arr[j] = arr[j-1]
			arr[j-1] = temp
			j = j - 1
		}
	}
	return arr
}