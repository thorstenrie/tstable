# tstable
Go package for tables with a simple API


[![Go Report Card](https://goreportcard.com/badge/github.com/thorstenrie/tstable)](https://goreportcard.com/report/github.com/thorstenrie/tstable)
[![CodeFactor](https://www.codefactor.io/repository/github/thorstenrie/tstable/badge)](https://www.codefactor.io/repository/github/thorstenrie/tstable)
![OSS Lifecycle](https://img.shields.io/osslifecycle/thorstenrie/tstable)

[![PkgGoDev](https://pkg.go.dev/badge/mod/github.com/thorstenrie/tstable)](https://pkg.go.dev/mod/github.com/thorstenrie/tstable)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/thorstenrie/tstable)
![Libraries.io dependency status for GitHub repo](https://img.shields.io/librariesio/github/thorstenrie/tstable)

![GitHub release (latest by date)](https://img.shields.io/github/v/release/thorstenrie/tstable)
![GitHub last commit](https://img.shields.io/github/last-commit/thorstenrie/tstable)
![GitHub commit activity](https://img.shields.io/github/commit-activity/m/thorstenrie/tstable)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/thorstenrie/tstable)
![GitHub Top Language](https://img.shields.io/github/languages/top/thorstenrie/tstable)
![GitHub](https://img.shields.io/github/license/thorstenrie/tstable)

The Go package tstable provides a simple interface for tables. A new instance of a table can be retrieved with New and providing a table header. Table rows can be added
with AddRow. The table visualization can be altered with SetGrid and SetPadding. The package provides a set of grids or a grid can be customized. The string representation of a table is retrieved with Print. A table is sorted alphabetically by the first column. It can be sorted by other columns with SortBy.

- **Simple**: Without configuration, just function calls
- **Easy to use**: Just define the header of a table and add rows
- **Tested**: Unit tests with high code coverage.
- **Dependencies**: Only depends on the [Go Standard Library](https://pkg.go.dev/std), [tserr](https://github.com/thorstenrie/tserr), [lpstats](https://github.com/thorstenrie/lpstats) and [tsfio](https://github.com/thorstenrie/tsfio)

````
┌─────────────────────┬────────────────────────────────┬────────────────┐
│  Fellowship member  │  Title                         │  Weapon        │
├─────────────────────┼────────────────────────────────┼────────────────┤
│  Aragorn            │  King of Gondor                │  Sword         │
│  Boromir            │  Captain of the White Tower    │  Sword         │
│  Gandalf            │  The Grey                      │  Wizard staff  │
│  Gimli              │  Lord of the Glittering Caves  │  Axe           │
│  Legolas            │  Prince of the Woodland Realm  │  Bow           │
└─────────────────────┴────────────────────────────────┴────────────────┘
````

## Usage

The package is installed with 

````go
go get github.com/thorstenrie/tstable
````

In the Go app, the package is imported with

````go
import "github.com/thorstenrie/tstable"
````

## Table grid

A table grid has an outside border. The header row is separated from the table rows by a horizontal grid line. Table rows do not have a grid line between the rows. Columns are divided by an inside grid line. The package provides a set of grids for table string representation. A grid can be used by providing its reference to SetGrid, for example:

````go
tbl.SetGrid(&tstable.DoubleBorderGrid)
````

<details>
  <summary>See all included grids</summary>
	
  ````
DoubleBorderGrid
  ╔═════════════════════╤════════════════════════════════╤════════════════╗
  ║  Fellowship member  │  Title                         │  Weapon        ║
  ╟─────────────────────┼────────────────────────────────┼────────────────╢
  ║  Aragorn            │  King of Gondor                │  Sword         ║
  ║  Boromir            │  Captain of the White Tower    │  Sword         ║
  ║  Gandalf            │  The Grey                      │  Wizard staff  ║
  ║  Gimli              │  Lord of the Glittering Caves  │  Axe           ║
  ║  Legolas            │  Prince of the Woodland Realm  │  Bow           ║
  ╚═════════════════════╧════════════════════════════════╧════════════════╝
DoubleHorizontalGrid
  ╒═════════════════════╤════════════════════════════════╤════════════════╕
  │  Fellowship member  │  Title                         │  Weapon        │
  ╞═════════════════════╪════════════════════════════════╪════════════════╡
  │  Aragorn            │  King of Gondor                │  Sword         │
  │  Boromir            │  Captain of the White Tower    │  Sword         │
  │  Gandalf            │  The Grey                      │  Wizard staff  │
  │  Gimli              │  Lord of the Glittering Caves  │  Axe           │
  │  Legolas            │  Prince of the Woodland Realm  │  Bow           │
  ╘═════════════════════╧════════════════════════════════╧════════════════╛
DoubleGrid
  ╔═════════════════════╦════════════════════════════════╦════════════════╗
  ║  Fellowship member  ║  Title                         ║  Weapon        ║
  ╠═════════════════════╬════════════════════════════════╬════════════════╣
  ║  Aragorn            ║  King of Gondor                ║  Sword         ║
  ║  Boromir            ║  Captain of the White Tower    ║  Sword         ║
  ║  Gandalf            ║  The Grey                      ║  Wizard staff  ║
  ║  Gimli              ║  Lord of the Glittering Caves  ║  Axe           ║
  ║  Legolas            ║  Prince of the Woodland Realm  ║  Bow           ║
  ╚═════════════════════╩════════════════════════════════╩════════════════╝
RoundGrid
  ╭─────────────────────┬────────────────────────────────┬────────────────╮
  │  Fellowship member  │  Title                         │  Weapon        │
  ├─────────────────────┼────────────────────────────────┼────────────────┤
  │  Aragorn            │  King of Gondor                │  Sword         │
  │  Boromir            │  Captain of the White Tower    │  Sword         │
  │  Gandalf            │  The Grey                      │  Wizard staff  │
  │  Gimli              │  Lord of the Glittering Caves  │  Axe           │
  │  Legolas            │  Prince of the Woodland Realm  │  Bow           │
  ╰─────────────────────┴────────────────────────────────┴────────────────╯
SimpleGrid
  ┌─────────────────────┬────────────────────────────────┬────────────────┐
  │  Fellowship member  │  Title                         │  Weapon        │
  ├─────────────────────┼────────────────────────────────┼────────────────┤
  │  Aragorn            │  King of Gondor                │  Sword         │
  │  Boromir            │  Captain of the White Tower    │  Sword         │
  │  Gandalf            │  The Grey                      │  Wizard staff  │
  │  Gimli              │  Lord of the Glittering Caves  │  Axe           │
  │  Legolas            │  Prince of the Woodland Realm  │  Bow           │
  └─────────────────────┴────────────────────────────────┴────────────────┘
BoldGrid
  ┏━━━━━━━━━━━━━━━━━━━━━┯━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┯━━━━━━━━━━━━━━━━┓
  ┃  Fellowship member  │  Title                         │  Weapon        ┃
  ┠─────────────────────┼────────────────────────────────┼────────────────┨
  ┃  Aragorn            │  King of Gondor                │  Sword         ┃
  ┃  Boromir            │  Captain of the White Tower    │  Sword         ┃
  ┃  Gandalf            │  The Grey                      │  Wizard staff  ┃
  ┃  Gimli              │  Lord of the Glittering Caves  │  Axe           ┃
  ┃  Legolas            │  Prince of the Woodland Realm  │  Bow           ┃
  ┗━━━━━━━━━━━━━━━━━━━━━┷━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┷━━━━━━━━━━━━━━━━┛
EmptyGrid
  
    Fellowship member    Title                           Weapon        
  
    Aragorn              King of Gondor                  Sword         
    Boromir              Captain of the White Tower      Sword         
    Gandalf              The Grey                        Wizard staff  
    Gimli                Lord of the Glittering Caves    Axe           
    Legolas              Prince of the Woodland Realm    Bow           
  
DoubleVerticalGrid
  ╓─────────────────────╥────────────────────────────────╥────────────────╖
  ║  Fellowship member  ║  Title                         ║  Weapon        ║
  ╟─────────────────────╫────────────────────────────────╫────────────────╢
  ║  Aragorn            ║  King of Gondor                ║  Sword         ║
  ║  Boromir            ║  Captain of the White Tower    ║  Sword         ║
  ║  Gandalf            ║  The Grey                      ║  Wizard staff  ║
  ║  Gimli              ║  Lord of the Glittering Caves  ║  Axe           ║
  ║  Legolas            ║  Prince of the Woodland Realm  ║  Bow           ║
  ╙─────────────────────╨────────────────────────────────╨────────────────╜
InterruptedGrid
  ┏╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍┯╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍┯╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍┓
  ╏  Fellowship member  ╎  Title                         ╎  Weapon        ╏
  ┠╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌┼╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌┼╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌┨
  ╏  Aragorn            ╎  King of Gondor                ╎  Sword         ╏
  ╏  Boromir            ╎  Captain of the White Tower    ╎  Sword         ╏
  ╏  Gandalf            ╎  The Grey                      ╎  Wizard staff  ╏
  ╏  Gimli              ╎  Lord of the Glittering Caves  ╎  Axe           ╏
  ╏  Legolas            ╎  Prince of the Woodland Realm  ╎  Bow           ╏
  ┗╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍┷╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍┷╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍╍┛
DashedGrid
  ┏┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┯┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┯┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┓
  ┇  Fellowship member  ┆  Title                         ┆  Weapon        ┇
  ┠┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┼┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┼┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┄┨
  ┇  Aragorn            ┆  King of Gondor                ┆  Sword         ┇
  ┇  Boromir            ┆  Captain of the White Tower    ┆  Sword         ┇
  ┇  Gandalf            ┆  The Grey                      ┆  Wizard staff  ┇
  ┇  Gimli              ┆  Lord of the Glittering Caves  ┆  Axe           ┇
  ┇  Legolas            ┆  Prince of the Woodland Realm  ┆  Bow           ┇
  ┗┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┷┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┷┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┅┛
DottedGrid
  ┏┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┯┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┯┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┓
  ┋  Fellowship member  ┊  Title                         ┊  Weapon        ┋
  ┠┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┼┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┼┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┨
  ┋  Aragorn            ┊  King of Gondor                ┊  Sword         ┋
  ┋  Boromir            ┊  Captain of the White Tower    ┊  Sword         ┋
  ┋  Gandalf            ┊  The Grey                      ┊  Wizard staff  ┋
  ┋  Gimli              ┊  Lord of the Glittering Caves  ┊  Axe           ┋
  ┋  Legolas            ┊  Prince of the Woodland Realm  ┊  Bow           ┋
  ┗┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┷┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┷┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┉┛
  ````
</details>

A custom grid can also be provided to SetGrid. A custom grid is defined with the Grid struct type. The Grid struct type contains the runes to define the grid format of a table. A table grid is defined by thirteen runes. A rune is allowed to be empty.

````go
type Grid struct {
	Hi, Hb, Vi, Vb, Hvi, Hvl, Hvr, Hvt, Hvb, Hvtl, Hvbl, Hvtr, Hvbr rune
}
//	Hi:   	horizontal inside, separation between header and the rest of the table rows
//	Hb:	horizontal border, at the top and bottom of the table
//	Vi:	vertical inside, separation between table columns
//	Vb:	vertical border, at the left and right side of the table
//	Hvi:	horizontal vertical inside
//	Hvl:	horizontal vertical left
//	Hvr:	horizontal vertical right
//	Hvt:	horizontal vertical top
//	Hvb:	horizontal vertical bottom
//	Hvtl:	horizontal vertical top left
//	Hvbl:	horizontal vertical bottom left
//	Hvtr:	horizontal vertcial top right
//	Hvbr:	horizontal vertcial bottom right
````

| `Hvtl` | `Hb`       | `Hvt` | `Hb`       | `Hvtr` |
|------|----------|-----|----------|------|
| `Vb`   | header_1 | `vi`  | header_2 | `Vb`   |
| `Hvl`  | `Hi`       | `Hvi` | `Hi`       | `Hvr`  |
| `Vb`   | cell_11  | `Vi`  | cell_21  | `Vb`   |
| `Vb`   | cell_21  | `Hvb` | cell_22  | `Hvbr` |
| `Hvbl` | `Hb`       | `Hvb` | `Hb`       | `Hvbr` |

An example with a custom table grid is included in [example/example.go](https://github.com/thorstenrie/tstable/blob/main/example/example.go)

## Example

````go
package main

import (
	"fmt"

	"github.com/thorstenrie/tstable"
)

var (
	header = []string{"Fellowship member", "Title", "Weapon"}
	rows   = [][]string{
		{"Gandalf", "The Grey", "Wizard staff"},
		{"Aragorn", "King of Gondor", "Sword"},
		{"Legolas", "Prince of the Woodland Realm", "Bow"},
		{"Gimli", "Lord of the Glittering Caves", "Axe"},
		{"Boromir", "Captain of the White Tower", "Sword"},
	}
	sortby = "Weapon"
)

func main() {
	tbl, _ := tstable.New(header)
	for _, r := range rows {
		tbl.AddRow(r)
	}
	for n, g := range tstable.AllGrids {
		tbl.SetGrid(g)
		fmt.Println(n)
		fmt.Print(tbl)
	}
}

````
[Go Playground](https://go.dev/play/p/XsuobSpC0Di)

## Links

[Godoc](https://pkg.go.dev/github.com/thorstenrie/tstable)

[Go Report Card](https://goreportcard.com/report/github.com/thorstenrie/tstable)

[Open Source Insights](https://deps.dev/go/github.com%2Fthorstenrie%2Ftstable)
