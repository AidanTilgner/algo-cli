package helpers

import (
	t "algo-cli/types"
)

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func QuickSort(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = QuickSort(arr, low, p-1)
		arr = QuickSort(arr, p+1, high)
	}
	return arr
}

func partitionResults(results []t.Result, low int, high int) ([]t.Result, int) {
	pivot := results[high].Rank
	i := low
	for j := low; j < high; j++ {
		if results[j].Rank < pivot {
			results[i], results[j] = results[j], results[i]
			i++
		}
	}
	results[i], results[high] = results[high], results[i]
	return results, i
}

func QuickSortResults(results []t.Result, low int, high int) []t.Result {
	if low < high {
		var p int
		results, p = partitionResults(results, low, high)
		results = QuickSortResults(results, low, p-1)
		results = QuickSortResults(results, p+1, high)
	}
	return results
}

func ReverseResults(results []t.Result) []t.Result {
	for i := 0; i < len(results)/2; i++ {
		j := len(results) - i - 1
		results[i], results[j] = results[j], results[i]
	}
	return results
}