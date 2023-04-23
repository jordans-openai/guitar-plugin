package tablature

import (
	"fmt"
	"github.com/dustmason/guitar-plugin/chord"
	"github.com/dustmason/guitar-plugin/fingering"
	"github.com/dustmason/guitar-plugin/voicing"
	"os"
	"strings"
)

const minLines = 3

type Tablature struct {
	fingering fingering.Fingering
	chordName string
}

// NewTablature is the entry point for the tablature package. It accepts a chord name and returns
// a Tablature which can be used to print a tablature for the chord.
func NewTablature(chordName string) (Tablature, error) {
	c, err := chord.NewChord(chordName)
	if err != nil {
		return Tablature{}, err
	}
	components, err := c.Components()
	if err != nil {
		return Tablature{}, err
	}
	voicings := voicing.Voicings(components)
	fingerings := fingering.GenerateFingerings(voicings[0])

	if os.Getenv("DEV") != "" {
		componentNames, _ := c.ComponentNames()
		fmt.Println("component names", componentNames, "components", components, "c", chordName)
		for i, v := range voicings {
			fmt.Println(" - v", i, v)
		}
		fmt.Println("fingerings for c", chordName, fingerings[0])
	}

	return Tablature{fingering: fingerings[0], chordName: chordName}, nil
}

func (c Tablature) String() string {
	var min, max int

	if len(c.fingering) > 0 {
		min, max = c.fingering[0].Fret, c.fingering[0].Fret
		for _, fret := range c.fingering {
			if fret.Fret < min {
				min = fret.Fret
			}
			if fret.Fret > max {
				max = fret.Fret
			}
		}
	}
	if min < 0 {
		min = 0
	}
	if max < 0 {
		max = 0
	}

	frets := make([]int, len(c.fingering))
	fingers := make([]int, len(c.fingering))
	for i, f := range c.fingering {
		frets[i] = f.Fret
		if f.Finger > 0 {
			fingers[i] = f.Finger
		}
	}

	lim := maxInt(max-min, minLines)
	var lines []string
	for i := 0; i <= lim; i++ {
		lines = append(lines, rowToString(i, min, frets, c.chordName))
		if i != lim {
			lines = append(lines, "   ├┼┼┼┼┤")
		}
	}

	lines = appendFingering(lines, fingers)
	return strings.Join(lines, "\n")
}

func rowToString(offset, base int, frets []int, chordName string) string {
	var row []string
	row = append(row, leftGutter(offset, base+offset))
	for _, fret := range frets {
		row = append(row, fretToString(fret, base+offset))
	}
	row = append(row, rightGutter(offset, chordName))
	return strings.Join(row, "")
}

func fretToString(fret int, fretNumber int) string {
	if fret == fretNumber && fret > -1 {
		return "●"
	} else if fretNumber == 0 {
		return "┬"
	} else {
		return "│"
	}
}

func leftGutter(offset, fret int) string {
	if offset == 0 {
		return fmt.Sprintf("%2d ", fret)
	}
	return "   "
}

func rightGutter(offset int, chordName string) string {
	if offset == 0 && chordName != "" {
		return fmt.Sprintf(" %s", chordName)
	}
	return ""
}

func appendFingering(lines []string, fingers []int) []string {
	var f []string
	f = append(f, "   ")
	for _, fret := range fingers {
		if fret == 0 {
			f = append(f, " ")
		} else {
			f = append(f, fmt.Sprint(fret))
		}
	}
	return append(lines, strings.Join(f, ""))
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
