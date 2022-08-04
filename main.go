package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"

	"github.com/rewizor/life/world"
	"github.com/rewizor/life/worm"
)

//----------------------------------->

func main() {

	var WormWorld world.World
	var TmpWormWorld world.World
	var NewWormWorld world.World

	w1 := worm.NewWorm(10, 10)

	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= rand.Intn(1000); i++ {
		w1 = worm.NewWorm(uint16(rand.Intn(10)), uint16(rand.Intn(10)))
		worm.PutWormToWorld(w1, &WormWorld)
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		r1 = r1
		//time.Sleep(time.Second / 4)
		//fmt.Println(w1)
	}

	for i := 1; i <= 500; i++ {

		//fmt.Println(WormWorld.Maps)
		//time.Sleep(4 * time.Second)
		var line string = ""
		for x, element := range WormWorld.Maps {
			line = ""
			for y, _ := range element {
				if WormWorld.Maps[x][y] == 0 {
					line += " "
				} else {
					line += "@"
				}
			}
			fmt.Println(line)
		}

		fmt.Println(i)
		time.Sleep(time.Second / 2)

		for x, element := range WormWorld.Maps {
			for y, _ := range element {
				TmpWormWorld.Maps[x][y] = world.CheckWormNeigbours(int(x), int(y), WormWorld)
			}

		}

		for x, element := range TmpWormWorld.Maps {
			for y, _ := range element {
				switch TmpWormWorld.Maps[x][y] {

				case 2:
					//life
					NewWormWorld.Maps[x][y] = WormWorld.Maps[x][y]
				case 3:
					//newborn
					NewWormWorld.Maps[x][y] = 1
				default:
					//death
					NewWormWorld.Maps[x][y] = 0
				}
			}

		}

		WormWorld = NewWormWorld

		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()

	}
}
