package sortcharacter

import "testing"

func TestSortCharacter(t *testing.T) {
	testCases := []struct {
		desc      string
		input     string
		vowel     string
		consonant string
	}{
		{
			desc:      "test",
			input:     "Sample Case",
			vowel:     "aaee",
			consonant: "ssmplc",
		}, {
			desc:      "test",
			input:     "Next Case",
			vowel:     "eeaa",
			consonant: "nxtcs",
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

		})
	}
}
