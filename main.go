package main

import "fmt"

func main() {
	fmt.Println("Welcome to Connect 4!")
	game := Game{}
	InitBoard(&game)
	var playerToken Token

	for {
		var column int
		if playerToken == PlayerOneToken {
			playerToken = PlayerTwoToken
		} else {
			playerToken = PlayerOneToken
		}
		game.PrintBoard()

		for valid := false; !valid; {
			fmt.Printf("[player %v] enter column: ", playerToken)
			_, err := fmt.Scanln(&column)
			if err != nil {
				fmt.Printf("error: %s\n", err)
			} else {
				valid = true
			}
		}

		err := PlaceToken(&game, column, playerToken)
		if err != nil {
			fmt.Printf("error: %s", err)
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
