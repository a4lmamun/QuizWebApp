package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	var q string
	num1, num2 := numberGenerator()

	switch operatorGenerator() {
	case 0:
		q = strconv.Itoa(num1) + "+" + strconv.Itoa(num2)
	case 1:
		q = strconv.Itoa(num1) + "-" + strconv.Itoa(num2)
	case 2:
		q = strconv.Itoa(num1) + "*" + strconv.Itoa(num2)
	}

	fmt.Println(q)
}
