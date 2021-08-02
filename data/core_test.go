package data

import (
	"strconv"
	"testing"
)

func TestSignum(t *testing.T) {
	result := Signum(-0.001)
	if result != -1 {
		t.Errorf("Signum returned %d, but should have returned -1", result)
	}
	result = Signum(55555)
	if result != 1 {
		t.Errorf("Signum returned %d, but should have returned 1", result)
	}
	result = Signum(0)
	if result != 0 {
		t.Errorf("Signum returned %d, but should have returned 0", result)
	}
}

func TestGetOrientation(t *testing.T) {
	result := GetOrientation(Point{0, 0, 0}, Point{1, 0, 0}, Point{0, 1, 0})
	if result != 1 {
		t.Errorf("Orientation should be 1, but is %d", result)
	}
	result = GetOrientation(Point{0, 0, 0}, Point{0, 1, 0}, Point{1, 0, 0})
	if result != -1 {
		t.Errorf("Orientation should be -1, but is %d", result)
	}
	result = GetOrientation(Point{0, 0, 0}, Point{1, 0, 0}, Point{2, 0, 0})
	if result != 0 {
		t.Errorf("Orientation should be 0, but is %d", result)
	}
}

func TestGetGeneralizedOrientation(t *testing.T) {
	result := GetGeneralizedOrientation(Point{0, 0, 0}, Point{1, 0, 0}, Point{0, 1, 0})
	if result != 1 {
		t.Errorf("Orientation should be 1, but is %d", result)
	}
	result = GetGeneralizedOrientation(Point{0, 0, 0}, Point{0, 1, 0}, Point{1, 0, 0})
	if result != -1 {
		t.Errorf("Orientation should be -1, but is %d", result)
	}
	result = GetGeneralizedOrientation(Point{0, 0, 0}, Point{1, 0, 0}, Point{2, 0, 0})
	if result != 1 {
		t.Errorf("Orientation should be 1, but is %d", result)
	}
	result = GetGeneralizedOrientation(Point{1, 0, 0}, Point{0, 0, 0}, Point{2, 0, 0})
	if result != -1 {
		t.Errorf("Orientation should be -1, but is %d", result)
	}
	result = GetGeneralizedOrientation(Point{0, 0, 0}, Point{2, 0, 0}, Point{1, 0, 0})
	if result != 0 {
		t.Errorf("Orientation should be 0, but is %d", result)
	}
}

func TestTriangle_GetArea(t *testing.T) {
	triangle := Triangle{Point{0, 0, 0}, Point{10, 0, 0}, Point{0, 10, 0}}
	area := triangle.GetArea()
	if area != 50 {
		t.Errorf("Area should be 50, bit is %f", area)
	}
}

func TestTriangle_IsPointInside(t *testing.T) {
	triangle := Triangle{Point{0, 0, 0}, Point{10, 0, 0}, Point{0, 10, 0}}
	isPointInside := triangle.IsPointInside(Point{4,4,0})
	if !isPointInside {
		t.Errorf("Should be true, but is %s",strconv.FormatBool(isPointInside))
	}
}

func TestAreLineSegmentsIntersected(t *testing.T) {
	lineSegment1, lineSegment2 := LineSegment{Point{0,0,0},Point{10,10,0}},LineSegment{Point{5,0,0},Point{0,5,0}}
	areLineSegmentsIntersected := AreLineSegmentsIntersected(lineSegment1,lineSegment2)
	if !areLineSegmentsIntersected {
		t.Errorf("Should be true, but is %s", strconv.FormatBool(areLineSegmentsIntersected))
	}
}
