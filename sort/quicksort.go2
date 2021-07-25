package sort

import (
	"math/rand"
)

func quickSortImpl[T Ordered](slice []T, begin, end int) {
	if end-begin <= 1 {
		return
	}
	pi := rand.Intn(end-begin) + begin
	pivot := slice[pi]
	i := begin
	j := end - 1
	for {
		for slice[i] < pivot {
			i++
		}
		for slice[j] > pivot {
			j--
		}
		if i >= j {
			break
		}
		slice[i], slice[j] = slice[j], slice[i]
		i++
		j--
	}
	quickSortImpl(slice, begin, i)
	quickSortImpl(slice, j+1, end)
}

func QuickSort[T Ordered](slice []T) {
	quickSortImpl(slice, 0, len(slice))
}
