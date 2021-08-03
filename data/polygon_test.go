package data

import "testing"

func TestPolygon_IsSimple(t *testing.T) {
	polygon := Polygon{points: []Point{
		{0, 0, 0},
		{5, 0, 0},
		{5, 5, 0},
		{0, 5, 0},
	}}
	isPolygonSimple := polygon.IsSimple()
	if !isPolygonSimple {
		t.Errorf("Should be simple, but is not")
	}
}

func TestPolygon_IsSimplePolygonConvex(t *testing.T) {
	polygon := Polygon{points: []Point{
		{0, 0, 0},
		{5, 0, 0},
		{5, 5, 0},
		{0, 5, 0},
	}}
	isSimplePolygonConvex := polygon.IsSimplePolygonConvex()
	if !isSimplePolygonConvex {
		t.Errorf("Should be convex, but is not")
	}
}

func TestPolygon_GetConvexOrientation(t *testing.T) {
	polygon := Polygon{points: []Point{
		{0, 0, 0},
		{5, 0, 0},
		{5, 5, 0},
		{0, 5, 0},
	}}
	convexOrientation := polygon.GetConvexOrientation()
	if convexOrientation != 1 {
		t.Errorf("Orientation should be 1, but is %d", convexOrientation)
	}
}

func TestPolygon_GetArea(t *testing.T) {
	polygon := Polygon{points: []Point{
		{0, 0, 0},
		{5, 0, 0},
		{5, 5, 0},
		{0, 5, 0},
	}}
	area := polygon.GetArea()
	if area != 25 {
		t.Errorf("Area should be 25, but is %f", area)
	}
}