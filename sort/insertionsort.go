package sort

func InsertionSort[T Ordered](slice []T) {
	for i := 1; i < len(slice); i++ {
		for j := i; j > 0 && slice[j] < slice[j-1]; j-- {
			slice[j], slice[j-1] = slice[j-1], slice[j]
		}
	}
}
