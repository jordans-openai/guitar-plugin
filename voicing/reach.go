package voicing

func (chord Voicing) Reach() float64 {
	chordInts := []int{}
	for _, v := range chord {
		if v > 0 {
			chordInts = append(chordInts, v)
		}
	}
	if len(chordInts) == 0 {
		return 0
	}
	min, max := minMax(chordInts)
	return float64(max-min) / 5
}

func minMax(arr []int) (int, int) {
	min, max := arr[0], arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		} else if v > max {
			max = v
		}
	}
	return min, max
}
