package main

import (
	"fmt"
	"prime/config"
	"prime/world"
)

func main() {
	container := world.Container{}
	container.InitLife()

	// loop lifeCycle (reprduce*X, armagedon)
	for i := 0; i < 100000; i++ {
		if container.LifeCycle() {
			fmt.Println(config.WantedResult, "at", i)
			return
		}
	}
	fmt.Println("not found :(")
	fmt.Println("Best fit", container.GetBestFitValue())
}
