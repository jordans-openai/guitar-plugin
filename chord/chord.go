package chord

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Chord struct {
	chord      string
	root       string
	components []int
	appended   []string
	on         string
}

func NewChord(chord string) (Chord, error) {
	root, components, on, err := parse(chord)
	if err != nil {
		return Chord{}, err
	}

	c := Chord{
		chord:      chord,
		root:       root,
		components: components,
		on:         on,
	}

	c.appendOnChord()

	return c, nil
}

func (c *Chord) Components() ([]int, error) {
	rootVal, err := noteToVal(c.root)
	if err != nil {
		return nil, err
	}

	components := make([]int, len(c.components))
	for i, v := range c.components {
		comp := (v + rootVal) % 12
		components[i] = comp
	}

	return components, nil
}

func (c *Chord) ComponentNames() ([]string, error) {
	components, err := c.Components()
	if err != nil {
		return nil, err
	}

	names := make([]string, len(components))
	for i, comp := range components {
		names[i] = valToNote(comp, c.root)
	}

	return names, nil
}

func (c *Chord) appendOnChord() {
	if c.on != "" {
		onVal, err := noteToVal(c.on)
		if err != nil {
			panic(err)
		}
		rootVal, err := noteToVal(c.root)
		if err != nil {
			panic(err)
		}
		c.components = append(c.components, (onVal-rootVal+12)%12)
	}
}

var inversionRe = regexp.MustCompile("/([0-9]+)")

func parse(chord string) (string, []int, string, error) {
	root, rest := "", ""
	if len(chord) > 1 && strings.Contains("b#", string(chord[1])) {
		root = chord[:2]
		rest = chord[2:]
	} else {
		root = chord[:1]
		rest = chord[1:]
	}

	if _, ok := notesMap[root]; !ok {
		return "", nil, "", fmt.Errorf("invalid note %s", root)
	}

	inversion := 0
	if inversionRe.MatchString(rest) {
		inversionStr := inversionRe.FindStringSubmatch(rest)[1]
		inversion, _ = strconv.Atoi(inversionStr)
		rest = inversionRe.ReplaceAllString(rest, "")
	}

	onChordIdx := strings.Index(rest, "/")
	on := ""
	if onChordIdx >= 0 {
		on = rest[onChordIdx+1:]
		rest = rest[:onChordIdx]
		if _, ok := notesMap[on]; !ok {
			return "", nil, "", fmt.Errorf("invalid note %s", on)
		}
	}

	components, err := getQuality(rest, inversion)
	if err != nil {
		return "", nil, "", err
	}
	return root, components, on, nil
}

func getQuality(name string, inversion int) ([]int, error) {
	q, ok := defaultQualities[name]
	if !ok {
		return nil, fmt.Errorf("unknown quality: %s", name)

	}

	components := make([]int, len(q))
	copy(components, q)

	for i := 0; i < inversion; i++ {
		n := components[0]
		for n < components[len(components)-1] {
			n += 12
		}
		components = append(components[1:], n)
	}

	return components, nil
}

func noteToVal(note string) (int, error) {
	val, ok := notesMap[note]
	if !ok {
		return 0, fmt.Errorf("unknown note %s", note)
	}
	return val, nil
}

func valToNote(val int, scale string) string {
	scaleDict, ok := scaleValDict[scale]
	if !ok {
		panic(fmt.Errorf("unknown scale %s", scale))
	}
	return scaleDict[val%12]
}
