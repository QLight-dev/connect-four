package main

import (
	"fmt"
)

type Game struct {
	board [][]Token
	Player string
}

// ANSI colours
const (
	esc            = "\u001b"
	clearScreen    = esc + "[2J"
	ResetColor     = esc + "[0m"
	YellowText     = esc + "[33m"
	RedText        = esc + "[31m"
	blackText      = esc + "[30m"
	blueBackground = esc + "[48;2;0;0;80m"
)

type Token string

const PlayerOneToken Token = YellowText + "⬤"
const PlayerTwoToken Token = RedText + "⬤"
const emptyToken Token = blackText + "⬤"

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
	if len(game.board[column]) == 6 {
		return fmt.Errorf("%v is full\n", column)
	}
	game.board[column] = append(game.board[column], token)
	return nil
}

func (game Game) PrintBoard() {
	// to make terminal output clearer
	fmt.Print(clearScreen + esc + "[H")

	for row := 5; row >= 0; row-- {
		fmt.Print(blueBackground)
		fmt.Print(" ")
		for column := 0; column < 7; column++ {
			if column > len(game.board)-1 || row > len(game.board[column])-1 {
				fmt.Print(emptyToken)
				fmt.Print(ResetColor)
				fmt.Print(blueBackground)
				fmt.Print(" ")
				continue
			}
			fmt.Print(game.board[column][row])
			fmt.Print(" ")
			fmt.Print(ResetColor)
			fmt.Print(blueBackground)
		}
		fmt.Print("\n")
	}
	fmt.Print(ResetColor)

	for i := 0; i < 7; i++ {
		fmt.Print(i)
	}
	fmt.Print("\n")
}

func (game Game) CheckWin(lastTokenPlacedCol int, lastTokenPlacedRow int, playerToken Token) (bool, string) {
	tokensInARow := 0

	// column checking
	for _, row := range game.board[lastTokenPlacedCol] {
		if row == playerToken {
			tokensInARow++
		} else {
			tokensInARow = 0
		}
		if tokensInARow == 4 {
			return true, game.Player
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
			return true, game.Player
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
				return true, game.Player
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
				return true, game.Player
			}
		} else {
			tokensInARow = 0
		}

		col--
		row++
	}

	return false, ""
}
