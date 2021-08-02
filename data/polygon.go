package data

type Polygon struct {
	points []Point
}

func (p *Polygon) IsSimple() bool {
	numberOfPoints := len(p.points)
	for i := 0; i < numberOfPoints-2; i++ {
		for j := i + 2; j < (numberOfPoints+i-2)%numberOfPoints; j++ {
			if AreLineSegmentsIntersected(LineSegment{p.points[i], p.points[i+1]}, LineSegment{p.points[j], p.points[j+1]}) {
				return false
			}
		}
	}
	return true
}

func (p *Polygon) IsSimplePolygonConvex() bool {
	orientation := 0
	numberOfPoints := len(p.points)
	for i := 0; i < numberOfPoints; i++ {
		nextPointIndex := (i+1)%numberOfPoints
		previousPointIndex := (i+numberOfPoints-1)%numberOfPoints
		cross := ThreePointsCross(p.points[previousPointIndex],p.points[i],p.points[nextPointIndex])
		if cross != 0 {
			if orientation == 0 {
				orientation = Signum(cross)
			} else if float64(orientation)*cross < 0 {
				return false
			}
		}
	}
	return true
}

func (p *Polygon) GetConvexOrientation() int {
	numberOfPoints := len(p.points)
	for i:=1;i<numberOfPoints-1;i++ {
		orientation := GetOrientation(p.points[i-1],p.points[i],p.points[i+1])
		if orientation != 0 {
			return orientation
		}
	}
	return 0
}