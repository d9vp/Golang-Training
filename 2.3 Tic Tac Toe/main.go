package main

import (
	"fmt"
	"strconv"
)

type Game interface {
	play(parameters ...interface{})
}

type TicTacToe struct {
	board   [9]string
	player1 string
	player2 string
	turn    int
}

// factory
func NewTicTacToe(player1Name, player2Name string) *TicTacToe {
	if player1Name == "" || player2Name == "" {
		panic("Player Names cannot be empty.")
	}
	ttt := &TicTacToe{
		player1: player1Name,
		player2: player2Name,
		turn:    1,
	}

	for i := 0; i < 9; i++ {
		ttt.board[i] = "-"
	}

	return ttt
}

// play method for interface
func (g *TicTacToe) play(parameters ...interface{}) {
	row, ok1 := parameters[0].(int)
	col, ok2 := parameters[1].(int)

	if !ok1 || !ok2 {
		fmt.Println("Invalid parameters")
		return
	}

	player := g.player1
	if g.turn%2 == 0 {
		player = g.player2
	}

	index := (row * 3) + col

	if g.isValidMove(index) {
		g.board[index] = g.getPlayerSymbol(player)
		g.printBoard()
		if g.checkWinner() != "" {
			fmt.Println(player, "wins!")
			return
		}
		g.turn++
	} else {
		fmt.Println("Invalid move. Try again.")
	}
}

func (g *TicTacToe) getPlayerSymbol(player string) string {
	if player == g.player1 {
		return "X"
	}
	return "O"
}

func (g *TicTacToe) isValidMove(index int) bool {
	return index >= 0 && index < 9 && g.board[index] == "-"
}

func (g *TicTacToe) checkWinner() string {
	for i := 0; i < 3; i++ {
		if g.board[i*3] == g.board[i*3+1] && g.board[i*3+1] == g.board[i*3+2] && g.board[i*3] != "-" {
			return g.board[i*3]
		}

		if g.board[i] == g.board[i+3] && g.board[i+3] == g.board[i+6] && g.board[i] != "-" {
			return g.board[i]
		}
	}

	if g.board[0] == g.board[4] && g.board[4] == g.board[8] && g.board[0] != "-" {
		return g.board[0]
	}
	if g.board[2] == g.board[4] && g.board[4] == g.board[6] && g.board[2] != "-" {
		return g.board[2]
	}

	return ""
}

func (g *TicTacToe) printBoard() {
	for i := 0; i < 9; i += 3 {
		fmt.Println(g.board[i : i+3])
	}
}

func getUserInput() (int, int) {
	var input string
	fmt.Print("Enter your move (row,col): ")
	fmt.Scan(&input)

	row, _ := strconv.Atoi(string(input[0]))
	col, _ := strconv.Atoi(string(input[2]))

	row -= 1
	col -= 1

	return row, col
}

// Main function demonstrating the facade pattern
func main() {
	// Initialize and play TicTacToe
	var tttGame Game = NewTicTacToe("Player1", "Player2")

	for i := 0; i < 9; i++ {
		// Get user input
		row, col := getUserInput()

		// Play the move for the current player
		tttGame.play(row, col)

		// After 5 moves, start checking for a winner
		if i >= 4 {
			winner := tttGame.(*TicTacToe).checkWinner()
			if winner != "" {
				fmt.Println(winner, "wins!")
				return
			}
		}
	}

	fmt.Println("It's a draw!")
}
