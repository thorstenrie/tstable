// Copyright (c) 2023 thorstenrie.
// All Rights Reserved. Use is governed with GNU Affero General Public LIcense v3.0
// that can be found in the LICENSE file.
package lpstr_test

import (
	"fmt"
	"testing"

	"github.com/thorstenrie/lpstr"
)

var (
	header      = []string{"Fellowship member", "Title", "Weapon"}
	gandalf     = []string{"Gandalf", "The Grey", "Wizard staff"}
	aragorn     = []string{"Aragorn", "King of Gondor", "Sword"}
	legolas     = []string{"Legolas", "Prince of the Woodland Realm", "Bow"}
	gimli       = []string{"Gimli", "Lord of the Glittering Caves", "Axe"}
	boromir     = []string{"Boromir", "Captain of the White Tower", "Sword"}
	sortby      = "Weapon"
	padding     = 1
	refWithGrid = " ╔═══════════════════╤══════════════════════════════╤══════════════╗\n" +
		" ║ Fellowship member │ Title                        │ Weapon       ║\n" +
		" ╟───────────────────┼──────────────────────────────┼──────────────╢\n" +
		" ║ Gimli             │ Lord of the Glittering Caves │ Axe          ║\n" +
		" ║ Legolas           │ Prince of the Woodland Realm │ Bow          ║\n" +
		" ║ Aragorn           │ King of Gondor               │ Sword        ║\n" +
		" ║ Boromir           │ Captain of the White Tower   │ Sword        ║\n" +
		" ║ Gandalf           │ The Grey                     │ Wizard staff ║\n" +
		" ╚═══════════════════╧══════════════════════════════╧══════════════╝\n"
	refWithoutGrid = " Fellowship member Title                        Weapon      \n" +
		" Gimli             Lord of the Glittering Caves Axe         \n" +
		" Legolas           Prince of the Woodland Realm Bow         \n" +
		" Aragorn           King of Gondor               Sword       \n" +
		" Boromir           Captain of the White Tower   Sword       \n" +
		" Gandalf           The Grey                     Wizard staff\n"
)

func loadTable(t *testing.T) *lpstr.Table {
	tbl, _ := lpstr.New(header)
	tbl.AddRow(gandalf)
	tbl.AddRow(aragorn)
	tbl.AddRow(legolas)
	tbl.AddRow(gimli)
	tbl.SetPadding(padding)
	tbl.SortBy(sortby)
	tbl.AddRow(boromir)
	return tbl
}

func printTable(t *testing.T, tbl *lpstr.Table, wanted string) {
	str, _ := tbl.Print()
	fmt.Print(tbl)
	if str != wanted {
		t.Error("test failed")
	}
}

func TestTableWithoutGrid(t *testing.T) {
	tbl := loadTable(t)
	tbl.SetGrid(&lpstr.EmptyGrid)
	printTable(t, tbl, refWithoutGrid)

}

func TestTableWithGrid(t *testing.T) {
	tbl := loadTable(t)
	tbl.SetGrid(&lpstr.RoundGrid)
	printTable(t, tbl, refWithGrid)
}

func TestMinTable(t *testing.T) {
	tbl, _ := lpstr.New([]string{" ", " "})
	tbl.AddRow([]string{"", ""})
	tbl.SetPadding(0)
	printTable(t, tbl, refWithGrid)
}
