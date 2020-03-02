package main

// MergeSort performs the merge sort algorithm.
// Please supplement this function to accomplish the home work.
func MergeSort(src []int64) {
	result := mergeSort(src)
	copy(src, result)
}

func mergeSort(data []int64) []int64 {
	if len(data) <= 1 {
		return data
	}
	middle := len(data) / 2

	leftChan := make(chan []int64)
	rightChan := make(chan []int64)
	go func() {
		leftChan <- mergeSort(data[:middle])
	}()
	go func() {
		rightChan <- mergeSort(data[middle:])
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
