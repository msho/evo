package world

import (
	"fmt"
	"prime/config"
)

//Container of life
type Container struct {
	Population []Spirit
}

// InitLife inside the container
func (c *Container) InitLife() {
	fmt.Println("Starting life")
	c.Population = MakeSpirits(config.MinimumPopulation)
}

// LifeCycle reproduce and kill the weakest. Only the fitest survive
func (c *Container) LifeCycle() bool {
	for i := 0; i < config.ReproduceBeforeDeath; i++ {
		if c.Reproduce() {
			return true
		}
	}
	c.Armagedon()

	return false
}

// Reproduce population
func (c *Container) Reproduce() bool {
	populationSize := len(c.Population)
	for i := 0; i < populationSize; i++ {
		newSpirit := c.Population[i].Split()
		if newSpirit.Value.Tokens[0].ScalarValue == 0 { // TODO
			return true
		}
		c.Population = append(c.Population, newSpirit)
	}
	return false
}

// GetBestFitValue get the best spirit yet
func (c *Container) GetBestFitValue() Expression {
	spiritScores := calculateScore(c.Population)
	maxScore := float32(0)
	chosenSpirit := c.Population[0]
	for pos, score := range spiritScores {
		if maxScore > score {
			maxScore = score
			chosenSpirit = c.Population[pos]
		}
	}
	return chosenSpirit.Value
}

// Armagedon for natrual selection
func (c *Container) Armagedon() {
	c.Population = Armagedon(c.Population)
}
