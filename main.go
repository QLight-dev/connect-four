package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Connect 4!")
	game := Game{}
	InitBoard(&game)
	var playerToken Token

	for {
		var column int
		if playerToken == PlayerOneToken {
			playerToken = PlayerTwoToken
			game.Player = RedText + "Red" + ResetColor
		} else {
			playerToken = PlayerOneToken
			game.Player = YellowText + "Yellow" + ResetColor
		}
		game.PrintBoard()

		for valid := false; !valid; {
			fmt.Printf("[%v] enter column: ", game.Player)
			_, err := fmt.Scanln(&column)
			if err != nil {
				fmt.Printf("error: %s\n", err)
				continue
			}
			err = PlaceToken(&game, column, playerToken)
			if err != nil {
				fmt.Printf("error: %s", err)
				continue
			}
			valid = true
		}

		// player 1 checking
		if won, playerWon := game.CheckWin(column, len(game.board[column])-1, PlayerOneToken); won {
			game.PrintBoard()
			fmt.Printf("player %s won\n", playerWon)
			break
		}

		// player 2 checking
		if won, playerWon := game.CheckWin(column, len(game.board[column])-1, PlayerTwoToken); won {
			game.PrintBoard()
			fmt.Printf("player %s won\n", playerWon)
			break
		}
	}
}
