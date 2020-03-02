package main

// MergeSort performs the merge sort algorithm.
// Please supplement this function to accomplish the home work.
func MergeSort(src []int64) {
	result := mergeSort(src)
	copy(src, result)
}

func mergeSort(data []int64) []int64 {
	return mergeSortAdapted(data)
}

func mergeSortAdapted(data []int64) []int64 {
	if len(data) < 1024 {
		return mergeSortSerialized(data)
	} else {
		return mergeSortParalleled(data)
	}
}

func mergeSortSerialized(data []int64) []int64 {
	if len(data) <= 1 {
		return data
	}
	middle := len(data) / 2

	left := mergeSortSerialized(data[:middle])
	right := mergeSortSerialized(data[middle:])

	return merge(left, right)
}

func mergeSortParalleled(data []int64) []int64 {
	if len(data) <= 1 {
		return data
	}
	middle := len(data) / 2

	leftChan := make(chan []int64)
	rightChan := make(chan []int64)
	go func() {
		leftChan <- mergeSortAdapted(data[:middle])
	}()
	go func() {
		rightChan <- mergeSortAdapted(data[middle:])
	}()

	return merge(<-leftChan, <-rightChan)
}

func merge(left, right []int64) []int64 {
	result := make([]int64, 0, len(left)+len(right))
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] > right[r] {
			result = append(result, right[r])
			r++
		} else {
			result = append(result, left[l])
			l++
		}
	}

	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return result
}
