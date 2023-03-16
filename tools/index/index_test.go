package index

import (
	"math"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
)

func TestIndexInit(t *testing.T) {
	t.Parallel()

	randNegative := -1 - int(gofakeit.Uint8())
	randPositive := 1 + int(gofakeit.Uint8())

	tests := []struct {
		caseName  string
		valueInit int
		exp       int
	}{
		{
			caseName:  "init_by_negative_value",
			valueInit: randNegative,
			exp:       0,
		},
		{
			caseName:  "init_by_positive_value",
			valueInit: randPositive,
			exp:       randPositive,
		},
	}

	for _, testCase := range tests {
		tc := testCase

		t.Run(tc.caseName, func(t *testing.T) {
			t.Parallel()

			require.Equal(t, tc.exp, MakeIndex(tc.valueInit).Int())
		})
	}
}

func TestIndexInc(t *testing.T) {
	t.Parallel()

	randPositive := 1 + int(gofakeit.Uint8())

	tests := []struct {
		caseName  string
		valueInit int
		exp       int
	}{
		{
			caseName:  "common",
			valueInit: randPositive,
			exp:       randPositive + 1,
		},
		{
			caseName:  "overflow",
			valueInit: math.MaxInt32,
			exp:       0,
		},
	}

	for _, testCase := range tests {
		tc := testCase

		t.Run(tc.caseName, func(t *testing.T) {
			t.Parallel()

			require.Equal(t, tc.exp, Inc(MakeIndex(tc.valueInit)).Int())
		})
	}
}

func TestAppend(t *testing.T) {
	t.Parallel()

	randInit := math.MaxUint8 + int(gofakeit.Uint8())
	randPositive := int(gofakeit.Uint8())
	randNegative := -1 - int(gofakeit.Uint8())

	tests := []struct {
		caseName    string
		valueInit   int
		valueAppend int
		exp         int
	}{
		{
			caseName:    "positive",
			valueInit:   randInit,
			valueAppend: randPositive,
			exp:         randInit + randPositive,
		},
		{
			caseName:    "negative",
			valueInit:   randInit,
			valueAppend: randNegative,
			exp:         randInit + randNegative,
		},
		{
			caseName:    "negative_to_zero",
			valueInit:   int(gofakeit.Uint8()),
			valueAppend: -math.MaxUint8 + randNegative,
			exp:         0,
		},
	}

	for _, testCase := range tests {
		tc := testCase

		t.Run(tc.caseName, func(t *testing.T) {
			t.Parallel()

			require.Equal(t, tc.exp, Append(MakeIndex(tc.valueInit), tc.valueAppend).Int())
		})
	}
}
