package fingering

import (
	"testing"
)

func TestLevenshteinDistance(t *testing.T) {
	testCases := []struct {
		name     string
		chordA   Fingering
		chordB   Fingering
		expected int
	}{
		{
			name:     "empty chords",
			chordA:   Fingering{},
			chordB:   Fingering{},
			expected: 0,
		},
		{
			name:     "one empty chord",
			chordA:   Fingering{{3, 1}, {5, 3}},
			chordB:   Fingering{},
			expected: 2,
		},
		{
			name:     "same chords",
			chordA:   Fingering{{3, 1}, {5, 3}},
			chordB:   Fingering{{3, 1}, {5, 3}},
			expected: 0,
		},
		{
			name:     "different chords",
			chordA:   Fingering{{3, 1}, {5, 3}},
			chordB:   Fingering{{2, 1}, {4, 3}},
			expected: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := fingeringDistance(tc.chordA, tc.chordB)
			if result != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, result)
			}
		})
	}
}
