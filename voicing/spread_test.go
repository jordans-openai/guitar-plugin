package voicing

import (
	"testing"
)

func TestSpread(t *testing.T) {
	testCases := []struct {
		name     string
		chord    Voicing
		expected float64
	}{
		{
			name:     "empty chord",
			chord:    Voicing{0, 0, 0, 0, 0, 0},
			expected: 0,
		},
		{
			name:     "single note",
			chord:    Voicing{8, -1, -1, -1, -1, -1},
			expected: 1,
		},
		{
			name:     "non-spread chord",
			chord:    Voicing{8, 7, 9, -1, -1, -1},
			expected: 1,
		},
		{
			name:     "spread chord",
			chord:    Voicing{8, -1, -1, -1, 12, 12},
			expected: 2,
		},
		{
			name:     "very spread chord",
			chord:    Voicing{1, -1, -1, -1, -1, 1},
			expected: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.chord.Spread()
			if result != tc.expected {
				t.Errorf("expected %f, got %f", tc.expected, result)
			}
		})
	}
}
