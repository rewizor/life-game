package worm

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

const x_world_size uint = 5
const y_world_size uint = 5

// ----------------------------------->
type Worm struct {
	x, y       int
	neighbours int8 //if -1 state need to be calculate
	live       bool
}

// ----------------------------------->
type World struct {
	Sheet [x_world_size][y_world_size]Worm
}

// ----------------------------------->
// ----------------------------------->   WORLD
// ----------------------------------->

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

// ----------------------------------->
func (w World) NextGeneration() World {
	for x, row := range w.Sheet {
		for y, worm := range row {
			worm.CheckNeigbours(w)
			if !worm.WillLive() {
				worm.Kill()
			}
			w.Sheet[x][y] = worm
			if worm.neighbours == 3 {
				w.Sheet[x][y].live = true
				w.Sheet[x][y].neighbours = -1
				fmt.Println("Urodziny")

			}
		}
	}
	fmt.Println(w)
	time.Sleep(1 * time.Second)
	return w
}

// ----------------------------------->
func (w *World) PutWormToWorld(worm *Worm) bool {

	if worm.x < w.WorldXSize() && worm.y < w.WorldYSize() {
		w.Sheet[worm.x][worm.y] = *worm
	} else {
		log.Fatalf("World map is too small for worm: %v", *worm)
	}

	return true

}

// ----------------------------------->
// ----------------------------------->  WORMS
// ----------------------------------->

// ----------------------------------->
func (w *Worm) CheckNeigbours(wrld World) {
	var neighbours int8 = 0
	var xx, yy, xxx, yyy int

	//fmt.Println(xx)
	xx = w.x - 1
	yy = w.y - 1
	xxx = w.x + 1
	yyy = w.y + 1
	if (xx) < 0 {
		xx = int(wrld.WorldXSize() - 1)
	}

	if (yy) < 0 {
		yy = int(wrld.WorldYSize() - 1)
	}

	xxx = xxx % wrld.WorldXSize()
	yyy = yyy % wrld.WorldYSize()
	//fmt.Println(wrld.Sheet[6][7].live)

	if wrld.Sheet[w.x][yy].live {
		neighbours++
	}
	if wrld.Sheet[w.x][yyy].live {
		neighbours++
	}
	if wrld.Sheet[xx][yy].live {
		neighbours++
	}
	if wrld.Sheet[xx][w.y].live {
		neighbours++
	}
	if wrld.Sheet[xx][yyy].live {
		neighbours++
	}
	if wrld.Sheet[xxx][yy].live {
		neighbours++
	}
	if wrld.Sheet[xxx][w.y].live {
		neighbours++
	}
	if wrld.Sheet[xxx][yyy].live {
		neighbours++
	}
	w.neighbours = neighbours
}

// ----------------------------------->
func (w *Worm) WillLive() bool {

	if w.neighbours >= 2 && w.neighbours <= 3 {
		return true
	} else {
		return false
	}
}

// ----------------------------------->
func (w *Worm) IsLive() bool {
	return w.live
}

// ----------------------------------->
func (w *Worm) Print() {
	opis := "Nie wiem ilu mam sąsiadów."
	fmt.Println("Cześć, jestem robal")
	fmt.Printf("Siedzę sobie w norce o współrzędnych: %v, %v \n", w.x, w.y)
	if w.neighbours > 0 {
		opis = "Mam " + strconv.FormatUint(uint64(w.neighbours), 10) + " sąsiad/a/ów"
	} else {
		if w.neighbours == 0 {
			opis = "Jestem samotny"
		} else {
			opis = "Nie wiem ilu mam sąsiadów"
		}
	}
	fmt.Println(opis)
}

// ----------------------------------->
func (w *Worm) Kill() {
	w.live = false
}

// ----------------------------------->
// x,y - coordinates
// n - amount of neigbours
func NewWorm(x, y int, n int8) *Worm {
	//True - worm is live
	return &Worm{x, y, n, true}
}
