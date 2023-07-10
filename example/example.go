// Copyright (c) 2023 thorstenrie.
// All Rights Reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package main

import (
	"fmt"

	"github.com/thorstenrie/tstable"
)

func main() {

	var (
		fancyGrid = tstable.Grid{
			Hi:   '\u2504',
			Hb:   '\u2504',
			Vi:   '\u2506',
			Vb:   '\u2506',
			Hvi:  '\u253C',
			Hvl:  '\u251C',
			Hvr:  '\u2524',
			Hvt:  '\u252C',
			Hvb:  '\u2534',
			Hvtl: '\u256D',
			Hvbl: '\u2570',
			Hvtr: '\u256E',
			Hvbr: '\u256F',
		}
		header = []string{"Fellowship member", "Title", "Weapon"}
		rows   = [][]string{
			{"Gandalf", "The Grey", "Wizard staff"},
			{"Aragorn", "King of Gondor", "Sword"},
			{"Legolas", "Prince of the Woodland Realm", "Bow"},
			{"Gimli", "Lord of the Glittering Caves", "Axe"},
			{"Boromir", "Captain of the White Tower", "Sword"},
		}
		sortby  = "Weapon"
		padding = 2
	)

	tbl, _ := tstable.New(header)
	for _, r := range rows {
		tbl.AddRow(r)
	}
	tbl.SetGrid(&fancyGrid)
	tbl.SetPadding(padding)
	tbl.SortBy(sortby)
	fmt.Print(tbl)
}
