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
	fmt.Printf("Red: %s\n", r)
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
	err := genetic.WriteMonochromePNG("random.png",
		r.CreateDataArray(genetic.TargetHeight, genetic.TargetWidth),
		genetic.RED)
	if err != nil {
		fmt.Println(err)
	}
	/*
		p := genetic.NewPopulation(100)
		fmt.Println(p)
	*/
}
