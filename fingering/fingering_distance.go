package fingering

import (
	"math"
)

func fingeringDistance(chordA, chordB Fingering) int {
	if len(chordA) == 0 {
		return countFingersUsed(chordB)
	}
	if len(chordB) == 0 {
		return countFingersUsed(chordA)
	}

	noteA := chordA[0]
	noteB := chordB[0]
	cost := noteDistance(noteA, noteB)
	return int(math.Min(
		float64(fingeringDistance(chordA[1:], chordB))+1,
		math.Min(
			float64(fingeringDistance(chordA, chordB[1:]))+1,
			float64(fingeringDistance(chordA[1:], chordB[1:]))+float64(cost),
		),
	))
}

func countFingersUsed(chord Fingering) int {
	count := 0
	for _, note := range chord {
		if note.Finger != 0 {
			count++
		}
	}
	return count
}

func noteDistance(noteA, noteB GuitarString) int {
	if noteA == noteB {
		return 0
	}
	if noteA.Fret == 0 {
		return 1
	}
	if noteB.Fret == 0 {
		return 1
	}
	return int(math.Abs(float64(noteA.Fret - noteB.Fret)))
}
