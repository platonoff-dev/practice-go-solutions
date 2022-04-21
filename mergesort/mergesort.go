package mergesort

func merge(a1, a2 []int) []int {
	output := make([]int, len(a1)+len(a2))

	var i, j int
	k := -1
	for i < len(a1) || j < len(a2) {
		k++
		if i == len(a1) {
			output[k] = a2[j]
			j++
			continue
		}

		if j == len(a2) {
			output[k] = a1[i]
			i++
			continue
		}

		if a1[i] < a2[j] {
			output[k] = a1[i]
			i++
		} else {
			output[k] = a2[j]
			j++
		}
	}

	return output
}

// MergeSort is used to sort an array of integer
func MergeSort(input []int) []int {
	// Array with 1 element already sorted
	if len(input) <= 1 {
		return input
	}

	left := input[:len(input)/2]
	right := input[len(input)/2:]
	return merge(MergeSort(left), MergeSort(right))
}
