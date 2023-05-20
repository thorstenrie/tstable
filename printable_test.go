// Copyright (c) 2023 thorstenrie.
// All Rights Reserved. Use is governed with GNU Affero General Public LIcense v3.0
// that can be found in the LICENSE file.
package lpstr_test

import (
	"testing"

	"github.com/thorstenrie/lpstr"
	"github.com/thorstenrie/tserr"
)

var (
	testStrP  = "Andúril"
	testStrNp = testStrP + "\n"
)

func TestPrintable1(t *testing.T) {
	str := lpstr.Printable(testStrP)
	if str != testStrP {
		t.Error(tserr.Return(&tserr.ReturnArgs{Op: "Printable", Actual: str, Want: testStrP}))
	}
}

func TestPrintable2(t *testing.T) {
	str := lpstr.Printable(testStrNp)
	if str != testStrP {
		t.Error(tserr.Return(&tserr.ReturnArgs{Op: "Printable", Actual: str, Want: testStrP}))
	}
}

func TestIsPrintable1(t *testing.T) {
	str := []string{testStrP, testStrP}
	b, e := lpstr.IsPrintable(str)
	if e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "IsPrintable", Fn: "slice of strings", Err: e}))
	}
	if !b {
		t.Error(tserr.Return(&tserr.ReturnArgs{Op: "IsPrintable", Actual: "false", Want: "true"}))
	}
}

func TestIsPrintable2(t *testing.T) {
	str := []string{testStrNp, testStrP}
	b, e := lpstr.IsPrintable(str)
	if e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "IsPrintable", Fn: "slice of strings", Err: e}))
	}
	if b {
		t.Error(tserr.Return(&tserr.ReturnArgs{Op: "IsPrintable", Actual: "true", Want: "false"}))
	}
}

func TestIsPrintableEmpty1(t *testing.T) {
	str := []string{}
	testIsPrintableEmpty(t, str)
}

func TestIsPrintableEmpty2(t *testing.T) {
	var str []string = nil
	testIsPrintableEmpty(t, str)
}

func testIsPrintableEmpty(t *testing.T, str []string) {
	b, e := lpstr.IsPrintable(str)
	if e == nil {
		t.Error(tserr.NilFailed("IsPrintable"))
	}
	if b {
		t.Error(tserr.Return(&tserr.ReturnArgs{Op: "IsPrintable", Actual: "true", Want: "false"}))
	}
}
