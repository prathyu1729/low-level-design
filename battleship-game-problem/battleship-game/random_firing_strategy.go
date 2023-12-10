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

func (r *randomFiringStrategy) getPos(xmin, xhigh, ymin, yhigh int) (int, int) {
	x := rand.Intn(xhigh-xmin+1) + xmin
	y := rand.Intn(yhigh-ymin+1) + ymin
	return x, y
}
