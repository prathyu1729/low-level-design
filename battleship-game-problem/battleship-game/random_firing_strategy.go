package battleship_game

import (
	"math/rand"
	"time"
)

type randomFiringStrategy struct {
}

func newRandomFiringStrategy() *randomFiringStrategy {
	rand.Seed(time.Now().UnixNano())
	return &randomFiringStrategy{}
}

func (r *randomFiringStrategy) getPos(boundary boundary) Point {
	xmin := boundary.topLeft.x
	xhigh := boundary.bottomRight.x
	yhigh := boundary.bottomRight.y
	ymin := boundary.topLeft.y
	x := rand.Intn(xhigh-xmin+1) + xmin
	y := rand.Intn(yhigh-ymin+1) + ymin
	return NewPoint(x, y)
}
