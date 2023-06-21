// Copyright (c) 2023 thorstenrie.
// All Rights Reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package lpstr_test

// Import Go standard library package testing as well as lpstr, tserr and tsfio
import (
	"testing" // testing

	"github.com/thorstenrie/lpstr" // lpstr
	"github.com/thorstenrie/tserr" // tserr
	"github.com/thorstenrie/tsfio" // tsfio
)

var (
	header  = []string{"Fellowship member", "Title", "Weapon"}           // Test table header
	gandalf = []string{"Gandalf", "The Grey", "Wizard staff"}            // Test table row
	aragorn = []string{"Aragorn", "King of Gondor", "Sword"}             // Test table row
	legolas = []string{"Legolas", "Prince of the Woodland Realm", "Bow"} // Test table row
	gimli   = []string{"Gimli", "Lord of the Glittering Caves", "Axe"}   // Test table row
	boromir = []string{"Boromir", "Captain of the White Tower", "Sword"} // Test table row
	sortby  = "Weapon"                                                   // Test table sort by column
	padding = 1                                                          // Test table padding
)

func loadTable(t *testing.T) *lpstr.Table {
	tbl, e := lpstr.New(header)
	if e != nil {
		t.Fatal(tserr.Op(&tserr.OpArgs{Op: "New", Fn: "table", Err: e}))
	}
	if e := tbl.AddRow(gandalf); e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "AddRow", Fn: "table", Err: e}))
	}
	if e := tbl.AddRow(aragorn); e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "AddRow", Fn: "table", Err: e}))
	}
	if e := tbl.AddRow(legolas); e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "AddRow", Fn: "table", Err: e}))
	}
	if e := tbl.AddRow(gimli); e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "AddRow", Fn: "table", Err: e}))
	}
	if e := tbl.SetPadding(padding); e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "SetPadding", Fn: "table", Err: e}))
	}
	if e := tbl.SortBy(sortby); e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "SortBy", Fn: "table", Err: e}))
	}
	if e := tbl.AddRow(boromir); e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "AddRow", Fn: "table", Err: e}))
	}
	return tbl
}

func goldenFn(name string) tsfio.Filename {
	return tsfio.Filename("testdata/" + name + ".golden")
}

func testData(name string, t *testing.T) string {
	if t == nil {
		panic(tserr.NilPtr())
	}
	fn := goldenFn(name)
	b, e := tsfio.ReadFile(fn)
	if e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "ReadFile", Fn: string(fn), Err: e}))
	}
	s := string(b)
	return s
}
