package fingering

import (
	"github.com/dustmason/guitar-plugin/voicing"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFingerings(t *testing.T) {
	singleNoteVoicing := voicing.Voicing{5, 0, 0, 0, 0, 0}
	fingerings := GenerateFingerings(singleNoteVoicing)
	assert.True(t, len(fingerings) > 0)

	threeNoteVoicing := voicing.Voicing{3, 2, 0, 0, 0, 0}
	fingerings = GenerateFingerings(threeNoteVoicing)
	assert.True(t, len(fingerings) > 0)

	sixNoteVoicing := voicing.Voicing{3, 2, 0, 4, 5, 3}
	fingerings = GenerateFingerings(sixNoteVoicing)
	assert.True(t, len(fingerings) > 0)
}
