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
NewMember creates a new member of the population with 2 genes
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
