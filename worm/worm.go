package worm

import (
	"fmt"
)

const x_world_size uint = 50
const y_world_size uint = 50

// ----------------------------------->
type Worm struct {
	//x, y       int
	neighbours int8 //if -1 state need to be calculate
	live       bool
}

// ----------------------------------->
type World struct {
	Sheet [x_world_size][y_world_size]Worm
}

// ----------------------------------->
func (w *World) NextGeneration() {
	for x, row := range w.Sheet {
		for y, _ := range row {
			if w.Sheet[x][y].neighbours < 2 {
				w.Sheet[x][y].KillWorm()
			}
			if w.Sheet[x][y].neighbours > 3 {
				w.Sheet[x][y].KillWorm()
			}
			if w.Sheet[x][y].neighbours == 3 {
				w.Sheet[x][y].NewWorm(x, y, w)
			}
		}
	}

}

// ----------------------------------->
func (w *World) SumNeighbours() {
	var neighbours int = 0
	var xx, yy, xxx, yyy int

	for x, row := range w.Sheet {
		for y, _ := range row {

			neighbours = 0

			xx = x - 1
			yy = y - 1
			xxx = x + 1
			yyy = y + 1

			if (xx) < 0 {
				xx = int(w.WorldXSize() - 1)
			}

			if (yy) < 0 {
				yy = int(w.WorldYSize() - 1)
			}

			xxx = xxx % w.WorldXSize()
			yyy = yyy % w.WorldYSize()

			if w.Sheet[x][yy].IsLive() {
				neighbours++
			}
			if w.Sheet[x][yyy].IsLive() {
				neighbours++
			}
			if w.Sheet[xx][yy].IsLive() {
				neighbours++
			}
			if w.Sheet[xx][y].IsLive() {
				neighbours++
			}
			if w.Sheet[xx][yyy].IsLive() {
				neighbours++
			}
			if w.Sheet[xxx][yy].IsLive() {
				neighbours++
			}
			if w.Sheet[xxx][y].IsLive() {
				neighbours++
			}
			if w.Sheet[xxx][yyy].IsLive() {
				neighbours++
			}
			w.Sheet[x][y].neighbours = int8(neighbours)

		}
	}
}

// ----------------------------------->
func (w Worm) IsLive() bool {
	return w.live
}

// ----------------------------------->
func (w *Worm) KillWorm() {
	w.live = false
	w.neighbours = -1
}

// ----------------------------------->
func (w Worm) NewWorm(x, y int, world *World) {
	world.Sheet[x][y].live = true
	world.Sheet[x][y].neighbours = -1
}

// ----------------------------------->
func (w World) WorldXSize() int {
	return len(w.Sheet)
}

// ----------------------------------->
func (w World) WorldYSize() int {
	return len(w.Sheet[0])
}

// ----------------------------------->
func (w World) PrintWorld() {
	for _, row := range w.Sheet {
		line := ""
		for _, w := range row {
			if w.IsLive() {
				line += "@"
			} else {
				line += " "
			}
		}
		fmt.Println(line)
	}
}
