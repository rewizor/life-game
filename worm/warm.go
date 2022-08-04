package worm

import (
	"log"

	"github.com/rewizor/life/world"
)

//----------------------------------->
type Worm struct {
	x, y       uint16
	neighbours int8 //if -1 state need to be calculate
}

//----------------------------------->
func NewWorm(x, y uint16) *Worm {
	return &Worm{x, y, -1}
}

//----------------------------------->
func PutWormToWorld(w *Worm, wrld *world.World) bool {

	if w.x < uint16(world.WorldXSize(*wrld)) && w.y < uint16(world.WorldYSize(*wrld)) {
		wrld.Maps[w.x][w.y] = 1
	} else {
		log.Fatalf("World map is too small for it: %v", *w)
	}

	return true

}
