// board/board.go
package board

import (
	"errors"
	"fmt"
	"games/cell"
)

type Board struct {
	cells [9]*cell.Cell
}

// NewBoard factory function
func NewBoard() (*Board, error) {
	b := &Board{}
	for i := 0; i < 9; i++ {
		b.cells[i] = cell.NewCell()
	}
	return b, nil
}

// GetCell returns the cell at a specific index
func (b *Board) GetCell(index int) *cell.Cell {
	if index < 0 || index > 8 {
		return nil
	}
	return b.cells[index]
}

// SetCellValue sets the value for a specific cell on the board
func (b *Board) SetCellValue(index int, val string) error {
	if index < 0 || index >= 9 {
		return errors.New("invalid index")
	}
	return b.cells[index].SetValue(val)
}

// PrintBoard prints the current state of the board
func (b *Board) PrintBoard() {
	for i := 0; i < 9; i += 3 {
		fmt.Println(b.cells[i].GetValue(), b.cells[i+1].GetValue(), b.cells[i+2].GetValue())
	}
}
