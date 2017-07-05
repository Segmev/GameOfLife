package main

import (
	"bytes"
	"image/color"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten"
)

var maxX, maxY = 800, 600
var cells = 1000

func displayTerm(tab [][]int) {
	var buffer bytes.Buffer

	for j := 0; j < maxY; j++ {
		buffer.WriteString("   ")
		for i := 0; i < maxX; i++ {
			if tab[j][i] == 0 {
				buffer.WriteString("  ")
			} else if tab[j][i] < 3 {
				buffer.WriteString("¤ ")
			} else if tab[j][i] < 6 {
				buffer.WriteString("* ")
			} else if tab[j][i] < 10 {
				buffer.WriteString("% ")
			} else {
				buffer.WriteString("# ")
			}
		}
		buffer.WriteString("\n")
	}
}

func display(tab [][]int, screen *ebiten.Image) {

	for j := 0; j < maxY; j++ {
		for i := 0; i < maxX; i++ {
			if tab[j][i] == 0 {
			} else if tab[j][i] < 3 {
				drawRect(screen, j, i, color.White)
			} else if tab[j][i] < 6 {
				drawRect(screen, j, i, color.White)
			} else if tab[j][i] < 10 {
				drawRect(screen, j, i, color.White)
			} else {
				drawRect(screen, j, i, color.White)
			}
		}
	}
}

func updateTab(tab [][]int) [][]int {
	buff := make([][]int, maxY)
	for i := 0; i < maxY; i++ {
		buff[i] = make([]int, maxX)
	}

	for j := 0; j < maxY; j++ {
		for i := 0; i < maxX; i++ {
			c := 0

			for a := -1; a <= 1; a++ {
				for b := -1; b <= 1; b++ {
					if !(a == 0 && b == 0) {
						if j+a >= 0 && j+a < maxY && i+b >= 0 && i+b < maxX {
							if tab[j+a][i+b] > 0 {
								c++
							}
						}
					}
				}
			}

			if tab[j][i] > 15 {
				for a := 0; a <= 1; a++ {
					for b := -1; b <= 0; b++ {
						if !(a == 0 && b == 0) {
							if j+a >= 0 && j+a < maxY && i+b >= 0 && i+b < maxX {
								buff[j+a][i+b]++
							}
						}
					}
				}
			}

			if c == 3 {
				buff[j][i] = tab[j][i] + 1
			} else if tab[j][i] > 0 && c == 2 {
				buff[j][i] = tab[j][i] + 1
			}
		}
	}
	return buff
}

func fillTab(tab [][]int) {
	rand.Seed(time.Now().UTC().UnixNano())
	for z := 0; z < cells; z++ {
		tab[rand.Intn(maxY)][rand.Intn(maxX)] = 1
	}
}

var square *ebiten.Image

func drawRect(screen *ebiten.Image, x, y int, color color.Color) {
	opts := &ebiten.DrawImageOptions{}
	square.Fill(color)
	opts.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(square, opts)
}

func update(screen *ebiten.Image) error {

	test = updateTab(test)
	display(test[:], screen)
	time.Sleep(120 * time.Millisecond)

	return nil
}

var test [][]int

func main() {
	args := os.Args[1:]
	if len(args) > 1 {
		if nb, err := strconv.Atoi(args[0]); err == nil {
			maxX = nb
		}
		if nb, err := strconv.Atoi(args[1]); err == nil {
			maxY = nb
		}
	}
	if len(args) > 2 {
		if nb, err := strconv.Atoi(args[2]); err == nil {
			cells = nb
		}
	}
	test = make([][]int, maxY)
	for i := 0; i < maxY; i++ {
		test[i] = make([]int, maxX)
	}

	fillTab(test)
	square, _ = ebiten.NewImage(1, 1, ebiten.FilterNearest)

	ebiten.Run(update, maxX, maxY, 1, "Hello world!")

	//for true {
	//	test = updateTab(test)
	//	display(test[:])
	//	time.Sleep(120 * time.Millisecond)
	//}
}
