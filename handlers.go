package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		fmt.Println(r.FormValue("option"))
	}

	rand.Seed(time.Now().UTC().UnixNano())
	num1, num2 := numberGenerator()

	var q string
	options := make([]int, 4)

	key := []int{0, 1, 2, 3}
	mutexAdder := rand.Intn(3)
	mutex := []int{1 + mutexAdder, 3 + mutexAdder, 4 + mutexAdder}
	rand.Shuffle(len(key), func(i, j int) { key[i], key[j] = key[j], key[i] })
	rand.Shuffle(len(mutex), func(i, j int) { mutex[i], mutex[j] = mutex[j], mutex[i] })

	switch operatorGenerator() {
	case 1:
		q = strconv.Itoa(num1) + "+" + strconv.Itoa(num2)
		res := num1 + num2
		values := []int{res, res + mutex[0], res + mutex[1], res + mutex[2]}
		options[0], options[1], options[2], options[3] = values[key[3]], values[key[0]], values[key[2]], values[key[1]]
	case 2:
		q = strconv.Itoa(num1) + "-" + strconv.Itoa(num2)
		res := num1 - num2
		values := []int{res, res + mutex[0], res + mutex[1], res + mutex[2]}
		options[0], options[1], options[2], options[3] = values[key[1]], values[key[2]], values[key[3]], values[key[0]]
	case 3:
		q = strconv.Itoa(num1) + "*" + strconv.Itoa(num2)
		res := num1 * num2
		values := []int{res, res + mutex[0], res + mutex[1], res + mutex[2]}
		options[0], options[1], options[2], options[3] = values[key[2]], values[key[1]], values[key[0]], values[key[3]]
	}

	err := tpl.ExecuteTemplate(w, "question.gohtml", Question{q, options})
	if err != nil {
		log.Fatalln(err)
	}
}
