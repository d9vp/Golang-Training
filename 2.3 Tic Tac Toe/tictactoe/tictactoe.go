// tictactoe/tictactoe.go
package tictactoe

import (
	"errors"
	"fmt"
	"games/board"
	"os"
	"strings"
)

type TicTacToe struct {
	board   *board.Board
	player1 string
	player2 string
	turn    int
}

// NewTicTacToe factory function
func NewTicTacToe(player1Name, player2Name string) (*TicTacToe, error) {
	if player1Name == "" || player2Name == "" {
		return nil, errors.New("player names cannot be empty")
	}
	b, err := board.NewBoard()
	if err != nil {
		return nil, err
	}
	return &TicTacToe{
		board:   b,
		player1: player1Name,
		player2: player2Name,
		turn:    1,
	}, nil
}

func GetPlayerNames() (string, string) {
	var player1, player2 string

	for {
		fmt.Print("Enter Player 1 name: ")
		fmt.Scan(&player1)
		if player1 != "" {
			break
		}
		fmt.Println("Player 1 name cannot be empty. Please try again.")
	}

	for {
		fmt.Print("Enter Player 2 name: ")
		fmt.Scan(&player2)
		if player2 != "" {
			break
		}
		fmt.Println("Player 2 name cannot be empty. Please try again.")
	}

	return player1, player2
}

func GetUserInput() (string, int, int) {
	var input string
	fmt.Print("Enter your move (row,col) or 'R' to reset: ")
	fmt.Scanln(&input)

	if strings.ToLower(input) == "exit" {
		os.Exit(0)
	}

	//resetting
	if strings.ToUpper(input) == "R" {
		return "R", -1, -1
	}

	input = strings.Replace(input, ",", " ", -1)
	parts := strings.Fields(input)

	//validation
	if len(parts) != 2 {
		fmt.Println("Invalid input. Please enter the row and column as 'row,col' or 'row col'.")
		return GetUserInput()
	}

	var row, col int
	_, err1 := fmt.Sscanf(parts[0], "%d", &row)
	_, err2 := fmt.Sscanf(parts[1], "%d", &col)

	if err1 != nil || err2 != nil || row < 1 || row > 3 || col < 1 || col > 3 {
		fmt.Println("Invalid move. Please enter valid row and column numbers between 1 and 3.")
		return GetUserInput()
	}

	return "", row - 1, col - 1
}

func (g *TicTacToe) Reset() {
	newBoard, _ := board.NewBoard()
	g.board = newBoard
	g.turn = 1
	fmt.Println("The game has been reset.")
	g.board.PrintBoard()
}

func (g *TicTacToe) GetCurrentPlayer() string {
	if g.turn%2 == 1 {
		return g.player1
	}
	return g.player2
}

func (g *TicTacToe) GetPlayerSymbol(player string) string {
	if player == g.player1 {
		return "X"
	}
	return "O"
}

// IsValidMove checks if a move is valid
func (g *TicTacToe) IsValidMove(index int) bool {
	cell := g.board.GetCell(index)
	return cell.GetValue() == "-"
}

// Play implements the Game interface
func (g *TicTacToe) Play(parameters ...interface{}) error {
	row, ok1 := parameters[0].(int)
	col, ok2 := parameters[1].(int)

	if !ok1 || !ok2 || row < 0 || row > 2 || col < 0 || col > 2 {
		return errors.New("invalid parameters")
	}

	index := (row * 3) + col
	player := g.GetCurrentPlayer()

	if g.IsValidMove(index) {
		symbol := g.GetPlayerSymbol(player)
		g.board.SetCellValue(index, symbol)
		g.board.PrintBoard()

		if winner := g.CheckWinner(); winner != "" {
			fmt.Println(player, "wins!")
		}
		g.turn++
		return nil
	} else {
		return errors.New("invalid move try again")
	}
}

// CheckWinner checks for a winning condition
func (g *TicTacToe) CheckWinner() string {
	for i := 0; i < 3; i++ {
		// Check rows
		if g.board.GetCell(i*3).GetValue() == g.board.GetCell(i*3+1).GetValue() &&
			g.board.GetCell(i*3+1).GetValue() == g.board.GetCell(i*3+2).GetValue() &&
			g.board.GetCell(i*3).GetValue() != "-" {
			return g.board.GetCell(i * 3).GetValue()
		}

		// Check columns
		if g.board.GetCell(i).GetValue() == g.board.GetCell(i+3).GetValue() &&
			g.board.GetCell(i+3).GetValue() == g.board.GetCell(i+6).GetValue() &&
			g.board.GetCell(i).GetValue() != "-" {
			return g.board.GetCell(i).GetValue()
		}
	}

	// Check diagonals
	if g.board.GetCell(0).GetValue() == g.board.GetCell(4).GetValue() &&
		g.board.GetCell(4).GetValue() == g.board.GetCell(8).GetValue() &&
		g.board.GetCell(0).GetValue() != "-" {
		return g.board.GetCell(0).GetValue()
	}
	if g.board.GetCell(2).GetValue() == g.board.GetCell(4).GetValue() &&
		g.board.GetCell(4).GetValue() == g.board.GetCell(6).GetValue() &&
		g.board.GetCell(2).GetValue() != "-" {
		return g.board.GetCell(2).GetValue()
	}

	return ""
}
