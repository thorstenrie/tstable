// Copyright (c) 2023 thorstenrie.
// All Rights Reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package lpstr_test

import (
	"testing"

	"github.com/thorstenrie/lpstr"
	"github.com/thorstenrie/tserr"
)

func TestGrids(t *testing.T) {
	tbl := loadTable(t)
	for name, grid := range lpstr.AllGrids {
		e := tbl.SetGrid(grid)
		if e != nil {
			t.Error(tserr.Op(&tserr.OpArgs{Op: "SetGrid", Fn: name, Err: e}))
		}
		s, e := tbl.Print()
		if e != nil {
			t.Error(tserr.Op(&tserr.OpArgs{Op: "Print", Fn: name, Err: e}))
		}
		/*fn := goldenFn(name)
		if err := tsfio.WriteSingleStr(fn, s); err != nil {
			t.Error(tserr.Op(&tserr.OpArgs{Op: "WriteSingleStr", Fn: string(fn), Err: err}))
		}*/
		want := testData(name, t)
		if s != want {
			t.Error(tserr.NotEqualStr(&tserr.NotEqualStrArgs{X: s, Y: want}))
		}
	}
}
