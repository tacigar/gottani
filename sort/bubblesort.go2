package sort

func BubbleSort[T Ordered](slice []T) {
	for i := 0; i < len(slice)-1; i++ {
		done := true
		for j := 1; j < len(slice)-i; j++ {
			if slice[j] < slice[j-1] {
				slice[j], slice[j-1] = slice[j-1], slice[j]
				done = false
			}
		}
		if done {
			return
		}
	}
}
