package voicing

import (
	"testing"
)

func TestCalculateReach(t *testing.T) {
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
			name:     "non-stretched chord",
			chord:    Voicing{8, 0, 8, 8, 8, 0},
			expected: 0,
		},
		{
			name:     "stretched chord",
			chord:    Voicing{0, 10, 10, 8, 11, 0},
			expected: 0.6,
		},
		{
			name:     "chord with negative values",
			chord:    Voicing{-1, 10, 10, 8, 11, -1},
			expected: 0.6,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.chord.Reach()
			if result != tc.expected {
				t.Errorf("expected %f, got %f", tc.expected, result)
			}
		})
	}
}
