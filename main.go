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
	var worm worm.Worm

	worm.NewWorm(25, 5, &swiat)
	worm.NewWorm(25, 6, &swiat)
	worm.NewWorm(25, 7, &swiat)
	worm.NewWorm(25, 8, &swiat)
	worm.NewWorm(25, 9, &swiat)
	worm.NewWorm(25, 10, &swiat)
	worm.NewWorm(25, 11, &swiat)
	worm.NewWorm(25, 12, &swiat)
	worm.NewWorm(25, 13, &swiat)
	worm.NewWorm(25, 14, &swiat)

	for i := 1; i <= 40; i++ {
		swiat.PrintWorld()
		fmt.Println("Generation : ", i)
		time.Sleep(1 * time.Second)
		swiat.SumNeighbours()
		swiat.NextGeneration()
	}

}
