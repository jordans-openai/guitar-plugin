package fingering

import "github.com/dustmason/guitar-plugin/voicing"

// Fingering represents which fret each finger should be on to produce a chord. each index is a string.
// each item is a fret, finger pair
type Fingering [][2]int

func GenerateFingerings(chord voicing.Voicing) []Fingering {
	chordWithFingers := attachPossibleFingers(chord)
	fingerings := generateFingeringsInner(chordWithFingers, Fingering{}, []Fingering{})
	return fingerings
}

func attachPossibleFingers(chord voicing.Voicing) [][]int {
	var chordWithFingers [][]int
	for _, fret := range chord {
		chordWithFingers = append(chordWithFingers, []int{fret, 0, 1, 2, 3, 4})
	}
	return chordWithFingers
}

func generateFingeringsInner(chord [][]int, current Fingering, results []Fingering) []Fingering {
	if len(chord) == 0 {
		results = append(results, current)
		return results
	}

	head, rest := chord[0], chord[1:]
	for _, finger := range head[1:] {
		if !isInvalidFingering(head[0], finger, current) {
			newFingering := append(Fingering{}, current...)
			newFingering = append(newFingering, [2]int{head[0], finger})
			results = generateFingeringsInner(rest, newFingering, results)
		}
	}

	return results
}

func isInvalidFingering(fret, finger int, fingerings Fingering) bool {
	if len(fingerings) == 0 {
		return false
	}

	newFret, newFinger := fingerings[len(fingerings)-1][0], fingerings[len(fingerings)-1][1]
	return (fret > newFret && finger < newFinger) ||
		(fret < newFret && finger > newFinger) ||
		(finger == newFinger && fret != newFret) ||
		playableBar(fret, finger, fingerings)
}

func playableBar(fret, finger int, fingerings Fingering) bool {
	newFret := fingerings[len(fingerings)-1][0]

	for _, f := range fingerings {
		if f[0] > newFret && f[1] == finger {
			return true
		}
	}
	return false
}
