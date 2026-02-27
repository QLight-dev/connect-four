package main

import (
	"fmt"
)

type Game struct {
	board [][]Token
	turn  int
}

type Token string

const PlayerOneToken Token = "1"
const PlayerTwoToken Token = "2"
const emptyToken Token = " "

func InitBoard(game *Game) {
	game.board = make([][]Token, 7)
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
	for row := 5; row >= 0; row-- {
		for column := 0; column < 7; column++ {
			if column > len(game.board)-1 || row > len(game.board[column])-1 {
				fmt.Print(emptyToken)
				continue
			}
			fmt.Print(game.board[column][row])
		}
		fmt.Print("\n")
	}
	for i := 0; i < 7; i++ {
		fmt.Print(i)
	}
	fmt.Print("\n")
}

func (game Game) CheckWin(lastTokenPlacedCol int, lastTokenPlacedRow int, playerToken Token) (bool, Token) {
	tokensInARow := 0

	// column checking
	for _, row := range game.board[lastTokenPlacedCol] {
		if row == playerToken {
			tokensInARow++
		} else {
			tokensInARow = 0
		}
		if tokensInARow == 4 {
			return true, playerToken
		}
	}
	tokensInARow = 0

	// row checking
	for i := range game.board {
		if lastTokenPlacedRow >= len(game.board[i]) {
			tokensInARow = 0
			continue
		}
		if game.board[i][lastTokenPlacedRow] == playerToken {
			tokensInARow++
		} else {
			tokensInARow = 0
		}
		if tokensInARow == 4 {
			return true, playerToken
		}
	}
	tokensInARow = 0

	// =============================
	// DIAGONAL 1 (bottom-left → top-right)
	// =============================

	var diagonalRootRow int
	var diagonalRootCol int

	{
		row := lastTokenPlacedRow
		col := lastTokenPlacedCol

		for row > 0 && col > 0 {
			row--
			col--
		}

		diagonalRootRow = row
		diagonalRootCol = col
	}

	// Player 1
	col := diagonalRootCol
	row := diagonalRootRow
	tokensInARow = 0

	for {
		if col >= len(game.board) || row >= 6 {
			break
		}

		if row >= len(game.board[col]) {
			tokensInARow = 0
		} else if game.board[col][row] == playerToken {
			tokensInARow++
			if tokensInARow == 4 {
				return true, playerToken
			}
		} else {
			tokensInARow = 0
		}

		col++
		row++
	}

	// =============================
	// DIAGONAL 2 (bottom-right → top-left)
	// =============================

	{
		row := lastTokenPlacedRow
		col := lastTokenPlacedCol

		for row > 0 && col < len(game.board)-1 {
			row--
			col++
		}

		diagonalRootRow = row
		diagonalRootCol = col
	}

	// Player 1
	col = diagonalRootCol
	row = diagonalRootRow
	tokensInARow = 0

	for {
		if col < 0 || row >= 6 {
			break
		}

		if row >= len(game.board[col]) {
			tokensInARow = 0
		} else if game.board[col][row] == playerToken {
			tokensInARow++
			if tokensInARow == 4 {
				return true, playerToken
			}
		} else {
			tokensInARow = 0
		}

		col--
		row++
	}

	return false, emptyToken
}
