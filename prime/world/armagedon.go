package world

import (
	"fmt"
	"prime/config"
)

// Armagedon make sure only the fitness would survive, kill the rest.
// Natrual selection
func Armagedon(population []Spirit) []Spirit {

	biggestScores := calculateScore(population)

	//get best MinimumPopulation score kill the rest
	return newSprices(population, biggestScores)
}

func newSprices(oldPopulation []Spirit, biggestScores map[int]float32) []Spirit {
	newPop := make([]Spirit, config.MinimumPopulation)
	i := 0
	for pos := range biggestScores {
		newPop[i] = oldPopulation[pos]
		fmt.Println("now at ", newPop[i].Value)
		i = i + 1
	}
	return newPop
}

/*********/

func calculateScore(population []Spirit) map[int]float32 {
	biggestScores := make(map[int]float32, config.MinimumPopulation)
	var totalScore float32 = 0
	for i := 0; i < len(population); i++ {
		spiritValue := population[i].Value
		totalScore = getLengthScore(spiritValue) + getCharactarsScore(spiritValue)

		//MinimumPopulation biggest score
		chooseBiggestScores(biggestScores, totalScore, i)
	}
	return biggestScores
}

func getLengthScore(value Expression) float32 {
	valLen := float32(len(value.Tokens))
	wantedLen := float32(len(config.WantedResult))
	var scoreMulty float32

	if valLen > wantedLen {
		scoreMulty = wantedLen / valLen
	} else {
		scoreMulty = valLen / wantedLen
	}

	return scoreMulty * config.MaxLengthScore
}

func getCharactarsScore(value Expression) float32 {
	var score float32 = 0
	for i := 0; i < len(value.Tokens); i++ {
		if value.Tokens[i].ScalarValue == 0 { // TODO
			score += config.MaxCharactarScore
		}
	}

	return score
}

func chooseBiggestScores(biggestScores map[int]float32, newScore float32, pos int) {
	if len(biggestScores) < config.MinimumPopulation {
		biggestScores[pos] = newScore
	} else {
		minVal, minKey := getLowestScore(biggestScores)
		if newScore > minVal {
			delete(biggestScores, minKey)
			biggestScores[pos] = newScore
		}
	}
}

func getLowestScore(biggestScores map[int]float32) (float32, int) {
	var lowestVal float32 = 9999999
	var lowestPos int
	for pos, score := range biggestScores {
		if lowestVal > score {
			lowestVal = score
			lowestPos = pos
		}
	}
	return lowestVal, lowestPos
}
