package objects

import (
	"errors"
	"fmt"
	"strconv"
)

type Grid struct {
	cells       [][]*Cell
	groups      []*Group
	lastEdition *Edition
}

type Edition struct {
	cell        *Cell
	ov          uint8 // old value (before)
	prevEdition *Edition
}

func (g *Grid) CellAt(posx int, posy int) *Cell {
	return g.cells[posy][posx]
}

func (g *Grid) GroupIterator() GroupIterator {
	return NewGroupIterator(g.groups)
}

func (g *Grid) IsComplete() bool {
	for posy := 0; posy < len(g.cells); posy++ {
		for posx := 0; posx < len(g.cells[posy]); posx++ {
			if g.CellAt(posx, posy).Value() == EMPTY_CELL_VALUE {
				return false
			}
		}
	}
	return true
}

func (g *Grid) String() string {
	var str string = ""
	for posy := 0; posy < len(g.cells); posy++ {
		for posx := 0; posx < len(g.cells[posy]); posx++ {
			var c *Cell = g.CellAt(posx, posy)
			var v uint8 = c.Value()
			if c.IsReadOnly() {
				str += fmt.Sprintf("\x1b[31;1m")
			}
			if c.IsEmpty() {
				str += fmt.Sprintf("▢ ")
			} else {
				str += fmt.Sprintf("%d ", v)
			}
			if c.IsReadOnly() {
				str += fmt.Sprintf("\x1b[0m")
			}
		}
		str += fmt.Sprintf("\n")
	}
	return str
}

func (g *Grid) State() *Edition {
	return g.lastEdition
}

func (g *Grid) Edit(c *Cell, nv uint8) {
	if c == nil {
		panic("Invalid edition, the cell is nil")
	}
	// c must be a cell of grid:
	if g.CellAt(int(c.posx), int(c.posy)) != c {
		panic("Invalid edition, editor can only edit the cells in the grid")
	}

	if c.IsReadOnly() {
		panic("Invalid edition, the cell is read-only")
	}

	e := &Edition{
		cell: c,
		ov:   c.Value(),
	}
	c.setValue(nv, false)

	e.prevEdition = g.lastEdition
	g.lastEdition = e
}

func (g *Grid) HasUndo() bool {
	return g.lastEdition != nil
}

func (g *Grid) Undo() {
	if !g.HasUndo() {
		panic("no undo possible")
	}

	e := g.lastEdition
	g.lastEdition = e.prevEdition
	e.prevEdition = nil

	e.cell.setValue(e.ov, true)
}

func NewGrid9x9() *Grid {
	var g *Grid = &Grid{}

	var lines = make([]*Group, 9)
	var columns = make([]*Group, 9)
	for i := 0; i < 9; i++ {
		lines[i] = &Group{}
		columns[i] = &Group{}
	}
	g.groups = append(g.groups, lines...)
	g.groups = append(g.groups, columns...)

	var regions = make([][]*Group, 3)
	for i := 0; i < 3; i++ {
		regions[i] = make([]*Group, 3)
		regions[i][0] = &Group{}
		regions[i][1] = &Group{}
		regions[i][2] = &Group{}
		g.groups = append(g.groups, regions[i]...)
	}

	g.cells = make([][]*Cell, 9)
	for posy := 0; posy < 9; posy++ {
		g.cells[posy] = make([]*Cell, 9)
		for posx := 0; posx < 9; posx++ {
			var cell *Cell = NewCell(uint8(posx), uint8(posy))
			g.cells[posy][posx] = cell
			lines[posy].Insert(cell)
			columns[posx].Insert(cell)
			regions[posy/3][posx/3].Insert(cell)
		}
	}
	return g
}

func LoadGrid9x9(values string) (*Grid, error) {
	if len(values) != 81 {
		return nil, errors.New("Invalid input format")
	}

	var grid *Grid = NewGrid9x9()
	for i := 0; i < len(values); i++ {
		n, _ := strconv.ParseUint(values[i:i+1], 10, 8)
		if uint8(n) != EMPTY_CELL_VALUE {
			var c *Cell = grid.CellAt(i%9, i/9)
			c.value = uint8(n)
			c.MarkReadOnly()
		}
	}
	return grid, nil
}
