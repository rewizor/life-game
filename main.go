package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/rewizor/life/worm"
)

// ----------------------------------->
func cls() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

// ----------------------------------->
func main() {

	cls()

	var swiat worm.World
	//var swiat_ng worm.World

	swiat.PutWormToWorld(worm.NewWorm(2, 1, -1))
	swiat.PutWormToWorld(worm.NewWorm(2, 2, -1))
	swiat.PutWormToWorld(worm.NewWorm(2, 3, -1))

	for i := 0; i < 10; i++ {
		swiat.PrintWorld()
		fmt.Println("Generation : ", i)
		swiat = swiat.NextGeneration()
		time.Sleep(1 * time.Second)
	}

}
