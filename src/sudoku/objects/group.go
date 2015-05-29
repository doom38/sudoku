package objects

type Group struct {
	cells             []*Cell
	onChangeCallbacks []func(c *Cell, old uint8, undo bool)
}

func (g *Group) Insert(c *Cell) {
	g.cells = append(g.cells, c)
	c.groups = append(c.groups, g)
	c.AddCellChangeCallback(g.onCellChange)
}

func (g *Group) HasValue(v uint8) bool {
	var l int = len(g.cells)
	for i := 0; i < l; i++ {
		if g.cells[i].Value() == v {
			return true
		}
	}
	return false
}

func (g *Group) CellCount() int {
	return len(g.cells)
}

func (g *Group) CellIterator() CellIterator {
	return NewCellIterator(g.cells)
}

func (g *Group) AddGroupChangeCallback(callback func(c *Cell, old uint8, undo bool)) {
	g.onChangeCallbacks = append(g.onChangeCallbacks, callback)
}

func (g *Group) onCellChange(c *Cell, old uint8, undo bool) {
	for _, cb := range g.onChangeCallbacks {
		cb(c, old, undo)
	}
}
