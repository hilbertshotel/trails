package lib

import "testing"

type Test struct {
	Dates    []string
	Expected int
}

func TestParseStreak(t *testing.T) {

	tests := map[string]Test{
		"test 0": {[]string{"12 Apr 22", "13 Apr 22"}, 2},
		"test 1": {[]string{"12 Apr 22", "13 May 22"}, 1},
		"test 3": {[]string{"12 Apr 22", "15 Apr 22", "16 Apr 22"}, 2},
		"test 4": {[]string{"12 Apr 22", "13 Apr 22", "16 Apr 22"}, 2},
		"test 5": {[]string{"14 Apr 22", "15 Apr 22", "16 Apr 22"}, 3},
		"test 6": {[]string{"30 Apr 22", "01 May 22"}, 2},
		"test 7": {[]string{"30 Apr 22", "01 May 22", "12 May 22", "12 May 22"}, 2},
		"test 8": {[]string{"01 May 22", "01 May 22", "02 May 22", "02 May 22"}, 4},
		"test 9": {[]string{
			"30 Apr 22",
			"01 May 22",
			"12 May 22",
			"12 May 22",
			"12 May 22",
			"13 May 22",
			"13 May 22",
			"28 May 22",
			"29 May 22",
			"30 May 22",
			"31 May 22",
			"01 Jun 22",
			"02 Jun 22",
		}, 6},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := ParseStreak(test.Dates)
			if result != test.Expected {
				t.Errorf("expected: %v, got: %v", test.Expected, result)
			}
		})
	}

}
