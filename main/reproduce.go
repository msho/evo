package main

import (
	"math/rand"
)

// MutateValue change a given value string randomly
func MutateValue(value string) string {
	if needAdd(value) {
		value = addToValue(value)
	}
	if needRemove(value) {
		value = removeFromValue(value)
	}

	return changeValue(value)
}

func randomNum(min, max int) int {
	//rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

func randomChar() rune {
	return rune(randomNum(11, 999))
}

func needAdd(value string) bool {
	return len(value) < len(WantedResult)
}

func needRemove(value string) bool {
	return len(value) > len(WantedResult)
}

func addToValue(value string) string {
	valToAdd := ""
	for i := 0; i < randomNum(0, 10); i++ {
		valToAdd += string(randomChar())
	}
	return value + valToAdd
}

func removeFromValue(value string) string {
	trimLen := randomNum(1, len(value))

	return value[:trimLen]
}

func changeValue(value string) string {
	mutatedValue := ""
	for i := 0; i < len(value) && i < len(WantedResult); i++ {
		if value[i] != WantedResult[i] {
			mutatedValue += string(randomChar())
		} else {
			mutatedValue += string(value[i])
		}
	}

	return mutatedValue
}
