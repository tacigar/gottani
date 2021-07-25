package sort

func merge[T Ordered](slice1, slice2, slice []T) {
	i, j := 0, 0
	for i < len(slice1) || j < len(slice2) {
		if j >= len(slice2) || (i < len(slice1) && slice1[i] < slice2[j]) {
			slice[i+j] = slice1[i]
			i++
		} else {
			slice[i+j] = slice2[j]
			j++
		}
	}
}

func MergeSort[T Ordered](slice []T) {
	l := len(slice)
	if l <= 1 {
		return
	}
	m := l / 2
	n := l - m
	slice1, slice2 := make([]T, m), make([]T, n)
	for i := 0; i < m; i++ {
		slice1[i] = slice[i]
	}
	for i := 0; i < n; i++ {
		slice2[i] = slice[m+i]
	}
	MergeSort(slice1)
	MergeSort(slice2)
	merge(slice1, slice2, slice)
}
