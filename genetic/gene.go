package genetic

import (
	"fmt"
)

/**
The likelihood of a given gene mutating. (e.g. 0.06 is a 6% chance)
*/
var mutationRate float64

/**
Gene is atructure intended to represent a single gene
A, B, and C are constant coefficients that are used by f
f is a function that evalutes the three constant coefficients
	and two inputs to return an integer value representing a color
	level (RGB)
format is a format-string representation of the function body
*/
type Gene struct {
	A, B, C int
	f       func(int, int, int, int, int) int
	format  string
}

/**
NewGene creates a new gene with a random set of coefficients and a
random function
*/
func NewGene() *Gene {
	g := Gene{
		randIntWithNeg(255),
		randIntWithNeg(255),
		randIntWithNeg(255),
		nil,
		"",
	}
	(&g).generateNewGeneFunction()

	return &g
}

/**
randIntWithNeg returns an integer between -numRange and numRange
*/
func randIntWithNeg(numRange int) int {
	return numRange - murphy.Intn(numRange)
}

/**
randCoefficient generates a random number between -255 and 255
for use as a coefficient in a Gene
*/
func randCoefficient() int {
	return randIntWithNeg(255)
}

/**
generateNewGeneFunction creates a function appropriate for a Gene
struct randomly.
*/
func (gene *Gene) generateNewGeneFunction() {
	switch murphy.Intn(10) {
	case 0:
		gene.f = func(a, b, c, x, y int) int {
			return (a * x / (b*y + 1)) + c
		}
		gene.format = "(%d * x / (%d*y + 1)) + %d"
	case 1:
		gene.f = func(a, b, c, x, y int) int {
			return a + b + c + x + y
		}
		gene.format = "%d + %d + %d + x + y"
	case 2:
		gene.f = func(a, b, c, x, y int) int {
			return a * x * b * y * c
		}
		gene.format = "%d * x * %d * y * %d"
	default:
		gene.f = func(a, b, c, x, y int) int {
			return a*x + b*y + c
		}
		gene.format = "%d * x + %d * y + %d"
	}
}

/**
Mutate the gene by changing the constants involved
*/
func (gene *Gene) Mutate() {
	gene.A = randCoefficient()
	gene.B = randCoefficient()
	gene.C = randCoefficient()
}

/**
Evaluates the gene's function for the given x and y
*/
func (gene *Gene) EvalWith(x, y int) int {
	return gene.f(gene.A, gene.B, gene.C, x, y)
}

/**
Returns a string representation of the Gene
*/
func (gene *Gene) String() string {
	functionString := fmt.Sprintf(gene.format, gene.A, gene.B, gene.C)
	return fmt.Sprintf("f(x,y) = %s", functionString)
}
