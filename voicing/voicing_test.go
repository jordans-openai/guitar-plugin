package voicing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVoicings(t *testing.T) {
	// empty list
	notes := []int{}
	voicings := Voicings(notes)
	assert.Equal(t, 0, len(voicings))

	// a single note
	notes = []int{0}
	voicings = Voicings(notes)
	assert.True(t, len(voicings) > 0)

	// required notes
	notes = []int{0, 4, 7} // C major triad aka 0,4,7
	voicings = Voicings(notes)
	assert.True(t, len(voicings) > 0)

	// optional notes
	notes = []int{0, 4, -7}
	voicings = Voicings(notes)
	assert.True(t, len(voicings) > 0)

	// more than six required notes
	notes = []int{0, 2, 4, 5, 7, 9, 11}
	voicings = Voicings(notes)
	assert.Equal(t, 0, len(voicings))

	// duplicate notes
	notes = []int{0, 0, 4, 7}
	voicings = Voicings(notes)
	assert.True(t, len(voicings) > 0)
}

func TestGenerateAllNoteSets(t *testing.T) {
	notes := []int{0, 4, 7, 11}
	voicings, _ := generateAllNoteSets(notes)

	// ensure all the sets have the required notes
	for _, v := range voicings {
		containsRequiredNotes := true
		for _, note := range notes {
			containsNote := false
			for _, finger := range v {
				if finger%12 == note || finger == note {
					containsNote = true
					break
				}
			}
			if !containsNote {
				containsRequiredNotes = false
				break
			}
		}
		assert.True(t, containsRequiredNotes)
	}
}
