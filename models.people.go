package main

import (
	"time"
	"math/rand"
)

type people struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var (
	// TODO: fetch this from a DB maybe? i guess in mem list is fine
	personList = []string{"Saralina", "Ashley", "Ben", "Ale", "Emma", "Shane", "Alec", "JoJo"}
)

// Return a list of all the people in secret santa
func getSecretSantaPick() (person string) {
	// randomly choose someone from the list return/remove them
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator

	index := rand.Intn(len(personList))
	secretPick := personList[index]

	// remove them from the list we are keeping track of
	personList = updatePersonList(personList, secretPick)
	return secretPick
}

func updatePersonList(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}
