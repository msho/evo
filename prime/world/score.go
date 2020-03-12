package world

import (
	"fmt"
	"math"
	"prime/config"
)

// CalculateScore by prime numbers found
func CalculateScore(population []Spirit) map[int]int {
	biggestScores := make(map[int]int, config.MinimumPopulation)
	var totalScore int = 0
	for i := 0; i < len(population); i++ {
		spiritValue := population[i].Value
		totalScore = calculateSpiritScore(spiritValue)

		//MinimumPopulation biggest score
		chooseBiggestScores(biggestScores, totalScore, i)
	}
	return biggestScores
}

func calculateSpiritScore(value Expression) int {
	return expressionPrimesFound(value, false) * config.PrimeScore
}

func expressionPrimesFound(value Expression, isLog bool) int {
	primeFound := 0
	primesFound := make(map[int]bool)
	stopLoop := int(config.HighestPrime / 3)
	for i := 0; i < stopLoop; i++ {
		if isPrimeFound(value.Tokens, i, primesFound) {
			primeFound++
		}
	}

	if isLog && primeFound > 20 {
		fmt.Println(value.ToString())
	}
	return primeFound
}

func isPrimeFound(tokens []Token, iteration int, primesFound map[int]bool) bool {
	expressionResult := operateExpression(tokens, iteration)
	if _, ok := primesFound[expressionResult]; ok {
		return false
	}
	if _, ok := config.WantedResult2[expressionResult]; ok {
		primesFound[expressionResult] = true
		return true
	}
	return false
}

func operateExpression(tokens []Token, iteration int) int {
	totalValue := float32(0)
	lastOperation := plus
	for _, token := range tokens {
		if int(token.Type) > 1 {
			lastOperation = token.Type
			continue
		}

		var currValue float32
		if token.Type == scalar {
			currValue = token.ScalarValue
		} else { // variable
			currValue = float32(math.Pow(float64(iteration), float64(token.VariableValue)))
		}
		totalValue = operateValue(totalValue, lastOperation, currValue)
	}
	return int(totalValue)
}

func operateValue(leftVal float32, operation TokenType, rightVal float32) float32 {
	switch operation {

	case plus:
		return leftVal + rightVal

	case minus:
		return leftVal - rightVal

	case times:
		return leftVal * rightVal

	case divide:
		if rightVal == 0 {
			return 0
		}
		return leftVal / rightVal

	case power:
		return float32(math.Pow(float64(leftVal), float64(rightVal)))

	case squareRoot:
		return float32(math.Abs(math.Sqrt(float64(leftVal))))

	case cubeRoot:
		return float32(math.Abs(math.Cbrt(float64(leftVal))))

	default:
		return float32(0)
	}

}

func chooseBiggestScores(biggestScores map[int]int, newScore int, pos int) {
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

func getLowestScore(biggestScores map[int]int) (int, int) {
	var lowestVal int = 9999999
	var lowestPos int
	for pos, score := range biggestScores {
		if lowestVal > score {
			lowestVal = score
			lowestPos = pos
		}
	}
	return lowestVal, lowestPos
}
