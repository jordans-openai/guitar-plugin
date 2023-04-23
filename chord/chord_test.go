package chord

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewChord(t *testing.T) {
	testCases := []struct {
		chord   string
		root    string
		on      string
		visible bool
		notes   []string
	}{
		{
			chord:   "Cmaj7",
			root:    "C",
			on:      "",
			visible: true,
			notes:   []string{"C", "E", "G", "B"},
		},
		{
			chord:   "Cm7",
			root:    "C",
			on:      "",
			visible: true,
			notes:   []string{"C", "Eb", "G", "Bb"},
		},
		{
			chord:   "Dm7/G",
			root:    "D",
			on:      "G",
			visible: true,
			notes:   []string{"D", "F", "A", "C", "G"},
		},
		{
			chord:   "F#maj7",
			root:    "F#",
			on:      "",
			visible: true,
			notes:   []string{"F#", "A#", "C#", "F"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.chord, func(t *testing.T) {
			chord, err := NewChord(tc.chord)
			assert.NoError(t, err)

			assert.Equal(t, tc.root, chord.root)
			assert.Equal(t, tc.on, chord.on)

			components, err := chord.ComponentNames()
			assert.NoError(t, err)
			assert.Equal(t, tc.notes, components)
		})
	}
}
