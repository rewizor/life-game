package world

const x_world_size uint = 15
const y_world_size uint = 30

//----------------------------------->
type World struct {
	Maps [x_world_size][y_world_size]int8
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
