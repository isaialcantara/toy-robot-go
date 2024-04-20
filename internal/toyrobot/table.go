package toyrobot

type Table struct {
	objectStore

	Width  int
	Length int
}

func NewTable(width int, length int) Table {
	store := newStore()

	return Table{
		objectStore: &store,
		Width:       width,
		Length:      length,
	}
}

func (t *Table) inBounds(pos Position) bool {
	return (pos.X >= 0 && pos.X < t.Width) && (pos.Y >= 0 && pos.Y < t.Length)
}
