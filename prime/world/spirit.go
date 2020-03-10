package world

import "fmt"

type tokenType int

const (
	scalar   tokenType = 0
	variable tokenType = 1

	plus   tokenType = 2
	minus  tokenType = 3
	times  tokenType = 4
	divide tokenType = 5
	power  tokenType = 6
	root   tokenType = 7

	operationsCount = 6
)

type token struct {
	Type tokenType

	ScalarValue   float32
	VariableValue int
}

func (t *token) ToString() string {
	var strVal string

	switch t.Type {
	case scalar:
		strVal = fmt.Sprint(t.ScalarValue)
	case variable:
		strVal = fmt.Sprintf("x%f", t.VariableValue)
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
	case root:
		strVal = "root"
	}
	return strVal
}

// Expression symbol a mathematic expression with variable and scalars
// Example (2^x1)-1
type Expression struct {
	Tokens         []token
	VariablesCount int
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
