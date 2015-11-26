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
	return &Member{genes}
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
createDataArray uses the genes of this member to create a two-dimensional
array of data by using indicies into that array as the input to the Member's
Gene functions
*/
func (mem *Member) CreateDataArray(arrayHeight, arrayWidth int) [][]uint8 {
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
