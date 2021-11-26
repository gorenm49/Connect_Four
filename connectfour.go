package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type panel struct {
	turn     int
	position [][]string
}

func main() {

	var p1 panel = panel{

		turn: 0,
		position: [][]string{
			[]string{"| - |", "| - |", "| - |", "| - |", "| - |", "| - |", "| - |"},
			[]string{"| - |", "| - |", "| - |", "| - |", "| - |", "| - |", "| - |"},
			[]string{"| - |", "| - |", "| - |", "| - |", "| - |", "| - |", "| - |"},
			[]string{"| - |", "| - |", "| - |", "| - |", "| - |", "| - |", "| - |"},
			[]string{"| - |", "| - |", "| - |", "| - |", "| - |", "| - |", "| - |"},
			[]string{"| - |", "| - |", "| - |", "| - |", "| - |", "| - |", "| - |"},
		},
	}

	p1.show()
	for {
		if p1.turn%2 == 0 {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Please Enter column number: ")
			text, _ := reader.ReadString('\n')
			text = strings.TrimSpace(text)

			if text == "q" {
				break
			}

			i1, err := strconv.Atoi(text)
			if err == nil {
				if i1 > 7 {
					fmt.Println("No such column exists, kindly choose a value between 1 and 7")
					continue
				}
				_, err := p1.add(1, i1)

				gameOver := p1.endGame(1)
				if gameOver {
					p1.show()
					fmt.Println("Congratulation!!! PLAYER 1 WINS!")
					break
				}
				if !err {
					p1.turn++
					p1.show()
				}
				continue
			}
		}

		if p1.turn%2 == 1 {
			p := Search1("R")
			_, err := p1.add(2, p)
			gameOver := p1.endGame(2)
			if gameOver {
				p1.show()
				fmt.Println("Congratulation!!! PLAYER 2 WINS")
				break
			}
			if !err {
				p1.turn++
				p1.show()
			}
			continue

		}

	}
}

func (p1 panel) getTile(p int) string {
	tile := ""
	if p == 1 {
		tile = "| X |"
	} else {
		tile = "| O |"
	}
	return tile
}

func (p1 panel) endGame(p int) bool {
	rows := 7
	cols := 6
	position := p1.position
	tile := p1.getTile(p)

	//  horizontal check
	for j := 0; j < rows-3; j++ {
		for i := 0; i < cols; i++ {
			if position[i][j] == tile && position[i][j+1] == tile && position[i][j+2] == tile && position[i][j+3] == tile {
				return true
			}
		}
	}

	//  vertical check
	for i := 0; i < cols-3; i++ {
		for j := 0; j < rows; j++ {

			if position[i][j] == tile && position[i+1][j] == tile && position[i+2][j] == tile && position[i+3][j] == tile {
				return true
			}
		}
	}

	// +ve diag
	for i := 3; i < cols; i++ {
		for j := 0; j < rows-3; j++ {
			if position[i][j] == tile && position[i-1][j+1] == tile && position[i-2][j+2] == tile && position[i-3][j+3] == tile {
				return true
			}
		}
	}

	// -ve diag
	for i := 0; i < cols-3; i++ {
		for j := 3; j < rows; j++ {
			if position[i][j] == tile && position[i-1][j-1] == tile && position[i-2][j-2] == tile && position[i-3][j-3] == tile {
				return true
			}
		}
	}

	return false

}

func (p1 panel) add(p int, pos int) (panel, bool) {

	if p1.position[0][pos-1] != "| - |" {
		fmt.Println("Can't add it here")
		return p1, true
	}

	tile := p1.getTile(p)

	idx := 1
	for {
		if p1.position[len(p1.position)-idx][pos-1] == "| - |" {
			p1.position[len(p1.position)-idx][pos-1] = tile
			break
		} else {
			idx++
		}
	}
	return p1, false
}

func (p1 panel) show() {
	for _, element := range p1.position {
		fmt.Println(element)
	}
}

func Search1(difficulty string) int {
	if difficulty == "R" {
		fmt.Println(" ... ")
		v := rand.Intn(7-1) + 1
		return v
	}
	return -1
}
