package qwerty

import (
	"math"
)

type Location struct {
	x float64
	y float64
}

var distances = map[rune]Location{
	// number row
	'0': Location{0,0},
	'1': Location{1,0},
	'2': Location{2,0},
	'3': Location{3,0},
	'4': Location{4,0},
	'5': Location{5,0},
	'6': Location{6,0},
	'7': Location{7,0},
	'8': Location{8,0},
	'9': Location{9,0},

	// first row
	'q': Location{0,1},
	'w': Location{1,1},
	'e': Location{2,1},
	'r': Location{3,1},
	't': Location{4,1},
	'y': Location{5,1},
	'u': Location{6,1},
	'i': Location{7,1},
	'o': Location{8,1},
	'p': Location{9,1},

	// second row
	'a': Location{0,2},
	's': Location{1,2},
	'd': Location{2,2},
	'f': Location{3,2},
	'g': Location{4,2},
	'h': Location{5,2},
	'j': Location{6,2},
	'k': Location{7,2},
	'l': Location{8,2},

	// third row
	'z': Location{0,3},
	'x': Location{1,3},
	'c': Location{2,3},
	'v': Location{3,3},
	'b': Location{4,3},
	'n': Location{5,3},
	'm': Location{6,3},
}

func Compare(s1, s2 string) float64 {
	// todo: We can likely compare with just another string metric
	if len(s1) != len(s2) {
		return -1
	}

	distance := float64(0)

	for i := range s1 {
		if s1[i] != s2[i] {
			distance += euclideanDistance(rune(s1[i]), rune(s2[i]))
		}
	}

	return distance
}

func euclideanDistance(s1, s2 rune) float64 {
	x := math.Pow(distances[s1].x - distances[s2].x, 2)
	y := math.Pow(distances[s1].y - distances[s2].y, 2)

	return math.Sqrt(float64(x+y))
}
