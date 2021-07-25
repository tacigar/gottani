package sort

func SelectionSort[T Ordered](slice []T) {
	for i := 0; i < len(slice); i++ {
		minIndex := i
		for j := i; j < len(slice); j++ {
			if slice[j] < slice[minIndex] {
				minIndex = j
			}
		}
		slice[i], slice[minIndex] = slice[minIndex], slice[i]
	}
}
