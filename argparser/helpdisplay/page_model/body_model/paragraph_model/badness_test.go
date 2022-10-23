package paragraph_model

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBedness(t *testing.T) {
	t.Parallel()

	b := makeBadness(10, 1, 45)
	require.False(t, b.isBadnessByOptimumDropped())
	b.dropBadnessByOptimum()
	require.True(t, b.isBadnessByOptimumDropped())
	b.update(50)
	require.Equal(t, 40, b.getBadByOptimum())
}

func TestBadnessWorse(t *testing.T) {
	t.Parallel()

	testData := []struct {
		caseData string
		b1       badness
		b2       badness
		worse    bool
	}{
		{
			caseData: "b1_equal_b2",
			b1:       makeBadness(10, 0, 20),
			b2:       makeBadness(10, 0, 20),
			worse:    true,
		},
		{
			caseData: "b1_less_b2",
			b1:       makeBadness(10, 0, 20),
			b2:       makeBadness(20, 0, 20),
			worse:    true,
		},

		{
			caseData: "negative_aboveRowBadnessByOptimum_1",
			b1:       makeBadness(30, -1, 20),
			b2:       makeBadness(10, -1, 20),
			worse:    false,
		},
		{
			caseData: "negative_aboveRowBadnessByOptimum_2",
			b1:       makeBadness(30, -1, 20),
			b2:       makeBadness(30, -1, 20),
			worse:    true,
		},
		{
			caseData: "negative_aboveRowBadnessByOptimum_3",
			b1:       makeBadness(30, -20, 20),
			b2:       makeBadness(30, -1, 20),
			worse:    false,
		},
		{
			caseData: "negative_aboveRowBadnessByOptimum_4",
			b1:       makeBadness(10, -20, 20),
			b2:       makeBadness(30, -1, 20),
			worse:    false,
		},

		{
			caseData: "positive_aboveRowBadnessByOptimum_and_b1_equal_b2",
			b1:       makeBadness(30, 1, 20),
			b2:       makeBadness(30, 1, 20),
			worse:    true,
		},
		{
			caseData: "positive_aboveRowBadnessByOptimum_1",
			b1:       makeBadness(30, 1, 20),
			b2:       makeBadness(5, 1, 20),
			worse:    false,
		},
		{
			caseData: "positive_aboveRowBadnessByOptimum_2",
			b1:       makeBadness(30, 1, 20),
			b2:       makeBadness(35, 1, 20),
			worse:    false,
		},
		{
			caseData: "positive_aboveRowBadnessByOptimum_1",
			b1:       makeBadness(10, 1, 20),
			b2:       makeBadness(5, 1, 20),
			worse:    false,
		},
		{
			caseData: "positive_aboveRowBadnessByOptimum_1",
			b1:       makeBadness(10, 25, 20),
			b2:       makeBadness(5, 1, 20),
			worse:    false,
		},
	}

	for _, td := range testData {
		t.Run(td.caseData, func(t *testing.T) {
			require.Equal(t, td.worse, td.b1.worse(td.b2))
		})
	}
}
