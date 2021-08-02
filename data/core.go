package data

import (
	"math"
)

type Point struct {
	X float64
	Y float64
	Z float64
}

type LineSegment struct {
	Start Point
	End   Point
}

type Triangle struct {
	A Point
	B Point
	C Point
}

func Signum(number float64) int {
	if number > 0 {
		return 1
	} else if number < 0 {
		return -1
	}
	return 0
}

func ThreePointsCross(A, B, C Point) float64 {
	return (B.X-A.X)*(C.Y-A.Y) - (C.X-A.X)*(B.Y-A.Y)
}

func GetOrientation(A, B, C Point) int {
	doubleArea := ThreePointsCross(A, B, C)
	return Signum(doubleArea)
}

func GetGeneralizedOrientation(A, B, C Point) int {
	a := B.X - A.X
	b := B.Y - A.Y
	c := C.X - A.X
	d := C.Y - A.Y
	doubleAreaOfTriangle := a*d - b*c
	if doubleAreaOfTriangle != 0 {
		return Signum(doubleAreaOfTriangle)
	}
	if a*c < 0 || b*d < 0 {
		return -1 // LEFT ORIENTED
	}
	if a*a+b*b < c*c+d*d {
		return 1 // RIGHT ORIENTED
	}
	return 0 // COLLINEAR
}

func (t *Triangle) GetArea() float64 {
	return math.Abs(ThreePointsCross(t.A, t.B, t.C)) / 2
}

func (t *Triangle) IsPointInside(p Point) bool {
	triangleOrientation := GetOrientation(t.A, t.B, t.C)
	pabCross := ThreePointsCross(p, t.A, t.B)
	pbcCross := ThreePointsCross(p, t.B, t.C)
	pcaCross := ThreePointsCross(p, t.C, t.A)
	return float64(triangleOrientation)*pabCross >= 0 && float64(triangleOrientation)*pbcCross >= 0 && float64(triangleOrientation)*pcaCross >= 0
}

func AreLineSegmentsIntersected(a, b LineSegment) bool {
	return (GetGeneralizedOrientation(a.Start, a.End, b.Start)*GetGeneralizedOrientation(a.Start, a.End, b.End) <= 0) && (GetGeneralizedOrientation(b.Start, b.End, a.Start)*GetGeneralizedOrientation(b.Start, b.End, a.End) <= 0)
}
