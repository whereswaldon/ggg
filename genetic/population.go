package genetic

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

/**
Population is an array of Members that can evolve.
Members is a slice of Members that is sorted from best fitness to
	worst
*/
type Population struct {
	Members      []*Member
	mutationRate float64
}

/**
NewPopulation creates a new random population of the
specified size.
*/
func NewPopulation(size int, mutate float64) *Population {
	m := make([]*Member, size*10)
	for i := range m {
		m[i] = NewMember()
	}
	sort.Sort(ByFitnessDesc(m))
	return &Population{m[:size], mutate}
}

/**
String returns a string representation of this population.
*/
func (pop *Population) String() string {
	stringPop := make([]string, len(pop.Members))
	for i, v := range pop.Members {
		stringPop[i] = v.String()
	}
	return fmt.Sprintf("Population:\t(mutation rate: %f)\n%s",
		pop.mutationRate,
		strings.Join(stringPop, "\n"))
}

/**
GetBestForPrinting returns a string representation of the most fit
members in the population. The number of members it returns depends on
the number given.
*/
func (pop *Population) GetBestForPrinting(howMany int) string {
	best := make([]string, howMany)
	for i, v := range pop.Members[:howMany] {
		best[i] = v.String()
	}
	return strings.Join(best, "\n")
}

/**
Returns the number of members in the population
*/
func (pop *Population) Size() int {
	return len(pop.Members)
}

/**
Evolve puts the population through a single level of generational
change. All members may mutate, and two of them will be crossed to add
a new Member to the population.
*/
func (pop *Population) Evolve() {
	pop.mutateAll()
	popSize := pop.Size()
	index1 := int(math.Abs(getRandomWeightedFloat()) * float64(popSize))
	index2 := int(math.Abs(getRandomWeightedFloat()) * float64(popSize))
	//fmt.Printf("index1: %d\nindex2: %d\n", index1, index2)
	newMem := Cross(pop.Members[index1], pop.Members[index2])
	pop.Members[popSize-1] = newMem
	sort.Sort(ByFitnessDesc(pop.Members))
}

/**
mutateAll gives every member of the population a chance to mutate.
*/
func (pop *Population) mutateAll() {
	for _, v := range pop.Members {
		if murphy.Float64() < pop.mutationRate {
			v.Mutate()
		}
	}
}
