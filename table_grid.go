// Copyright (c) 2023 thorstenrie.
// All Rights Reserved. Use is governed with GNU Affero General Public LIcense v3.0
// that can be found in the LICENSE file.
package lpstr

// Import Go standard library package strings and tserr
import (
	"strings" // strings

	"github.com/thorstenrie/tserr" // tserr
)

const (
	hmchar   = "\u2500" // horizontal 						─
	htchar   = "\u2550" // horizontal top					═
	hbchar   = "\u2550" // horizontal bottom				═
	vmchar   = "\u2502" // vertical 						│
	vlchar   = "\u2551" // vertical left					║
	vrchar   = "\u2551" // vertical right					║
	hvmchar  = "\u253C" // horizontal vertical 				┼
	hvlchar  = "\u255F" // horizontal vertical left			╟
	hvrchar  = "\u2562" // horizontal vertical right		╢
	hvtchar  = "\u2564" // horizontal vertical top			╤
	hvbchar  = "\u2567" // horizontal vertical bottom		╧
	hvtlchar = "\u2554" // horizontal vertical top left		╔
	hvblchar = "\u255A" // horizontal vertical bottom left	╚
	hvtrchar = "\u2557" // horizontal vertical top right	╗
	hvbrchar = "\u255D" // horizontal vertical bottom right ╝
)

// hline returns a horizontal grid line for table t as a string.
func (t *Table) hline(r int) (string, error) {
	// Initialize return value as empty string
	text := ""
	// Return an empty string and an error if t is nil
	if t == nil {
		return text, tserr.NilPtr()
	}
	// Return an empty string and nil if the table grid is switched off
	if !t.grid {
		return text, nil
	}
	// Return an empty string and an error if width is nil
	if t.width == nil {
		return text, tserr.NilPtr()
	}
	// Return an empty string and an error if rows is nil
	if t.rows == nil {
		return "", tserr.NilPtr()
	}
	// Maximum number of horizontal lines
	rmax := len(t.rows) + 1
	// Return an empty string and an error if r is higher than the maximum number of horizontal lines
	if r > rmax {
		return "", tserr.Lower(&tserr.LowerArgs{Var: "row index", Actual: int64(r), HigherBound: int64(rmax)})
	}
	// Return an empty string and an error if padding is negative
	if t.padding < 0 {
		return text, tserr.Higher(&tserr.HigherArgs{Var: "padding", Actual: int64(t.padding), LowerBound: 0})
	}
	// Add initial padding to the horizontal line
	text += strings.Repeat(" ", t.padding)
	// Iterate width column by column
	for i, w := range t.width {
		// Return an empty string and an error if width is negative
		if w < 0 {
			return "", tserr.Higher(&tserr.HigherArgs{Var: "width", Actual: int64(w), LowerBound: 0})
		}
		// Retrieve the horizontal vertical rune
		hv, e := t.hvchar(r, i)
		// Return an empty string and an error if hvchar fails
		if e != nil {
			return "", tserr.Op(&tserr.OpArgs{Op: "hvchar", Fn: "table", Err: e})
		}
		// Retrieve the horizontal line rune
		h, e := t.hchar(r)
		// Return an empty string and an error if hchar fails
		if e != nil {
			return "", tserr.Op(&tserr.OpArgs{Op: "hchar", Fn: "table", Err: e})
		}
		// Add a horizontal line runes for each column of the table
		text += hv + strings.Repeat(h, t.padding+t.padding+w)
	}
	// Retrieve the horizontal vertical rune at the end of the horizontal line
	hv, e := t.hvchar(r, len(t.header))
	// Return an empty string and an error if hvchar fails
	if e != nil {
		return "", tserr.Op(&tserr.OpArgs{Op: "hvchar", Fn: "table", Err: e})
	}
	// End rune of the horizontal line
	text += hv + "\n"
	// Return the horizontal grid line and nil
	return text, nil
}

