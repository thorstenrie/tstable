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

## Usage

The package is installed with 

```
go get github.com/thorstenrie/tstable
```

In the Go app, the package is imported with

```
import "github.com/thorstenrie/tstable"
```

## Table grid

| `hvtl` | `hb`       | `hvt` | `hb`       | `hvtr` |
|------|----------|-----|----------|------|
| `vb`   | header_1 | vi  | header_2 | `vb`   |
| `hvl`  | `hi`       | `hvi` | `hi`       | `hvr`  |
| `vb`   | cell_11  | `vi`  | cell_21  | `vb`   |
| `vb`   | cell_21  | `hvb` | cell_22  | `hvbr` |
| `hvbl` | `hb`       | `hvb` | `hb`       | `hvbr` |

## Example

```
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
	tbl.SetGrid(&tstable.RoundGrid)
	fmt.Print(tbl)
}
```
[Go Playground](https://go.dev/play/p/a-JNpKkRSxH)

An example with a custom table grid is included in [example/example.go](https://github.com/thorstenrie/tstable/blob/main/example/example.go)

## Links

[Godoc](https://pkg.go.dev/github.com/thorstenrie/tstable)

[Go Report Card](https://goreportcard.com/report/github.com/thorstenrie/tstable)

[Open Source Insights](https://deps.dev/go/github.com%2Fthorstenrie%2Ftstable)
