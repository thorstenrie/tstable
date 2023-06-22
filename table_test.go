// Copyright (c) 2023 thorstenrie.
// All Rights Reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package lpstr_test

// Import Go standard library packages and lpstr
import (
	// fmt
	"testing" // testing

	"github.com/thorstenrie/lpstr" // lpstr
	"github.com/thorstenrie/tserr" // tserr
	// tsfio
)

// TestMinTable1 tests the string representation of a table with one column and one row with empty strings as contents. The test fails
// if the retrieved string does not equal to the testdata golden file.
func TestMinTable1(t *testing.T) {
	// Set name of test table
	name := "MinTable1"
	// Retrieve new test table with one column and an empty string in the header
	tbl, e := lpstr.New([]string{""})
	// The test fails, if New returns an error
	if e != nil {
		t.Fatal(tserr.Op(&tserr.OpArgs{Op: "New", Fn: "table", Err: e}))
	}
	// Add one row with an empty string to the table
	if err := tbl.AddRow([]string{""}); err != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "AddRow", Fn: "table", Err: err}))
	}
	// Set padding to zero
	if err := tbl.SetPadding(0); err != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "SetPadding", Fn: "table", Err: err}))
	}
	// Evaluate test table
	evalTable(name, tbl, t)
}

// TestMinTable2 tests the string representation of a table with two columns and one row with empty strings as contents. The test fails
// if the retrieved string does not equal to the testdata golden file.
func TestMinTable2(t *testing.T) {
	// Set name of test table
	name := "MinTable2"
	// Retrieve new test table with two columns and empty strings in the header
	tbl, e := lpstr.New([]string{"", ""})
	// The test fails, if New returns an error
	if e != nil {
		t.Fatal(tserr.Op(&tserr.OpArgs{Op: "New", Fn: "table", Err: e}))
	}
	// Add one row with empty strings to the table
	if err := tbl.AddRow([]string{"", ""}); err != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "AddRow", Fn: "table", Err: err}))
	}
	// Set padding to zero
	if err := tbl.SetPadding(0); err != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "SetPadding", Fn: "table", Err: err}))
	}
	// Evaluate test table
	evalTable(name, tbl, t)
}

// TestMinTable3 tests the string representation of a table with one columns, only an empty string as header and no rows. The test fails
// if the retrieved string does not equal to the testdata golden file.
func TestMinTable3(t *testing.T) {
	// Set name of test table
	name := "MinTable3"
	// Retrieve new test table with two columns and empty strings in the header
	tbl, e := lpstr.New([]string{""})
	// The test fails, if New returns an error
	if e != nil {
		t.Fatal(tserr.Op(&tserr.OpArgs{Op: "New", Fn: "table", Err: e}))
	}
	// Set padding to zero
	if err := tbl.SetPadding(0); err != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "SetPadding", Fn: "table", Err: err}))
	}
	// Evaluate test table
	evalTable(name, tbl, t)
}

// TestMinTable4 tests the string representation of a table with two columns, only containing empty strings in the header and no rows. The test fails
// if the retrieved string does not equal to the testdata golden file.
func TestMinTable4(t *testing.T) {
	// Set name of test table
	name := "MinTable4"
	// Retrieve new test table with two columns and empty strings in the header
	tbl, e := lpstr.New([]string{"", ""})
	// The test fails, if New returns an error
	if e != nil {
		t.Fatal(tserr.Op(&tserr.OpArgs{Op: "New", Fn: "table", Err: e}))
	}
	// Set padding to zero
	if err := tbl.SetPadding(0); err != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "SetPadding", Fn: "table", Err: err}))
	}
	// Evaluate test table
	evalTable(name, tbl, t)
}
