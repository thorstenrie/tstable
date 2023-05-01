package lpstr

import "errors"

func (t *Table) find(h string) (int, error) {
	if t == nil {
		return -1, errors.New("nil pointer")
	}
	if (t.header == nil) || (len(t.header) == 0) {
		return -1, errors.New("empty header")
	}
	for i, v := range t.header {
		if v == h {
			return i, nil
		}
	}
	return -1, errors.New("not found")
}
