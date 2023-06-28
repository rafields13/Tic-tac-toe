package main

import (
	"fmt"
	"strings"
)

type Player struct {
	Symbol   string
	Nickname string
}

type Board struct {
	Cells  [3][3]string
	Player Player
}

func (b *Board) Display() {
	for _, row := range b.Cells {
		fmt.Println(strings.Join(row[:], ""))
	}
}

func (b *Board) Update(row, col int) {
	b.Cells[row][col] = b.Player.Symbol
}

func (b *Board) CheckWinner() bool {
	s := b.Player.Symbol
	c := b.Cells

	for i := 0; i < 3; i++ {
		if c[i][0] == s && c[i][1] == s && c[i][2] == s {
			return true
		}

		if c[0][i] == s && c[1][i] == s && c[2][i] == s {
			return true
		}
	}

	if c[0][0] == s && c[1][1] == s && c[2][2] == s {
		return true
	}

	if c[0][2] == s && c[1][1] == s && c[2][0] == s {
		return true
	}

	return false
}

func (b *Board) CheckDraw() bool {
	for _, row := range b.Cells {
		for _, cell := range row {
			if cell == "" {
				return false
			}
		}
	}

	return true
}

func main() {
	var p1, p2 Player
	p1.Symbol = "X"
	p2.Symbol = "O"

	fmt.Print("Player1, type your nickname, please: ")
	_, _ = fmt.Scanln(&p1.Nickname)

	fmt.Print("Player2, type your nickname, please: ")
	_, _ = fmt.Scanln(&p2.Nickname)

	board := Board{
		Player: p1,
	}

	for {
		board.Display()

		fmt.Printf("%s, digite a posição (linha coluna): ", board.Player.Nickname)
		var row, col int
		_, _ = fmt.Scanln(&row, &col)

		if row < 0 || row >= 3 || col < 0 || col >= 3 || board.Cells[row][col] != "" {
			fmt.Println("Invalid position. Try again.")
			continue
		}

		board.Update(row, col)

		if board.CheckWinner() {
			board.Display()
			fmt.Printf("Congrats, %s! You won!\n", board.Player.Nickname)
			break
		} else if board.CheckDraw() {
			board.Display()
			fmt.Println("Draw!")
			break
		}

		if board.Player == p1 {
			board.Player = p2
		} else {
			board.Player = p1
		}
	}
}
