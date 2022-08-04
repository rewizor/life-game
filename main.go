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
func cls() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

//----------------------------------->
func main() {

	cls()
	var WormWorld world.World
	var TmpWormWorld world.World
	var NewWormWorld world.World

	w1 := worm.NewWorm(10, 10)

	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= rand.Intn(1000); i++ {
		w1 = worm.NewWorm(uint16(rand.Intn(10)), uint16(rand.Intn(10)))
		worm.PutWormToWorld(w1, &WormWorld)
		//time.Sleep(time.Second / 4)
		//fmt.Println(w1)
	}

	for i := 1; i <= 500; i++ {

		world.PrintWorld(WormWorld)

		fmt.Println("Generation: ", i)
		time.Sleep(time.Second / 2)

		world.SumNeigbours(WormWorld, &TmpWormWorld)
		world.NextGeneration(WormWorld, TmpWormWorld, &NewWormWorld)

		WormWorld = NewWormWorld

		cls()

	}
}
