package genetic

import (
	"fmt"
	"math"
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
	F       func(int, int, int, int, int) int
	Format  string
}

/**
NewGene creates a new gene with a random set of coefficients and a
random function
*/
func NewGene() *Gene {
	g := Gene{
		randCoefficient(),
		randCoefficient(),
		randCoefficient(),
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
	return numRange - murphy.Intn(numRange*2)
}

/**
randCoefficient generates a random number between -255 and 255
for use as a coefficient in a Gene
*/
func randCoefficient() int {
	return randIntWithNeg(max(TargetWidth, TargetHeight))
}

/**
generateNewGeneFunction creates a function appropriate for a Gene
struct randomly.
*/
func (gene *Gene) generateNewGeneFunction() {
	switch murphy.Intn(2) {
	/*
		case 0:
			gene.F = func(a, b, c, x, y int) int {
				return (a * x / (b*y + 1)) + c
			}
			gene.Format = "(%d * x / (%d*y + 1)) + %d"
		case 1:
			gene.F = func(a, b, c, x, y int) int {
				return a + b + c + x + y
			}
			gene.Format = "%d + %d + %d + x + y"
		case 2:
			gene.F = func(a, b, c, x, y int) int {
				return a * x * b * y * c
			}
			gene.Format = "%d * x * %d * y * %d"
		case 3:
			gene.F = func(a, b, c, x, y int) int {
				return a*x + b*y + c
			}
			gene.Format = "%d * x + %d * y + %d"
	*/
	case 0:
		gene.F = func(a, b, c, x, y int) int {
			return -1 * (int(math.Sqrt(float64((a-x)*(a-x)+(b-y)*(b-y)))) + c*c)
		}
		gene.Format = "-1*(((%d - x)^2 + (%d - y)^2)^.5 + %d^2)"
	default:
		gene.F = func(a, b, c, x, y int) int {
			return int(math.Sqrt(float64((a-x)*(a-x)+(b-y)*(b-y)))) + c*c
		}
		gene.Format = "((%d - x)^2 + (%d - y)^2)^.5 + %d^2"
	}
}

/**
Copy returns an exact duplicate of this gene.
*/
func (gene *Gene) Copy() *Gene {
	return &Gene{
		gene.A,
		gene.B,
		gene.C,
		gene.F,
		gene.Format,
	}
}

/**
Mutate the gene by changing the constants involved
*/
func (gene *Gene) Mutate() {
	victim := murphy.Intn(3)
	switch victim {
	case 0:
		gene.A = randCoefficient()
	case 1:
		gene.B = randCoefficient()
	case 2:
		gene.C = randCoefficient()
	}
}

/**
Evaluates the gene's function for the given x and y
*/
func (gene *Gene) EvalWith(x, y int) int {
	return gene.F(gene.A, gene.B, gene.C, x, y)
}

/**
Returns a string representation of the Gene
*/
func (gene *Gene) String() string {
	functionString := fmt.Sprintf(gene.Format, gene.A, gene.B, gene.C)
	return fmt.Sprintf("f(x,y) = %s", functionString)
}
