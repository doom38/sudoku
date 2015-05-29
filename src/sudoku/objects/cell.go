package objects

const (
	EMPTY_CELL_VALUE uint8 = 0
)

type Cell struct {
	value, posx, posy uint8
	readonly          bool
	groups            []*Group
	onChangeCallbacks []func(c *Cell, old uint8, undo bool)
}

func NewCell(posx uint8, posy uint8) *Cell {
	return &Cell{
		posx:     posx,
		posy:     posy,
		readonly: false,
		value:    EMPTY_CELL_VALUE,
	}
}

func (c *Cell) Value() uint8 {
	return c.value
}

func (c *Cell) PosX() uint8 {
	return c.posx
}

func (c *Cell) PosY() uint8 {
	return c.posy
}

func (c *Cell) setValue(v uint8, undo bool) {

	var old uint8 = c.value
	c.value = v

	// Fire change events:
	for _, cb := range c.onChangeCallbacks {
		cb(c, old, undo)
	}
}

func (c *Cell) IsReadOnly() bool {
	return c.readonly
}

func (c *Cell) MarkReadOnly() {
	c.readonly = true
}

func (c *Cell) IsEmpty() bool {
	return c.value == EMPTY_CELL_VALUE
}

func (c *Cell) GroupIterator() GroupIterator {
	return NewGroupIterator(c.groups)
}

func (c *Cell) AddCellChangeCallback(callback func(c *Cell, old uint8, undo bool)) {
	c.onChangeCallbacks = append(c.onChangeCallbacks, callback)
}
