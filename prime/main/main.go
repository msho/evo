package main

import (
	"fmt"
	"prime/world"
)

func main() {
	container := world.Container{}
	container.InitLife()

	// loop lifeCycle (reprduce*X, armagedon)
	for i := 0; i < 9000000; i++ {
		if container.LifeCycle() {
			fmt.Println("!! Fount at", i)

			break
		}
	}
	bestFit := container.GetBestFitValue()
	fmt.Println("Best fit", bestFit.ToString())
}
