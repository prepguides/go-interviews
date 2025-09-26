package algorithms

// QuickSort sorts a slice using the quicksort algorithm
func QuickSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	left, right := 0, len(arr)-1
	pivot := partition(arr, left, right)

	QuickSort(arr[:pivot])
	QuickSort(arr[pivot+1:])
}

// partition partitions the array around a pivot
func partition(arr []int, left, right int) int {
	pivot := arr[right]
	i := left - 1

	for j := left; j < right; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[right] = arr[right], arr[i+1]
	return i + 1
}
