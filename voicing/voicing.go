package voicing

import (
	"fmt"
	"github.com/dustmason/guitar-plugin/permutation"
	"strings"
)

var openStringNotes = []int{40, 45, 50, 55, 59, 64}

type Voicing [6]int

func Voicings(notes []int) []Voicing {
	if len(notes) == 0 {
		return []Voicing{}
	}
	allNoteSets, requiredNotes := generateAllNoteSets(notes)
	var voicings []Voicing
	for _, noteSet := range allNoteSets {
		voicings = buildChords(noteSet, Voicing{}, voicings, requiredNotes)
	}
	return removeDuplicates(voicings)
}

func generateAllNoteSets(notes []int) ([][]int, map[int]bool) {
	requiredNotes := make(map[int]bool)
	_allNotes := make([]int, 0)

	for _, note := range notes {
		if note >= 0 { // negative numbers are "optional" notes, but we ignore them for now
			requiredNotes[note] = true
		}
		_allNotes = append(_allNotes, abs(note))
	}

	if len(requiredNotes) > 6 {
		return [][]int{}, map[int]bool{}
	}

	k := 6 - len(requiredNotes)
	perms := permutation.Generate(_allNotes, &k, true)
	noteSets := make([][]int, 0)

	_reqNotes := make([]int, 0)
	for note := range requiredNotes {
		_reqNotes = append(_reqNotes, note)
	}
	for _, tail := range perms {
		noteSet := append([]int{}, _reqNotes...)
		noteSet = append(noteSet, tail...)
		noteSets = append(noteSets, noteSet)
	}

	return noteSets, requiredNotes
}

func buildChords(noteSet []int, chord Voicing, chords []Voicing, requiredNotes map[int]bool) []Voicing {
	type stackItem struct {
		noteSet []int
		chord   Voicing
	}

	stack := []stackItem{{noteSet: noteSet, chord: chord}}
	seen := make(map[string]bool)

	for len(stack) > 0 {
		item := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		_noteSet, _chord := item.noteSet, item.chord

		if len(_noteSet) == 0 {
			key := voicingKey(_chord)
			if _, exists := seen[key]; !exists {
				seen[key] = true
				chords = append(chords, _chord)
			}
			continue
		}

		note, rest := _noteSet[0], _noteSet[1:]
		for _, position := range allNotes(note) {
			stringIndex, fret := position[0], position[1]
			if fret == 0 && isWrongOpenNote(stringIndex, requiredNotes) {
				continue
			}
			minFret, maxFret := minMaxFretWithNewPosition(_chord, stringIndex, fret)
			if maxFret-minFret <= 4 {
				chord[stringIndex] = fret
				stack = append(stack, stackItem{noteSet: rest, chord: chord})
			}
		}
	}

	return chords
}

func isWrongOpenNote(stringIndex int, requiredNotes map[int]bool) bool {
	ok, _ := requiredNotes[openStringNotes[stringIndex]%12]
	return !ok
}

func minMaxFretWithNewPosition(chord Voicing, stringIndex, fret int) (int, int) {
	minFret := 1<<63 - 1
	maxFret := -1 << 63

	for i, f := range chord {
		if i == stringIndex {
			f = fret
		}
		if f < minFret {
			minFret = f
		}
		if f > maxFret {
			maxFret = f
		}
	}

	return minFret, maxFret
}

func removeDuplicates(voicings []Voicing) []Voicing {
	uniqueVoicings := make([]Voicing, 0)
	seen := make(map[string]bool)

	for _, voicing := range voicings {
		key := voicingKey(voicing)
		if _, exists := seen[key]; !exists {
			seen[key] = true
			uniqueVoicings = append(uniqueVoicings, voicing)
		}
	}

	return uniqueVoicings
}

func voicingKey(voicing Voicing) string {
	keys := make([]string, len(voicing))

	for i, finger := range voicing {
		keys[i] = fmt.Sprintf("%d", finger)
	}

	return strings.Join(keys, ",")
}

var notePositionsCache = make(map[int][][2]int)

func allNotes(targetNote int) [][2]int {
	if positions, ok := notePositionsCache[targetNote]; ok {
		return positions
	}
	frets := 18
	var positions [][2]int
	for fret := 0; fret <= frets; fret++ {
		for stringIndex, openNote := range openStringNotes {
			note := openNote + fret

			if note%12 == targetNote || note == targetNote {
				positions = append(positions, [2]int{stringIndex, fret})
			}
		}
	}
	notePositionsCache[targetNote] = positions
	return positions
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
