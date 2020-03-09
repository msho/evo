package main

import "fmt"

func main() {
	container := Container{}
	container.InitLife()

	// loop lifeCycle (reprduce*X, armagedon)
	for i := 0; i < 100000; i++ {
		if container.LifeCycle() {
			fmt.Println(WantedResult, "at", i)
			return
		}
	}
	fmt.Println("not found :(")
	fmt.Println("Best fit", container.GetBestFitValue())
}
