package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

type Element struct {
	alive  bool
	symbol rune
}

func clear() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

const xi = 20
const yi = 50

var a [xi][yi]Element
var b [xi][yi]Element

func populate() {
	for i := 0; i < xi; i++ {
		for j := 0; j < yi; j++ {
			a[i][j].symbol = '-'
			a[i][j].alive = false
		}
	}
}

func countAliveNeighbors(x int, y int) int {
	count := 0

	start_x := x - 1
	start_y := y - 1
	end_x := x + 1
	end_y := y + 1

	if y == 0 {
		start_y = 0
		end_y = 1
	}
	if y == yi-1 {
		start_y = y - 1
		end_y = y
	}
	if x == 0 {
		start_x = 0
		end_x = x + 1
	}

	if x == xi-1 {
		start_x = x - 1
		end_x = x
	}

	for i := start_x; i <= end_x; i++ {
		for j := start_y; j <= end_y; j++ {
			if i == x && j == y {
				continue
			}
			if b[i][j].alive == true {
				count++
			}
		}
	}
	return count
}

func update() {
	for i := 0; i < xi; i++ {
		for j := 0; j < yi; j++ {
			if b[i][j].alive == true {
				switch countAliveNeighbors(i, j) {
				case 0:
				case 1:
					a[i][j].symbol = '-'
					a[i][j].alive = false
					break
				case 2:
				case 3:
					break
				default:
					a[i][j].symbol = '-'
					a[i][j].alive = false
				}
			} else {
				if countAliveNeighbors(i, j) == 3 {
					a[i][j].symbol = '#'
					a[i][j].alive = true
				}
			}
		}
	}
}

func draw() {
	for i := 0; i < xi; i++ {
		for j := 0; j < yi; j++ {
			b[i][j] = a[i][j]
			fmt.Printf("%c", a[i][j].symbol)
		}
		fmt.Println()
	}
	update()
}

func readFile() bool {
	file, err := os.Open("input.csv")
	if err != nil {
		return false
	}

	reader := bufio.NewReader(file)

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Fatal(err)
		}

		x := string(line)[0] - 48
		y := string(line)[2] - 48
		fmt.Println(x,y)
		if (x < xi && x > 0 && y < yi && y > 0) {
			a[x][y].symbol = '#'
			a[x][y].alive = true 
		}

		time.Sleep(1 * time.Second)
	}
	return true
}

func main() {
	populate()
	readFile()

	for {
		clear()
		draw()
		time.Sleep(1 * time.Second)
	}
}
