package main

import (
	"fmt"
	"github.com/whereswaldon/ggg/genetic"
	//"math"
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

	/*
		g := &genetic.Gene{250, 250, 10,
			func(a, b, c, x, y int) int {
				return int(math.Sqrt(float64((a-x)*(a-x)+(b-y)*(b-y)))) + c*c
			},
			"%d*x - %d*y + %d"}
		m := genetic.NewMember()
		m.Genes = make([]*genetic.Gene, 1)
		m.Genes[0] = g
		fmt.Println(m)
		genetic.WriteMonochromePNG("out.png", m.GetData(), genetic.RED)
	*/
	//test out new population
	p := genetic.NewPopulation(200, 0.15)
	//fmt.Println(p.GetBestForPrinting(1))
	bestScore := 0
	runs := 100
	for i := 0; i < runs; i++ {
		p.Evolve()
		tempScore := p.Members[0].GetFitness()
		if tempScore > bestScore {
			fmt.Println(p.GetBestForPrinting(1))
			bestScore = tempScore
			genetic.WriteMonochromePNG(fmt.Sprintf("solution-%d.png", i), p.Members[0].GetData(), genetic.RED)
		} else {
			fmt.Printf("%d/%d:\t%d <= %d\n", i, runs, tempScore, bestScore)
		}
	}
}
