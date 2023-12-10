package battleship_game

type Point struct {
	x int
	y int
}

func NewPoint(x, y int) Point {
	return Point{x: x, y: y}
}

type boundary struct {
	topLeft     Point
	bottomRight Point
}

func (b boundary) contains(point Point) bool {
	xMin := b.topLeft.x
	xMax := b.bottomRight.x
	yMin := b.topLeft.y
	yMax := b.bottomRight.y
	x := point.x
	y := point.y

	return (x >= xMin && x <= xMax) || (y >= yMin && y <= yMax)
}
