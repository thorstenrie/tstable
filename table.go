package lpstr

import (
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/thorstenrie/lpstats"
)

type Table struct {
	header  []string
	rows    [][]string
	width   []int
	key     int
	padding int
	grid    bool
}

const (
	DefaultPadding = 2
)

func New(h []string) (*Table, error) {
	if (h == nil) || (len(h) == 0) {
		return nil, errors.New("empty header")
	}
	p, e := IsPrintable(h)
	if e != nil {
		return nil, e
	}
	if !p {
		return nil, errors.New("non-printable characters in header")
	}
	t := &Table{
		padding: DefaultPadding,
		grid:    true,
		header:  h,
		rows:    make([][]string, 0),
		width:   make([]int, len(h)),
		key:     0,
	}
	for i, c := range h {
		t.width[i] = utf8.RuneCountInString(c)
	}
	return t, nil
}

func (t *Table) AddRow(r []string) error {
	if t == nil {
		return errors.New("nil pointer")
	}
	if (r == nil) || (len(r) == 0) {
		return errors.New("empty row")
	}
	if len(r) != len(t.header) {
		return fmt.Errorf("expected %d columns, but row has %d columns", len(t.header), len(r))
	}
	p, e := IsPrintable(r)
	if e != nil {
		return e
	}
	if !p {
		return errors.New("row contains non-printable characters")
	}
	t.rows = append(t.rows, r)
	for i, c := range r {
		t.width[i] = lpstats.Max(t.width[i], utf8.RuneCountInString(c))
	}
	return nil
}

func (t *Table) Print() (string, error) {
	text := ""
	if t == nil {
		return text, errors.New("nil pointer")
	}
	if (t.header == nil) || (t.rows == nil) {
		return text, errors.New("nil pointer")
	}
	if len(t.header) != len(t.width) {
		return text, errors.New("invalid columns")
	}
	if t.padding < 0 {
		return text, errors.New("invalid padding")
	}
	if err := t.sort(); err != nil {
		return text, err
	}
	hline, e := t.hline()
	vline, e := t.vline()
	text += hline
	for i, h := range t.header {
		if !t.grid {
			h = "[" + h + "]"
		}
		if t.width[i]-len(h) < 0 {
			e = errors.New("invalid width")
			break
		}
		text += vline + strings.Repeat(" ", t.padding) + h + strings.Repeat(" ", t.width[i]-len(h))
	}
	text += vline + "\n" + hline
	for _, r := range t.rows {
		if len(t.width) != len(r) {
			e = errors.New("incompatible row")
			break
		}
		for j, c := range r {
			text += vline + strings.Repeat(" ", t.padding) + c + strings.Repeat(" ", t.width[j]-len(c))
		}
		text += vline + "\n"
	}
	return text + hline, e
}

func (t *Table) SortBy(h string) error {
	if t == nil {
		return errors.New("nil pointer")
	}
	if (t.header == nil) || (len(t.header) == 0) {
		return errors.New("empty header")
	}
	i, e := t.find(h)
	if e != nil {
		return e
	}
	t.key = i
	return nil
}

func (t *Table) SetPadding(p int) error {
	if t == nil {
		return errors.New("nil pointer")
	}
	if p < 0 {
		return errors.New("padding must be positive")
	}
	t.padding = p
	return nil
}

func (t *Table) WithoutGrid() error {
	if t == nil {
		return errors.New("nil pointer")
	}
	if (t.width == nil) || (t.header == nil) {
		return errors.New("nil pointer")
	}
	if len(t.width) != len(t.header) {
		return errors.New("invalid number of columns")
	}
	for i, w := range t.width {
		if len(t.header[i]) == w {
			t.width[i] += 2
		}
	}
	t.grid = false
	return nil
}
