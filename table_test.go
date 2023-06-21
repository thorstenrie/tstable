// Copyright (c) 2023 thorstenrie.
// All Rights Reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package lpstr_test

// Import Go standard library packages and lpstr
import (
	"fmt"     // fmt
	"testing" // testing

	"github.com/thorstenrie/lpstr" // lpstr
)

// TestMinTable1 tests the string representation of a table with one column and one row. The test fails
// if the retrieved string does not equal to the testdata golden file.
func TestMinTable1(t *testing.T) {
	tbl, _ := lpstr.New([]string{""})
	tbl.AddRow([]string{""})
	tbl.SetPadding(0)
	s, _ := tbl.Print()
	fmt.Println(s)
}

func TestMinTable2(t *testing.T) {
	tbl, _ := lpstr.New([]string{"", ""})
	tbl.AddRow([]string{"", ""})
	tbl.SetPadding(0)
	s, _ := tbl.Print()
	fmt.Println(s)
}

func TestMinTable3(t *testing.T) {
	tbl, _ := lpstr.New([]string{""})
	tbl.SetPadding(0)
	s, _ := tbl.Print()
	fmt.Println(s)
}
