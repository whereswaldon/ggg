package main

import (
	"fmt"
	"github.com/whereswaldon/ggg/genetic"
)

func main() {
	m := genetic.NewMember()
	fmt.Printf("%s\n", m)
	fmt.Printf("%#v\n", m.CreateDataArray(23, 23))
}
