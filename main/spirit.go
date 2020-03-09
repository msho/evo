package main

// Spirit is a living thinking thing
type Spirit struct {
	Value string
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
