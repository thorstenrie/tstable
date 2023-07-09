// Copyright (c) 2023 thorstenrie.
// All Rights Reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package tstable

// A Grid contains the runes to define the grid format of a table. A table grid is defined by thirteen runes.
// A rune is allowed to be empty.
//
//	hi:   	horizontal inside, separation between header and the rest of the table rows
//	hb:		horizontal border, at the top and bottom of the table
//	vi:		vertical inside, separation between table columns
//	vb:		vertical border, at the left and right side of the table
//	hvi:	horizontal vertical inside
//	hvl:	horizontal vertical left
//	hvr:	horizontal vertical right
//	hvt:	horizontal vertical top
//	hvb:	horizontal vertical bottom
//	hvtl:	horizontal vertical top left
//	hvbl:	horizontal vertical bottom left
//	hvtr:	horizontal vertcial top right
//	hvbr:	horizontal vertcial bottom right
//
// A table grid has an outside border. The header row is separated from the table rows by a horizontal grid line.
// Table rows do not have a grid line between the rows. Columns are divided by an inside grid line.
type Grid struct {
	hi, hb, vi, vb, hvi, hvl, hvr, hvt, hvb, hvtl, hvbl, hvtr, hvbr rune
}

var (
	// The EmptyGrid defines an empty table grid. A table with the EmptyGrid does not have a grid.
	EmptyGrid = Grid{}

	// The DoubleBorderGrid has a double-lined border.
	DoubleBorderGrid = Grid{
		hi:   '\u2500', // horizontal inside				─
		hb:   '\u2550', // horizontal border				═
		vi:   '\u2502', // vertical inside					│
		vb:   '\u2551', // vertical border					║
		hvi:  '\u253C', // horizontal vertical inside		┼
		hvl:  '\u255F', // horizontal vertical left			╟
		hvr:  '\u2562', // horizontal vertical right		╢
		hvt:  '\u2564', // horizontal vertical top			╤
		hvb:  '\u2567', // horizontal vertical bottom		╧
		hvtl: '\u2554', // horizontal vertical top left		╔
		hvbl: '\u255A', // horizontal vertical bottom left	╚
		hvtr: '\u2557', // horizontal vertical top right	╗
		hvbr: '\u255D', // horizontal vertical bottom right	╝
	}

	// The DoubleHorizontalGrid has double-lined horizontal lines.
	DoubleHorizontalGrid = Grid{
		hi:   '\u2550', // horizontal inside				─
		hb:   '\u2550', // horizontal border				═
		vi:   '\u2502', // vertical inside					│
		vb:   '\u2502', // vertical border					║
		hvi:  '\u256A', // horizontal vertical inside		┼
		hvl:  '\u255E', // horizontal vertical left			╟
		hvr:  '\u2561', // horizontal vertical right		╢
		hvt:  '\u2564', // horizontal vertical top			╤
		hvb:  '\u2567', // horizontal vertical bottom		╧
		hvtl: '\u2552', // horizontal vertical top left		╔
		hvbl: '\u2558', // horizontal vertical bottom left	╚
		hvtr: '\u2555', // horizontal vertical top right	╗
		hvbr: '\u255B', // horizontal vertical bottom right	╝
	}

	// The DoubleVerticalGrid has double-lined vertical lines.
	DoubleVerticalGrid = Grid{
		hi:   '\u2500', // horizontal inside				─
		hb:   '\u2500', // horizontal border				─
		vi:   '\u2551', // vertical inside					║
		vb:   '\u2551', // vertical border					║
		hvi:  '\u256B', // horizontal vertical inside		╫
		hvl:  '\u255F', // horizontal vertical left			╟
		hvr:  '\u2562', // horizontal vertical right		╢
		hvt:  '\u2565', // horizontal vertical top			╥
		hvb:  '\u2568', // horizontal vertical bottom		╨
		hvtl: '\u2553', // horizontal vertical top left		╓
		hvbl: '\u2559', // horizontal vertical bottom left	╙
		hvtr: '\u2556', // horizontal vertical top right	╖
		hvbr: '\u255C', // horizontal vertical bottom right	╜
	}

	// The DoubleGrid has a double-lined grid.
	DoubleGrid = Grid{
		hi:   '\u2550', // horizontal inside				═
		hb:   '\u2550', // horizontal border				═
		vi:   '\u2551', // vertical inside					║
		vb:   '\u2551', // vertical border					║
		hvi:  '\u256C', // horizontal vertical inside		╬
		hvl:  '\u2560', // horizontal vertical left			╠
		hvr:  '\u2563', // horizontal vertical right		╣
		hvt:  '\u2566', // horizontal vertical top			╦
		hvb:  '\u2569', // horizontal vertical bottom		╩
		hvtl: '\u2554', // horizontal vertical top left		╔
		hvbl: '\u255A', // horizontal vertical bottom left	╚
		hvtr: '\u2557', // horizontal vertical top right	╗
		hvbr: '\u255D', // horizontal vertical bottom right	╝
	}

	// The RoundGrid has rounded corners.
	RoundGrid = Grid{
		hi:   '\u2500', // horizontal inside				─
		hb:   '\u2500', // horizontal border				─
		vi:   '\u2502', // vertical inside					│
		vb:   '\u2502', // vertical border					│
		hvi:  '\u253C', // horizontal vertical inside		┼
		hvl:  '\u251C', // horizontal vertical left			├
		hvr:  '\u2524', // horizontal vertical right		┤
		hvt:  '\u252C', // horizontal vertical top			┬
		hvb:  '\u2534', // horizontal vertical bottom		┴
		hvtl: '\u256D', // horizontal vertical top left		╭
		hvbl: '\u2570', // horizontal vertical bottom left	╰
		hvtr: '\u256E', // horizontal vertical top right	╮
		hvbr: '\u256F', // horizontal vertical bottom right	╯
	}

	// The SimpleGrid has single grid lines
	SimpleGrid = Grid{
		hi:   '\u2500', // horizontal inside				─
		hb:   '\u2500', // horizontal border				─
		vi:   '\u2502', // vertical inside					│
		vb:   '\u2502', // vertical border					│
		hvi:  '\u253C', // horizontal vertical inside		┼
		hvl:  '\u251C', // horizontal vertical left			├
		hvr:  '\u2524', // horizontal vertical right		┤
		hvt:  '\u252C', // horizontal vertical top			┬
		hvb:  '\u2534', // horizontal vertical bottom		┴
		hvtl: '\u250C', // horizontal vertical top left		┌
		hvbl: '\u2514', // horizontal vertical bottom left	└
		hvtr: '\u2510', // horizontal vertical top right	┐
		hvbr: '\u2518', // horizontal vertical bottom right	┘
	}

	// The BoldGrid has bold border lines.
	BoldGrid = Grid{
		hi:   '\u2500', // horizontal inside				─
		hb:   '\u2501', // horizontal border				━
		vi:   '\u2502', // vertical inside					│
		vb:   '\u2503', // vertical border					┃
		hvi:  '\u253C', // horizontal vertical inside		┼
		hvl:  '\u2520', // horizontal vertical left			┠
		hvr:  '\u2528', // horizontal vertical right		┨
		hvt:  '\u252F', // horizontal vertical top			┯
		hvb:  '\u2537', // horizontal vertical bottom		┷
		hvtl: '\u250F', // horizontal vertical top left		┏
		hvbl: '\u2517', // horizontal vertical bottom left	┗
		hvtr: '\u2513', // horizontal vertical top right	┓
		hvbr: '\u251B', // horizontal vertical bottom right	┛
	}

	// The InterruptedGrid has interrupted grid lines.
	InterruptedGrid = Grid{
		hi:   '\u254C', // horizontal inside				╌
		hb:   '\u254D', // horizontal border				╍
		vi:   '\u254E', // vertical inside					╎
		vb:   '\u254F', // vertical border					╏
		hvi:  '\u253C', // horizontal vertical inside		┼
		hvl:  '\u2520', // horizontal vertical left			┠
		hvr:  '\u2528', // horizontal vertical right		┨
		hvt:  '\u252F', // horizontal vertical top			┯
		hvb:  '\u2537', // horizontal vertical bottom		┷
		hvtl: '\u250F', // horizontal vertical top left		┏
		hvbl: '\u2517', // horizontal vertical bottom left	┗
		hvtr: '\u2513', // horizontal vertical top right	┓
		hvbr: '\u251B', // horizontal vertical bottom right	┛
	}

	// The DashedGrid has dashed grid lines.
	DashedGrid = Grid{
		hi:   '\u2504', // horizontal inside				┄
		hb:   '\u2505', // horizontal border				┅
		vi:   '\u2506', // vertical inside					┆
		vb:   '\u2507', // vertical border					┇
		hvi:  '\u253C', // horizontal vertical inside		┼
		hvl:  '\u2520', // horizontal vertical left			┠
		hvr:  '\u2528', // horizontal vertical right		┨
		hvt:  '\u252F', // horizontal vertical top			┯
		hvb:  '\u2537', // horizontal vertical bottom		┷
		hvtl: '\u250F', // horizontal vertical top left		┏
		hvbl: '\u2517', // horizontal vertical bottom left	┗
		hvtr: '\u2513', // horizontal vertical top right	┓
		hvbr: '\u251B', // horizontal vertical bottom right	┛
	}

	// The DottedGrid has dotted grid lines.
	DottedGrid = Grid{
		hi:   '\u2508', // horizontal inside				┈
		hb:   '\u2509', // horizontal border				┉
		vi:   '\u250A', // vertical inside					┊
		vb:   '\u250B', // vertical border					┋
		hvi:  '\u253C', // horizontal vertical inside		┼
		hvl:  '\u2520', // horizontal vertical left			┠
		hvr:  '\u2528', // horizontal vertical right		┨
		hvt:  '\u252F', // horizontal vertical top			┯
		hvb:  '\u2537', // horizontal vertical bottom		┷
		hvtl: '\u250F', // horizontal vertical top left		┏
		hvbl: '\u2517', // horizontal vertical bottom left	┗
		hvtr: '\u2513', // horizontal vertical top right	┓
		hvbr: '\u251B', // horizontal vertical bottom right	┛
	}

	// AllGrids is a map which contains all Grids of the package. The map returns the *Grid when using the name of the Grid as key.
	AllGrids = map[string]*Grid{
		"EmptyGrid":            &EmptyGrid,
		"DoubleBorderGrid":     &DoubleBorderGrid,
		"DoubleHorizontalGrid": &DoubleHorizontalGrid,
		"DoubleVerticalGrid":   &DoubleVerticalGrid,
		"DoubleGrid":           &DoubleGrid,
		"RoundGrid":            &RoundGrid,
		"SimpleGrid":           &SimpleGrid,
		"BoldGrid":             &BoldGrid,
		"InterruptedGrid":      &InterruptedGrid,
		"DashedGrid":           &DashedGrid,
		"DottedGrid":           &DottedGrid,
	}
)
