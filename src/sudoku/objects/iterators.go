package objects

type CellIterator struct {
	cells  []*Cell
	cursor int
}

func NewCellIterator(cells []*Cell) CellIterator {
	return CellIterator{
		cells:  cells,
		cursor: 0,
	}
}

func (ci *CellIterator) HasNext() bool {
	return ci.cursor < len(ci.cells)
}

func (ci *CellIterator) Next() *Cell {
	if ci.HasNext() {
		c := ci.cells[ci.cursor]
		ci.cursor++
		return c
	}
	return nil
}

type GroupIterator struct {
	groups []*Group
	cursor int
}

func NewGroupIterator(groups []*Group) GroupIterator {
	return GroupIterator{
		groups: groups,
		cursor: 0,
	}
}

func (gi *GroupIterator) HasNext() bool {
	return gi.cursor < len(gi.groups)
}

func (gi *GroupIterator) Next() *Group {
	if gi.HasNext() {
		g := gi.groups[gi.cursor]
		gi.cursor++
		return g
	}
	return nil
}
