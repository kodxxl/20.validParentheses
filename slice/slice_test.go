package slice

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		got  string
		want bool
	}{
		{"(){}[]", true},
		{"({[]})", true},
		{"({[}))", false},
		{"]", false},
	}

	for i, test := range tests {
		if test.want != IsValid(test.got) {
			t.Errorf("Invalid algorythm on %d", i)
		}
	}
}
