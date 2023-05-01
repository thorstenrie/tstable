package lpstr_test

import (
	"fmt"
	"testing"

	"github.com/thorstenrie/lpstr"
)

var (
	header  = []string{"Fellowship member", "Title", "Weapon"}
	gandalf = []string{"Gandalf", "The Grey", "Wizard staff"}
	aragorn = []string{"Aragorn", "King of Gondor", "Sword"}
	legolas = []string{"Legolas", "Prince of the Woodland Realm", "Bow"}
	gimli   = []string{"Gimli", "Lord of the Glittering Caves", "Axe"}
	boromir = []string{"Boromir", "Captain of the White Tower", "Sword"}
	sortby  = "Weapon"
)

func TestTable(t *testing.T) {
	tbl, _ := lpstr.New(header)
	tbl.AddRow(gandalf)
	tbl.AddRow(aragorn)
	tbl.AddRow(legolas)
	tbl.AddRow(gimli)
	tbl.SetPadding(1)
	//tbl.WithoutGrid()
	tbl.SortBy(sortby)
	tbl.AddRow(boromir)
	str, _ := tbl.Print()
	fmt.Printf("%s", str)
}
