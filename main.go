package main

import (
	"fmt"
	"log"
	"time"
	//"github.com/rivo/tview"
)

const x_world_size uint = 50
const y_world_size uint = 100

//----------------------------------->
type Worm struct {
	x, y       uint16
	neighbours int8 //if -1 state need to be calculate
}

//----------------------------------->
type World struct {
	maps [x_world_size][y_world_size]int8
}

//----------------------------------->
func WorldXSize(w World) int {
	return len(w.maps)
}

//----------------------------------->
func WorldYSize(w World) int {
	return len(w.maps[0])
}

//----------------------------------->
func newWorm(x, y uint16) *Worm {
	return &Worm{x, y, -1}
}

//----------------------------------->
func checkWormNeigbours(x, y int, w World) int8 {
	var neighbours int8 = 0
	var xx, yy, xxx, yyy int

	//fmt.Println(xx)
	xx = x - 1
	yy = y - 1
	xxx = x + 1
	yyy = y + 1
	if (xx) < 0 {
		xx = int(WorldXSize(w) - 1)
	}

	if (yy) < 0 {
		yy = int(WorldYSize(w) - 1)
	}

	xxx = xxx % WorldXSize(w)
	yyy = yyy % WorldYSize(w)

	if w.maps[x][yy] != 0 {
		neighbours++
	}
	if w.maps[x][yyy] != 0 {
		neighbours++
	}
	if w.maps[xx][yy] != 0 {
		neighbours++
	}
	if w.maps[xx][y] != 0 {
		neighbours++
	}
	if w.maps[xx][yyy] != 0 {
		neighbours++
	}
	if w.maps[xxx][yy] != 0 {
		neighbours++
	}
	if w.maps[xxx][y] != 0 {
		neighbours++
	}
	if w.maps[xxx][yyy] != 0 {
		neighbours++
	}

	return neighbours
}

//----------------------------------->
func nextWorldGeneration(world World) {

}

//----------------------------------->
func putWormToWorld(w *Worm, world *World) bool {

	if w.x < uint16(WorldXSize(*world)) && w.y < uint16(WorldYSize(*world)) {
		world.maps[w.x][w.y] = 1
	} else {
		log.Fatalf("World map is too small for it: %v", *w)
	}
	return true

}

//----------------------------------->

func main() {

	var WormWorld World
	var TmpWormWorld World
	var NewWormWorld World
	w1 := newWorm(10, 10)
	putWormToWorld(w1, &WormWorld)
	w1 = newWorm(10, 11)
	putWormToWorld(w1, &WormWorld)
	w1 = newWorm(10, 12)
	putWormToWorld(w1, &WormWorld)
	w1 = newWorm(10, 13)
	putWormToWorld(w1, &WormWorld)
	w1 = newWorm(10, 14)
	putWormToWorld(w1, &WormWorld)
	w1 = newWorm(10, 15)
	putWormToWorld(w1, &WormWorld)
	w1 = newWorm(10, 16)
	putWormToWorld(w1, &WormWorld)
	w1 = newWorm(10, 17)
	putWormToWorld(w1, &WormWorld)
	w1 = newWorm(10, 18)
	putWormToWorld(w1, &WormWorld)
	w1 = newWorm(10, 19)
	putWormToWorld(w1, &WormWorld)
	w1 = newWorm(10, 20)
	putWormToWorld(w1, &WormWorld)

	w1 = newWorm(11, 15)
	putWormToWorld(w1, &WormWorld)
	w1 = newWorm(11, 16)
	putWormToWorld(w1, &WormWorld)
	w1 = newWorm(11, 17)
	putWormToWorld(w1, &WormWorld)
	w1 = newWorm(11, 18)
	putWormToWorld(w1, &WormWorld)
	w1 = newWorm(11, 19)
	putWormToWorld(w1, &WormWorld)
	w1 = newWorm(11, 20)
	putWormToWorld(w1, &WormWorld)

	for i := 1; i <= 500; i++ {

		for x, element := range WormWorld.maps {
			for y, _ := range element {
				TmpWormWorld.maps[x][y] = checkWormNeigbours(int(x), int(y), WormWorld)
			}

		}

		for x, element := range TmpWormWorld.maps {
			for y, _ := range element {
				switch TmpWormWorld.maps[x][y] {

				case 2:
					//life
					NewWormWorld.maps[x][y] = WormWorld.maps[x][y]
				case 3:
					//newborn
					NewWormWorld.maps[x][y] = 1
				default:
					//death
					NewWormWorld.maps[x][y] = 0
				}
			}

		}

		WormWorld = NewWormWorld

		for _, element := range NewWormWorld.maps {
			fmt.Println(element)
		}
		fmt.Println(i)
		time.Sleep(time.Second / 2)
	}
}
