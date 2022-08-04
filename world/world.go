package world

import "fmt"

const x_world_size uint = 50
const y_world_size uint = 50

//----------------------------------->
type World struct {
	Maps [x_world_size][y_world_size]int8
}

//----------------------------------->
func PrintWorld(w World) {
	var line string = ""
	for x, element := range w.Maps {
		line = ""
		for y := range element {
			if w.Maps[x][y] == 0 {
				line += " "
			} else {
				line += "@"
			}
		}
		fmt.Println(line)
	}
}

//----------------------------------->
func NextGeneration(wormWorld World, tmpWormWorld World, newWormWorld *World) {
	for x, element := range tmpWormWorld.Maps {
		for y := range element {
			switch tmpWormWorld.Maps[x][y] {

			case 2:
				//life
				newWormWorld.Maps[x][y] = wormWorld.Maps[x][y]
			case 3:
				//newborn
				newWormWorld.Maps[x][y] = 1
			default:
				//death
				newWormWorld.Maps[x][y] = 0
			}
		}

	}
}

//----------------------------------->
func SumNeigbours(currentWorld World, sumWorld *World) {
	for x, element := range currentWorld.Maps {
		for y := range element {
			sumWorld.Maps[x][y] = CheckWormNeigbours(int(x), int(y), currentWorld)
		}

	}

}

//----------------------------------->
func WorldXSize(w World) int {
	return len(w.Maps)
}

//----------------------------------->
func WorldYSize(w World) int {
	return len(w.Maps[0])
}

//----------------------------------->
func CheckWormNeigbours(x, y int, w World) int8 {
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

	if w.Maps[x][yy] != 0 {
		neighbours++
	}
	if w.Maps[x][yyy] != 0 {
		neighbours++
	}
	if w.Maps[xx][yy] != 0 {
		neighbours++
	}
	if w.Maps[xx][y] != 0 {
		neighbours++
	}
	if w.Maps[xx][yyy] != 0 {
		neighbours++
	}
	if w.Maps[xxx][yy] != 0 {
		neighbours++
	}
	if w.Maps[xxx][y] != 0 {
		neighbours++
	}
	if w.Maps[xxx][yyy] != 0 {
		neighbours++
	}

	return neighbours
}
