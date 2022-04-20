package coins

var partitions []int

// Returns next element in sequence 1, -1, 2, -2, 3, -3, ...
func nextK(k int) int {
	if k > 0 {
		return -k
	} else {
		return -k + 1
	}
}

// Generalized pentagonal number https://en.wikipedia.org/wiki/Pentagonal_number_theorem
func gpn(k int) int {
	return k * (3*k - 1) / 2
}

// Number partitions https://en.wikipedia.org/wiki/Partition_(number_theory)
func Partition(n int) int {
	// Memoization
	if n < len(partitions) && partitions[n] != 0 {
		return partitions[n]
	}

	var sign int

	s := 0
	k := 1
	for n >= gpn(k) {
		if k%2 == 0 {
			sign = -1
		} else {
			sign = 1
		}

		s += sign * Partition(n-gpn(k))
		k = nextK(k)
	}

	if n < len(partitions) {
		partitions[n] = s
	}
	return s
}

func Piles(n int) int {
	partitions = make([]int, n+2)
	partitions[0] = 1
	partitions[1] = 1

	return Partition(n)
}
