package main

import "math/rand"

func processShort(numberOfIterations int) {
	s := ""
	for i := 0; i < numberOfIterations; i++ {
		s += getRandomString(500)
	}
}

func processLong(numberOfIterations int) {
	s := ""
	for i := 0; i < numberOfIterations; i++ {
		s += getRandomString(3000)
	}
}

func getRandomString(stringLength int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, stringLength)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
