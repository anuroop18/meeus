// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

package apsis_test

import (
	"fmt"
	"math"
	"testing"
	"time"

	"github.com/soniakeys/meeus/apsis"
	"github.com/soniakeys/meeus/julian"
	"github.com/soniakeys/meeus/moonposition"
	"github.com/soniakeys/sexagesimal"
)

func ExampleMeanApogee() {
	// Example 50.a, p. 357.
	fmt.Printf("JDE = %.4f\n", apsis.MeanApogee(1988.75))
	// Output:
	// JDE = 2447442.8191
}

func ExampleApogee() {
	// Example 50.a, p. 357.
	j := apsis.Apogee(1988.75)
	fmt.Printf("JDE = %.4f\n", j)
	y, m, d := julian.JDToCalendar(j)
	d, f := math.Modf(d)
	fmt.Printf("%d %s %d, at %m TD\n", y, time.Month(m), int(d),
		sexa.TimeFromDay(f).Fmt())
	// Output:
	// JDE = 2447442.3543
	// 1988 October 7, at 20ʰ30ᵐ TD
}

func ExampleApogeeParallax() {
	// Example 50.a, p. 357.
	p := apsis.ApogeeParallax(1988.75)
	fmt.Printf("%.3f\n", sexa.Angle(p).Sec())
	fmt.Printf("%0.3d\n", sexa.Angle(p).Fmt())
	// Output:
	// 3240.679
	// 54′00″.679
}

// Test cases from p. 361.
func TestPerigee(t *testing.T) {
	for _, c := range []struct {
		y, m  int
		d, dy float64
	}{
		{1997, 12, 9 + 16.9/24, 1997.93},
		{1998, 1, 3 + 8.5/24, 1998.01},
		{1990, 12, 2 + 10.8/24, 1990.92},
		{1990, 12, 30 + 23.8/24, 1991},
	} {
		ref := julian.CalendarGregorianToJD(c.y, c.m, c.d)
		j := apsis.Perigee(c.dy)
		if math.Abs(j-ref) > .1 {
			t.Fatal("got", j, "expected", ref)
		}
	}
}

// Lacking a worked example from the text, test using meeus/moonposition.
func TestPerigeeParallax(t *testing.T) {
	got := apsis.PerigeeParallax(1997.93)
	_, _, d := moonposition.Position(apsis.Perigee(1997.93))
	want := moonposition.Parallax(d)
	Δ := math.Abs(got-want) / math.Pi * 180 * 3600 // difference in arc seconds
	// for this case anyway it's within a tenth of an arc second
	if Δ > .1 {
		t.Fatal("got", got, "want (about)", want)
	}
}
