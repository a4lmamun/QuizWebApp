package main

import (
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func index(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UTC().UnixNano())
	num1, num2 := numberGenerator()

	var q string
	switch operatorGenerator() {
	case 0:
		q = strconv.Itoa(num1) + "+" + strconv.Itoa(num2)
	case 1:
		q = strconv.Itoa(num1) + "-" + strconv.Itoa(num2)
	case 2:
		q = strconv.Itoa(num1) + "*" + strconv.Itoa(num2)
	}

	err := tpl.ExecuteTemplate(w, "question.gohtml", q)
	if err != nil {
		log.Fatalln(err)
	}
}
