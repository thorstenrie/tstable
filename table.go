// Copyright (c) 2023 thorstenrie.
// All Rights Reserved. Use is governed with GNU Affero General Public LIcense v3.0
// that can be found in the LICENSE file.
package lpstr

// Import Go standard packages and lpstats
import (
	"errors"       // errors
	"fmt"          // fmt
	"strings"      // strings
	"unicode/utf8" // utf8

	"github.com/thorstenrie/lpstats" // lpstats
	"github.com/thorstenrie/tserr"
)

// Table holds the header of the table and all rows of the table. It also contains
// information on the width of each row, the row index for sorting, padding and whether
// the table is printed with a grid. Per default, a table has padding 2, a grid and
// is sorted by its first row.
type Table struct {
	header  []string   // Header as a slice of strings
	rows    [][]string // Rows as a slice of slices of strings
	width   []int      // Width of each row
	key     int        // Row index for sorting (default first column)
	padding int        // Padding (default 2)
	grid    bool       // Table grid (default true)
}

// defaultPadding defines the default padding for a new Table
const (
	defaultPadding = 2
)

// New returns a pointer to a new Table. It expects the header of the table
// h as a slice of strings. It returns nil and an error, if h is nil, has
// zero length or contains non-printable runes.
func New(h []string) (*Table, error) {
	// Return nil and an error if h is nil or h has zero length
	if (h == nil) || (len(h) == 0) {
		return nil, tserr.Empty("header")
	}
	// Retrieve whether h contains only printable characters with IsPrintable
	p, e := IsPrintable(h)
	// Return nil and an error if IsPrintable fails
	if e != nil {
		return nil, tserr.Op(&tserr.OpArgs{Op: "IsPrintable", Fn: "header", Err: e})
	}
	if !p {
		return nil, errors.New("non-printable characters in header")
	}
	t := &Table{
		padding: defaultPadding,
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
