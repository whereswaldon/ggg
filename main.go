package main

import (
	"fmt"
	"github.com/whereswaldon/ggg/genetic"
	"os"
)

func main() {
	//Process the source image
	args := os.Args
	if len(args) > 1 {
		imageName := os.Args[1]
		err := genetic.SetTarget(imageName)
		if err != nil {
			fmt.Println(err)
		}
		err = genetic.WriteMonochromePNG("red.png", genetic.Target, genetic.RED)
		if err != nil {
			fmt.Println(err)
		}
	}

	//test out new population
	p := genetic.NewPopulation(100, 0.06)
	fmt.Println(p.GetBestForPrinting(3))
	for i := 0; i < 10; i++ {
		p.Evolve()
		fmt.Println(p.GetBestForPrinting(3))
	}
}
