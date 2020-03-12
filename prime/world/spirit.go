package world

import "fmt"

type TokenType int

const (
	scalar   TokenType = 0
	variable TokenType = 1

	plus       TokenType = 2
	minus      TokenType = 3
	times      TokenType = 4
	divide     TokenType = 5
	power      TokenType = 6
	squareRoot TokenType = 7
	cubeRoot   TokenType = 8

	operationsCount = 7
)

// Token is nice
type Token struct {
	Type TokenType

	ScalarValue   float32
	VariableValue int
}

// ToString of a token. Yeah!
func (t *Token) ToString() string {
	var strVal string

	switch t.Type {
	case scalar:
		strVal = fmt.Sprint(t.ScalarValue)
	case variable:
		strVal = fmt.Sprintf("x%d", t.VariableValue)
	case plus:
		strVal = "+"
	case minus:
		strVal = "-"
	case times:
		strVal = "*"
	case divide:
		strVal = "/"
	case power:
		strVal = "^"
	case squareRoot:
		strVal = "square-root"
	case cubeRoot:
		strVal = "cube-root"
	}
	return strVal
}

// Expression symbol a mathematic expression with variable and scalars
// Example (2^x1)-1
type Expression struct {
	Tokens         []Token
	VariablesCount int
}

// ToString write expression in human thing
func (e *Expression) ToString() string {
	strVal := ""
	for _, token := range e.Tokens {
		strVal += " " + token.ToString()
	}
	return strVal
}

// Spirit is a living thinking thing
type Spirit struct {
	Value Expression
}

// Split that spirit
func (s *Spirit) Split() Spirit {
	return Spirit{
		Value: MutateValue(s.Value),
	}
}

// MakeSpirits make number 'num' of Spirits
func MakeSpirits(num int) []Spirit {
	spirits := make([]Spirit, num)

	return spirits
}
