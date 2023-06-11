// Copyright (c) 2023 thorstenrie.
// All Rights Reserved. Use is governed with GNU Affero General Public LIcense v3.0
// that can be found in the LICENSE file.
package lpstr

type Grid struct {
	hi, hb, vi, vb, hvi, hvl, hvr, hvt, hvb, hvtl, hvbl, hvtr, hvbr rune
}

var (
	EmptyGrid = Grid{}

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
)
