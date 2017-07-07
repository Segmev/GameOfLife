package main

import (
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten"
)

var (
	maxX, maxY = 150, 100
	scale      = 4.0
	cells      = 0
	tab        [][]int
	pixels     []uint8
)

func display(tab [][]int, screen *ebiten.Image) {
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			pos := maxX*y*4 + 4*x
			if tab[y][x] > 0 {
				pixels[pos] = 30 + uint8(tab[y][x])
				pixels[pos+1] = 30 + uint8(tab[y][x])
				pixels[pos+2] = 0xFF
				pixels[pos+3] = 0xff
			} else {
				pixels[pos] = 0x20
				pixels[pos+1] = 0x23
				pixels[pos+2] = 0x20
				pixels[pos+3] = 0xAA
			}
		}
	}
	screen.ReplacePixels(pixels)
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
			if tab[j][i] > 200 {
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

// func drawRect(screen *ebiten.Image, x, y int, color color.Color) {
// 	opts := &ebiten.DrawImageOptions{}
// 	square.Fill(color)
// 	opts.GeoM.Translate(float64(x), float64(y))
// 	screen.DrawImage(square, opts)
// }

var count = 0

func handleInputs() bool {
	inputDetected := false
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		inputDetected = true
		cursorX, cursorY := ebiten.CursorPosition()
		if cursorX >= 0 && cursorX < maxX && cursorY >= 0 && cursorY < maxY {
			tab[cursorY][cursorX]++
		}
	}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		inputDetected = true
		cursorX, cursorY := ebiten.CursorPosition()
		if cursorX >= 0 && cursorX < maxX && cursorY >= 0 && cursorY < maxY {
			tab[cursorY][cursorX]--
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		inputDetected = true
		os.Exit(0)
	}

	if ebiten.IsKeyPressed(ebiten.KeyR) {
		inputDetected = true
		tab = make([][]int, maxY)
		for i := 0; i < maxY; i++ {
			tab[i] = make([]int, maxX)
		}
	}
	return inputDetected
}

func update(screen *ebiten.Image) error {
	if !handleInputs() {
		tab = updateTab(tab)
	}
	if ebiten.IsRunningSlowly() {
		return nil
	}
	display(tab[:], screen)
	return nil
}

func main() {
	args := os.Args[1:]
	if len(args) > 1 {
		if nb, err := strconv.Atoi(args[0]); err == nil {
			if nb > 50 {
				maxX = nb
			}
		}
		if nb, err := strconv.Atoi(args[1]); err == nil {
			if nb > 50 {
				maxY = nb
			}
		}
	}
	if len(args) > 2 {
		if nb, err := strconv.Atoi(args[2]); err == nil {
			cells = nb
		}
	}
	tab = make([][]int, maxY)
	pixels = make([]uint8, maxX*maxY*4)
	for i := 0; i < maxY; i++ {
		tab[i] = make([]int, maxX)
	}

	fillTab(tab)
	ebiten.Run(update, maxX, maxY, scale, "Game of Life - Revisited by Segmev")
}
