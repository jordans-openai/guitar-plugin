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

	cost := noteDistance(noteA, noteB, chordA, chordB)

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
		if note[1] != 0 { // todo note[1] is the finger - is this right?
			count++
		}
	}
	return count
}

func noteDistance(noteA, noteB [2]int, chordA, chordB Fingering) int {
	if noteA == noteB {
		return 0
	}
	if noteA[0] == 0 {
		return 1
	}
	if noteB[0] == 0 {
		return 1
	}
	stringA := indexOf(noteA, chordA)
	stringB := indexOf(noteB, chordB)
	if stringA == stringB {
		return 1
	}
	return int(math.Abs(float64(noteA[0]-noteB[0])) + math.Abs(float64(stringA-stringB)))
}

func indexOf(note [2]int, chord Fingering) int {
	for i, n := range chord {
		if n == note {
			return i
		}
	}
	return -1
}
