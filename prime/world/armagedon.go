package world

import (
	"prime/config"
)

// Armagedon make sure only the fitness would survive, kill the rest.
// Natrual selection
func Armagedon(population []Spirit) []Spirit {

	biggestScores := CalculateScore(population)

	//get best MinimumPopulation score kill the rest
	return newSprices(population, biggestScores)
}

func newSprices(oldPopulation []Spirit, biggestScores map[int]int) []Spirit {
	newPop := make([]Spirit, config.MinimumPopulation)
	i := 0
	for pos := range biggestScores {
		newPop[i] = oldPopulation[pos]
		//fmt.Println("now at ", newPop[i].Value)
		i = i + 1
	}
	return newPop
}
