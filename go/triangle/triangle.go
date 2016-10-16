package triangle

import "math"

const testVersion = 3

type Kind string

const (
	NaT Kind = "notatriangle"
	Equ Kind = "equilateral"
	Iso Kind = "isoceles"
	Sca Kind = "scalene"
)

type triangle struct {
	a, b, c float64
	kind    Kind
}

func (t *triangle) setKind(k Kind) {
	t.kind = k
}

func (t triangle) isPositive() bool {
	return t.a < 0 || t.b < 0 || t.c < 0
}

func (t triangle) inequal() bool {
	return (t.a+t.b >= t.c) && (t.a+t.c >= t.b) && (t.b+t.c >= t.a)
}

func (t triangle) isInf(sign int) bool {
	return math.IsInf(t.a, sign) || math.IsInf(t.b, sign) || math.IsInf(t.c, sign)
}

func (t triangle) isNaN() bool {
	return math.IsNaN(t.a) || math.IsNaN(t.b) || math.IsNaN(t.c)
}

func (t triangle) isEqu() bool {
	return t.a == t.b && t.b == t.c

}

func (t triangle) isIso() bool {
	return (t.a == t.b && t.b != t.c) || (t.a == t.c && t.b != t.c) || (t.b == t.c && t.a != t.c)
}

func (t triangle) isSca() bool {
	return t.a != t.b && t.b != t.c && t.a != t.c
}

func KindFromSides(a, b, c float64) Kind {
	t := triangle{
		a: a,
		b: b,
		c: c,
	}

	if !t.isPositive() {
		t.setKind(NaT)
		return t.kind
	}

	if !t.inequal() {
		t.setKind(NaT)
		return t.kind
	}

	if t.isInf(1) || t.isInf(-1) {
		t.setKind(NaT)
		return t.kind
	}

	if t.isEqu() {
		t.setKind(Equ)
		return t.kind
	}

	if t.isIso() {
		t.setKind(Iso)
		return t.kind
	}
	if (a == b && b != c) || (a == c && b != c) || (b == c && a != c) {
		return Iso
	}

	if a != b && b != c && a != c {
		return Sca
	}

	return t.kind
}

// inequality tests Triangle Inequlity (https://en.wikipedia.org/wiki/Triangle_inequality).
func inequality(a, b, c float64) bool {
	return (a+b >= c) && (a+c >= b) && (b+c >= a)
}

// inf tests whether parameters are equal to +Inf and -Inf.
// Coincidentally, math.Inf() is equal enough to math.NaN() to pass test suite.
func inf(a, b, c float64) bool {
	switch pinf := math.Inf(0); {
	case a == pinf:
		return true
	case b == pinf:
		return true
	case c == pinf:
		return true
	}

	switch ninf := math.Inf(-1); {
	case a == ninf:
		return true
	case b == ninf:
		return true
	case c == ninf:
		return true
	}

	return false
}
