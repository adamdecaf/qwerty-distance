package qwerty

import (
	"math"
)

type Location struct {
	x float64
	y float64
}

var distances = map[rune]Location{
	// first row
	'q': Location{0,0},
	'w': Location{1,0},
	'e': Location{2,0},
	'r': Location{3,0},
	't': Location{4,0},
	'y': Location{5,0},
	'u': Location{6,0},
	'i': Location{7,0},
	'o': Location{8,0},
	'p': Location{9,0},

	// second row
	'a': Location{0,1},
	's': Location{1,1},
	'd': Location{2,1},
	'f': Location{3,1},
	'g': Location{4,1},
	'h': Location{5,1},
	'j': Location{6,1},
	'k': Location{7,1},
	'l': Location{8,1},

	// third row
	'z': Location{0,2},
	'x': Location{1,2},
	'c': Location{2,2},
	'v': Location{3,2},
	'b': Location{4,2},
	'n': Location{5,2},
	'm': Location{6,2},
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
