// Copyright (c) 2023 thorstenrie.
// All Rights Reserved. Use is governed with GNU Affero General Public LIcense v3.0
// that can be found in the LICENSE file.
package lpstr_test

// Import Go standard packages, lpstr and tserr
import (
	"testing" // testing

	"github.com/thorstenrie/lpstr" // lpstr
	"github.com/thorstenrie/tserr" // tserr
)

// Testcases
var (
	testStrP  = "Andúril"       // Test string with only printable runes
	testStrNp = testStrP + "\n" // Test string with a non-printable rune
)

// TestPrintable1 tests Printable, if it returns the identical string, if the provided string a only
// contains printable runes. The test fails, if not.
func TestPrintable1(t *testing.T) {
	// Retrieve string from Printable for a test string with only printable runes
	str := lpstr.Printable(testStrP)
	// The test fails, if the retrived string does not equal the provided string with only printable runes
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
