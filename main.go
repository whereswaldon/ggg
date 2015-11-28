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
	m := genetic.NewMember()
	fmt.Printf("%s\n", m)
	fmt.Printf("%#v\n", m.CreateDataArray(23, 23))
	err := genetic.WriteMonochromePNG("random.png", m.CreateDataArray(genetic.TargetHeight, genetic.TargetWidth), genetic.RED)
	if err != nil {
		fmt.Println(err)
	}
}
