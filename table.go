// Copyright (c) 2023 thorstenrie.
// All Rights Reserved. Use is governed with GNU Affero General Public LIcense v3.0
// that can be found in the LICENSE file.
package lpstr

// Import Go standard packages and lpstats
import (
	"errors" // errors
	// fmt
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
	// Retrieve whether h contains only printable runes with IsPrintable
	p, e := IsPrintable(h)
	// Return nil and an error if IsPrintable fails
	if e != nil {
		return nil, tserr.Op(&tserr.OpArgs{Op: "IsPrintable", Fn: "header", Err: e})
	}
	// Return nil and an error if the header contains non-printable runes
	if !p {
		return nil, tserr.NonPrintable("header")
	}
	// Retrieve a new instance of struct Table
	t := &Table{
		padding: defaultPadding,      // default padding
		grid:    true,                // with table grid
		header:  h,                   // set header
		rows:    make([][]string, 0), // allocate and initialize rows
		width:   make([]int, len(h)), // allocate and initialize width
		key:     0,                   // set sort key to first column
	}
	// Iterate over elements of h
	for i, c := range h {
		// Set width of column to number of runes in element c of h
		t.width[i] = utf8.RuneCountInString(c)
	}
	// Return pointer to Table
	return t, nil
}

// AddRow appends a row r at the end of the rows of table t. The row r is provided by a slice of strings. Row r must
// contain the same number of elements as the table header. The order of elements must match
// the order of columns defined by the table header. It returns
// an error if t is nil, r is nil or empty or if the number of elements in r does not equal the
// number of elements in the table header.
func (t *Table) AddRow(r []string) error {
	// Return an error if t is nil
	if t == nil {
		return tserr.NilPtr()
	}
	// Return an error if r is nil or r is empty
	if (r == nil) || (len(r) == 0) {
		return tserr.Empty("row")
	}
	// Return an error if the number of elements in r does not equal to the number of elements of the table header
	if len(r) != len(t.header) {
		return tserr.Equal(&tserr.EqualArgs{Var: "row", Actual: int64(len(r)), Want: int64(len(t.header))})
	}
	// Return an error if the number of elements in r does not equal to the number of elements of width
	if len(r) != len(t.width) {
		return tserr.Equal(&tserr.EqualArgs{Var: "row", Actual: int64(len(r)), Want: int64(len(t.width))})
	}
	// Retrieve in p whether r only contains printable runes
	p, e := IsPrintable(r)
	// If IsPrintable returns an error, return that error
	if e != nil {
		return tserr.Op(&tserr.OpArgs{Op: "IsPrintable", Fn: "row", Err: e})
	}
	// Return an error if row r contains non-printable runes
	if !p {
		return tserr.NonPrintable("row")
	}
	// Append row r at the end of the rows of the table
	t.rows = append(t.rows, r)
	// Iterate all elements of row r
	for i, c := range r {
		// Set the width of column i to the maximum of current width of column i and number or runes in element i of row r
		// This adjusts the width of column i, if c contains more runes than previous elements of column i
		t.width[i] = lpstats.Max(t.width[i], utf8.RuneCountInString(c))
	}
	// Return nil
	return nil
}

// Print returns the contents of table t in a string representation. The formatting
// of the table can be altered by changing the padding with SetPadding or disabling the table grid
// with WithoutGrid. The rows are sorted in alphabetical order according to the selected column with
// SortBy. Per default, it is sorted by the first column.
func (t *Table) Print() (string, error) {
	// Initialize return value with an empty string
	text := ""
	// Return an empty string and an error, if t is nil
	if t == nil {
		return text, tserr.NilPtr()
	}
	// Return an empty string and an error, if header or rows are nil
	if (t.header == nil) || (t.rows == nil) {
		return text, tserr.NilPtr()
	}
	// Return an empty string and an error, if the number of elements in header does not equal the number of elements in width
	if len(t.header) != len(t.width) {
		return text, tserr.Equal(&tserr.EqualArgs{Var: "table width slice", Actual: int64(len(t.width)), Want: int64(len(t.header))})
	}
	// Return an empty string and an error, if padding is negative
	if t.padding < 0 {
		return text, tserr.Higher(&tserr.HigherArgs{Var: "padding", Actual: int64(t.padding), LowerBound: 0})
	}
	// Sort table by selected row, which is given by the row index in struct field key
	if err := t.sort(); err != nil {
		// Return an empty string and an error, if sorting fails
		return text, tserr.Op(&tserr.OpArgs{Op: "sort", Fn: "table", Err: err})
	}
	// Retrieve horizontal grid line
	hline, e := t.hline()
	// Return an empty string and an error, if hline fails
	if e != nil {
		return text, tserr.Op(&tserr.OpArgs{Op: "hline", Fn: "table", Err: e})
	}
	// Retrieve vertical grid lin
	vline, e := t.vline()
	// Return an empty string and an error, if vline fails
	if e != nil {
		return text, tserr.Op(&tserr.OpArgs{Op: "vline", Fn: "table", Err: e})
	}
	// Add horizontal grid line to string
	text += hline
	// Print header
	for i, h := range t.header {
		// Add header formatting if table is printed without grid
		if !t.grid {
			h = "[" + h + "]"
		}
		// Return an empty string and an error, if the difference of width of column i and length of h is negative
		if t.width[i]-len(h) < 0 {
			return "", tserr.Higher(&tserr.HigherArgs{Var: "width", Actual: int64(t.width[i]), LowerBound: int64(len(h))})
		}
		// Add header of column to return string
		text += vline + strings.Repeat(" ", t.padding) + h + strings.Repeat(" ", t.width[i]-len(h))
	}
	// Add horizontal grid line to return string
	text += vline + "\n" + hline
	// Print rows
	for _, r := range t.rows {
		// Return an empty string and an error, if the size of row r is not equal to the size of width
		if len(t.width) != len(r) {
			return "", tserr.Higher(&tserr.HigherArgs{Var: "sizte of row", Actual: int64(len(r)), LowerBound: int64(len(t.width))})
		}
		// Print row r
		for j, c := range r {
			// TODO
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
