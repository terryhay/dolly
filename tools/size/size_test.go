package size

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

			require.Equal(t, tc.exp, MakeWidth(tc.valueInit).Int())
			require.Equal(t, tc.exp, MakeHeight(tc.valueInit).Int())
		})
	}
}

func TestAppendHeight(t *testing.T) {
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

			require.Equal(t, tc.exp, AppendHeight(MakeHeight(tc.valueInit), tc.valueAppend).Int())
		})
	}
}

func TestDif(t *testing.T) {
	t.Parallel()

	l := math.MaxUint8 + int(gofakeit.Uint8())
	r := int(gofakeit.Uint8())

	tests := []struct {
		caseName string
		l        int
		r        int
		exp      int
	}{
		{
			caseName: "common",
			l:        l,
			r:        r,
			exp:      l - r,
		},
		{
			caseName: "zero",
			l:        l,
			r:        2*math.MaxUint8 + r,
			exp:      0,
		},
	}

	for _, testCase := range tests {
		tc := testCase

		t.Run(tc.caseName, func(t *testing.T) {
			t.Parallel()

			require.Equal(t, MakeWidth(tc.exp), Dif(MakeWidth(tc.l), MakeWidth(tc.r)))
			require.Equal(t, MakeHeight(tc.exp), Dif(MakeHeight(tc.l), MakeHeight(tc.r)))
		})
	}
}

func TestIndex(t *testing.T) {
	t.Parallel()

	valueInit := int(gofakeit.Uint8())
	require.Equal(t, valueInit, MakeHeight(valueInit).Index().Int())
}
