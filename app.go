package main

import "fmt"

type Game struct {
	board [][]Token
	turn  int
}

type Token string

const PlayerOneToken Token = "1"
const PlayerTwoToken Token = "2"
const emptyToken Token = " "

func InitBoard(game *Game) {
	// make columns
	game.board = make([][]Token, 7)
	// make rows
	for i := 0; i <= 6; i++ {
		game.board[i] = make([]Token, 0, 6)
	}
}

func PlaceToken(game *Game, column int, token Token) error {
	if column > 6 {
		return fmt.Errorf("%v is out of bounds\n", column)
	}
	if column < 0 {
		return fmt.Errorf("%v is out of bounds\n", column)
	}
	if len(game.board[column]) == 7 {
		return fmt.Errorf("%v is full\n", column)
	}
	game.board[column] = append(game.board[column], token)
	return nil
}

func (game Game) PrintBoard() {
	// rows
	for row := 7; row >= 0; row-- {
		for column := 0; column <= 7; column++ {
			if column > len(game.board)-1 || row > len(game.board[column])-1 {
				fmt.Print(emptyToken)
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

func (game Game) CheckWin(lastTokenPlacedCol int, lastTokenPlacedRow int) (bool, Token) {
	tokensInARow := 0

	// column checking
	// player 1
	for _, row := range game.board[lastTokenPlacedCol] {
		if row == PlayerOneToken {
			tokensInARow++
		} else {
			tokensInARow = 0
		}
		if tokensInARow == 4 {
			return true, PlayerOneToken
		}
	}
	tokensInARow = 0

	// player 2
	for _, row := range game.board[lastTokenPlacedCol] {
		if row == PlayerTwoToken {
			tokensInARow++
		} else {
			tokensInARow = 0
		}
		if tokensInARow == 4 {
			return true, PlayerTwoToken
		}
	}
	tokensInARow = 0

	// row checking
	// player 1
	for i := range game.board {
		if lastTokenPlacedRow >= len(game.board[i]) {
			tokensInARow = 0
			continue
		}

		if game.board[i][lastTokenPlacedRow] == PlayerOneToken {
			tokensInARow++
		} else {
			tokensInARow = 0
		}

		if tokensInARow == 4 {
			return true, PlayerOneToken
		}
	}
	tokensInARow = 0

	// player 2
	for i := range game.board {
		if lastTokenPlacedRow >= len(game.board[i]) {
			tokensInARow = 0
			continue
		}

		if game.board[i][lastTokenPlacedRow] == PlayerTwoToken {
			tokensInARow++
		} else {
			tokensInARow = 0
		}

		if tokensInARow == 4 {
			return true, PlayerTwoToken
		}
	}
	tokensInARow = 0
	return false, emptyToken
}
