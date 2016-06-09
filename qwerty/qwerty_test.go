package qwerty

import (
	"math"
	"testing"
)

var tolerance float64 = 0.0005

func TestQuertyDistance(t *testing.T) {
	compare("a", "a", 0, t)
	compare("a", "s", 1, t)

	compare("a", "w", 1.414214, t)
	compare("a", "x", 1.414214, t)

	compare("a", "d", 2, t)
	compare("a", "c", 2.236068, t)

	compare("a", "g", 4, t)
	compare("a", "b", 4.123106, t)

	compare("a", "h", 5, t)
	compare("a", "n", 5.099020, t)

	compare("a", "p", 9.055385, t)
	compare("a", "l", 8, t)
	compare("a", "m", 6.082763, t)
}

func TestQwertyCompareNumbers(t *testing.T) {
	compare("1", "1", 0, t)
	compare("1", "3", 2, t)
	compare("1", "9", 8, t)
}

func TestQwertyLongerStrings(t *testing.T) {
	compare("food", "food", 0, t)

	compare("foof", "food", 1, t)
	compare("food", "foof", 1, t)

	compare("food", "blue", 5.414214, t)
	compare("food", "blur", 5.828427, t)
}
func compare(s1, s2 string, answer float64, t *testing.T) {
	res := Compare(s1, s2)
	if math.Abs(res - answer) > tolerance {
		t.Fatalf("got %f, but expected %f, s1='%s', s2='%s'", res, answer, s1, s2)
	}
}