// hchar returns the horizontal grid line rune for row r. The returned rune depends on whether
// it is the top line with r = 0, the bottom line with r = rmax or a line in between. It returns an
// empty string and an error if r is negative or r is higher than rmax.
func (t *Table) hchar(r int) (string, error) {
	// Return an empty string and an error if t is nil
	if t == nil {
		return "", tserr.NilPtr()
	}
	// Return an empty string and an error if r is negative
	if r < 0 {
		return "", tserr.Higher(&tserr.HigherArgs{Var: "row index", Actual: int64(r), LowerBound: 0})
	}
	// Return an empty string and an error if rows is nil
	if t.rows == nil {
		return "", tserr.NilPtr()
	}
	// Maximum number of horizontal grid lines
	rmax := len(t.rows) + 1
	// Return an empty string and an error if r is higher tham rmax
	if r > rmax {
		return "", tserr.Lower(&tserr.LowerArgs{Var: "row index", Actual: int64(r), HigherBound: int64(rmax)})
	}
	// Return the top horizontal grid line rune for the top line
	if r == 0 {
		return htchar, nil
	}
	// Return the bottom horizontal grid line rune for the bottom line
	if r == rmax {
		return hbchar, nil
	}
	// Return the horizontal grid line rune
	return hmchar, nil
}

// hvchar returns the horizontal vertical grid line for row r and column c. It returns an
// an empty string and an error if r or c are negative or are higher than rmax or cmax.
func (t *Table) hvchar(r, c int) (string, error) {
	// Return an empty string and an error if t is nil
	if t == nil {
		return "", tserr.NilPtr()
	}
	// Return an empty string and an error if r is negative
	if r < 0 {
		return "", tserr.Higher(&tserr.HigherArgs{Var: "row index", Actual: int64(r), LowerBound: 0})
	}
	// Return an empty string and an error fi c is negative
	if c < 0 {
		return "", tserr.Higher(&tserr.HigherArgs{Var: "column index", Actual: int64(c), LowerBound: 0})
	}
	// Return an empty string and an error if rows is nil
	if t.rows == nil {
		return "", tserr.NilPtr()
	}
	// Return an empty string and an error if header is nil
	if t.header == nil {
		return "", tserr.NilPtr()
	}
	// Maximum number of horizontal grid lines
	rmax := len(t.rows) + 1
	// Maximum number of vertical grid lines
	cmax := len(t.header)
	// Return an empty string and an error if r is higher than rmax
	if r > rmax {
		return "", tserr.Lower(&tserr.LowerArgs{Var: "row index", Actual: int64(r), HigherBound: int64(rmax)})
	}
	// Return an empty string and an error if c is higher than cmax
	if c > cmax {
		return "", tserr.Lower(&tserr.LowerArgs{Var: "column index", Actual: int64(c), HigherBound: int64(cmax)})
	}
	// First horizontal grid line
	if r == 0 {
		// First column
		if c == 0 {
			return hvtlchar, nil
		}
		// Last column
		if c == cmax {
			return hvtrchar, nil
		}
		// Any other column
		return hvtchar, nil
	}
	// Last horizontal grid line
	if r == rmax {
		// First column
		if c == 0 {
			return hvblchar, nil
		}
		// Last column
		if c == cmax {
			return hvbrchar, nil
		}
		// Any other column
		return hvbchar, nil
	}
	// First column of any other horizontal line
	if c == 0 {
		return hvlchar, nil
	}
	// Last grid line of any other horizontal line
	if c == cmax {
		return hvrchar, nil
	}
	// Any other horizontal vertical grid line
	return hvmchar, nil
}

// vline returns a vertical grid line for table t as a string.
func (t *Table) vline(c int) (string, error) {
	// Return an empty string and an error if t is nil
	if t == nil {
		return "", tserr.NilPtr()
	}
	// Return an empty string and nil if the table grid is switched off
	if !t.grid {
		return "", nil
	}
	// Return an empty string and an error if padding is negative
	if t.padding < 0 {
		return "", tserr.Higher(&tserr.HigherArgs{Var: "padding", Actual: int64(t.padding), LowerBound: 0})
	}
	// Return an empty string and an error if c is negative
	if c < 0 {
		return "", tserr.Higher(&tserr.HigherArgs{Var: "row index", Actual: int64(c), LowerBound: 0})
	}
	// Return an empty string and an error if header is nil
	if t.header == nil {
		return "", tserr.NilPtr()
	}
	// Maximum number of vertical grid lines
	cmax := len(t.header)
	// Return an empty string and an error if c is higher than cmax
	if c > cmax {
		return "", tserr.Lower(&tserr.LowerArgs{Var: "column index", Actual: int64(c), HigherBound: int64(cmax)})
	}
	// Set vertical grid line
	vchar := vmchar
	// Update vertical grid line for the first column
	if c == 0 {
		vchar = vlchar
	}
	// Update vertical grid line for the last column
	if c == cmax {
		vchar = vrchar
	}
	// Return vertical grid line with padding and nil
	return strings.Repeat(" ", t.padding) + vchar, nil
}
