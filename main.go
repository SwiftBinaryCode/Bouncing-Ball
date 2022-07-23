package main

import (
	"fmt"
	"time"
)

func main() {

	const (
		width  = 50
		height = 10

		cellEmpty = ' '
		cellBall  = '⚾'

		maxFrames = 1200
		speed     = time.Second / 20
	)

	var (
		px, py int
		vx, vy = 1, 1
		cell   rune
	)

	board := make([][]bool, 50)

	for column := range board {
		board[column] = make([]bool, 10)

	}

	screen.Clear()

	for i := 0; i < maxFrames; i++ {

		px += vx
		py += vy

		if px <= 0 || px >= width-1 {

			vx *= -1
		}

		if py <= 0 || py >= height-1 {

			vy *= -1
		}

		//remove the previous ball
		for y := range board[0] {
			for x := range board {
				board[x][y] = false
			}
		}

		board[px][py] = true
		buf := make([]rune, 0, width*height)

		//draw the board
		// for y := 0; y < height; y++ {
		// 	for x := 0; x < width; x++ {

		// 	}
		// }

		buf = buf[:0]
		for y := range board[0] {
			for x := range board {
				cell = cellEmpty
				if board[x][y] {
					cell = cellBall
				}

				// fmt.Print(string(cell), " ")

				buf = append(buf, cell, ' ')
			}
			// fmt.Println()
			buf = append(buf, '\n')

		}

		screen.MoveTopLeft()

		fmt.Print(string(buf))

		time.Sleep(speed)

	}

}
