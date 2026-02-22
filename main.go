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

		fmt.Printf("[player %v] enter column: ", playerToken)
		_, err := fmt.Scanln(&column)
		if err != nil {
			fmt.Printf("error: %s", err)
		}

		err = PlaceToken(&game, column, playerToken)
		if err != nil {
			fmt.Printf("error: %s", err)
		}
		if won, playerWon := game.CheckWin(column, len(game.board[column])-1); won {
			game.PrintBoard()
			fmt.Printf("player %s won\n", playerWon)
			break
		}
	}
}
