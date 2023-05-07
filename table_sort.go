// Copyright (c) 2023 thorstenrie.
// All Rights Reserved. Use is governed with GNU Affero General Public LIcense v3.0
// that can be found in the LICENSE file.
package lpstr

// Import Go standard packages
import (
	"errors"
	"sort"
)

// ByRow implements sort.Interface fo Table based on
// the selected row, which is given by the row index in field key. Initially,
// the first row is selected. The selected row can be changed with function SortBy.
type byRow Table

// Len returns the number of rows in Table
func (t byRow) Len() int {
	// Return zero in case rows is nil
	if t.rows == nil {
		return 0
	}
	// Otherwise, return the length of rows
	return len(t.rows)
}

// Swap swaps the rows with indexes i and j.
func (t byRow) Swap(i, j int) {
	// Return without swap if rows is nil
	if t.rows == nil {
		return
	}
	// Return without swap if i or j is higher or equal to the length of rows
	if (i >= len(t.rows)) || (j >= len(t.rows)) {
		return
	}
	// Swap rows with indexes i and j
	t.rows[i], t.rows[j] = t.rows[j], t.rows[i]
}

// Less reports whether the row with index i must
// sort before the row with index j.
func (t byRow) Less(i, j int) bool {
	if t.rows == nil {
		return false
	}
	if (i >= len(t.rows)) || (j >= len(t.rows)) {
		return false
	}
	if (t.key >= len(t.rows[i])) || (t.key >= len(t.rows[j])) {
		return false
	}
	return t.rows[i][t.key] < t.rows[j][t.key]
}

func (t *Table) sort() error {
	if t == nil {
		return errors.New("nil pointer")
	}
	sort.Sort(byRow(*t))
	return nil
}
