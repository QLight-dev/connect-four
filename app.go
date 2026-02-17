package main

import "fmt"

type Game struct {
	board [][]Token
	turn  int
}

type Token string

const PlayerOneToken Token = "1"
const PlayerTwoToken Token = "2"

func InitBoard(game *Game) {
	// make columns
	game.board = make([][]Token, 6) // 0-index will count as a column so its 7-column
	// make rows
	for i := 0; i < 6; i++ {
		game.board[i] = make([]Token, 0, 5) // 0-index will count as a row so its 6-row
	}
}

func PlaceToken(game *Game, column int, token Token) error {
	if column > 6 {
		return fmt.Errorf("%v is out of bounds", column)
	}
	if column < 0 {
		return fmt.Errorf("%v is out of bounds", column)
	}
	game.board[column] = append(game.board[column], token)
	return nil
}

func (game Game) PrintBoard() {
	// rows
	for row := 7; row > -1; row-- {
		for column := 0; column < 8; column++ {
			if column > len(game.board)-1 || row > len(game.board[column])-1 {
				fmt.Print(" ")
				continue
			}
			fmt.Print(game.board[column][row])
		}
		fmt.Print("\n")
	}
	// visual numbers so players don't get confused
	for i := 0; i < 7; i++ {
		fmt.Print(i)
	}
	fmt.Print("\n")
}
