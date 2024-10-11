package cell

import "errors"

type Cell struct {
	value string
}

// NewCell factory function
func NewCell() *Cell {
	return &Cell{value: "-"}
}

// GetValue returns the current value of the cell
func (c *Cell) GetValue() string {
	return c.value
}

// SetValue sets the value of the cell (X, O, or -)
func (c *Cell) SetValue(val string) error {
	if val != "X" && val != "O" && val != "-" {
		return errors.New("invalid cell value")
	}
	c.value = val
	return nil
}
