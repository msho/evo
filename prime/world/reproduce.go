package world

import (
	"math/rand"
)

// MutateValue change a given value string randomly
func MutateValue(value Expression) Expression {
	lastTokenIndex := len(value.Tokens) - 1
	if lastTokenIndex <= 0 {
		value = addScalarOrVariable(value)
		lastTokenIndex++
	}

	lastTokenType := value.Tokens[lastTokenIndex].Type
	if lastTokenType < 2 {
		value = addOperation(value)
	} else {
		value = addScalarOrVariable(value)
	}

	return value

}

func randomNum(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func randomAddVariable(variableCount int) bool {
	// make variable count rare
	return randomNum(0, (variableCount+1)^2) == 1
}

func randomOperation() int {
	return randomNum(2, operationsCount+1)
}

func addScalarOrVariable(value Expression) Expression {
	var newToken Token
	if randomAddVariable(value.VariablesCount) {
		// add variable
		newToken = Token{
			Type:          variable,
			VariableValue: value.VariablesCount + 1,
		}
		value.VariablesCount++
	} else {
		// add scalar
		newToken = Token{
			Type:        scalar,
			ScalarValue: float32(randomNum(0, 10)),
		}
	}
	value.Tokens = append(value.Tokens, newToken)
	return value
}

func addOperation(value Expression) Expression {
	operation := TokenType(randomOperation())
	value.Tokens = append(value.Tokens, Token{Type: operation})

	return value
}
