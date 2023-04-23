package permutation

func Generate(list []int, k *int, repetitions bool) [][]int {
	if len(list) == 0 || (k != nil && *k == 0) {
		return [][]int{[]int{}}
	}

	if k == nil {
		k = new(int)
		*k = -1
	}

	var perms [][]int

	for _, head := range list {
		tailList := nextList(list, head, repetitions)
		nextK := nextK(k)

		for _, tail := range Generate(tailList, nextK, repetitions) {
			perm := append([]int{head}, tail...)
			perms = append(perms, perm)
		}
	}

	return perms
}

func nextK(k *int) *int {
	if *k < 0 {
		return k
	}

	nextK := *k - 1
	return &nextK
}

func nextList(list []int, head int, repetitions bool) []int {
	if repetitions {
		return list
	}

	remaining := make([]int, 0, len(list)-1)
	for _, v := range list {
		if v != head {
			remaining = append(remaining, v)
		}
	}

	return remaining
}
