package voicing

func (chord Voicing) Spread() float64 {
	trimmedChord := trim(chord)
	notesCount := notes(chord)
	if trimmedChord == 0 || notesCount == 0 {
		return 0
	}
	return float64(trimmedChord) / float64(notesCount)
}

func trim(chord Voicing) int {
	chordSlice := chord[:]
	for len(chordSlice) > 0 && chordSlice[0] <= 0 {
		chordSlice = chordSlice[1:]
	}
	for len(chordSlice) > 0 && chordSlice[len(chordSlice)-1] <= 0 {
		chordSlice = chordSlice[:len(chordSlice)-1]
	}
	return len(chordSlice)
}

func notes(chord Voicing) int {
	count := 0
	for _, note := range chord {
		if note > -1 {
			count++
		}
	}
	return count
}
