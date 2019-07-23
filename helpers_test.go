package main

import "testing"

func TestSplit(t *testing.T) {
	str := "hello,/usr/local/bin"
	wantShortcut := "hello"
	wantDir := "/usr/local/bin"

	gotShortcut, gotDir := Split(str, ",")

	if gotShortcut != wantShortcut {
		t.Errorf("Got: %s, want: %s", gotShortcut, wantShortcut)
	}
	if gotDir != wantDir {
		t.Errorf("Got: %s, want: %s", gotDir, wantDir)
	}
}
