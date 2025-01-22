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

func TestRemoveSpace(t *testing.T) {
	testCases := []struct {
		desc      string
		input     string
		ans       string
		consonant string
	}{
		{
			desc:  "test",
			input: "Sample Case",
			ans:   "SampleCase",
		}, {
			desc:  "test",
			input: "Sample Case case",
			ans:   "SampleCasecase",
		}, {
			desc:  "test",
			input: "one two third forth",
			ans:   "onetwothirdforth",
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			res := removeSpace(tC.input)
			t.Log("res:`", res)
			if res != tC.ans {
				t.Fail()
			}
		})
	}
}
