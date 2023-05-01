package lpstr

import (
	"errors"
	"sort"
)

type byRow Table

func (t byRow) Len() int {
	if t.rows == nil {
		return 0
	}
	return len(t.rows)
}

func (t byRow) Swap(i, j int) {
	if t.rows == nil {
		return
	}
	if (i >= len(t.rows)) || (j >= len(t.rows)) {
		return
	}
	t.rows[i], t.rows[j] = t.rows[j], t.rows[i]
}

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
