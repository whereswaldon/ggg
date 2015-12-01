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
	A, B, C float64
	F       func(float64, float64, float64, int, int) int
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
randFloatWithNeg creates a random floating point number between the integer
and its negative equivalent
*/
func randFloatWithNeg(numRange int) float64 {
	return murphy.Float64() * float64(numRange)
}

/**
randCoefficient generates a random number between -255 and 255
for use as a coefficient in a Gene
*/
func randCoefficient() float64 {
	return randFloatWithNeg(max(TargetWidth, TargetHeight))
}

/**
generateNewGeneFunction creates a function appropriate for a Gene
struct randomly.
*/
func (gene *Gene) generateNewGeneFunction() {
	switch murphy.Intn(3) {
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
		gene.F = func(a, b, c float64, x, y int) int {
			xf, yf := float64(x), float64(y)
			return -1 * int(math.Sqrt(((a-xf)*(a-xf)+(b-yf)*(b-yf)))+c*c)
		}
		gene.Format = "-1*(((%f - x)^2 + (%f - y)^2)^.5 + %f^2)"
	case 1:
		gene.F = func(a, b, c float64, x, y int) int {
			xf, yf := float64(x), float64(y)
			return int(c - c*math.Abs(a-xf)/a - c*math.Abs(b-yf)/b)
		}
		gene.Format = "c-c*|a-x|/a - c*|b-yf|/b a=%f b=%f c=%f"
		/*
			case 2:
				gene.F = func(a, b, c float64, x, y int) int {
					return int(c + b + a)
				}
				gene.Format = "%f + %f + %f"
		*/
	default:
		gene.F = func(a, b, c float64, x, y int) int {
			xf, yf := float64(x), float64(y)
			return int(math.Sqrt(((a-xf)*(a-xf) + (b-yf)*(b-yf))) + c*c)
		}
		gene.Format = "((%f - x)^2 + (%f - y)^2)^.5 + %f^2"
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
