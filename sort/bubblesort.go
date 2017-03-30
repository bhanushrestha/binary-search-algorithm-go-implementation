package sort

type SortingAlgorithm interface {
	SortArray(unsortedArray []int)
}

type BubbleSort struct {
}

func NewBubbleSort() *BubbleSort {
	return new(BubbleSort)
}

//Implementation of the bubble sort algorithm to sort given unsortedArray
func (sort *BubbleSort) SortArray(unsortedArray []int) {
	sortComplete := true
	for i := 1; i < len(unsortedArray); i++ {
		if !sortComplete {
			return
		}
		sortComplete = false
		for j := 0; j < len(unsortedArray)-i; j++ {
			if unsortedArray[j] > unsortedArray[j+1] {
				temp := unsortedArray[j]
				unsortedArray[j] = unsortedArray[j+1]
				unsortedArray[j+1] = temp
				sortComplete = true
			}
		}
	}
}

/*func SortArray(unsortedArray []int) []int {
	sortedArray := make([]int, len(unsortedArray))
	sortComplete := true
	copy(sortedArray, unsortedArray)

	for i := 1; i < len(sortedArray); i++ {
		if !sortComplete {
			return sortedArray
		}
		sortComplete = false
		for j := 0; j < len(sortedArray)-i; j++ {
			if sortedArray[j] > sortedArray[j+1] {
				temp := sortedArray[j]
				sortedArray[j] = sortedArray[j+1]
				sortedArray[j+1] = temp
				sortComplete = true
			}
		}
	}

	return sortedArray
}*/
