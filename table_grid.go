// Copyright (c) 2023 thorstenrie.
// All Rights Reserved. Use is governed with GNU Affero General Public LIcense v3.0
// that can be found in the LICENSE file.
package lpstr

// Import Go standard library package strings and tserr
import (
	"strings" // strings

	"github.com/thorstenrie/tserr" // tserr
)

// hline returns a horizontal grid line for table t as a string.
func (t *Table) hline() (string, error) {
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
	// Return an empty string and an error if padding is negative
	if t.padding < 0 {
		return text, tserr.Higher(&tserr.HigherArgs{Var: "padding", Actual: int64(t.padding), LowerBound: 0})
	}
	// Add initial padding to the horizontal line
	text += strings.Repeat(" ", t.padding)
	// Iterate width
	for _, w := range t.width {
		// Return an empty string and an error if width is negative
		if w < 0 {
			return "", tserr.Higher(&tserr.HigherArgs{Var: "width", Actual: int64(w), LowerBound: 0})
		}
		// Add a horizontal line runes for each column of the table
		text += "+" + strings.Repeat("-", t.padding+t.padding+w)

	}
	// End rune of the horizontal line
	text += "+\n"
	// Return the horizontal grid line and nil
	return text, nil
}

// vline returns a vertical grid line for table t as a string.
func (t *Table) vline() (string, error) {
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
	// Return vertical grid line with padding and nil
	return strings.Repeat(" ", t.padding) + "|", nil
}
