package main

func maximumToys(prices []int32, k int32) int32 {
	pricesLength := int32(len(prices))
	// sortToys(prices)
	prices = mergeSortToys(prices)
	var tempTotal int32

	if prices[0] > k {
		return 0
	}
	for i := int32(0); i < pricesLength; i++ {
		tempTotal += prices[i]
		if tempTotal > k {
			return i
		}
	}

	return pricesLength
}

func mergeSortToys(arr []int32) []int32 {
	arrLength := int32(len(arr))

	if arrLength == 1 {
		return arr
	}

	arrMiddle := arrLength / 2
	var left = make([]int32, arrMiddle)
	var right = make([]int32, arrLength-arrMiddle)

	for i := int32(0); i < arrLength; i++ {
		if i < arrMiddle {
			left[i] = arr[i]
		} else {
			right[i-arrMiddle] = arr[i]
		}
	}

	return merge(mergeSortToys(left), mergeSortToys(right))
}

func merge(left, right []int32) []int32 {
	result := make([]int32, len(left)+len(right))

	i := int32(0)

	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}

	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return result
}

func sortToys(a []int32) {
	n := int32(len(a))
	var temp int32

	for i := int32(0); i < n; i++ {

		for j := int32(0); j < n-1; j++ {
			// Swap adjacent elements if they are in decreasing order
			if a[j] > a[j+1] {
				temp = a[j]
				a[j] = a[j+1]
				a[j+1] = temp
			}
		}
	}
}

func main() {}
