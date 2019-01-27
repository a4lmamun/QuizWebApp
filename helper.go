package main

import "math/rand"

func numberGenerator() (int, int) {
	return rand.Intn(15), rand.Intn(10)
}

func operatorGenerator() int {
	return rand.Intn(3)
}
