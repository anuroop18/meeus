package fit_test

import (
	"fmt"
	"math"

	"github.com/soniakeys/meeus/fit"
)

func ExampleLinear() {
	// Example 4.a, p. 37.
	a, b := fit.Linear([]struct{ X, Y float64 }{
		{.2982, 10.92},
		{.2969, 11.01},
		{.2918, 10.99},
		{.2905, 10.78},
		{.2707, 10.87},

		{.2574, 10.80},
		{.2485, 10.75},
		{.2287, 10.14},

		{.2238, 10.21},
		{.2156, 9.97},
		{.1992, 9.69},

		{.1948, 9.57},
		{.1931, 9.66},
		{.1889, 9.63},
		{.1781, 9.65},
		{.1772, 9.44},
		{.1770, 9.44},

		{.1755, 9.32},
		{.1746, 9.20},
	})
	fmt.Printf("a, b: %.2f, %.2f\n", a, b)
	// Output:
	// a, b: 13.67, 7.03
}

func ExampleCorrelationCoefficient() {
	// Example 4.b, p. 40.
	data := []struct{ X, Y float64 }{
		{73, 90.4},
		{38, 125.3},
		{35, 161.8},
		{42, 143.4},
		{78, 52.5},
		{68, 50.8},
		{74, 71.5},
		{42, 152.8},
		{52, 131.3},
		{54, 98.5},
		{39, 144.8},

		{61, 78.1},
		{42, 89.5},
		{49, 63.9},
		{50, 112.1},
		{62, 82.0},
		{44, 119.8},
		{39, 161.2},
		{43, 208.4},
		{54, 111.6},
		{44, 167.1},
		{37, 162.1},
	}
	a, b := fit.Linear(data)
	fmt.Printf("y = %.2f - %.2fx\n", b, -a)
	fmt.Printf("r = %.3f\n", fit.CorrelationCoefficient(data))
	// Output:
	// y = 244.18 - 2.49x
	// r = -0.767
}

// Meeus provides this as an example, but without exact results.
// In fact the answer doesn't seem close at all.  Correctness of the
// function Multiple3 is suspect.
func ExampleMultiple3() {
	data := []struct{ X, Y float64 }{
		{3, .0433},
		{20, .2532},
		{34, .3386},
		{50, .3560},
		{75, .4983},
		{88, .7577},
		{111, 1.4585},
		{129, 1.8628},
		{143, 1.8264},
		{160, 1.2431},
		{183, -.2043},
		{200, -1.2431},
		{218, -1.8422},
		{230, -1.8726},
		{248, -1.4889},
		{269, -.8372},
		{290, -.4377},
		{303, -.3640},
		{320, -.3508},
		{344, -.2126},
	}
	f0 := math.Sin
	f1 := func(x float64) float64 { return math.Sin(2 * x) }
	f2 := func(x float64) float64 { return math.Sin(3 * x) }
	a, b, c := fit.Multiple3(data, f0, f1, f2)
	fmt.Println(a, b, c)
}