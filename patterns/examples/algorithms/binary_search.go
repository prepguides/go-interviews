package algorithms

// BinarySearch performs binary search on a sorted slice
func BinarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

// BinarySearchRecursive performs binary search recursively
func BinarySearchRecursive(arr []int, target, left, right int) int {
	if left > right {
		return -1
	}

	mid := left + (right-left)/2

	if arr[mid] == target {
		return mid
	} else if arr[mid] < target {
		return BinarySearchRecursive(arr, target, mid+1, right)
	} else {
		return BinarySearchRecursive(arr, target, left, mid-1)
	}
}
