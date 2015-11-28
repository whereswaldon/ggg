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

	//Test out the Member implementation
	r := genetic.NewMember()
	g := genetic.NewMember()
	b := genetic.NewMember()
	for i := 0; i < 100; i++ {
		fmt.Printf("Red: %s\n", r)
		fmt.Printf("Blue: %s\n", b)
		fmt.Printf("Green: %s\n", g)
		/*	g := genetic.NewMember()
			fmt.Printf("Green: %s\n", g)
			b := genetic.NewMember()
			fmt.Printf("Blue: %s\n", b)
			a := genetic.NewMember()
			fmt.Printf("Alpha: %s\n", a)
			err := genetic.WriteMultichromePNG("random.png",
				r.CreateDataArray(genetic.TargetHeight, genetic.TargetWidth),
				g.CreateDataArray(genetic.TargetHeight, genetic.TargetWidth),
				b.CreateDataArray(genetic.TargetHeight, genetic.TargetWidth),
				a.CreateDataArray(genetic.TargetHeight, genetic.TargetWidth))
		*/
		err := genetic.WriteMultichromePNG(fmt.Sprintf("random-%d.png", i),
			r.GetData(), g.GetData(), b.GetData(), nil)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Fitness: %d\n", r.GetFitness())
		r.Mutate()
		g.Mutate()
		b.Mutate()
		/*
			p := genetic.NewPopulation(100)
			fmt.Println(p)
		*/
	}
}
