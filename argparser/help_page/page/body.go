package page

import "github.com/terryhay/dolly/tools/index"

// Body contains data for display page
type Body struct {
	rows []Row
}

// MakeBody creates page Body object
func MakeBody(rows []Row) Body {
	if len(rows) == 0 {
		return Body{}
	}

	return Body{
		rows: rows,
	}
}

// RowCount returns row count
func (b *Body) RowCount() index.Index {
	if b == nil {
		return index.Zero
	}
	return index.MakeIndex(len(b.rows))
}

// Row returns row by index
func (b *Body) Row(i index.Index) Row {
	if b == nil || i >= b.RowCount() {
		return Row{}
	}
	return b.rows[i]
}
