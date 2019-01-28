package main

import "math/rand"

func numberGenerator() (int, int) {
	if a, b := 1+rand.Intn(10), 1+rand.Intn(10); a > b {
		return a, b
	} else {
		return b, a
	}
}

func operatorGenerator() int {
	return 1 + rand.Intn(3)
}
