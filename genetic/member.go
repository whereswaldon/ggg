package genetic

import (
	"fmt"
	"strings"
)

/**
Member represents a single member of the solution polulation.
Genes contains the set of genes of that member
*/
type Member struct {
	Genes []*Gene
	//data generated by gene functions, if any
	data [][]uint8
	//indicates whether the genes have mutated since last data was generated
	mutated bool
}

/**
NewMember creates a new member of the population with 1 to 10 genes
*/
func NewMember() *Member {
	//create between 1 and 10 genes
	genes := make([]*Gene, murphy.Intn(10)+1)

	for i := range genes {
		genes[i] = NewGene()
	}
	return &Member{genes, nil, false}
}

/**
String returns a string representation of the member with
each of its Gene functions printed on a new line
*/
func (mem *Member) String() string {
	stringGenes := make([]string, len(mem.Genes))
	for i, v := range mem.Genes {
		stringGenes[i] = v.String()
	}
	return fmt.Sprintf("Member:\n\t%s", strings.Join(stringGenes, "\n\t"))
}

/**
GetData retrieves the data array generated by this Member's gene functions.
*/
func (mem *Member) GetData() [][]uint8 {
	if mem.data == nil || mem.mutated {
		mem.data = mem.createDataArray(TargetHeight, TargetWidth)
		mem.mutated = false
		return mem.data
	}
	return mem.data
}

/**
Mutate changes this members genes by adding a gene function,
removing a gene function, or modifying an existing function.
*/
func (mem *Member) Mutate() {
	mem.mutated = true
	numGenes := len(mem.Genes)
	dieRoll := murphy.Intn(20)
	switch dieRoll {
	case 0:
		//delete a gene
		if numGenes > 0 {
			mem.deleteFirstGene()
			fmt.Println("Lost a gene.")
		} else {
			//add a gene
			mem.addGeneAtEnd()
			fmt.Println("Gained a gene.")
		}
	case 1:
		fallthrough
	case 2:
		//add a gene
		mem.addGeneAtEnd()
		fmt.Println("Gained a gene.")
	default:
		//mutate an existing gene
		victim := murphy.Intn(numGenes)
		mem.Genes[victim].Mutate()
		fmt.Printf("Changed gene %d\n", victim)
	}
}

/**
Removes the first Gene function from the member's genes
*/
func (mem *Member) deleteFirstGene() {
	mem.Genes = mem.Genes[1:]
}

/**
addGeneAtEnd inserts a new gene function at the end of a member's genes
*/
func (mem *Member) addGeneAtEnd() {
	mem.Genes = append(mem.Genes, NewGene())
}

/*
GetFitness returns an integer score representing how closely this member's
data aligns with the target data. Low scores are better.
*/
func (mem *Member) GetFitness() int {
	data := mem.GetData()
	var dataPoint, targetPoint uint8
	offBy := 0
	for y := range Target {
		for x := range Target[y] {
			dataPoint = data[y][x]
			targetPoint = Target[y][x]
			if dataPoint > targetPoint {
				offBy += int(dataPoint) - int(targetPoint)
			} else {
				offBy += int(targetPoint) - int(dataPoint)
			}
		}
	}

	return offBy
}

/**
createDataArray uses the genes of this member to create a two-dimensional
array of data by using indicies into that array as the input to the Member's
Gene functions
*/
func (mem *Member) createDataArray(arrayHeight, arrayWidth int) [][]uint8 {
	data := make([][]uint8, arrayHeight)
	for y := range data {
		data[y] = make([]uint8, arrayWidth)

		//insert data
		for x := range data[y] {
			data[y][x] = mem.EvalGenesAt(x, y)
		}
	}

	return data
}

/**
Eval genes at finds the sum of the gene functions at a particular coordinate pair
*/
func (mem *Member) EvalGenesAt(xCoord, yCoord int) uint8 {
	var val uint8
	for _, g := range mem.Genes {
		val += uint8(g.EvalWith(xCoord, yCoord))
	}
	return val
}
