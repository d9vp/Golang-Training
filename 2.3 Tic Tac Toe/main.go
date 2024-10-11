// main.go
package main

import (
	"fmt"
	"games/tictactoe"
)

func main() {
	//getting player names
	player1, player2 := tictactoe.GetPlayerNames()

	game, err := tictactoe.NewTicTacToe(player1, player2)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		for i := 0; i < 9; i++ {
			for {
				command, row, col := tictactoe.GetUserInput()

				if command == "R" {
					game.Reset()
					break
				}
				err := game.Play(row, col)
				if err == nil {
					break
				} else {
					fmt.Println(err)
				}
			}

			if i >= 4 {
				winner := game.CheckWinner()
				if winner != "" {
					fmt.Println(winner, "wins!")
					return
				}
			}
		}

		fmt.Println("It's a draw! Would you like to play again? Press 'R' to reset or exit to quit.")
	}
}
