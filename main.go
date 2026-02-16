package main

import "fmt"

func main() {
	fmt.Println("Welcome to Connect 4!")
	game := Game{}
	InitBoard(&game)

	for {
		var column int
		game.PrintBoard()

		fmt.Print("Enter column: ")
		_, err := fmt.Scanln(&column)
		if err != nil {
			fmt.Printf("error: %s", err)
		}

		err = PlaceToken(&game, column, PlayerOneToken)
		if err != nil {
			fmt.Printf("error: %s", err)
			return
		}

	}

}
