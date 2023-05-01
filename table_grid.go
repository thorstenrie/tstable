package lpstr

import (
	"errors"
	"strings"
)

func (t *Table) hline() (string, error) {
	text := ""
	if t == nil {
		return text, errors.New("nil pointer")
	}
	if !t.grid {
		return text, nil
	}
	if t.width == nil {
		return text, errors.New("nil pointer")
	}
	if t.padding < 0 {
		return text, errors.New("invalid padding")
	}
	text += strings.Repeat(" ", t.padding)
	var e error
	for _, w := range t.width {
		if w < 0 {
			e = errors.New("invalid width")
			break
		}
		text += "+" + strings.Repeat("-", t.padding+t.padding+w)
	}
	text += "+\n"
	return text, e
}

func (t *Table) vline() (string, error) {
	text := ""
	if t == nil {
		return text, errors.New("nil pointer")
	}
	if !t.grid {
		return text, nil
	}
	if t.padding < 0 {
		return text, errors.New("invalid padding")
	}
	text += strings.Repeat(" ", t.padding) + "|"
	return text, nil
}
