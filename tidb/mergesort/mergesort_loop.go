package main

func mergeSortLoop(src []int64) {
	length := len(src)
	for i := 1; i < length; i *= 2 {
		index := 0
		for 2*i+index <= length {
			index += 2 * i
			mergeLoop(src, index-2*i, index-i, index)
		}
		if index+i < length {
			mergeLoop(src, index, index+i, length)
		}
	}
}

func mergeLoop(src []int64, left, mid, right int) {
	tmpSlice := make([]int64, right-left)
	i := left
	j := mid
	index := 0

	for index+left < right {
		if i < mid && (j == right || src[i] < src[j]) {
			tmpSlice[index] = src[i]
			i++
		} else {
			tmpSlice[index] = src[j]
			j++
		}
		index++
	}

	copy(src[left:right], tmpSlice)
}
