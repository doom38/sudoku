package solverV1

import (
	"fmt"
	"sudoku/objects"
)

type SolverV1 struct {
	grid        *objects.Grid
	constraints []*UnicityConstraint
}

type UnicityConstraint struct {
	cell   *objects.Cell
	values []uint8
}

func New(grid *objects.Grid) *SolverV1 {
	s := &SolverV1{
		grid: grid,
	}

	var c *objects.Cell
	for posx := 0; posx < 9; posx++ {
		for posy := 0; posy < 9; posy++ {
			c = grid.CellAt(posx, posy)
			if !c.IsReadOnly() {
				s.constraints = append(s.constraints, newUnicityConstraint(c))
			}
		}
	}
	return s
}

func newUnicityConstraint(c *objects.Cell) *UnicityConstraint {
	uc := &UnicityConstraint{
		cell: c,
	}

	founds := make([]bool, 10)
	gi := c.GroupIterator()
	var ci objects.CellIterator
	var group *objects.Group
	for gi.HasNext() {
		group = gi.Next()
		group.AddGroupChangeCallback(uc.onCellChange)
		ci = group.CellIterator()
		for ci.HasNext() {
			c := ci.Next()
			if c.Value() != objects.EMPTY_CELL_VALUE {
				founds[c.Value()] = true
			}
		}
	}

	for i, found := range founds {
		if !found && uint8(i) != objects.EMPTY_CELL_VALUE {
			uc.values = append(uc.values, uint8(i))
		}
	}

	return uc
}

func (uc *UnicityConstraint) onCellChange(c *objects.Cell, old uint8, undo bool) {
	if undo && old != objects.EMPTY_CELL_VALUE {
		uc.values = append(uc.values, old)
	}
	if !undo && c.Value() != objects.EMPTY_CELL_VALUE {
		nv := make([]uint8, 0)
		for _, v := range uc.values {
			if v != c.Value() {
				nv = append(nv, v)
			}
		}
		uc.values = nv
	}
}

func (s *SolverV1) Solve() {
	var it int = 0
	var done bool = true
	for done {
		it++
		done = false
		fmt.Printf("#%d:\n", it)
		for _, c := range s.constraints {
			if len(c.values) == 1 {
				fmt.Printf("  [%d,%d] -> %d\n", c.cell.PosX(), c.cell.PosY(), c.values[0])
				s.grid.Edit(c.cell, c.values[0])
				done = true
				break
			}
		}
	}
}

func (s *SolverV1) Dump() {
	for i, c := range s.constraints {
		fmt.Printf("#%.2d [%d, %d] -> %v \n", i, c.cell.PosX(), c.cell.PosY(), c.values)
	}
}
