package main

import (
	"fmt"
	"github.com/whereswaldon/ggg/genetic"
	//	"math"
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
		g := &genetic.Gene{100.0, 300.0, 255.0,
			func(a, b, c float64, x, y int) int {
				xf, yf := float64(x), float64(y)
				return int(c - c*math.Abs(a-xf)/a - c*math.Abs(b-yf)/b)
			},
			"%f*x - %f*y + %f"}
		m := genetic.NewMember()
		m.Genes = make([]*genetic.Gene, 1)
		m.Genes[0] = g
		fmt.Println(m)
		genetic.WriteMonochromePNG("out.png", m.GetData(), genetic.RED)
	*/

	//test out new population
	p := genetic.NewPopulation(200, 0.03)
	//fmt.Println(p.GetBestForPrinting(1))
	bestScore := 0
	runs := 1000
	for i := 0; i < runs; i++ {
		p.Evolve()
		tempScore := p.Members[0].GetFitness()
		if tempScore > bestScore {
			fmt.Println(p.GetBestForPrinting(1))
			bestScore = tempScore
			genetic.WriteMonochromePNG(fmt.Sprintf("solution-%d.png", i), p.Members[0].GetData(), genetic.RED)
		} else {
			fmt.Printf("%d,\t%d,\t%d\n", i, runs, tempScore, bestScore)
		}
	}
}
