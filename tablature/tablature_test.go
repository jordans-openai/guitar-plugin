package tablature

import (
	"github.com/dustmason/guitar-plugin/fingering"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	emptyChord := Tablature{fingering: fingering.Fingering{
		{-1, 0},
		{-1, 0},
		{-1, 0},
		{-1, 0},
		{-1, 0},
		{-1, 0},
	}, chordName: ""}
	expected := " 0 ┬┬┬┬┬┬\n   ├┼┼┼┼┤\n   ││││││\n   ├┼┼┼┼┤\n   ││││││\n   ├┼┼┼┼┤\n   ││││││\n         "
	assert.Equal(t, expected, emptyChord.String())

	singleNoteChord := Tablature{fingering: fingering.Fingering{
		{0, 1},
		{-1, 0},
		{-1, 0},
		{-1, 0},
		{-1, 0},
		{-1, 0},
	}, chordName: "E"}
	expected = " 0 ●┬┬┬┬┬ E\n   ├┼┼┼┼┤\n   ││││││\n   ├┼┼┼┼┤\n   ││││││\n   ├┼┼┼┼┤\n   ││││││\n   1     "
	assert.Equal(t, expected, singleNoteChord.String())

	threeNoteChord := Tablature{fingering: fingering.Fingering{
		{0, 1},
		{2, 2},
		{2, 3},
		{-1, 0},
		{-1, 0},
		{-1, 0},
	}, chordName: "E"}
	expected = " 0 ●┬┬┬┬┬ E\n   ├┼┼┼┼┤\n   ││││││\n   ├┼┼┼┼┤\n   │●●│││\n   ├┼┼┼┼┤\n   ││││││\n   123   "
	assert.Equal(t, expected, threeNoteChord.String())

	sixNoteChord := Tablature{fingering: fingering.Fingering{
		{3, 2},
		{2, 1},
		{0, 0},
		{0, 0},
		{3, 3},
		{3, 4},
	}, chordName: "G"}
	expected = " 0 ┬┬●●┬┬ G\n   ├┼┼┼┼┤\n   ││││││\n   ├┼┼┼┼┤\n   │●││││\n   ├┼┼┼┼┤\n   ●│││●●\n   21  34"
	assert.Equal(t, expected, sixNoteChord.String())

	barChord := Tablature{fingering: fingering.Fingering{
		{5, 1},
		{7, 2},
		{7, 3},
		{6, 4},
		{5, 1},
		{5, 1},
	}, chordName: "A"}
	expected = " 5 ●│││●● A\n   ├┼┼┼┼┤\n   │││●││\n   ├┼┼┼┼┤\n   │●●│││\n   ├┼┼┼┼┤\n   ││││││\n   123411"
	assert.Equal(t, expected, barChord.String())
}
