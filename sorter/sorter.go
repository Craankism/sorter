package sorter

func BubbleSort(sorter []string) {
	n := len(sorter)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if sorter[j] > sorter[j+1] {
				sorter[j], sorter[j+1] = sorter[j+1], sorter[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}
func BubbleSortDesc(sorter []string) {
	n := len(sorter)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if sorter[j] < sorter[j+1] {
				sorter[j], sorter[j+1] = sorter[j+1], sorter[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}