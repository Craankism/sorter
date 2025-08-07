package sorter

import (
	"fmt"
)

func BubbleSort(sorter []string, debug bool){
	n := len(sorter)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if sorter[j] > sorter[j+1] {
				sorter[j], sorter[j+1] = sorter[j+1], sorter[j]
				swapped = true
				if debug {
					fmt.Printf("Swapped %s and %s\n", sorter[j], sorter[j+1])
				}
			}
		}
		if !swapped {
			break
		}
	}
}
func BubbleSortDesc(sorter []string, debug bool) {
	n := len(sorter)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if sorter[j] < sorter[j+1] {
				sorter[j], sorter[j+1] = sorter[j+1], sorter[j]
				swapped = true
				if debug {
					fmt.Printf("Swapped %s and %s\n", sorter[j], sorter[j+1])
				}
			}
		}
		if !swapped {
			break
		}
	}
}